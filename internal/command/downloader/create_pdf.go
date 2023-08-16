package downloader

import (
	"os"

	"github.com/pdfcpu/pdfcpu/pkg/api"
	"github.com/pdfcpu/pdfcpu/pkg/pdfcpu/types"
	"github.com/urfave/cli/v2"
)

type CreatePdf struct {
	BaseHandler
}

var _ Chain = (*CreatePdf)(nil)

func (c *CreatePdf) HandlerRequest(ctx *cli.Context, dl *Downloader) {
	dl.PrintLog("CreatePdf-HandlerRequest", "开始处理请求")
	if c.IsCanHandler(ctx, dl) {
		imp, err := api.Import("pos:full", types.POINTS)
		if err != nil {
			dl.PrintLog(`api.Import("pos:full", types.POINTS)`, "发生错误err=", err)

			dl.err = err
		} else {
			for index, names := range dl.images {
				savePath := dl.pdfDir + dl.pathKey
				name := savePath + "/" + index + ".pdf"
				bookmarkKey := dl.pathKey + "/" + index
				dl.PrintLog("name", name)
				if err = os.MkdirAll(savePath, 0777); err != nil {
					dl.PrintLog(`CreatePdf-HandlerRequest-os.MkdirAll`, "发生错误err=", err)

					dl.err = err
					break
				}

				_ = os.Remove(name)

				if err = api.ImportImagesFile(names, name, imp, nil); err != nil {
					dl.PrintLog(`api.ImportImagesFile(names, name, imp, nil)`, "发生错误err=", err)
					dl.err = err
					break
				}

				dl.pdfBookmark[name] = bookmarkKey

			}

		}

	}

	c.NextHandler.HandlerRequest(ctx, dl)
}

func (c *CreatePdf) IsCanHandler(ctx *cli.Context, dl *Downloader) bool {
	dl.PrintLog("CreatePdf-IsCanHandler", "dl.err=", dl.err, "dl.imags=", len(dl.images))

	if dl.err != nil || len(dl.images) == 0 {
		return false
	}

	return true
}
