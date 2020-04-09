package main
import "fmt"
type person struct {
	name string
}
func (p person) String1() string{ return "person is "+p.name }
func (p *person) String2() string{ return "person is "+p.name }
func (p person) modify1(){ p.name = "李四" }
func (p *person) modify2(){ p.name = "找吧" }

func main() {
	//实例化后是个指针类型，需要取指针获得变量值
	p := person{name: "张三"}
	fmt.Println("person:")
	p.modify1() //值传递，修改无效
	fmt.Println("__|__: "+ p.String1())
	fmt.Println("__|_*: "+ p.String2())
	fmt.Println("__|&_: "+ (&p).String1())
	fmt.Println("__|&*: "+ (&p).String2())
	(&p).modify1() //修改无效
	fmt.Println("_&|__: "+ p.String1())
	fmt.Println("_&|_*: "+ p.String2())
	fmt.Println("_&|&_: "+ (&p).String1())
	fmt.Println("_&|&*: "+ (&p).String2())
	fmt.Println()

	p = person{name: "张三"}
	fmt.Println("*person:")
	p.modify2() //修改成功
	fmt.Println("*_|__: "+ p.String1())
	fmt.Println("*_|_*: "+ p.String2())
	fmt.Println("*_|&_: "+ (&p).String1())
	fmt.Println("*_|&*: "+ (&p).String2())
	(&p).modify2() //修改成功
	fmt.Println("*&|__: "+ p.String1())
	fmt.Println("*&|_*: "+ p.String2())
	fmt.Println("*&|&_: "+ (&p).String1())
	fmt.Println("*&|&*: "+ (&p).String2())
}