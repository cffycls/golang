package main

import (
    "crypto/md5"
    json2 "encoding/json"
    "fmt"
    "io"
    "net/http"
    "os"
    "path/filepath"
    "strconv"
    "strings"
    "time"
)

//http框架梳理
//1.输出hello world
func sayHello(w http.ResponseWriter, r *http.Request)  {
    w.Write([]byte("hello world!\n"))
}

//跨域解耦 F(r,w) => C(F){ F(w,r) }
func cors(f http.HandlerFunc) http.HandlerFunc {
    return func(w http.ResponseWriter, r *http.Request) {
        // 允许访问所有域，可以换成具体url，注意仅具体url才能带cookie信息
        w.Header().Set("Access-Control-Allow-Origin", "*")
        //header的类型
        //w.Header().Add("Access-Control-Allow-Headers", "Content-Type,AccessToken,X-CSRF-Token, Authorization, Token")
        //设置为true，允许ajax异步请求带cookie信息
        w.Header().Add("Access-Control-Allow-Credentials", "true")
        //允许请求方法
        //w.Header().Add("Access-Control-Allow-Methods", "POST, GET, OPTIONS, PUT, DELETE")
        //返回数据格式是json
        //w.Header().Set("content-type", "application/json;charset=UTF-8")
        if r.Method == "OPTIONS" {
            w.WriteHeader(http.StatusNoContent)
            return
        }
        f(w, r)
    }
}

func main()  {
    //controller.Max()
    //2.注册路由处理函数
    http.HandleFunc("/sayHello", cors(sayHello))
    //curl http://127.0.0.1:8081/sayHello

    //2.1文件浏览请求，静态资源服务器
    fileHandler := http.FileServer(http.Dir("./video"))
    http.Handle("/video/", http.StripPrefix("/video/", fileHandler))
    //浏览器 http://127.0.0.1:8081/video/test.mp4

    //2.2文件上传请求
    http.HandleFunc("/upload", cors(uploadHandler))

    //2.3列表信息浏览请求
    http.HandleFunc("/api/videoList", cors(fileListHandler))

    //3.启动web服务
    http.ListenAndServe(":8081",nil)
}

func uploadHandler(w http.ResponseWriter, r *http.Request)  {
    /** a.限制上传大小为10M
        returns a non-EOF error for a Read beyond the limit
     */
    r.Body = http.MaxBytesReader(w, r.Body, 10*1024*1024)
    err := r.ParseMultipartForm(10*1024*1024)
    if err != nil {
        http.Error(w, err.Error(), http.StatusInternalServerError)
        return
    }

    //b.获取文件，并检验
    file, fileHandler, err := r.FormFile("uploadFile")
    defer file.Close()
    fmt.Println("上传文件： ", fileHandler.Filename)
    ret := strings.HasSuffix(fileHandler.Filename, ".mp4")
    if ret == false {
        http.Error(w, "file not mp4", http.StatusInternalServerError)
        return
    }

    //c.生成随机名称，并保存
    md5Byte := md5.Sum([]byte(fileHandler.Filename + time.Now().String()))
    newFilename := fmt.Sprintf("%x", md5Byte) + ".mp4" //md5byte转换为十进制
    dst, err := os.Create("./video/" + newFilename)
    defer dst.Close()
    if err != nil {
        http.Error(w, err.Error(), http.StatusInternalServerError)
        return
    }
    if _, err := io.Copy(dst, file); err != nil {
        http.Error(w, err.Error(), http.StatusInternalServerError)
        return
    }
    w.Write([]byte(fileHandler.Filename+": "+strconv.FormatInt(fileHandler.Size/1024/1024,10)+"M"))
}

func fileListHandler(w http.ResponseWriter, r *http.Request)  {
    files, _ := filepath.Glob("video/*")
    var vlist []string
    for _, file := range files {
        fmt.Println(file)
        vlist = append(vlist, "http://"+r.Host+"/"+file)
    }
    fmt.Println(w.Header())
    retJson, _ := json2.Marshal(vlist)
    w.Write(retJson)
    return
}
