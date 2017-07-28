package util

import (
	"encoding/json"
	"io/ioutil"
	"net/http"
	"net/url"
	"time"

	"github.com/astaxie/beego/logs"
)

type AlarmData struct {
	Alarm_title   string `json:"alarm_title"`
	WeixinAgengId int    `json:"weixinAgengId"`
	App_name      string `json:"app_name"`
	Key           string `json:"key"`
	Service       string `json:"service"`
	Send_time     string `json:"send_time"`
	Alarm_list    string `json:"alarm_list"`
	Alarm_content string `json:"alarm_content"`
	Shield        int    `json:"shield"`
	Dev_list      string `json:"dev_list"`
	Dept          string `json:"dept"`
}

type Alarm struct {
	App_mame   string     `json:"app_name"`
	Alarm_type string     `json:"alarm_type"`
	Ip         string     `json:"ip"`
	Data       *AlarmData `json:"data"`
}

var (
	ALARM_URL = "http://monitor.m.jd.local/tools/new_alarm_api/send_alarm?req="
)

func SendAdminAlarm(content string) {
	content += "[管理员告警]"
	//可在此处增加更多管理员
	SendAlarm(content, "chenjie19@jd.com")
}

func SendAlarm(content string, emails string) {

	t := time.Now()
	date := t.Format("2006-01-02 15:04")
	d := &AlarmData{
		Alarm_title:   "无线运维研发",
		WeixinAgengId: 3,
		App_name:      "jingdongAPIMon",
		Key:           "data_alarmer_nsq_consumer",
		Service:       "统一监控平台业务告警",
		Send_time:     date,
		Alarm_list:    emails,
		Alarm_content: content,
		Shield:        1,
		Dev_list:      "陈捷",
		Dept:          "平台技术服务部",
	}

	a := &Alarm{
		App_mame:   "jingdongAPIMon",
		Alarm_type: "weixin",
		Ip:         "10.8.134.93",
		Data:       d,
	}

	bts, err := json.Marshal(a)
	if err != nil {
		logs.Error("json encode error ,", err.Error())
		return
	}

	logs.Info(string(bts))

	u, _ := url.Parse(ALARM_URL + url.QueryEscape(string(bts)))
	q := u.Query()
	u.RawQuery = q.Encode()
	response, err := http.Get(u.String())
	if err != nil {
		logs.Error("http get alarm url error , ", err.Error())
		return
	}
	defer response.Body.Close()
	data, err := ioutil.ReadAll(response.Body)
	if err != nil {
		logs.Error("read response body error, ", err.Error())
	}
	if response.StatusCode != 200 {
		logs.Error("status code not 200")
	}
	logs.Info("response data body ,", string(data))
}
