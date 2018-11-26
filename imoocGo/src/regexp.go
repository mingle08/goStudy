package main

import (
    "fmt"
    "regexp"
)

const text = `my email is cccbbb@gmail.com@163.com
 email1 is aaa@163.com
 email2 is ddd@eyou.com.cn`

func main(){
//	reg := regexp.MustCompile(`[a-zA-Z0-9]+@[a-zA-Z0-9]+\.[a-zA-Z0-9]+`)
    reg := regexp.MustCompile(`([a-zA-Z0-9])+@([a-zA-Z0-9])+(\.[a-zA-Z0-9.]+)`)
//	match := reg.FindAllString(text, -1)
    match := reg.FindAllStringSubmatch(text, -1)
	fmt.Println(match)
	
}

