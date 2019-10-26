package net

import (
	"fmt"
	"github.com/gocolly/colly"
	"github.com/ifls/gocore/futil"
	"github.com/ifls/gocore/util"
	"go.uber.org/zap"
	"io"
	"os"
	"strconv"
	"strings"
	"testing"
	"time"
)

type ImgInfo struct {
	url        string
	data       []byte
	suffix     string
	downloaded bool
}

type HtmlInfo struct {
	url     string
	visited bool
}

var imgUrls map[string]*ImgInfo
var htmlUrls map[string]*HtmlInfo

var downlaodQueue []string
var visitQueue []string
var c *colly.Collector
var allowedDomains []string
var finalpath string

func init() {
	imgUrls = map[string]*ImgInfo{}
	htmlUrls = map[string]*HtmlInfo{}

	downlaodQueue = make([]string, 0)
	visitQueue = make([]string, 0)
	visitQueue = append(visitQueue, "http://pic.netbian.com/tupian/1.html")
	allowedDomains = make([]string, 0)
	allowedDomains = append(allowedDomains, "netbian.com", "pic.netbian.com")
	finalpath = "/Users/ifls/Downloads/logs/imgs/" + strconv.Itoa(int(util.NextId())) + "/"
	os.Mkdir(finalpath, os.ModePerm)
}

func onFindImgUrl(element *colly.HTMLElement) {
	url := element.Attr("src")
	url = isImageUrl(url)
	if "" == url {
		return
	}

	if imgUrls[url] == nil {
		imgUrls[url] = &ImgInfo{
			url:        url,
			data:       nil,
			suffix:     "",
			downloaded: false,
		}
		downlaodQueue = append(downlaodQueue, url)
		//util.DevInfo("imgs len = %d\n", len(imgUrls))
	}
}

func isHTMLUrl(url string) string {
	if strings.Contains(url, "http") {
		return url
	}

	if strings.Contains(url, "html") {
		return "http://pic.netbian.com" + url
	}

	return ""
}

func onFindHtmlUrl(e *colly.HTMLElement) {
	url := e.Attr("href")
	url = isHTMLUrl(url)
	if url == "" {
		return
	}
	if htmlUrls[url] == nil {
		htmlUrls[url] = &HtmlInfo{
			url:     url,
			visited: false,
		}
		util.DevInfo("ONHTML Link found: %s\n", url)
		visitQueue = append(visitQueue, url)
		util.DevInfo("html len = %d\n", len(htmlUrls))
		return
	}
}

func scheDownload() {
	if len(downlaodQueue) >= 1 {
		fmt.Printf("download url img = %s\n", downlaodQueue[0])
		go downloadImgUrl(downlaodQueue[0])
		downlaodQueue = downlaodQueue[1:]
	} else {
		time.Sleep(100 * time.Millisecond)
	}
	scheDownload()
}

func downloadImgUrl(url string) {
	if strings.Contains(url, "http") {
		c.Visit(url)
	} else {
		c.Visit("http://pic.netbian.com" + url)
	}
}

func OnResponse(response *colly.Response) {
	url := response.Request.URL.String()
	util.DevInfo("onResponse = %v", url)
	url = isImageUrl(url)
	if url == "" {
		//util.DevInfo("onResponse url is not imgurl")
		return
	}

	if imgUrls[url] == nil {
		//util.DevInfo("onResponse imgurls[url] == nil")
		return
	}

	if imgUrls[url].downloaded == false {
		filname := strings.ReplaceAll(url, "/", "_")
		path := finalpath + filname
		go response.Save(path)
		go WriteToGcpOss(response, url)
		imgUrls[url].downloaded = true
	}
}

func isImageUrl(url string) string {
	if !strings.Contains(url, "http") {
		url = "http://pic.netbian.com" + url
	}

	if strings.Contains(url, "jpg") {
		//util.DevInfo("onResponse url is jpg")
		return url
	}
	if strings.Contains(url, "png") {
		//util.DevInfo("onResponse url is png")
		return url
	}
	return ""
}

func scheVisit() {
	if len(visitQueue) >= 1 {
		c.Visit(visitQueue[0])
		visitQueue = visitQueue[1:]
	} else {
		util.LogError("visited end ")
		time.Sleep(100 * time.Millisecond)
		util.DevInfo("htmlurl length = %d\n imgurl length = %d\n", len(htmlUrls), len(imgUrls))
	}

	scheVisit()
}

func start() {
	go scheDownload()
	scheVisit()
}

func copy_file(w io.Reader, url string) {
	filname := strings.ReplaceAll(url, "/", "_")
	path := "/Users/ifls/Downloads/logs/imgs/" + filname
	f, err := os.Create(path)
	if err != nil {
		util.LogErr(err, zap.String("reason", "file create error"))
	}
	defer f.Close()

	if _, err = io.Copy(f, w); err != nil {
		util.LogErr(err, zap.String("reason", "data copy error"))
	}
	if err := f.Close(); err != nil {
		util.LogErr(err, zap.String("reason", "file close error"))
	}
}

func TestScrapy(t *testing.T) {
	// Instantiate default collector
	c = colly.NewCollector(
		colly.AllowedDomains(allowedDomains...),
		colly.AllowURLRevisit(),
		colly.UserAgent("Mozilla/5.0 (Windows NT 6.1) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/41.0.2228.0 Safari/537.36"),
	)

	// Before making a request print "Visiting ..."
	c.OnRequest(func(r *colly.Request) {
		fmt.Printf("start request url = %s\n", r.URL.String())
		//r.Ctx.Put("url", r.URL.String())
	})

	c.OnError(func(response *colly.Response, e error) {
		util.DevInfo("onError=%v", e)
	})

	c.OnResponse(func(response *colly.Response) {
		OnResponse(response)
	})

	// On every a element which has href attribute call callback
	c.OnHTML("a[href]", func(element *colly.HTMLElement) {
		onFindHtmlUrl(element)
	})

	c.OnHTML("img", func(element *colly.HTMLElement) {
		onFindImgUrl(element)
	})

	c.OnXML("//h1", func(element *colly.XMLElement) {
		//util.DevInfo("OnXML = %v", element.Text)
	})

	c.OnScraped(func(response *colly.Response) {
		url := response.Request.URL.String()
		util.DevInfo("visit finished = %v\n", url)
		if htmlUrls[url] != nil {
			htmlUrls[url].visited = true
		}
	})

	start()
}

func WriteToGcpOss(response *colly.Response, imgUrl string) {
	data := response.Body

	strs := strings.Split(imgUrl, "/")
	lastStr := strs[len(strs)-1]
	subfixs := strings.Split(lastStr, ".")
	subfix := subfixs[len(subfixs)-1]

	hashKey := util.Sha256Hash(data)
	name := util.Base64Encoding(hashKey)
	filename := string(name)

	filename = filename + "." + subfix
	object := futil.GetDir(subfix, name) + filename

	err := futil.WriteGcpOss(data, futil.TEST_BUCKET, object, func(ossUrl string) {
		//fileStruct := gostruct.FileCore{
		//	Name:       filename,
		//	Suffix:     subfix,
		//	Size:       len(data),
		//	OssUrl:     ossUrl,
		//	CreateTime: time.Now().Format(util.TIME_FORMAT),
		//}
		//
		//err := orm.InsertRecord(fileStruct)
		//if err != nil {
		//	util.LogErr(err, zap.String("reason", "insert filestruct error"))
		//	return
		//}
	})
	if err != nil {
		util.LogErr(err)
		return
	}

	//picStruct := gostruct.Picture{
	//	Fid:   0,
	//	Title: "",
	//	Tags:  nil,
	//}
}
