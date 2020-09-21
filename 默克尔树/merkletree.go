package main

import (
	"crypto/sha256"
	"encoding/hex"
	"fmt"
)

func min(a int, b int)int{
	if a>b{
		return b
	}else{
		return a
	}
}


//默克尔书
type MerkleTree struct {
	RootNode *MerkleNode
}

//默克尔输的节点
type MerkleNode struct {
	Left     *MerkleNode
	Right    *MerkleNode
	HashData []byte
}

//叶子节点 非叶子节点
func NewMerkleNode(left, right *MerkleNode, hashdata []byte) *MerkleNode {
	//创建结构体
	mynode := new(MerkleNode)
	if left == nil && right == nil {
		//赋值数据
		mynode.HashData = hashdata
	} else {
		//叠加数据
		prehashes := append(left.HashData, right.HashData...)
		//计算哈希
		hashleftright := sha256.Sum256(prehashes)
		//截取数据
		hash := sha256.Sum256(hashleftright[:])
		mynode.HashData = hash[:]
	}

	mynode.Left = left
	mynode.Right = right
	return mynode
}

//数组反转
func ReverseByte(data []byte)[]byte{
	for i,j := 0,len(data)-1;i<j;i,j=i+1,j-1{
		data[i],data[j] = data[j],data[i]
	}
	return data

}

//构造默克尔树
func NewMerkleTree(dataxx [][]byte) *MerkleTree {
	//叶子结合
	var nodes []MerkleNode
	for _, datax := range dataxx {
		node := NewMerkleNode(nil,nil,datax)
		nodes = append(nodes, *node)
	}
	//每一层的第一个元素
	j:=0
	//每次折半处理
	for length := len(dataxx);length>1;length=(length+1)/2{
		for i:=0;i<length;i+=2{
			half := min(i+1, length-1)
			node := NewMerkleNode(&nodes[j+i],&nodes[j+half], nil)
			nodes = append(nodes, *node)
		}
		//完成一个长度的叠加
		j += length

	}

	myTree := MerkleTree{&nodes[len(nodes)-1]}
	return &myTree
}

func main() {
	data1,_ := hex.DecodeString("5cd1e336b2a0e4b7491c020aedbb9a51211f0fe996337e351430dd7345deb55a")
	data2,_ := hex.DecodeString("4ea2a51f1753103c07af75fb9f5b8aaf60a7ba32f60a9fa8a441f9991102ce5b")
	data3,_ := hex.DecodeString("58f003250c11f842b5e84ddfb44ada61046a21e257f2bec558cb228b2ee3e9f3")
	data4,_ := hex.DecodeString("ae2dbac0e69a9fff6e17dbcba8d281a201d22ce3e2cff1a71a0375ca5b873d04")
	data5,_ := hex.DecodeString("9d5bdd9aeea6b7a02b1f68fe20d81a567c35ef32a4e5bed6409a447491067fd3")
	datax := [][]byte{data1,data2,data3,data4,data5}
	datay := [][]byte{ReverseByte(data1),data2,data3,data4,data5}

	myroot1:=NewMerkleTree(datax)
	myroot2:=NewMerkleTree(datay)
	fmt.Printf("%x\n", myroot1.RootNode.HashData)
	fmt.Printf("%x\n", myroot2.RootNode.HashData)

}