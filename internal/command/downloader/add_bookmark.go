package downloader

import (
	"os"
	"strings"

	"github.com/pdfcpu/pdfcpu/pkg/api"

	"github.com/maogou/pep-ebook/internal/bookmark"
	"github.com/urfave/cli/v2"
)

type AddBookmark struct {
	BaseHandler
}

var _ Chain = (*AddBookmark)(nil)

func (a *AddBookmark) HandlerRequest(ctx *cli.Context, dl *Downloader) {
	dl.PrintLog("AddBookmark-HandlerRequest", "开始处理请求")
	if a.IsCanHandler(ctx, dl) {
		for filename, bookmarkKey := range dl.pdfBookmark {
			if _, err := os.Stat(filename); err != nil {
				dl.fail[bookmarkKey] = filename
				continue
			}

			if bms, ok := bookmark.Bookmark[bookmarkKey]; ok {
				outName := strings.ReplaceAll(filename, ".pdf", "_bookmark.pdf")
				if err := api.AddBookmarksFile(filename, outName, bms, true, nil); err != nil {
					dl.fail[bookmarkKey] = filename
					dl.PrintLog("AddBookmark-HandlerRequest", "api.AddBookmarksFile添加书签失败err=", err)
				}

			}

			dl.success[bookmarkKey] = filename

		}
	}

	a.NextHandler.HandlerRequest(ctx, dl)

}

func (a *AddBookmark) IsCanHandler(ctx *cli.Context, dl *Downloader) bool {
	dl.PrintLog("AddBookmark-IsCanHandler", "dl.err=", dl.err, "dl.pdfBookmark", dl.pdfBookmark)

	if dl.err != nil || len(dl.pdfBookmark) == 0 {
		return false
	}
	return true
}
