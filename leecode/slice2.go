package main

import "fmt"

func main()  {
    s3 := [][]int{
        {1,1,1,1,1},
        {2,2,2,2,2},
        {3,3,3,3,3},
        {4,4,4,4,4},
        {5,5,5,5,5},
    }
    fmt.Println()
    fmt.Println(s3)
    fmt.Println(rotate(s3))
    rotate2(s3)
    fmt.Println(s3)
    rotate22(s3)
    fmt.Println(s3)
    s3 = [][]int{
        {1,2,3},
        {4,5,6},
        {7,8,9},
    }
    fmt.Println()
    fmt.Println(s3)
    fmt.Println(rotate(s3))
    rotate2(s3)
    fmt.Println(s3)
    rotate22(s3)
    fmt.Println(s3)

    fmt.Println()
    s := []int{2, 7, 11, 15}
    s = []int{3,2,4}
    fmt.Print(twoSum(s, 6))

}

func rotate(matrix [][]int) [][]int {
    length := len(matrix)
    var n [][]int
    for i:=0; i<=length-1; i++ {
        row := make([]int, 0)
        for j:=length-1; j>=0; j-- {
            row = append(row, matrix[j][i])
        }
        //fmt.Println(n)
        n = append(n, row)
    }
    return n
}
func rotate2(matrix [][]int) {
    length := len(matrix)
    var n [][]int
    for i:=0; i<=length-1; i++ {
        row := make([]int, 0)
        for j:=length-1; j>=0; j-- {
            row = append(row, matrix[j][i])
        }
        n = append(n, row)
    }
    matrix = n
}
func rotate22(matrix [][]int) {
    fmt.Println("======")
    n := matrix
    fmt.Println(n)

    fmt.Println("======")
    length := len(n)
    var rows [][]int
    for i:=0; i<=length-1; i++ {
        row := make([]int, 0)
        for j:=length-1; j>=0; j-- {
            row = append(row, n[j][i])
        }
        rows = append(rows, row)
    }
    for i:=0; i<=length-1; i++ {
        for k:=0; k<=length-1; k++ {
            matrix[i][k] = rows[i][k]
        }
    }
    fmt.Println(matrix)
}
func twoSum(nums []int, target int) []int {
    for i:=0; i<=len(nums)-1; i++ {
        for j:=i; j<=len(nums)-1; j++ {
            fmt.Println(i,j,nums[i] + nums[j], target)
            if nums[i] + nums[j] == target {
                return []int {i, j}
            }
        }
    }
    return nil
}