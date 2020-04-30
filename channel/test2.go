package main

import (
    "fmt"
    "math/rand"
    "strconv"
    "time"
)

type array2j struct {
    a []string
    b string

}

func main() {
    ch := make(chan string, 3)
    c2 := make(chan string)
    var queue array2j
    for i:=1; i<=5; i++ {
        go func(i int) {
            time.Sleep(time.Duration(rand.Intn(100)) * time.Microsecond)
            fmt.Println("go func:" + strconv.Itoa(i))
            ch <- strconv.Itoa(i) + "_ch_" + strconv.Itoa(rand.Int())
        }(i)
    }
    for j:=1; j<=2; j++ {
        go func() {
            time.Sleep(1 * time.Second)
            ch <- "c2"
        }()
    }
    time.Sleep(1 * time.Second)
    for {
        select {
        case a,e := <-ch:
            fmt.Println(a,e)
            queue.a = append(queue.a, a)
        case b,e := <-c2:
            fmt.Println(b,e)
            queue.b = b
        }
        if len(ch) + len(c2) == 0 {
            fmt.Println("queue", queue)
            break
        }
        res, err :=  <- ch
        fmt.Println(res, err)
    }

    fmt.Println("hello go!")
}
/**

go func:1
go func:4
go func:2
go func:3
go func:5
1_ch_5577006791947779410 true
4_ch_8674665223082153551 true
2_ch_6129484611666145821 true
3_ch_4037200794235010051 true
5_ch_3916589616287113937 true
c2 true
c2 true
queue {[1_ch_5577006791947779410 4_ch_8674665223082153551 2_ch_6129484611666145821 3_ch_4037200794235010051 5_ch_3916589616287113937 c2 c2] }
hello go!

*/