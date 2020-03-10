package main

import (
	"./demo_interface"
	"./demo_json"
	"./demo_struct" //GOPATH ?? GOROOT /src/demo_struct 都可以运行，后者为编辑器所接受。或直接使用相对路径
	"./demo_sugar"
	"fmt"
	"reflect"
)

func main()  {
	//1.
	makeSlice()
	fmt.Println()
	makeMap()
	newMap()
	fmt.Println()
	makeChan()

	//2.
	fmt.Println()
	receivePanic()

	//3.
	fmt.Println()
	fmt.Print("跨文件夹调用: ")
	fmt.Println("GOPATH/src/demo_struct/demo_struct.go")
	demo_struct.TestStruct()

	//4.
	dog := new(demo_struct.Dog)
	dog.Id = 2
	dog.Name = "Tom"
	dog.Age = 6
	dog.Run()
	dog.Eat()
	//dog.see()

	//5.
	fmt.Println()
	//接口测试
	//var b demo_interface.Behavior
	//b = new(demo_struct.Cat)
	//b.Run()
	//b.Eat()
	catt := new(demo_struct.Cat)
	dogg := new(demo_struct.Dog)
	action(catt)
	action(dogg)

	//6.
	//序列化
	demo_json.Serialize()
	demo_json.SerializeMap()
	demo_json.UnSerialize()
	demo_json.UnSerializeMap()

	//7.
	//测试语法糖
	demo_sugar.Sugar("A", "b", "C")
}

func action(b demo_interface.Behavior) string {
	b.Eat()
	b.Run()
	return ""
}

//makeSlice 创建切片
func makeSlice()  {
	mSlice := make([] string,3)
	mSlice[0] = "dog"
	mSlice[1] = "cat"
	mSlice[2] = "mouse"
	fmt.Println(mSlice)
	fmt.Println("len长度 = ", len(mSlice))
	fmt.Println("cap容量 = ", cap(mSlice))

	mSlice = append(mSlice, "yahaha")
	fmt.Println("append: ", mSlice)
	fmt.Println("after +1 len = ", len(mSlice))
	fmt.Println("after +1 cap = ", cap(mSlice))

	mSlice = append(mSlice, "wahaha")
	mSlice = append(mSlice, "yahaha")
	mSlice = append(mSlice, "xahaha")
	fmt.Println(mSlice)
	fmt.Println("after +4 len = ", len(mSlice))
	fmt.Println("after +4 cap = ", cap(mSlice))

	//delete(mSlice,1)
	fmt.Println("irst argument to delete must be map; have []string")

	mIdSliceDst := make([] string,2)
	mIdSliceDst[0] = "id-dst-0"
	mIdSliceDst[1] = "id-dst-1"
	mIdSliceSrc := make([] string,3)
	mIdSliceSrc[0] = "id-src-0"
	mIdSliceSrc[1] = "id-src-1"
	ref := copy(mIdSliceDst, mIdSliceSrc)
	fmt.Println("copy 切片复制影响数：", ref)
	fmt.Println(mIdSliceDst)
}

//makeMap 创建map
func makeMap()  {
	mMap := make(map[int] string)
	mMap[10] = "dog"
	mMap[11] = "cat"
	mMap[12] = "mouse"
	fmt.Println(mMap)
	fmt.Println("makeMap 引用类型", reflect.TypeOf(mMap))
	delete(mMap,1)
	fmt.Println("delete: ", mMap)
	delete(mMap,11)
	delete(mMap,12)
	fmt.Println("delete: ", mMap)
}

//makeChan 创建有2个缓存的channel
func makeChan()  {
	//mChan := make(chan int,3) //缓存数3
	//mChan := make(chan int)
	mChan := make(chan int, 2)
	//defer close(mChan)
	mChan <- 1
	mChan <- 336
	fmt.Println("mChan: ", mChan)
	close(mChan)
	fmt.Println("mChan closed")
}


/**
	new 部分
 */
//newMap 创建map
func newMap()  {
	nMap := new(map[int] string)
	fmt.Println("newMap 指针类型", reflect.TypeOf(nMap))
}

/**
	panic + recover异常处理
 */
func receivePanic()  {
	fmt.Print("anic + recover异常处理: ")
	defer recoverPanic()
	//panic(errors.New("I am a panic"))  //优先级最高
	//panic("I am a panic")  //优先级高
	panic(1)
}
func recoverPanic()  {
	massage := recover()
	switch massage.(type) {
	case error:
		fmt.Println("error: ", massage)
	case string:
		fmt.Println("string: ", massage)
	default:
		fmt.Println("unknown: ", massage)
	}
}