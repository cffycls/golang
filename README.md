go
====

8、go module命令 
---- 
```
go mod init golang[任意：项目名称]  
————初始化，生go.mod文件[类库引用记录]  

go mod graph  
————打印依赖  
go mod download  
————下载依赖    
go mod tidy  
————整理[删除、下载]依赖  
go mod verify  
————验证依赖版本是否正确安装  
go mod why  
————查看包被哪个模块使用  
go list -m all 列出所有包和依赖

go mod vendor
————把依赖放到vendor目录
go mod edit -require github.com/xx/xxx@v1.2.2 添加依赖/-droprequire
go mod edit -excluded github.com/xx/xxx@v1.2.1 排除依赖/-dropexclude
go mod edit -fmt 格式化

go get/ go build也会添加依赖
```
