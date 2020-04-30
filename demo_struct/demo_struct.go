package demo_struct

import "fmt"

type Animal struct {
    Color string
}

//定义dog结构体
type Dog struct {
    Animal
    Id int
    Name string
    Age int
}

func (a *Animal)Eat() string {
    fmt.Println("Animal is Eatting")
    return "Eat yummy yummy!!"
}
func (d *Dog)Run() string {
    fmt.Println("dog is running: ", d.Id, d.Name,d.Age)
    d.see()
    return "Run happy happy!!"
}

//定义catog结构体
type Cat struct {
    Animal
    Id int
    Name string
    Age int
}

func (a *Cat)Eat() string {
    fmt.Println("Cat-- is Eatting")
    return "Cat-- Eat yummy yummy!!"
}
func (d *Cat)Run() string {
    fmt.Println("Cat-- is running: ", d.Id, d.Name,d.Age)
    d.see()
    return "Cat-- Run happy happy!!"
}
//小写外部不能执行
func (d *Animal)see() string {
    d.Color = "default"
    fmt.Println("Animal is seeing: ", d.Color)
    return "see happy happying!!"
}


func TestStruct()  {
    var dog Dog
    dog.Id = 1
    dog.Name = "aHuang"
    dog.Age = 2
    dog.Color = "yellow"

    fmt.Println("dog: ", dog)
}