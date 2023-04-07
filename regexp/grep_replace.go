package main

import (
	"fmt"
	"regexp"
	"strings"
)

func main() {
	s := "hello 123 haha 456"
	reg := regexp.MustCompile("[0-9]+")
	res := reg.ReplaceAllString(s, "- ")
	fmt.Println(res) // <hello -  haha - >

	name := "邮政"
	mark := "|"
	reg = regexp.MustCompile("\\s|\\r\\n，|/|,")
	name = reg.ReplaceAllString(name, mark)
	fmt.Println("name:", name)

	arr := strings.Split(name, mark)
	fmt.Println("arr:", arr)
	reg = regexp.MustCompile("\\(.+")
	var arr2 = make([]string, len(arr))
	for i := 0; i < len(arr); i++ {
		fmt.Println("arr[i]:", arr[i])
		arr2[i] = reg.ReplaceAllString(arr[i], "")
		fmt.Println("arr2[i]:", arr2[i])
	}
	result := strings.Join(arr2, mark)
	fmt.Println("result:", result)

}
