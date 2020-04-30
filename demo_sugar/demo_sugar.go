package demo_sugar

import "fmt"

//可变长度的参数，:=声明、赋值、类型推断
func Sugar(values ...string)  {
    for _, v := range values {
        fmt.Println("v ---> ", v)
    }
}

