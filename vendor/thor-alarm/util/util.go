package util

import (
	"encoding/json"
	"fmt"
	"math"
	"reflect"
	"strconv"
	"time"
)

func CheckTime(from, to string) bool {
	timestamp := time.Now().Unix()
	tm := time.Unix(timestamp, 0)
	now := tm.Format("15:04:05")
	if now < from || now > to {
		return false
	} else {
		return true
	}
}

func CheckDimension(dataD map[string]interface{}, strategyD string) bool {
	//dataD数据维度和strategyD策略维度 维度的数字也就是解析后的map的key数量要一样
	//并且dataD里的值必须在存在在strategy对应的key的value里，value应该是一个切片，
	//那就要求原始httpapi或者的json就是个维度用数组方式json而来的，这里需要特别注意
	var p map[string]interface{}
	json.Unmarshal([]byte(strategyD), &p)
	if len(dataD) != len(p) {
		return false
	}
	for k, xv := range dataD {
		if yv, ok := p[k]; !ok || !Contains(xv, yv) {
			return false
		}
	}
	return true
}

//判断一个值是否在slice或者array里
func Contains(obj interface{}, target interface{}) bool {
	targetValue := reflect.ValueOf(target)
	switch reflect.TypeOf(target).Kind() {
	case reflect.Slice, reflect.Array:
		for i := 0; i < targetValue.Len(); i++ {
			if targetValue.Index(i).Interface() == obj {
				return true
			}
		}
	case reflect.Map:
		if targetValue.MapIndex(reflect.ValueOf(obj)).IsValid() {
			return true
		}
	}
	return false
}

func GetMinuteTime(t int64) string {
	tt := t - t%60
	return fmt.Sprintf("%d", tt)
}

func GetContinueDiffTime(t int64, st int64) string {
	tt := (t - t%60 - st) / 60
	return fmt.Sprintf("%d", tt)
}

func GetRedisKey(rType string, appid string, value string) string {
	//basedata           value为整分钟时间戳
	//lastdata           不带value
	//continue           value为策略id
	//baseline           value为整分钟时间戳
	//baselineDay        value为天格式时间
	//continueTimestamp  value为策略id
	//brokendata         value为整分钟时间戳
	if rType == "lastdata" {
		return "Talarm_" + appid + "_" + rType
	}
	return "Talarm_" + appid + "_" + rType + "_" + value
}

func ConvergenceCount(count int64, convergence int64, rs int64, sInt int64) int64 {
	switch convergence {
	case 1: //比例 指数
		i, frac := math.Modf(math.Log2(float64(sInt - count + 1)))
		if frac == 0 {
			return int64(i)
		}
		return 0
	case 2: //差值 倍数
		ii := sInt - count + 1
		if ii%rs == 0 {
			return ii / rs
		}
		return 0
	}
	return 0
}

func AlarmContent(cType string, value float64) (content string) {
	if value < 0 {
		content = "降低"
		value = -value * 100
		switch cType {
		case "eq":
			content += "达到" + fmt.Sprintf("%g", value)
		case "lt":
			content += "高于" + fmt.Sprintf("%g", value)
		case "gt":
			content += "低于" + fmt.Sprintf("%g", value)
		}
	} else {
		content = "升高"
		value = value * 100
		switch cType {
		case "eq":
			content += "达到" + fmt.Sprintf("%g", value)
		case "lt":
			content += "低于" + fmt.Sprintf("%g", value)
		case "gt":
			content += "高于" + fmt.Sprintf("%g", value)
		}
	}
	content += "%,"
	return
}

func StrToFloat64round(str string, prec int, round bool) float64 {
	f, _ := strconv.ParseFloat(str, 64)
	return FloatPrecision(f, prec, round)
}

// FloatPrecision float指定精度; round为true时, 表示支持四舍五入
func FloatPrecision(f float64, prec int, round bool) float64 {
	pow10N := math.Pow10(prec)
	if round {
		return math.Trunc((f+0.5/pow10N)*pow10N) / pow10N
	}
	return math.Trunc((f)*pow10N) / pow10N
}
