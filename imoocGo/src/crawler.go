package main

import (
    "fmt"
    "io"
    "io/ioutil"
    "bufio"
    "net/http"
    "golang.org/x/text/transform"
    "golang.org/x/text/encoding"
    "golang.org/x/net/html/charset"
)

func main(){
	resp, err := http.Get("http://www.zhenai.com/zhenghun")
	if err != nil {
		panic(err)
	}
	
	defer resp.Body.Close()
	
	if resp.StatusCode != http.StatusOK {
		fmt.Println("Error:status code:", resp.StatusCode)
	}
	
	e := determineEncoding(resp.Body)
	utf8Reader := transform.NewReader(resp.Body, 
		e.NewDecoder())
    all, err := ioutil.ReadAll(utf8Reader)
    if err != nil {
    	panic(err)
    }
    fmt.Printf("%s\n", all)
	
	
}

func determineEncoding(r io.Reader) encoding.Encoding {
    bytes, err := bufio.NewReader(r).Peek(1024)
    if err != nil {
    	panic(err)
    }
    
    e, _, _ := charset.DetermineEncoding(bytes, "")
    return e;
}
