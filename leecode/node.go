package main

import (
    "fmt"
    "strconv"
)

//打印列表
func PrintList(list *ListNode) {
    p := list
    for p != nil {
        //fmt.Printf("%d-%v-%p ", p.Val, p, p.Next)
        fmt.Print(strconv.Itoa(p.Val)+"->")
        p = p.Next
    }
    fmt.Println()
}
func create (d int, node *ListNode) {
    n := new(ListNode)
    n.Val = d
    n.Next = node.Next
    node.Next = n
    //fmt.Println(d)
    //PrintList(node)
}
func makeNode(s []int) *ListNode {
    ln := &ListNode{
        Val:0,
        Next:nil,
    }
    for i:=len(s)-1; i>=0; i-- {
        create(s[i], ln)
    }
    return ln.Next
}

func main()  {
    s := [] int {1,2,3,4,5,6,7,8}
    //s = [] int {5}
    n1 := makeNode(s)
    PrintList(n1)

    s = [] int {2,3,4,5,6,7,8,9}
    //s = [] int {5}
    n2 := makeNode(s)
    PrintList(n2)
    result := addTwoNumbers(n1, n2)
    PrintList(result)
    fmt.Println("==========")
    c1 := []int {1,4,5}
    m1 := makeNode(c1)
    c2 := []int {1,3,4}
    m2 := makeNode(c2)
    c3 := []int {2,6}
    m3 := makeNode(c3)
    PrintList(m1)
    PrintList(m2)
    PrintList(m3)

    nodes := []*ListNode {m1, m2, m3}
    ret := mergeKLists(nodes)
    PrintList(ret)
}

type ListNode struct {
    Val int
    Next *ListNode
}
func addTwoNumbers(l1 *ListNode, l2 *ListNode) *ListNode {
    //副本
    node1 := l1
    node2 := l2
    //初始链表节点，指针才能与nil相比较
    headNode := &ListNode{
        Val: 0,
        Next: nil,
    }
    nowNode := headNode
    //溢出位
    flag := 0
    for (node1 != nil) || (node2 != nil) {
        v1 := 0
        if node1 != nil {
            v1 = node1.Val
            node1 = node1.Next
        }
        v2 := 0
        if node2 != nil {
            v2 = node2.Val
            node2 = node2.Next
        }

        sv := v1 + v2 + flag
        fmt.Print(strconv.Itoa(v1)+"+"+strconv.Itoa(v2)+"+"+strconv.Itoa(flag)+"="+strconv.Itoa(sv)+", \n")
        if sv < 10 {
            flag = 0
        }else{
            sv = sv - 10
            flag = 1
        }
        //1. 副本找到最新节点 的指针，逐个连接
        if nowNode.Next != nil {
            nowNode = nowNode.Next
        }
        //2. 新建插入节点 的指针
        join := &ListNode{ Val: sv, Next: nil, }
        //3. 替换空指针l
        nowNode.Next = join
        fmt.Printf("nowNode: %d-%v-%p ", nowNode.Val, nowNode, nowNode.Next)

        fmt.Println(sv, "??")
        PrintList(headNode)
        PrintList(nowNode)
        fmt.Println()
    }

    //数据有溢出位
    PrintList(headNode)
    PrintList(nowNode)
    if flag == 1 {
        fmt.Printf("nowNode+: %d-%v-%p ", nowNode.Val, nowNode, nowNode.Next)
        fmt.Println(1, "??")
        join := &ListNode{ Val: 1, Next: nil, }
        nowNode = nowNode.Next
        nowNode.Next = join
        PrintList(nowNode)
    }
    fmt.Println(flag)
    PrintList(headNode)
    return headNode.Next
}

func mergeKLists(lists []*ListNode) *ListNode {
    headNode := &ListNode{
        Val: 0,
        Next: nil,
    }
    nowNode := headNode
    //合并，搜集排序值到slice
    fmt.Println("******")
    var s []int
    for _,nodes := range(lists) {
        PrintList(nodes)
        for nodes.Next != nil {
            if nowNode.Next!=nil {
                nowNode = nowNode.Next
            }
            nowNode.Next = &ListNode{
                Val: nodes.Val,
                Next: nil,
            }
            s = append(s, nodes.Val)
            nodes = nodes.Next
        }
    }
    comNode := headNode.Next
    fmt.Println("~~~~~~~~~~~~~\ncomNode:\n")
    PrintList(comNode)

    //排序
    newHeadNode := &ListNode{
        Val: 0,
        Next: nil,
    }
    nowNode = newHeadNode
    for i:=0; i<len(s)-1; i++ {
        for j:=i+1; j<len(s); j++ {
            if s[i]>s[j] {
                temp := s[i]
                s[i] = s[j]
                s[j] = temp
            }
        }
    }
    fmt.Println(s)
    //重新组合
    for i:=0; i<len(s); i++ {
        for {
            if comNode.Val == s[i] {
                if nowNode.Next!=nil {
                    nowNode = nowNode.Next
                }
                nowNode.Next = &ListNode{
                    Val: comNode.Val,
                    Next: nil,
                }
                break;
            }
            //指针回拨
            if comNode.Next==nil {
                comNode = headNode.Next
            }else{
                comNode = comNode.Next
            }
        }
    }
    PrintList(newHeadNode)
    return newHeadNode.Next
}