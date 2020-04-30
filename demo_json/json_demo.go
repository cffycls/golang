package demo_json

import (
    "encoding/json"
    "fmt"
)

type Server struct {
    ServerName string `json:"name"`
    ServerIp string `json:"ip"`
    ServerPort int `json:"port"`
}

//序列化结构体
func Serialize()  {
    server := new(Server)
    server.ServerName = "local"
    server.ServerIp = "127.0.0.1"
    server.ServerPort = 8080

    b,err := json.Marshal(server) //序列化
    if err != nil {
        fmt.Println("Marshal json error: ", err.Error())
        return
    }
    fmt.Println("序列化结构体 Marshal json: ", b, "\n", string(b))
}
//序列化map
func SerializeMap()  {
    server := make(map[string] interface{})
    server["ServerName"] = "local"
    server["ServerIp"] = "127.0.0.1"
    server["ServerPort"] = 8080

    b,err := json.Marshal(server) //序列化
    if err != nil {
        fmt.Println("Marshal json error: ", err.Error())
        return
    }
    fmt.Println("序列化map Marshal map json: ", string(b))
}

//反序列化结构体
func UnSerialize()  {
    //jsonString := `{"ServerName":"local","ServerIp":"127.0.0.1","ServerPort":8080}`
    jsonString := `{"name":"local","ip":"127.0.0.1","port":8080}`
    server := new(Server) //内容指定到结构体容器
    err := json.Unmarshal([] byte(jsonString), &server) //Unmarshal(data []byte, v interface{})
    if err != nil {
        fmt.Println("UnMarshal json error: ", err.Error())
        return
    }
    fmt.Println("反序列化结构体 UnMarshal json sring to struct: ", server)
}
//反序列化map
func UnSerializeMap()  {
    jsonString := `{"name":"local","ip":"127.0.0.1","port":8080}`
    server := make(map[string] interface{}) //内容指定到map
    err := json.Unmarshal([] byte(jsonString), &server) //Unmarshal(data []byte, v interface{})
    if err != nil {
        fmt.Println("UnMarshal json error: ", err.Error())
        return
    }
    fmt.Println("反序列化map UnMarshal json sring to struct: ", server)
}