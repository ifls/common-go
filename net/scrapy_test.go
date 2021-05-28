package net

import (
	"fmt"
	"github.com/gocolly/colly"
	"github.com/ifls/gocore/utils"
	log2 "github.com/ifls/gocore/utils/log"
	"github.com/ifls/gocore/x/io/file"
	"log"
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

var downloadQueue []string
var visitQueue []string
var c *colly.Collector
var allowedDomains []string
var finalPath string

func init() {
	imgUrls = map[string]*ImgInfo{}
	htmlUrls = map[string]*HtmlInfo{}

	downloadQueue = make([]string, 0)
	visitQueue = make([]string, 0)
	visitQueue = append(visitQueue, "http://pic.netbian.com/tupian/1.html")
	allowedDomains = make([]string, 0)
	allowedDomains = append(allowedDomains, "netbian.com", "pic.netbian.com")
	finalPath = "/Users/ifls/Downloads/logs/imgs/" + strconv.Itoa(int(utils.NextId())) + "/"
	err := os.Mkdir(finalPath, os.ModePerm)
	if err != nil {
		log.Println(err)
	}
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
		downloadQueue = append(downloadQueue, url)
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
		log2.DevInfo("ONHTML Link found: %s\n", url)
		visitQueue = append(visitQueue, url)
		log2.DevInfo("html len = %d\n", len(htmlUrls))
		return
	}
}

func scheDownload() {
	if len(downloadQueue) >= 1 {
		fmt.Printf("download url img = %s\n", downloadQueue[0])
		go downloadImgUrl(downloadQueue[0])
		downloadQueue = downloadQueue[1:]
	} else {
		time.Sleep(100 * time.Millisecond)
	}
	scheDownload()
}

func downloadImgUrl(url string) {
	if strings.Contains(url, "http") {
		_ = c.Visit(url)
	} else {
		_ = c.Visit("http://pic.netbian.com" + url)
	}
}

func OnResponse(response *colly.Response) {
	url := response.Request.URL.String()
	log2.DevInfo("onResponse = %v", url)
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
		path := finalPath + filname
		go func(path string) {
			response.Save(path)
		}(path)
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
		_ = c.Visit(visitQueue[0])
		visitQueue = visitQueue[1:]
	} else {
		log2.LogError("visited end ")
		time.Sleep(100 * time.Millisecond)
		log2.DevInfo("htmlurl length = %d\n imgurl length = %d\n", len(htmlUrls), len(imgUrls))
	}

	scheVisit()
}

func start() {
	go scheDownload()
	scheVisit()
}

func TestCopyFile(t *testing.T) {
	CopyFile(nil, "")
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
		log2.DevInfo("onError=%v", e)
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
		log2.DevInfo("visit finished = %v\n", url)
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

	hashKey := utils.Sha256Hash(data)
	name := utils.Base64Encoding(hashKey)
	filename := name

	filename = filename + "." + subfix
	object := file.GetDir(subfix, name) + filename

	err := file.WriteGcpOss(data, file.TestBucket, object, func(ossUrl string) {
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
		log2.LogErr(err)
		return
	}

	//picStruct := gostruct.Picture{
	//	Fid:   0,
	//	Title: "",
	//	Tags:  nil,
	//}
}

func TestYouling(t *testing.T) {
	// 设置运行时默认操作界面，并开始运行
	// 运行软件前，可设置 -a_ui 参数为"web"、"gui"或"cmd"，指定本次运行的操作界面
	// 其中"gui"仅支持Windows系统

}
