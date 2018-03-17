package main

import (
	"database/sql"
	"errors"
	"fmt"
	"time"

	_ "github.com/go-sql-driver/mysql"
)

func main() {
	// res, err := dbSelecter()
	// fmt.Println(err, res)
	go dbSelecter()
	go dbSelecter()
	time.Sleep(time.Second)
}

func dbSelecter() (string, error) {
	/*DSN数据源名称
	  [username[:password]@][protocol[(address)]]/dbname[?param1=value1¶mN=valueN]
	  user@unix(/path/to/socket)/dbname
	  user:password@tcp(localhost:5555)/dbname?charset=utf8&autocommit=true
	  user:password@tcp([de:ad:be:ef::ca:fe]:80)/dbname?charset=utf8mb4,utf8
	  user:password@/dbname
	  无数据库: user:password@/
	*/
	db, err := sql.Open("mysql", "root:devop@tcp(127.0.0.1:3306)/jd_idc_out?charset=utf8") //第一个参数为驱动名
	if err != nil {
		return "", err
	}
	defer db.Close()
	query, err := db.Query("select count(*) as A from jf_config")
	if err != nil {
		return "", err
	}

	//fmt.Println(v)
	column, _ := query.Columns() //读出查询出的列字段名
	if len(column) > 1 {
		return "", errors.New("column not only one ,too more")
	}
	values := make([][]byte, len(column))     //values是每个列的值，这里获取到byte里
	scans := make([]interface{}, len(column)) //因为每次查询出来的列是不定长的，用len(column)定住当次查询的长度
	for i := range values {                   //让每一行数据都填充到[][]byte里面
		scans[i] = &values[i]
	}
	results := make(map[int]map[string]string) //最后得到的map
	i := 0
	for query.Next() { //循环，让游标往下移动
		if err := query.Scan(scans...); err != nil { //query.Scan查询出来的不定长值放到scans[i] = &values[i],也就是每行都放在values里
			return "", err
		}
		row := make(map[string]string) //每行数据
		for k, v := range values {     //每行数据是放在values里面，现在把它挪到row里
			key := column[k]
			row[key] = string(v)
		}
		results[i] = row //装入结果集中
		i++
	}

	fmt.Println(results)
	return results[0][column[0]], nil
}
