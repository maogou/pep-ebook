package downloader

import (
	"errors"
	"fmt"
	"io"
	"net/http"
	"os"
	"time"

	"github.com/urfave/cli/v2"
)

type Download struct {
	BaseHandler
}

var _ Chain = (*Download)(nil)

func (d *Download) HandlerRequest(ctx *cli.Context, dl *Downloader) {
	dl.PrintLog("Download-HandlerRequest", "开始执行")
	if d.IsCanHandler(ctx, dl) {
		dl.ZLog.Info().Msg("🚨🚨🚨正在生成中请稍等.....🚨🚨🚨")
		for _, item := range dl.paths {
			var images []string
			path := dl.imagesTmpDir + dl.pathKey + "/" + item.Remark + "/"
			if err := os.MkdirAll(path, 0777); err != nil {
				dl.PrintLog("os.MkdirAll", "创建目录"+path+"失败,请重试!", err)
				dl.err = errors.New("创建目录" + path + "失败,请重试!")
				return
			}

			for i := 1; i <= item.Pages; i++ {
				bookUrl := fmt.Sprintf(item.QueryUrl, i, time.Now().Unix())
				resp, err := http.Get(bookUrl)
				if err != nil {
					dl.PrintLog("http.Get", "请求"+bookUrl+"失败", err)

					dl.err = errors.New("请求" + bookUrl + "失败")
					return
				}
				defer func() {
					resp.Body.Close()
				}()

				contents, err := io.ReadAll(resp.Body)

				if err != nil {
					dl.PrintLog("io.ReadAll", "读取电子书url"+bookUrl+"内容失败", err)

					dl.err = errors.New("读取电子书url" + bookUrl + "内容失败")
					return
				}

				filename := fmt.Sprintf("%s%d.jpg", path, i)

				if err = os.WriteFile(filename, contents, 0666); err != nil {
					dl.PrintLog(" os.WriteFile", "写入电子书内容到"+filename+"失败", err)

					dl.err = errors.New("写入电子书内容到" + filename + "失败")
					return
				}

				images = append(images, filename)

				dl.PrintLog("download ok", bookUrl, "第", i, "页处理完成")

			}

			if len(images) > 0 {
				dl.images[item.Remark] = images
			}
		}
	}

	d.NextHandler.HandlerRequest(ctx, dl)
}

func (d *Download) IsCanHandler(ctx *cli.Context, dl *Downloader) bool {
	dl.PrintLog("Download-IsCanHandler", "dl.err=", dl.err, "dl.image=", len(dl.images))

	return dl.err == nil
}
