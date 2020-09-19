package main

import (
	"container/heap"
	"fmt"
)

type HuffmanTree interface {
	//哈夫曼树接口
	Freq()int

}

type HuffmanLeaf struct {
	//频率
	freq int
	//int 32
	value rune
}


//哈夫曼树类型
type HuffmanNode struct {
	freq int
	left, right HuffmanTree
}

func (self HuffmanLeaf)Freq()int{
	return self.freq
}

func (self HuffmanNode)Freq()int{
	return self.freq
}

type treeHeap [] HuffmanTree


//求长度
func (th treeHeap) Len() int{
	return len(th)
}


//比较函数
func (th treeHeap) Less(i int, j int)bool{
	return th[i].Freq() < th[j].Freq()
}

//压入
func (th * treeHeap)Push(ele interface{}){
	*th = append(*th, ele.(HuffmanTree))
}

//弹出
func (th * treeHeap)Pop()(po interface{}){
	po = (*th)[len(*th)-1]
	*th = (*th)[:len(*th)-1]
	return
}

func (th treeHeap) Swap(i,j int){
	th[i],th[j] = th[j],th[i]
}

func BuildTree(symFreqs map[rune]int)HuffmanTree{
	var trees treeHeap
	for c,f := range symFreqs{
		trees = append(trees, HuffmanLeaf{f,c})
	}
	heap.Init(&trees)
	for trees.Len()>1{
		a:=heap.Pop(&trees).(HuffmanTree)
		b:=heap.Pop(&trees).(HuffmanTree)
		heap.Push(&trees, HuffmanNode{a.Freq()+b.Freq(),a,b})

	}
	//构造哈夫曼树
	return heap.Pop(&trees).(HuffmanTree)
}


func showtimes(tree HuffmanTree, prefix []byte){
	switch i:=tree.(type) {
	case HuffmanLeaf:
		//打印数据与频率
		//fmt.Println(i.value,i.freq)
		fmt.Printf("%c\t%d\n",i.value, i.freq)
	case HuffmanNode:
		prefix = append(prefix, '0')
		//递归到左子树
		showtimes(i.left, prefix)
		//删除最后一个
		prefix=prefix[:len(prefix)-1]

		prefix = append(prefix, '1')
		showtimes(i.right, prefix)
		prefix=prefix[:len(prefix)-1]
	}
}

func showcodes(tree HuffmanTree, prefix []byte){
	switch i:=tree.(type) {
	case HuffmanLeaf:
		//打印数据与频率
		//fmt.Println(i.value,i.freq)
		fmt.Printf("%s\n",string(prefix))
	case HuffmanNode:
		prefix = append(prefix, '0')
		//递归到左子树
		showcodes(i.left, prefix)
		//删除最后一个
		prefix=prefix[:len(prefix)-1]

		prefix = append(prefix, '1')
		showcodes(i.right, prefix)
		prefix=prefix[:len(prefix)-1]
	}
}

func main() {
	stringcode := "aaaabbbccceefff"

	fmt.Println("stringcode", stringcode)

	symFreqs := make(map[rune]int)

	for _,c := range stringcode{
		//统计频率
		symFreqs[c]++
	}

	//fmt.Println(symFreqs)
	//fmt.Println("a",int('a'))

	trees := BuildTree(symFreqs)
	showtimes(trees, []byte{})
	fmt.Println("----------------------")
	showcodes(trees, []byte{})
}


