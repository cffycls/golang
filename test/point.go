package main

import (
	"fmt"
	"reflect"
)
type person struct {
	name string
}
func (p person) String1() string{ return "person is "+p.name }
func (p *person) String2() string{ return "*person is "+p.name }
func (p *person) String3() string{
	fmt.Println("*TypeOf(p) 3: " + reflect.TypeOf(p).String())
	p.name = "*new 3" + p.name
	return "person is "+p.name
}
func (p person) String4() string {
	fmt.Println("TypeOf(p) 4: " + reflect.TypeOf(p).String())
	p.name = "*new 4" + p.name
	return "person is "+p.name
}
func (p person) modify1(){ p.name = "李四" }
func (p *person) modify2(){ p.name = "王五" }

func main() {
	//实例化后是个指针类型，需要取指针获得变量值
	p := person{name: "张三"}
	fmt.Println("person modify:")
	p.modify1() //值接收者，修改无效
	fmt.Println("pp1: "+ p.String1())
	fmt.Println("p*2: "+ p.String2())
	p = person{name: "张三"}
	(&p).modify1() //修改无效
	fmt.Println("pp1: "+ p.String1())
	fmt.Println("p*2: "+ p.String2())
	fmt.Println("p*3: "+ p.String3())
	fmt.Println("pp4: "+ p.String4())
	fmt.Println()

	p = person{name: "张三"}
	fmt.Println("*person modify:")
	p.modify2() //指针接收者，修改成功
	fmt.Println("*p: "+ p.String1())
	fmt.Println("**: "+ p.String2())
	p = person{name: "张三"}
	(&p).modify2() //修改成功
	fmt.Println("*p: "+ p.String1())
	fmt.Println("**: "+ p.String2())
	fmt.Println("**3: "+ p.String3())
	fmt.Println("*p4: "+ p.String4())
}