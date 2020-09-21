package main

import "fmt"

func main1() {

	s := NewKeyWordTree()
	s.Put(1, "ba")
	s.Put(2, "ab")
	s.Put(3, "abc")
	s.Put(4, "abcd")
	s.Put(5, "bcd")
	//fmt.Println(s.Search("a", 4))
	fmt.Println(s.Sugg("b",4))
	fmt.Println(s.Sugg("a",4))


}

func rev(keyword string)string{
	runes := []rune(keyword)
	for from, to := 0, len(runes)-1;from<to;from,to=from+1,to-1{
		runes[from], runes[to] = runes[to], runes[from]
	}
	return string(runes)
}


func main2() {
	fmt.Println(rev("abcd123"))
}

func main() {
	s := NewKeyWordTree()
	s.Put(1, "ba")
	s.Put(2, "abd")
	s.Put(3, "abcd")
	s.Put(4, "abcd")
	s.Put(5, "bcd")

	fmt.Println(s.Sugg("d",4))

}