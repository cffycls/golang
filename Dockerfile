#
# BUILD 阶段
# 
FROM golang AS gobuild 

ENV GOROOT /usr/local/go
ENV GOPATH /home/go
ENV PATH $PATH:$GOROOT/bin:$GOPATH/bin 

#安装beego依赖
RUN go get github.com/astaxie/beego
RUN go get github.com/beego/bee
RUN go get github.com/astaxie/beego/orm
RUN go get github.com/go-sql-driver/mysql

# 设置我们应用程序的工作目录
WORKDIR $GOPATH/src/beeapi

# 添加所有需要编译的应用代码
ADD src/beeapi .

# 编译一个静态的go应用（在二进制构建中包含C语言依赖库）
RUN CGO_ENABLED=0 GOOS=linux go build -a -installsuffix cgo .

# 设置我们应用程序的启动命令
CMD ["./beeapi"]

RUN pwd && ls -l
#beeapi端口，
#EXPOSE 9601


#
# 生产阶段
# 
FROM scratch AS prod

# 从buil阶段拷贝二进制文件
COPY --from=gobuild /home/go/src/beeapi/beeapi .
CMD ["./beeapi"]
