package downloader

import (
	"errors"
	"fmt"
	"io"
	"math/rand"
	"net/http"
	"os"
	"regexp"
	"strings"
	"time"

	"github.com/urfave/cli/v2"
)

type Download struct {
	BaseHandler
}

var _ Chain = (*Download)(nil)
var (
	curlRegex      = regexp.MustCompile(`curl '([^']+)'`)
	bookIDRegex    = regexp.MustCompile(`book\.pep\.com\.cn/(\d+)`)
	urlFormatRegex = regexp.MustCompile(`(\d+)/files/mobile/(\d+)\.jpg`)
	headerRegex    = regexp.MustCompile(`-H '([^:]+):\s*([^']+)'`)
	cookieRegex    = regexp.MustCompile(`-(?:-cookie|b) '([^']+)'`)
)

type CurlData struct {
	PageURLFormat string
	ReferFormat   string
	Headers       http.Header
}

func MakeBookURL(bookID string) string {
	return fmt.Sprintf("https://book.pep.com.cn/%s/mobile/index.html", bookID)
}

func ParseCURL(curl string) (CurlData, error) {
	data := CurlData{Headers: make(http.Header)}

	// 提取 URL 和 book_id
	urlMatch := curlRegex.FindStringSubmatch(curl)
	if len(urlMatch) < 2 {
		return data, fmt.Errorf("invalid URL format")
	}
	originalURL := urlMatch[1]

	// 提取 book_id
	bookID := ""
	bookIDMatch := bookIDRegex.FindStringSubmatch(originalURL)
	if len(bookIDMatch) >= 2 {
		bookID = bookIDMatch[1]
	}

	// 生成 PageURLFormat
	data.PageURLFormat = urlFormatRegex.ReplaceAllString(originalURL, "%s/files/mobile/%d.jpg")

	// 解析 Headers
	headerMatches := headerRegex.FindAllStringSubmatch(curl, -1)

	for _, match := range headerMatches {
		if len(match) >= 3 {
			key := strings.TrimSpace(match[1])
			value := strings.TrimSpace(strings.Trim(match[2], `"`)) // 移除值中的双引号
			data.Headers.Add(key, value)
		}
	}

	// 解析 Cookie 参数
	cookieMatches := cookieRegex.FindAllStringSubmatch(curl, -1)
	for _, match := range cookieMatches {
		if len(match) >= 2 {
			cookieStr := strings.TrimSpace(match[1])
			if existingCookie := data.Headers.Get("Cookie"); existingCookie != "" {
				data.Headers.Add("Cookie", existingCookie+"; "+cookieStr)
			} else {
				data.Headers.Add("Cookie", cookieStr)
			}
		}
	}

	// 替换 Referer
	if referer := data.Headers.Get("Referer"); referer != "" {
		data.Headers.Set("Referer", strings.ReplaceAll(referer, bookID, "%s"))
	}

	return data, nil
}

func NewRequestViaCurl(data CurlData, bookID string, page int) (*http.Request, error) {
	bookUrl := fmt.Sprintf(data.PageURLFormat, bookID, page)

	req, err := http.NewRequest(http.MethodGet, bookUrl, nil)
	if err != nil {
		return nil, err
	}

	headers := data.Headers.Clone()
	if referer := headers.Get("Referer"); referer != "" {
		headers.Set("Referer", fmt.Sprintf(referer, bookID))
	}
	req.Header = headers

	return req, nil
}

func (d *Download) HandlerRequest(ctx *cli.Context, dl *Downloader) {
	dl.PrintLog("Download-HandlerRequest", "开始执行")
	if d.IsCanHandler(ctx, dl) {
		dl.ZLog.Info().Msg("🚨🚨🚨正在生成中请稍等.....🚨🚨🚨")
		curl := dl.authenticatedCurl
		curlData, err := ParseCURL(strings.TrimSpace(curl))
		if err != nil {
			dl.PrintLog("ParseCURL", "解析认证请求\n"+curl+"\n失败", err)

			dl.err = errors.New("解析认证请求\n" + curl + "\n失败")
			return
		}
		for _, item := range dl.paths {
			var images []string
			path := dl.imagesTmpDir + dl.pathKey + "/" + item.Remark + "/"
			if err := os.MkdirAll(path, 0777); err != nil {
				dl.PrintLog("os.MkdirAll", "创建目录"+path+"失败,请重试!", err)
				dl.err = errors.New("创建目录" + path + "失败,请重试!")
				return
			}

			for i := 1; i <= item.Pages; i++ {
				bookID := item.BookID
				// check file exists
				filename := fmt.Sprintf("%s%d.jpg", path, i)
				if _, err := os.Stat(filename); err == nil {
					images = append(images, filename)
					continue
				}

				bookPageURL := fmt.Sprintf(curlData.PageURLFormat, bookID, i)
				req, err := NewRequestViaCurl(curlData, bookID, i)
				if err != nil {
					dl.PrintLog("http.NewRequest", "创建请求"+bookPageURL+"失败", err)

					dl.err = errors.New("请求" + bookPageURL + "失败")
					return
				}
				resp, err := http.DefaultClient.Do(req)
				if err != nil {
					dl.PrintLog("http.Get", "请求"+bookPageURL+"失败", err)

					dl.err = errors.New("请求" + bookPageURL + "失败")
					return
				}
				defer func() {
					resp.Body.Close()
				}()

				// check image
				if !strings.Contains(resp.Header.Get("Content-Type"), "image") {
					dl.PrintLog("Content-Type", "电子书url返回"+bookPageURL+"非图片,认证请求已过期", err)

					dl.err = errors.New("电子书url返回" + bookPageURL + "非图片,认证请求已过期")
					return
				}

				contents, err := io.ReadAll(resp.Body)

				if err != nil {
					dl.PrintLog("io.ReadAll", "读取电子书url"+bookPageURL+"内容失败", err)

					dl.err = errors.New("读取电子书url" + bookPageURL + "内容失败")
					return
				}

				if err = os.WriteFile(filename, contents, 0666); err != nil {
					dl.PrintLog("os.WriteFile", "写入电子书内容到"+filename+"失败", err)

					dl.err = errors.New("写入电子书内容到" + filename + "失败")
					return
				}

				images = append(images, filename)

				dl.PrintLog("download ok", bookPageURL, "第", i, "页处理完成")

				time.Sleep(time.Duration(1+rand.Intn(3)) * time.Second)
			}

			if len(images) > 0 {
				if len(item.Remark) == 0 {
					dl.images[dl.subject] = images
				} else {
					dl.images[item.Remark] = images
				}
			}
		}
	}

	d.NextHandler.HandlerRequest(ctx, dl)
}

func (d *Download) IsCanHandler(ctx *cli.Context, dl *Downloader) bool {
	dl.PrintLog("Download-IsCanHandler", "dl.err=", dl.err, "dl.image=", len(dl.images))

	return dl.err == nil
}
