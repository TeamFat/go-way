package main

import (
	"crypto/md5"
	"crypto/sha1"
	"fmt"
	"io"
)

//对字符串进行MD5哈希
func a(data string) int64 {
	t := md5.New()
	io.WriteString(t, data)
	return int64(t.Sum(nil)[0]) % 60
}

//对字符串进行SHA1哈希
func b(data string) string {
	t := sha1.New()
	io.WriteString(t, data)
	return fmt.Sprintf("%x", t.Sum(nil))
}
func main() {
	var data string = "abc"
	fmt.Println(a(data))
	fmt.Println(a("adc"))
	fmt.Printf("SHA1 : %s\n", b(data))
	fmt.Println(59%60, 60%60, 61%60)
}
