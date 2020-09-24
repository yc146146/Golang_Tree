package main

import (
	"errors"
	"fmt"
)

type HashNode struct {
	Key    int
	Value  int
	Childs map[int]*HashNode
}

//插入数据
func (c *HashNode) AddValueRecursize(keys []int) {

	//记录层次
	c.Value += 1
	if len(keys) == 0 {
		return
	}

	childNode, ok := c.Childs[keys[0]]

	if !ok {
		//新建一个节点
		childNode = &HashNode{Key: keys[0]}
		if c.Childs == nil {
			c.Childs = make(map[int]*HashNode)
		}
		c.Childs[keys[0]] = childNode
	}

	if len(keys) > 1 {
		childNode.AddValueRecursize(keys[1:])
	} else if len(keys) == 1 {
		childNode.Value += 1
	}

}

//插入数据 无需创建
func (c *HashNode) AddValueWithoutCreate(keys []int) error {
	c.Value += 1
	if len(keys) == 0 {
		return nil
	}

	childNode, ok := c.Childs[keys[0]]

	if !ok {

		return errors.New("no key for node")
	}

	if len(keys) > 1 {
		childNode.AddValueRecursize(keys[1:])
	} else if len(keys) == 1 {
		childNode.Value += 1
	}
	return nil
}

//插入数据不存储次数
func (c *HashNode) AddNodeWithoutValue(keys []int) error {

	if len(keys) == 0 {
		return nil
	}


	childNode, ok := c.Childs[keys[0]]

	if !ok {

		//新建一个节点
		childNode = &HashNode{Key: keys[0]}
		if c.Childs == nil {
			c.Childs = make(map[int]*HashNode)
		}
		c.Childs[keys[0]] = childNode
	}

	if len(keys) > 1 {
		childNode.AddValueRecursize(keys[1:])
	} else if len(keys) == 1 {
		childNode.Value += 1
	}


	return nil
}

//提取数据
func (c *HashNode) GetValueRecursive(keys []int) (int , error){
	if len(keys) == 0 {
		return c.Value,nil
	}

	childNode, ok := c.Childs[keys[0]]

	if !ok {

		return 0, errors.New("节点不存在")
	}else{
		return childNode.GetValueRecursive(keys[1:])
	}
}

func main() {
	root := &HashNode{}
	root.AddValueRecursize([]int{0,2,3,4,5})
	fmt.Println(root.GetValueRecursive([]int{1}))
}
