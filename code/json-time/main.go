package main

import (
	"encoding/json"
	"time"
	//"fmt"
	"fmt"
)

type Student struct {
	Name  string    `json:"name"`
	Brith time.Time `json:"brith"`
}

type JsonTime time.Time

// 实现它的json序列化方法
func (this JsonTime) MarshalJSON() ([]byte, error) {
	var stamp = fmt.Sprintf("\"%s\"", time.Time(this).Format("2006-01-02 15:04:05"))
	return []byte(stamp), nil
}

type Student1 struct {
	Name  string   `json:"name"`
	Brith JsonTime `json:"brith"`
}

type Student2 struct {
	Name string `json:"name"`
	// 一定要将json的tag设置忽略掉不解析出来
	Brith time.Time `json:"-"`
}

// 实现它的json序列化方法
func (this Student2) MarshalJSON() ([]byte, error) {
	// 定义一个该结构体的别名
	type AliasStu Student2
	// 定义一个新的结构体
	tmpStudent := struct {
		AliasStu
		Brith string `json:"brith"`
	}{
		AliasStu: (AliasStu)(this),
		Brith:    this.Brith.Format("2006-01-02 15:04:05"),
	}
	return json.Marshal(tmpStudent)
}

func main() {
	stu := Student{
		Name:  "qiangmzsx",
		Brith: time.Date(1993, 1, 1, 20, 8, 23, 28, time.Local),
	}

	b, err := json.Marshal(stu)
	if err != nil {
		println(err)
	}

	println(string(b)) //{"name":"qiangmzsx","brith":"1993-01-01T20:08:23.000000028+08:00"}

	println("===================")

	stu1 := Student1{
		Name:  "qiangmzsx",
		Brith: JsonTime(time.Date(1993, 1, 1, 20, 8, 23, 28, time.Local)),
	}
	b1, err := json.Marshal(stu1)
	if err != nil {
		println(err)
	}

	println(string(b1)) //{"name":"qiangmzsx","brith":"1993-01-01 20:08:23"}

	println("===================")
	stu2 := Student2{
		Name:  "qiangmzsx",
		Brith: time.Date(1993, 1, 1, 20, 8, 23, 28, time.Local),
	}

	b2, err := json.Marshal(stu2)
	if err != nil {
		println(err)
	}

	println(string(b2)) //{"name":"qiangmzsx","brith":"1993-01-01 20:08:23"}
	//https: //gocn.io/article/388
}
