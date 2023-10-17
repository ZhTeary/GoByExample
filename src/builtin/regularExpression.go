package builtin

import (
	"bytes"
	"fmt"
	"regexp"
)

func RegularExp() {
	//tests
	//match, _ := regexp.MatchString("p([a-z]+)ch", "peach")
	//fmt.Println(match)

	//optimized Regexp struct
	r, _ := regexp.Compile("p([a-z]+)ch")

	//if match
	fmt.Println(r.MatchString("peach"))

	//finds the first match
	fmt.Println(r.FindString("peach punch"))

	//first match indexes
	fmt.Println("idx:", r.FindStringIndex("peach punch"))

	//whole-pattern matches contains sub matches. p([a-z]+)ch ([a-z]+)
	fmt.Println(r.FindStringSubmatch("peach punch"))

	//参数n 是控制匹配个数的 -1代表匹配所有
	fmt.Println(r.FindAllString("peach punch pinch", -1))
	fmt.Println(r.FindAllString("peach punch pinch", 2))

	fmt.Println("all:", r.FindAllStringSubmatchIndex("peach punch pinch", -1))

	//match []byte
	fmt.Println(r.Match([]byte("peach")))
	r = regexp.MustCompile("p([a-z]+)ch")
	fmt.Println("regexp:", r)

	fmt.Println(r.ReplaceAllString("a peach", "<fruit>"))

	in := []byte("peach")
	out := r.ReplaceAllFunc(in, bytes.ToUpper)
	fmt.Println(string(out))

}
