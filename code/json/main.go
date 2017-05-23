package main

import (
	"encoding/json"
	"fmt"
	"os"
)

type Book struct {
	Title       string
	Author      []string
	Publisher   string
	Price       float64
	IsPublished bool
}

func main() {
	b := []byte(`{
	    "Title":"go programming language",
	    "Author":["john","ada","alice"],
	    "Publisher":"qinghua",
	    "IsPublished":true,
	    "Price":99
	  }`)
	//先创建一个目标类型的实例对象，用于存放解码后的值
	var book Book
	//Json.Unmarshal()函数会根据一个约定的顺序查找目标结构中的字段，如果找到一个则进行匹配。这些字段在类型声明中必须都是以大写字母开头、可被导出的字段。
	//但是当JSON数据中的数据结构和GO中的目标类型不匹配时，会怎样呢？
	//如果JSON中的字段在GO目标类型中不存在，json.Unmarshal（）函数在解码过程中会丢弃该字段。
	//这个特性让我们可以从一段JSON数据中筛选出指定的值填充到多个GO语言类型中，当然，前提是已知JSON数据的字段结构。
	//这也意味着，目标类型中不可被导出的私有字段（非首字母大写）将不会受到解析转化的影响
	err := json.Unmarshal(b, &book)
	if err != nil {
		fmt.Println("error in translating,", err.Error())
		return
	}
	fmt.Println(book.Author)

	//=========================================================================
	//如果要解析一个配置文件，为了使程序端的改动能够灵活，大多数情况下是不知道配置文件中到底是怎样的结构，只是在需要的时候去取就可以了。
	//GO语言支持接口。在go中，接口是一组预定义方法的组合，任何一个类型均可通过实现接口预定义的方法来实现，且无需显示声明，所以没有任何方法的空接口代表任何类型。
	//换句话说，每一个类型其实都至少实现了一个空接口。GO内建这样灵活的类型系统，向我们传达了一个很有价值的信息：空接口是通用类型。
	//如果要解析一段未知结构的JSON，只需向这段JSON数据解码输出到一个空接口即可。
	//在GO的标准库encoding/json包中，允许使用map[string]interface{}和[]interface{}类型的值来分别存放未知结构的JSON对象或数组。
	//先创建一个目标类型的实例对象，用于存放解码后的值
	var inter interface{}
	err1 := json.Unmarshal(b, &inter)
	if err1 != nil {
		fmt.Println("error in translating,", err1.Error())
		return
	}
	//要访问解码后的数据结构，需要先判断目标结构是否为预期的数据类型
	book1, ok := inter.(map[string]interface{})
	//然后通过for循环一一访问解码后的目标数据
	if ok {
		for k, v := range book1 {
			switch vt := v.(type) {
			case float64:
				fmt.Println(k, " is float64 ", vt)
			case string:
				fmt.Println(k, " is string ", vt)
			case []interface{}:
				fmt.Println(k, " is an array:")
				for i, iv := range vt {
					fmt.Println(i, iv)
				}
			default:
				fmt.Println("illegle type")
			}
		}
	}
	os.Exit(0)

	//运行结果
	//[john ada alice]
	//Price  is float64  99
	//Title  is string  go programming language
	//Author  is an array:
	//0 john
	//1 ada
	//2 alice
	//Publisher  is string  qinghua
	//illegle type

	//更多 https://github.com/astaxie/build-web-application-with-golang/blob/master/zh/07.2.md
}
