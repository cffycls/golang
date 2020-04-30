package main

import (
    "bufio"
    "fmt"
    "io"
    "os"
    "strconv"
    "strings"
)
func main()  {
    scanDir := "./"
    scanDir = "/home/cffycls/Desktop/vd/"


    dstP,_ := os.Open(scanDir)
    defer dstP.Close()
    dir,_ := dstP.Readdir(0)     //获取文件夹下各个文件或文件夹的fileInfo
    for _,fileInfo := range dir{
        dstF,_ := os.Open(scanDir+"/"+fileInfo.Name())
        info,_ := dstF.Stat()
        fmt.Println(fileInfo.Name())
        if info.IsDir() {
            os.Mkdir(scanDir+"/"+fileInfo.Name()+"_new", os.ModePerm)
            combine(scanDir+"/"+fileInfo.Name()+"/", scanDir+"/"+fileInfo.Name()+"_new/")
        }
    }
}

func combine(scanDir, dstFile string)  {
    var scanFiles [][]string
    file, err := os.Open(scanDir + "index.m3u8")
    if err != nil {
        fmt.Println("read err", err.Error())
        return
    }
    defer file.Close()
    // 使用bufio读取
    r := bufio.NewReader(file)
    fileNum := 0
    var scanFile []string
    for {
        data,_,err := r.ReadLine()
        if err == io.EOF {
            scanFiles = append(scanFiles, scanFile)
            break
        }
        s := string(data)
        //data, _ := r.ReadBytes('\n')
        if strings.HasPrefix(s, "/storage") {
            sn := strings.Split(s, "/")
            //fmt.Println(sn[len(sn)-1])
            if fileNum == 50 {
                scanFiles = append(scanFiles, scanFile)
                scanFile = scanFile[0:0]
                fileNum = 0
            }
            scanFile = append(scanFile, sn[len(sn)-1])
            fileNum++
        }
    }
    fmt.Println("scanFiles: ", (len(scanFiles)-1)*50+fileNum)

    wt := 0
    fileNum = 0
    for _, scanFile := range scanFiles{
        for _,ff := range scanFile{
            src,_ := os.Open(scanDir + ff)
            defer src.Close()
            dst,_ := os.OpenFile(dstFile +"_" + strconv.Itoa(fileNum)+"__"+ ff[0:10] + ".ts", os.O_WRONLY|os.O_CREATE, 0644)
            defer dst.Close()
            w,_ := io.Copy(dst, src)
            wt += int(w)
            fileNum++
        }
    }
    fmt.Println("----", dstFile + ".ts", wt)
}
