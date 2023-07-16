package main

import (
	"fmt"
	//"os"
	"net/http"
	"log"
	"io/ioutil"
	//"unicode/utf8"
	//"strconv"
	"strings"
	//"bytes"
	//"golang.org/x/net/http/cookiejar"
	//"golang.org/x/net/html"
)

func IndexAll(s, t string) (index []int) {
    for pos := strings.Index(s, t); pos != -1; {
        index = append(index, pos)
        s = s[pos+1:]
        pos = strings.Index(s,t)
    }
	for i := 1; i < len(index); i++ {
		index[i] += index[i-1]
	}
    return
}

func main() {
	client := &http.Client{}
	req, err := http.NewRequest("GET","https://www.vilipix.com/user/9609632?p=1", nil)
	if err != nil {
		log.Print(err)
	}
	req.Header.Set("Content-Type", "text/html; charset=utf-8")
	req.Header.Set("User-Agent", "Mozilla/5.0 (Linux; Android 6.0; Nexus 5 Build/MRA58N) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/114.0.0.0 Mobile Safari/537.36 Edg/114.0.1823.82")
	//req.Header.Set("Accept-Encoding","gzip, deflate, br")
	req.Header.Set("Accept-Language","zh-CN,zh;q=0.9,en;q=0.8,en-GB;q=0.7,en-US;q=0.6")
	req.Header.Set("Referer","https://m.vilipix.com/")
	req.Header.Set("Sec-Ch-Ua",`"Not.A/Brand";v="8", "Chromium";v="114", "Microsoft Edge";v="114"`)
	req.Header.Set("Sec-Ch-Ua-Mobile","?1")
	req.Header.Set("Sec-Ch-Ua-Platform",`"Android"`)
	req.Header.Set("Sec-Fetch-Dest","image")
	req.Header.Set("Sec-Fetch-Mode","no-cors")
	req.Header.Set("Sec-Fetch-Site","same-site")
	req.Header.Set("X","4.2.2.2")
	req.Header.Set("Accept","*/*")
	resp, err := client.Do(req)
	defer resp.Body.Close()
	if err != nil {
		log.Print(err)
	}
	htm, err := ioutil.ReadAll(resp.Body)
	s := string(htm)
	//fmt.Printf("%s\n",string(htm))
	//fmt.Println(len(htm),string(htm),len(string(htm)))
	fmt.Printf("%v\n",string(htm))
	//s := string(htm)
	const t = "img9."
	index := IndexAll(s, t)
	fmt.Println(len(index))
	//for _,pos := range index {
	//	fmt.Println(pos)
	//}
	//fmt.Printf("-1")
	//}
}
