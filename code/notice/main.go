package main

import (
	"database/sql"
	"fmt"
	"log"
	"net/http"
	"time"

	_ "github.com/go-sql-driver/mysql"
)

func main() {
	db, err := sql.Open("mysql", "root:devop@/notice")
	if err != nil {
		log.Fatalf("Open database error: %s\n", err)
	}
	defer db.Close()

	err = db.Ping()
	if err != nil {
		log.Fatal(err)
	}
	now := time.Now()
	year, mon, day := now.Date()
	hour, min, _ := now.Clock()
	cur := fmt.Sprintf("%04d-%02d-%02d %02d:%02d:00",
		year, mon, day, hour, min)
	//log.Println(cur)
	rows, err := db.Query("select title, content from notice where task_time = ?", cur)
	if err != nil {
		log.Println(err)
	}
	defer rows.Close()
	var title string
	var content string
	var url string
	for rows.Next() {
		err := rows.Scan(&title, &content)
		if err != nil {
			log.Fatal(err)
		}
		//log.Println(title, content)
		//需要替换key
		url = fmt.Sprintf("http://sc.ftqq.com/key.send?text=%s&desp=%s",
			title, content)
		_, err = http.Get(url)
		if err != nil {
			log.Fatal(err)
		}
	}
	err = rows.Err()
	if err != nil {
		log.Fatal(err)
	}
}
