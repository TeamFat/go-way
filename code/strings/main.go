package main

import s "strings"
import "fmt"

//为打印函数起个小名儿，比较有特点的用法
var p = fmt.Println

func main() {

	p("Contains:  ", s.Contains("test", "es"))        //是否包含
	p("SubString:  ", "test"[1:2])                    //截取子字符串
	p("Count:     ", s.Count("test", "t"))            //函数第二个参数中指定字符串的个数
	p("HasPrefix: ", s.HasPrefix("test", "te"))       //判断前缀
	p("HasSuffix: ", s.HasSuffix("test", "st"))       //判断后缀
	p("Index:     ", s.Index("test", "e"))            //判断第一个出现的位置
	p("LastIndex:     ", s.LastIndex("test", "t"))    //判断最后一个出现的位置
	p("Join:      ", s.Join([]string{"a", "b"}, "-")) //用-把数组拼接成字符串
	p("Repeat:    ", s.Repeat("a", 5))                //重复5次
	p("Replace:   ", s.Replace("foo", "o", "0", -1))  //替换字符，后两个参数表示起始位置和长度，-1表示到结尾
	p("Replace:   ", s.Replace("foo", "o", "0", 1))   //替换字符，后两个参数表示起始位置和长度
	p("Split:     ", s.Split("a-b-c-d-e", "-"))       //分割字符串，与Join相反
	p("ToLower:   ", s.ToLower("TEST"))               //转成小写
	p("ToUpper:   ", s.ToUpper("test"))               //转成大写
	p()                                               //打印空行

	p("Len: ", len("hello")) //字符串长度
	p("Char:", "hello"[1])   //获取指定位置的字符(ascii码)
}

//输出
// Contains:   true
// SubString:   e
// Count:      2
// HasPrefix:  true
// HasSuffix:  true
// Index:      1
// LastIndex:      3
// Join:       a-b
// Repeat:     aaaaa
// Replace:    f00
// Replace:    f0o
// Split:      [a b c d e]
// ToLower:    test
