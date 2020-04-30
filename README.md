go
====
[2020.4.30 10:45] 整理2,3  
[2020.3.10 10:13] 建立1  

1、数据基础篇
---- 
int、string 等略。
a.切片 slice
```cgo
    mSlice := make([] string,3)
    mSlice[0] = "dog"
    mSlice[1] = "cat"
    mSlice[2] = "mouse"
    mSlice = append(mSlice, "yahaha")
    //len(mSlice)、 cap(mSlice)、 copy(mIdSliceDst, mIdSliceSrc)
```
b.关联 map
```cgo
    mMap := make(map[int] string)
    mMap[10] = "dog"
    mMap[11] = "cat"
    mMap[12] = "mouse"
    //delete(mMap,1)无有无效、 delete(mMap,11) 
	nMap := new(map[int] string)
```
c.结构体 struct
```cgo
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
```
d.通道 chan[nel]
```cgo
    mChan := make(chan int,3) //缓存数3
    mChan <- 336
    //...
	close(mChan)
```
错误捕捉
```cgo
func receivePanic()  {
	fmt.Print("panic + recover异常处理: ")
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
```

2、包定义、管理基础篇
---- 
使用版本go1.14.1.  

a.go默认配置，自定义文件|包引用
```cgo
GO111MODULE="on"

import (
    "main/demo_interface"
    "main/demo_struct" //current/package 

)
```
当前包的命名空间，起始于根目录，本目录下的文件或包的引用使用此路径。

b.其它包引用
```cgo
go get github.com/go-redis/redis 

import (
	"fmt"
	"github.com/go-redis/redis/v7"
)
```
命令参见 go help mod

3、网络应用基础篇
---- 
a.mysql操作
```cgo
    "database/sql"
	_"github.com/go-sql-driver/mysql"

type User struct {
	id 				int			`db:"id"`
	username 		string		`db:"username"`
	last_login 		string		`db:"last_login"`
}
	dsn := "root:123456@tcp(172.1.11.11:3306)/test?charset=utf8"
	mysqlDB, err := sql.Open("mysql", dsn) //1、连接

    //执行SQL语句
	err = mysqlDB.QueryRow("select count(id) from user").Scan(&sum) //2、num查询行数

    stmt, err := mysqlDB.Prepare("INSERT INTO user (username,passwd,last_login) VALUES(?,?,?)")
    _, err := stmt.Exec(username, passwd, last_login) //3、绑定参数并执行

    rows, _ := mysqlDB.Query(`SELECT * FROM user ORDER BY id ASC`) //4、查询多行，并逐行追加到结构体
    users := []User{}
    for rows.Next() {
        err := rows.Scan(&id, &username, &passwd, &last_login)
        users = append(users, User{id, username, last_login})
        if err != nil {
            log.Fatalln(err)
        }
    }
```
b.redis操作
```cgo
	"github.com/go-redis/redis"

    //1、连接：单机、集群
    client := redis.NewClient(&redis.Options{
		Addr:     addr,
	})
	redisCsDB = redis.NewClusterClient(&redis.ClusterOptions{
		Addrs: addr,
	})
	pong, err := redisCsDB.Ping().Result() //网络测试
	l,_ := time.LoadLocation("Asia/Shanghai")
    //2、获取数据吧
	val, err := redisCsDB.Get("set-10").Result()
```
c.curl
```cgo
	"io/ioutil"
	"net/http"

    func curlLength (url string) string {
        //1. 建立 http.Client 客户端对象
        client := &http.Client{
            Timeout: time.Second * 10,
        }
        //2. 建立 http.Requesst 请求对象
        req, _ := http.NewRequest("GET", url, strings.NewReader("wd=?"))
        headers := map[string]string{
            "User-Agent":    "Chrome/63.0.3239.132 (Windows NT 10.0; WOW64)",
            "Authorization": "Bearer access_token",
            "Content-Type":  "application/octet-stream",
        }
        for k,v := range headers {
            req.Header.Set(k, v)
        }
        //3. 执行请求 http.Client.Do(req)，并使用ioututil.ReadAll()读取执行结果
        resp, err := client.Do(req)
        if err != nil {
            fmt.Println(err.Error())
            return err.Error()
        }
        bodyText, err := ioutil.ReadAll(resp.Body)
        if err != nil {
            fmt.Println(err.Error())
        }
    
        if resp != nil {
            if resp.StatusCode == 200 {
                return strconv.Itoa(len(string(bodyText)))
            } else {
                fmt.Println(resp.Status)
                return resp.Status
            }
        }else{
            return "0"
        }
    }
```