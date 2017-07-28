// models project models.go
package models

type ConfigJson struct {
	AlarmUrl        string
	NsqlookupdAddrs []string
	JimdbInfo       string
	Apikey          string
	DbAddr          string
	DbPort          int64
	DbName          string
	DbUser          string
	DbPwd           string
	Table           string
}

type BigRemoteConfig struct {
	Code    int64        `json:code`
	Message string       `json:message`
	Data    RemoteConfig `json:"data"`
}

type RemoteConfig struct {
	AppId       string        `json:"app_id"`
	AppName     string        `json:"app_name"`
	Topic       string        `json:"topic"`
	Channal     string        `json:"channal"`
	Addrs       []string      `json:"addrs"`
	DbInfo      DbCfg         `json:"db_info"`
	JimdbInfo   string        `json:"jimdb_info"`
	CalBaseLine int64         `json:"cal_baseline"`
	Strategy    []AlarmConfig `json:"strategy"`
}

type DbCfg struct {
	DbAddr string `json:"db_addr"`
	DbPort int64  `json:"db_port"`
	DbName string `json:"db_name"`
	DbUser string `json:"db_user"`
	DbPwd  string `json:"db_pwd"`
	Table  string `json:"table"`
}

type AlarmConfig struct {
	AlarmId         int64    `json:"alarm_id"`
	AlarmName       string   `json:"alarm_name"`
	AlarmList       string   `json:"alarm_list"`
	AlarmDimension  string   `json:"alarm_dimension"`
	From            string   `json:"from"`
	To              string   `json:"to"`
	IsThreshold     int64    `json:"is_threshold"`
	Threshold       []string `json:"threshold"`
	IsBaseLine      int64    `json:"is_baseline"`
	BaseLine        []string `json:"baseline"`
	IsAnTime        int64    `json:"is_antime"`
	AnTimeSub       int64    `json:"antimesub"`
	AnTime          []string `json:"antime"`
	IsMon           int64    `json:"is_mon"`
	Mon             []string `json:"mon"`
	ContinuityCount int64    `json:"continuity_count"`
	Convergence     int64    `json:"convergence"`
	Ratio           int64    `json:"ratio"`
	Sub             int64    `json:"sub"`
	RmAlarmPoint    int64    `json:"rm_alarm_point"`
}

type ThresholdConfig struct {
	AppId     string   `json:"app_id"`
	Threshold []string `json:"threshold"`
	Value     float64  `json:"value"`
}

type MonConfig struct {
	AppId string   `json:"app_id"`
	Mon   []string `json:"mon"`
	Value float64  `json:"value"`
}

type AnTimeConfig struct {
	AppId     string   `json:"app_id"`
	AnTime    []string `json:"antime"`
	Value     float64  `json:"value"`
	AnTimeSub int64    `json:"antimesub"`
	Timestamp int64    `json:"timestamp"`
}

type BaseLineConfig struct {
	AppId     string   `json:"app_id"`
	BaseLine  []string `json:"baseline"`
	Value     float64  `json:"value"`
	Timestamp int64    `json:"timestamp"`
}

type AlarmLog struct {
	AppId        string `json:"app_id"`
	AlarmName    string `json:"alarm_name"`
	AlarmList    string `json:"alarm_list"`
	AlarmTime    int64  `json:"alarm_time"`
	AlarmContent string `json:"alarm_content"`
}

type MQdata struct {
	Type     string `json:"type"`
	Endpoint string `json:"endpoint"`
	Tags     string `json:"tags"`
	Payload  string `json:"payload"`
}

type BaseLineCalData struct {
	AppId     string `json:"app_id"`
	Timestamp int64  `json:"timestamp"`
}
