package main

import (
	"context"
	"fmt"
	"github.com/chromedp/chromedp"
	"log"
	"time"
	"strings"
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

func GetHttpHtmlContent(url string, selector string, sel interface{}) (string, error) {
	options := []chromedp.ExecAllocatorOption{
		chromedp.Flag("headless", true), // debug使用
		chromedp.Flag("blink-settings", "imagesEnabled=true"),
		chromedp.UserAgent(`Mozilla/5.0 (Windows NT 6.3; Win64; x64) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/73.0.3683.103 Safari/537.36`),
	}
	options = append(chromedp.DefaultExecAllocatorOptions[:], options...)

	c, _ := chromedp.NewExecAllocator(context.Background(), options...)

	chromeCtx, cancel := chromedp.NewContext(c, chromedp.WithLogf(log.Printf))
	_ = chromedp.Run(chromeCtx, make([]chromedp.Action, 0, 1)...)

	timeoutCtx, cancel := context.WithTimeout(chromeCtx, 40*time.Second)
	defer cancel()

	var htmlContent string
	err := chromedp.Run(timeoutCtx,
		chromedp.Navigate(url),
		chromedp.WaitVisible(selector),
		chromedp.OuterHTML(sel, &htmlContent, chromedp.ByJSPath),
	)
	if err != nil {
		log.Fatal("Run err : %v\n", err)
		return "", err
	}
	return htmlContent, nil
}


func main() {
	url := "https://www.vilipix.com/user/9609632?p=1"
	//selector := "body > div > div.banner > div.swiper-container-place > div > div.swiper-slide.swiper-slide-0.swiper-slide-visible.swiper-slide-active > a.item.item-big > div.item-bottom"
	selector := "body > div"
	param := `document.querySelector("body")`
	html, _ := GetHttpHtmlContent(url, selector, param)
	fmt.Printf("%s\n",html)
	s := string(html)
	const t = "img9."
	index := IndexAll(s, t)
	fmt.Println(len(index))
	return 
}

