package main

import (
	"bufio"
	"database/sql"
	"fmt"
	_ "fmt"
	_ "github.com/go-sql-driver/mysql"
	"io"
	_ "io"
	_ "io/ioutil"
	"os"
	"path"
	"strconv"
	_ "strings"
	"time"
)

func ReadFile(filePath string, handle func(string)) error {
	f, err := os.Open(filePath)
	defer f.Close()
	if err != nil {
		return err
	}
	//255*100 测试行大于4K时读取被截断 [:4K]
	buf := bufio.NewReaderSize(f, 25500)
	for {
		line, _, err := buf.ReadLine()
		statistic.readLine++

		handle(string(line))
		if err != nil {
			if err == io.EOF{
				//fmt.Println( "io.EOF:", err, string(line))
				return nil
			}
			return err
		}
		//return nil
	}
}

func buildQuery(line string){
	if len(line) == 0 {
		//结尾
		query := curSql[:len(curSql)-1]
		fmt.Println( "==> EOF队列：" + strconv.Itoa(statistic.sqlLineNum), line)
		go execQuery(query)
	}else{
		newSql := curSql + "(" + line +"),"

		if len(newSql) > maxSqlLen {
			query := curSql[:len(curSql)-1]
			fmt.Println( "任务队列：" + strconv.Itoa(statistic.sqlLineNum))
			go execQuery(query)
			//回归
			curSql = sqlBuild + "(" + line +"),"
		}else{
			curSql = newSql
		}
	}
}
func execQuery(query string) {
	statistic.sqlLineNum++
	res, err := myDb.Exec(query) //Result
	if err != nil {
		fmt.Println(err.Error()) //显示异常
		panic(err.Error()) //抛出异常
	}
	re, err := res.RowsAffected() //int64, error
	if err != nil {
		fmt.Println(err.Error()) //显示异常
		fmt.Println(err) //抛出异常
	}
	string := strconv.FormatInt(re, 10)

	rows, err := strconv.Atoi(string)
	if err != nil {
		fmt.Println(err) //抛出异常
	}
	channelResult <- rows
}

type statistics struct {
	execDoneNum int
	sqlLineNum int
	chanRecNum int
	readLine int
}

var sqlBuild string
var curSql string
var myDb *sql.DB
var maxSqlLen = 1024*1024*2
var statistic statistics = statistics{0,0,0,0}
//容器mysql的最大连接数是150 (200崩溃)
var channelResult = make(chan int, 100)

/**
 * 从PHP拼接的MySQL长语句导入测试
 */
func main(){
	var err error
	myDb, err = sql.Open("mysql", "root:123456@tcp(172.1.11.11:3306)/testdb?charset=utf8")
	if err != nil {
		fmt.Println(err.Error()) //显示异常
		panic(err.Error()) //抛出异常
	}
	defer myDb.Close()
	var count int
	rows, err := myDb.Query("SELECT COUNT(id) as count FROM t10_5")
	if err != nil {
		fmt.Println(err.Error()) //显示异常
		panic(err.Error()) //抛出异常
	}
	for rows.Next() {
		rows.Scan(&count)
	}
	fmt.Println(count)
	fmt.Println()

	//初始化sql
	sqlBuild = "INSERT INTO `t10_5` ("
	for i:=1; i<100; i++ {
		sqlBuild += "`field_"+ strconv.Itoa(i) +"`,"
	}
	sqlBuild = sqlBuild[:len(sqlBuild)-1] + ") VALUES "

	pwd, _ := os.Getwd()
	dataPath := path.Join(pwd, "sql.data")
	fmt.Println(dataPath)
	curSql = sqlBuild;
	ReadFile(dataPath, buildQuery)
	time.Sleep(time.Second)
	for {
		x, ok := <- channelResult
		statistic.execDoneNum += x
		statistic.chanRecNum++
		fmt.Println(statistic.sqlLineNum, statistic.execDoneNum, ok, len(channelResult))
		if len(channelResult)==0 {
			print("---------- 完成 ----------")
			fmt.Println(statistic.execDoneNum, statistic.sqlLineNum)
			break
		}
	}
	fmt.Println("chanRecNum=", statistic.chanRecNum, "sqlLineNum=", statistic.sqlLineNum, "readLine=", statistic.readLine, statistic.execDoneNum, len(channelResult))
	var count2 int
	rows, err = myDb.Query("SELECT COUNT(id) as count FROM t10_5")
	if err != nil {
		fmt.Println(err.Error()) //显示异常
		panic(err.Error()) //抛出异常
	}
	for rows.Next() {
		rows.Scan(&count2)
	}
	fmt.Println(count2)
	fmt.Println(count2-count, statistic.execDoneNum, "缺失行：", count2-count-statistic.execDoneNum)
}

/**
523 100000 true 0
---------- 完成 ----------100000 523
chanRecNum= 523 sqlLineNum= 523 readLine= 200001 100000 0
100000
100000 100000 缺失行： 0
 */