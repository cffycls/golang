package main

import (
    "encoding/hex"
    "crypto/md5"
    "fmt"
    "log"
    "strconv"
    "strings"
    "github.com/go-redis/redis"
    "reflect"
    "time"
)

func redisClient(addr string) {
    client := redis.NewClient(&redis.Options{
        Addr:     addr,
        Password: "", // no password set
        DB:       0,  // use default DB
    })

    pong, err := client.Ping().Result()
    fmt.Println(pong, err)
    // Output: PONG <nil>
}

func clusterClient(addr []string) {
    fmt.Println("type:", reflect.TypeOf(redisCsDB))

    redisCsDB = redis.NewClusterClient(&redis.ClusterOptions{
        Addrs: addr,
    })
    fmt.Println("type:", reflect.TypeOf(redisCsDB))

    // Output: PONG <nil>
    pong, err := redisCsDB.Ping().Result()
    fmt.Println(pong, err)
    if err != nil {
        panic(err)
    }
    //fmt.Println(redisCsDB)
}

func md5V(str string) string {
    h := md5.New()
    h.Write([]byte(str))
    return hex.EncodeToString(h.Sum(nil))
}

var redisDB *redis.Client
var redisCsDB *redis.ClusterClient

func main() {
    var servers string = "172.1.50.11:6379,172.1.50.12:6379,172.1.50.13:6379,172.1.30.11:6379,172.1.30.12:6379,172.1.30.13:6379"
    serverInfo := strings.Split(servers, ",")
    fmt.Println(serverInfo)

    clusterClient(serverInfo)
    //设置时区
    l,_ := time.LoadLocation("Asia/Shanghai")
    fmt.Println(time.Now().In(l))

    val, err := redisCsDB.Get("set-10").Result()
    fmt.Println(val)
    fmt.Println(err)
    if len(val) < 1 {
        //清空当前的
        redisCsDB.FlushDB()
        for i:=0; i<20000; i++ {
            //fmt.Println("set-"+ strconv.Itoa(i))
            key := "set-"+ strconv.Itoa(i)
            str :=  time.Now().Format("2006-01-02 15:04:05 ") + strconv.Itoa(i)
            val, err = redisCsDB.Set(key, md5V(str), 0).Result()
            fmt.Println(str)
            fmt.Println(val)
        }
    }

    t1 := time.Now()
    for i:=10; i<20000; i++ {
        key := "set-"+ strconv.Itoa(i)
        _, err := redisCsDB.Get(key).Result()
        if err != nil {
            panic(err)
        }
        //fmt.Println("set-"+ strconv.Itoa(i), val)
    }
    t2 := time.Now()
    fmt.Println()
    fmt.Println(t2.Sub(t1))

    //base. 批量获取redis中的缓存值 https://studygolang.com/articles/14605
    t1 = time.Now()
    pipe := redisCsDB.Pipeline()
    for i:=10; i<20000; i++ {
        key := "set-"+ strconv.Itoa(i)
        pipe.Get(key).Result()
    }
    _, err = pipe.Exec()
    t2 = time.Now()
    fmt.Println(t2.Sub(t1))
    fmt.Println()

    log.Println(err == redis.Nil)
    //执行命令
    slotsInfo, err := redisCsDB.Do("cluster", "slots").Result()
    fmt.Println("slotsInfo type:", reflect.TypeOf(slotsInfo), "\n", slotsInfo, err)
    info, err := redisCsDB.Do("cluster", "info").Result()
    fmt.Println("cluster info:", reflect.TypeOf(info), "\n", info, err)

    //这里被for阻塞
    fmt.Printf("main OK!\n")
}