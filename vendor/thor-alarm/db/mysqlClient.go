package db

import (
	"database/sql"
	"reflect"
	"strconv"
	"thor-alarm/models"
	"time"

	"github.com/astaxie/beego/logs"
	_ "github.com/go-sql-driver/mysql"
)

var (
	Db   *sql.DB
	stmt *sql.Stmt
	sw   bool
	ch   chan models.AlarmLog
)

func InitMysqlClient(DbCfg models.DbCfg) error {
	sw = true
	ch = make(chan models.AlarmLog, 100)
	var err error = nil
	Db, err = sql.Open("mysql", DbCfg.DbUser+":"+DbCfg.DbPwd+"@tcp("+DbCfg.DbAddr+":"+strconv.FormatInt(DbCfg.DbPort, 10)+")/"+DbCfg.DbName+"?charset=utf8")
	if err != nil {
		return err
	}
	Db.SetMaxOpenConns(2)
	Db.SetMaxIdleConns(1)
	err = Db.Ping()
	if err != nil {
		return err
	}
	LogModel := models.AlarmLog{}
	typeName := reflect.TypeOf(LogModel)
	columns := ""
	qmark := ""
	for k := 0; k < typeName.NumField(); k++ {
		columns = columns + typeName.Field(k).Tag.Get("json") + ","
		qmark = qmark + "?,"
	}
	if typeName.NumField() > 0 {
		columns = columns[0 : len(columns)-1]
		qmark = qmark[0 : len(qmark)-1]
	}
	prepare := "INSERT INTO " + DbCfg.Table + " (" + columns + ") values (" + qmark + ")"
	logs.Info(prepare)

	stmt, err = Db.Prepare(prepare)
	if err != nil {
		return err
	}
	return nil
}

func AddLog(AlarmLog models.AlarmLog) {
	if sw == true {
		ch <- AlarmLog
	}
}

func WriteLog() {
	for {
		log := <-ch
		if log.AlarmContent != "" {
			_, err := stmt.Exec(log.AppId, log.AlarmName, log.AlarmList, log.AlarmTime, log.AlarmContent)
			if err != nil {
				logs.Error(err)
			}
		}
	}
	logs.Info("DB:MySql WriteLog exit.")
}

func StopMysqlClient() {
	sw = false
	for {
		if len(ch) == 0 {
			close(ch)
			break
		}
		time.Sleep(time.Second)
	}
	if stmt != nil {
		stmt.Close()
	}
	if Db != nil {
		Db.Close()
	}
}
