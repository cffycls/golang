package main

import (
    "bufio"
    "crypto/aes"
    "crypto/cipher"
    "fmt"
    "io"
    "io/ioutil"
    "os"
    "strconv"
    "strings"
)

/**
 * 安卓小视屏播放器
 * ts 浏览器视频缓存：文件内容提取、处理，文件拷贝，按指定序列排序，AES解密
 */
func main()  {
    scanDir := "./"

    //、知识点1. 文件操作
    /**
     * os.Open || OP.Readdir || OP.Stat 等
     */
    dstP,_ := os.Open(scanDir)
    defer dstP.Close()
    dir,_ := dstP.Readdir(0)     //获取文件夹下各个文件或文件夹的列表
    for _,fileInfo := range dir{
        fmt.Print(scanDir+fileInfo.Name())
        dstF,_ := os.Open(scanDir+fileInfo.Name())
        info,_ := dstF.Stat()
        fmt.Println(fileInfo.Name())
        if info.IsDir() && !strings.HasSuffix(fileInfo.Name(), "_new") {
            os.Mkdir(scanDir+fileInfo.Name()+"_new", os.ModePerm)
            copyFile(scanDir+fileInfo.Name()+"/", scanDir+fileInfo.Name()+"_new/", fileInfo.Name())
        }
    }
}

func deAes(data, key []byte) []byte {
    block,err := aes.NewCipher(key)
    if err != nil {
        fmt.Println(err.Error(), key)
        return nil
    }

    blockSize := block.BlockSize()
    blockMode := cipher.NewCBCDecrypter(block, key[:blockSize])
    origData := make([]byte, len(data))
    blockMode.CryptBlocks(origData, data)
    if len(origData)<1 {
        return origData
    }
    origData = origData[:(len(origData) - int(origData[len(origData)-1]))]
    return origData
}

func copyFile(scanDir, dstFile, baseName string)  {
    file, err := os.Open(scanDir + "index.m3u8")
    if err != nil {
        fmt.Println("read err", err.Error())
        return
    }
    defer file.Close()
    // 使用bufio读取
    r := bufio.NewReader(file)
    fileNum := 0
    wt := 0
    var privateMethod map[string]string
    privateMethod = make(map[string]string)
    for {
        //知识点2. r.ReadLine等 文件内容提取、处理
        data,_,err := r.ReadLine()
        if err == io.EOF {
            break
        }
        s := string(data)
        if strings.HasPrefix(s, "#EXT-X-KEY") {
            privateMethod["METHOD"] = "AES-128"
            privateMethod["URI"] = "key.key"
            key,_ := ioutil.ReadFile(scanDir + privateMethod["URI"])
            privateMethod["KEY"] = string(key)
            fmt.Println(privateMethod)
        }

        //data, _ := r.ReadBytes('\n')
        if strings.HasPrefix(s, "/storage") {
            sn := strings.Split(s, "/")
            ff := sn[len(sn)-1] //ff[0:10]
            if len(privateMethod) == 0 {
                src,_ := os.Open(scanDir + ff)
                defer src.Close()
                dst,_ := os.OpenFile(dstFile + baseName + "_" + strconv.Itoa(fileNum)+ ".ts", os.O_WRONLY|os.O_CREATE, 0644)
                defer dst.Close()
                w,_ := io.Copy(dst, src)
                wt += int(w)
            }else{
                data,_ := ioutil.ReadFile(scanDir + ff)
                dst,_ := os.OpenFile(dstFile + baseName + "_" + strconv.Itoa(fileNum)+ ".ts", os.O_WRONLY|os.O_CREATE, 0644)
                defer dst.Close()

                //知识点3. []byte 字符串切片参数data
                data = deAes(data, []byte(privateMethod["KEY"]))
                w,_ := dst.Write(data)
                wt += int(w)
            }
            fileNum++
        }
    }
    fmt.Println("----", dstFile + ".ts total:", wt/1024/1024,"M")
}
