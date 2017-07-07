package main

import (
	"fmt"
	"time"

	"github.com/malisit/kolpa"
)

func main() {
	k := kolpa.C() // Initiate kolpa
	//k.SetLanguage("zh_CN") //https://github.com/andy-zhangtao/kolpa
	fmt.Println(k.FirstName())                                                // Prints John
	fmt.Println(k.Address())                                                  // Prints 729 Richmond Springs Suite 949, Luisborough, VT 85700-5554
	fmt.Println(k.UserAgent())                                                // Prints Mozilla/5.0 (compatible; MSIE 5.0; Windows 98; Win 9x 4.90; Trident/4.0)
	fmt.Println(k.Color())                                                    // Prints Lime #00FF00
	fmt.Println(k.DateTimeAfter(time.Date(2015, 1, 0, 0, 0, 0, 0, time.UTC))) // Prints 2015-09-08 15:34:29 +0300 EEST

}
