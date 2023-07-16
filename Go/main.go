package main

import (
	"context"
	"fmt"
	"github.com/chromedp/chromedp"
	"io/ioutil"
	"log"
	"net/http"
	"os"
	"path/filepath"
	"strconv"
	"strings"
	"sync"
	"time"
)

const (
	PreUserUrl = "https://www.vilipix.com/user/"
	PreArtUrl  = "https://www.vilipix.com/illust/"
)

func IndexAll(s, t string) (index []int) {
	for pos := strings.Index(s, t); pos != -1; {
		index = append(index, pos)
		s = s[pos+1:]
		pos = strings.Index(s, t)
	}
	for i := 1; i < len(index); i++ {
		index[i] += index[i-1] + 1
	}
	return
}

func GetPicId(s string) string {
	var st, ed int
	for st = 0; st < len(s) && (s[st] < '0' || s[st] > '9'); st++ {
	}
	for ed = st; ed < len(s) && (s[ed] >= '0' && s[ed] <= '9'); ed++ {
	}
	return s[st:ed]
}

func GetHttpHtmlContent(url string, selector string, sel interface{}) (string, error) {
	options := []chromedp.ExecAllocatorOption{
		chromedp.Flag("headless", true),
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

func GetPicIds(url string) map[string]bool {
	selector := "body > div"
	param := `document.querySelector("body")`
	var picid = make(map[string]bool)
	for i := 1; ; i++ {
		html, _ := GetHttpHtmlContent(url+"?p="+strconv.Itoa(i), selector, param)
		s := string(html)
		const t = `href="/illust`
		index := IndexAll(s, t)
		if len(index) == 0 {
			break
		}
		fmt.Printf("Fetch %d images on page %d!\n", len(index)/2, i)
		for _, pos := range index {
			picid[GetPicId(s[pos:pos+30])] = true
		}
	}
	return picid
}

func Gethtml(url string) []uint8 {
	resp, err := http.Get(url)
	if err != nil {
		fmt.Fprintf(os.Stderr, "Gethtml: %v\n", err)
		return nil
	}
	htm, err := ioutil.ReadAll(resp.Body)
	resp.Body.Close()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Gethtml: reading %s: %v\n", url, err)
		return nil
	}
	return htm
}

var tokens = make(chan struct{}, 20)

func down(id string, wg *sync.WaitGroup, dir string) {
	defer wg.Done()
	tokens <- struct{}{}
	fmt.Printf("Start downloading %s.jpg!\n", id)
	url := PreArtUrl + id
	s := string(Gethtml(url))
	st := strings.Index(s, "https://img9.vilipix.com")
	ed := st + 1 + strings.Index(s[st+1:], "jpg")
	url = s[st : ed+3]
	file, err := os.Create(filepath.Join(".", dir, id+".jpg"))
	if err != nil {
		fmt.Fprintf(os.Stderr, "down create:%v\n", err)
		return
	}
	defer file.Close()
	file.Write([]byte(Gethtml(url)))
	fmt.Printf("%s.jpg is downloaded now!\n", id)
	<-tokens
}

func download(picid map[string]bool, userid string) {
	fmt.Printf("There are %d pictures in total!\nStart downloading now!\n", len(picid))
	os.MkdirAll(userid+"artwork", 0755)
	var wg sync.WaitGroup
	for id := range picid {
		wg.Add(1)
		go down(id, &wg, userid+"artwork")
	}
	wg.Wait()
}

func main() {
	var userid string
	fmt.Printf("Please enter user ID!\n")
	fmt.Scanf("%s", &userid)
	url := PreUserUrl + userid
	download(GetPicIds(url), userid)
	return
}
