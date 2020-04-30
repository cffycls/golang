package main

import "fmt"

func minCount(coins []int) int {
    num := 0
    for _,v := range coins {
        num += v/2
        if v/2*2 < v {
            num++
        }
    }

    return num
}
func main()  {
    c := []int{4,2,1}
    fmt.Println(minCount(c))
}