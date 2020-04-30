package main

import "fmt"

func eat(n int) int {
    if n==10 {
        return 1
    }else{
        return (eat(n+1) + 1) * 2
    }
}

func main()  {
    // func(n) = (func(n+1) + 1) * 2
    date := 1
    fmt.Println(eat(date))

}
