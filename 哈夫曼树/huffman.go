package main

import (
	"bytes"
	"container/heap"
	"encoding/binary"
	"errors"
	"github.com/dgryski/go-bitstream"
	"io"
	"sort"
)

//结束
var EOF uint32 = 0xFFFFFFFF
var ErrInvaild = errors.New("cook book fails")

//处理哈夫曼树权重
type node struct {
	weight int
	child  [2]*node
	leaf   bool
	sym    uint32
}

type nodes []node

func (n nodes) Len() int { return len(n) }

func (n nodes) Swap(i, j int)       { n[i], n[j] = n[j], n[i] }
func (n nodes) Less(i, j int) bool  { return n[i].weight < n[j].weight }
func (n *nodes) Push(x interface{}) { *n = append(*n, x.(node)) }

//弹出一个数据
func (n *nodes) Pop() interface{} {
	old := *n
	length := len(old)
	x := old[length-1]
	*n = old[0 : length-1]
	return x
}

type symbol struct {
	s      uint32
	code   uint32
	length int
}

type codebook []symbol

func (c codebook) calcCodes() (symptrs, []uint32) {
	var sptrs symptrs
	for i := range c {
		if c[i].length != 0 {
			//数据叠加
			sptrs = append(sptrs, &c[i])
		}
	}
	sort.Sort(sptrs)
	numl := make([]uint32, sptrs[len(sptrs)-1].length+1)

	prevlen := -1
	var code uint32
	for i := range sptrs {
		if sptrs[i].length > prevlen {
			//代码叠加
			code <<= uint(sptrs[i].length - prevlen)
			//取出长度, 记录上一个
			prevlen = sptrs[i].length
		}
		numl[sptrs[i].length]++
		sptrs[i].code = code
		code++
	}

	return sptrs, numl

}

//自身类型转化为byte
func (c codebook) MarshalBin() ([]byte, error) {
	var b []byte
	var vbuf [binary.MaxVarintLen32]byte

	//长度加入
	l := binary.PutUvarint(vbuf[:], uint64(len(c)))
	//字符叠加
	b = append(b, vbuf[:l]...)

	for i := range c {
		l := binary.PutUvarint(vbuf[:], uint64(c[i].length))
		b = append(b, vbuf[:l]...)
	}

	return b, nil

}

//二进制数据转化为自身对象
func (c *codebook) UnMarshalBin(data []byte) error {
	//读取数据
	r := bytes.NewBuffer(data)
	//读取数据
	l, err := binary.ReadUvarint(r)

	if err != nil {
		return ErrInvaild
	}
	*c=make(codebook,l)
	for i := uint32(0); i < uint32(l); i++ {
		clen, err := binary.ReadUvarint(r)
		if err != nil {
			return ErrInvaild
		}
		//构造 对象
		(*c)[i] = symbol{s: i, length: int(clen)}
	}

	return nil

}

//数组 内部包含指针
type symptrs []*symbol

func (s symptrs) Len() int      { return len(s) }
func (s symptrs) Swap(i, j int) { s[i], s[j] = s[j], s[i] }
func (s symptrs) Less(i, j int) bool {
	return s[i].length < s[j].length || s[i].length == s[j].length && s[i].s < s[j].s
}

//编码
type Encoder struct {
	//结束
	eof uint32
	//处理编码的对象
	m codebook
	//编码的数组
	sym symptrs

	numl []uint32
}
type Writer struct {
	e *Encoder
	//二进制写入
	*bitstream.BitWriter
	closed bool
}

func NewEncoder(counts []int) *Encoder {
	var n nodes
	for i, v := range counts {
		if v != 0 {
			heap.Push(&n, node{weight: v, leaf: true, sym: uint32(i)})
		}
	}

	//按照长度读取
	eof := uint32(len(counts))

	heap.Push(&n, node{weight: 0, leaf: true, sym: eof})
	for n.Len() > 1 {
		n1 := heap.Pop(&n).(node)
		n2 := heap.Pop(&n).(node)
		heap.Push(&n, node{weight: n1.weight + n2.weight, child: [2]*node{&n2, &n1}})

	}

	m := make(codebook, eof+1)
	//循环比那里每个节点
	walk(&n[0], 0, m)
	sptrs, numl := m.calcCodes()

	return &Encoder{
		eof:  eof,
		m:    m,
		sym:  sptrs,
		numl: numl,
	}

}

//遍历整个树根
func walk(n *node, depth int, m codebook) {
	if n.leaf {
		//遍历
		m[n.sym] = symbol{s: n.sym, length: depth}
		return
	}

	walk(n.child[0], depth+1, m)
	walk(n.child[1], depth+1, m)
}

func (e *Encoder) Symbolen(s uint32) int {
	if s == EOF {
		s = e.eof
	}
	if s >= uint32(len(e.m)) {
		return 0
	}
	return e.m[s].length
}

//新建一个写入对象
func (e *Encoder) Write(w io.Writer) *Writer {
	return &Writer{e: e, BitWriter: bitstream.NewWriter(w)}

}

//返回二进制
func (e *Encoder) CodebookBytes() []byte {
	b, _ := e.m.MarshalBin()
	return b
}

//解码器
func (e *Encoder) Decoder() *Decoder {
	return &Decoder{eof: e.eof, numl: e.numl, sym: e.sym}
}

func (w *Writer) WriteSymbol(s uint32) (int, error) {
	if s == EOF {
		s = w.e.eof
	}

	if s > EOF {
		return 0, ErrInvaild
	}

	sym := w.e.m[s]
	w.WriteBits(uint64(sym.code), sym.length)
	return sym.length, nil

}

func (w *Writer) Close() {
	if w.closed {
		return
	}

	//刷新
	w.Flush(bitstream.Zero)
}

//解码
type Decoder struct {
	eof uint32
	sym symptrs

	numl []uint32
}

func NewDecoder(cb []byte) (*Decoder, error) {
	var c codebook
	//读取异常
	if err := c.UnMarshalBin(cb); err != nil {
		return nil, err
	}

	//空计算
	sptrs, numl := c.calcCodes()
	eof := uint32(len(c)) - 1
	return &Decoder{eof: eof, numl: numl, sym: sptrs}, nil

}

func (d *Decoder) ReadSymbol(br *bitstream.BitReader) (uint32, error) {
	var offset uint32
	var code uint32
	for i := 0; i < len(d.numl); i++ {
		b, err := br.ReadBit()
		if err != nil {
			return 0, err
		}
		code <<= 1
		if b {
			code |= 1
		}

		offset += d.numl[i]
		//添加便宜系数
		first := d.sym[offset].code

		if code-first<d.numl[i+1]{
			s := d.sym[code-first+offset].s
			if s==d.eof{
				s=EOF
			}
			return s,nil
		}

	}
	return 0,ErrInvaild
}
