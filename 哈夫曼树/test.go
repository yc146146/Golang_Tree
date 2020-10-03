package main

import (
	"fmt"
)

type person struct {
	name string
	age int
}

//var b *Books

//b = &Book1
//fmt.Println(b)    //&{Go 语言 www.runoob.com Go 语言教程 6495407}
//fmt.Println(*b)   //{Go 语言 www.runoob.com Go 语言教程 6495407}
//fmt.Println(&b)   //0xc000082018
//fmt.Println(Book1)    //{Go 语言 www.runoob.com Go 语言教程 6495407}

//var b *Books     //就是说b这个指针是Books类型的。
//b  = &Book1     //Book1是Books的一个实例化的结构，&Book1就是把这个结构体的内存地址赋给了b，
//*b         //那么在使用的时候，只要在b的前面加个*号，就可以把b这个内存地址对应的值给取出来了
//&b        // 就是取了b这个指针的内存地址，也就是b这个指针是放在内存空间的什么地方的。
//Book1       // 就是Book1这个结构体，打印出来就是它自己。也就是指针b前面带了*号的效果。

func (p *person)init(){

	p.name="a1"
	p.age = 10

	//fmt.Println(p.name)
	//fmt.Println((*p).name)

}

func main2() {

	var p *person
	p = new(person)

	p.init()

	fmt.Println(p)
	fmt.Println(*p)
	fmt.Println(&p)

	p2 := person{"b1",30}

	fmt.Println(p2)
	//fmt.Println(*p2)
	fmt.Println(&p2)

	var ptr *person

	ptr = &p2
	fmt.Println()
	fmt.Println(ptr)
	fmt.Println(*ptr)
	fmt.Println(&ptr)



}
