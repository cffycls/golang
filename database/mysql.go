package main

import (
    "crypto/md5"
    "crypto/sha1"
    "encoding/hex"
    "fmt"
    "log"
    "strconv"
    "time"
    "database/sql"
    _"github.com/go-sql-driver/mysql"
)

type User struct {
    id                 int            `db:"id"`
    username         string        `db:"username"`
    last_login         string        `db:"last_login"`
}

func md5V(str string) string {
    h := md5.New()
    h.Write([]byte(str))
    return hex.EncodeToString(h.Sum(nil))
}

func sha1V (str string) string {
    hs := sha1.New()
    hs.Write([]byte(str))
    return hex.EncodeToString((hs.Sum(nil)))
}

var mysqlDB *sql.DB

func main() {
    t1 := time.Now()
    //https://github.com/jmoiron/sqlx postgres
    //http://go-database-sql.org/prepared.html mysql
    dsn := "root:123456@tcp(172.1.11.11:3306)/leecode?charset=utf8"
    mysqlDB, err := sql.Open("mysql", dsn)
    //fmt.Println(mysqlDB, err)
    var (
        id int
        username string
        passwd string
        last_login string
        sum int
    )
    //id, username, passwd, last_login
    err = mysqlDB.QueryRow("select count(id) from user").Scan(&sum)
    if err != nil {
        log.Fatal(err)
    }
    fmt.Println(`目前行数：`, sum)
    if sum<10 {
        stmt, err := mysqlDB.Prepare("INSERT INTO user (username,passwd,last_login) VALUES(?,?,?)")
        if err != nil {
            log.Fatal(err)
        }
        for i:=0; i<20; i++ {
            username = "Dolly"+strconv.Itoa(i)
            last_login = "2000-01-01 10:00:00"
            passwd = sha1V(md5V(username + last_login))
            _, err := stmt.Exec(username, passwd, last_login)
            if err != nil {
                log.Fatal(err)
            }
        }
        err = mysqlDB.QueryRow("select count(id) from user").Scan(&sum)
        if err != nil {
            log.Fatal(err)
        }
        fmt.Println(`行数=`, sum)

    }else {
        rows, _ := mysqlDB.Query(`SELECT * FROM user ORDER BY id ASC`)
        users := []User{}
        for rows.Next() {
            err := rows.Scan(&id, &username, &passwd, &last_login)
            users = append(users, User{id, username, last_login})
            if err != nil {
                log.Fatalln(err)
            }
        }
        log.Println("\n", users)
    }
    t2 := time.Now()
    fmt.Println()
    fmt.Println(t2.Sub(t1))

    fmt.Printf("main OK!\n")
}