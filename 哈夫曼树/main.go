package main

import (
	"bytes"
	"fmt"
	"github.com/dgryski/go-bitstream"
	"io/ioutil"
)

func main() {
	strpath := "./word.txt"

	data, err := ioutil.ReadFile(strpath)

	if err != nil {
		fmt.Println(err)
	}

	//二进制
	fmt.Println(data)

	counts := make([]int, 256)
	for _, v := range data {
		counts[v]++
	}
	fmt.Println(counts)

	//编码
	e := NewEncoder(counts)

	fmt.Println(e)

	var b bytes.Buffer
	w := e.Write(&b)

	for _, v := range data {
		//写入
		w.WriteSymbol(uint32(v))
	}
	w.WriteSymbol(EOF)
	w.Close()

	fmt.Printf("未压缩的数据%s\n", data)
	fmt.Println("压缩的数据", b.Bytes())
	fmt.Printf("压缩的数据%s\n", b.Bytes())

	cbb := e.CodebookBytes()
	d, err := NewDecoder(cbb)

	br := bitstream.NewReader(bytes.NewReader(b.Bytes()))
	//未压缩的锁具
	var uncompressed []byte

	for {
		b, err:= d.ReadSymbol(br)
		if err != nil{
			return
		}
		if b == EOF{
			break
		}
		uncompressed = append(uncompressed,byte(b))
	}

	fmt.Println("解压缩",uncompressed)
	fmt.Printf("解压缩%s\n", uncompressed)

}
