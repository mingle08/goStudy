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
    "regexp"
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
	
	printCityList(all)
}

func determineEncoding(r io.Reader) encoding.Encoding {
    bytes, err := bufio.NewReader(r).Peek(1024)
    if err != nil {
    	panic(err)
    }
    
    e, _, _ := charset.DetermineEncoding(bytes, "")
    return e;
}

func printCityList(contents []byte){
	/*
	re := regexp.MustCompile(`<a href="http://www.zhenai.com/zhenghun/[0-9a-z]+"[^>]*>[^<]+</a>`)
    matches := re.FindAll(contents, -1)
    for _, m := range matches{
    	fmt.Printf("%s\n", m)
    }
    */
	
	// 用()选取结果
	re := regexp.MustCompile(`<a href="(http://www.zhenai.com/zhenghun/[0-9a-z]+)"[^>]*>([^<]+)</a>`)
    matches := re.FindAllSubmatch(contents, -1)
    for _, m := range matches {
//    	for _, submatch := range m {
//    		fmt.Printf("%s ", submatch)
//    	}
        fmt.Printf("city %s, url %s", m[2], m[1])
    	fmt.Println()
    }
    fmt.Printf("matches found: %d", len(matches))
}

