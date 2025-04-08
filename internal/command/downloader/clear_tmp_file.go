package downloader

import (
	"os"

	"github.com/maogou/pep-ebook/internal/constant"
	"github.com/urfave/cli/v2"
)

type ClearTmpFile struct {
	BaseHandler
}

func (c *ClearTmpFile) HandlerRequest(ctx *cli.Context, dl *Downloader) {
	dl.PrintLog("ClearTmpFile-HandlerRequest", "开始处理临时文件")

	_ = os.RemoveAll(constant.ImageCacheDir)

	dl.ZLog.Info().Msg("💯🏆下载完成,让学习成为一种习惯🚀")
}

func (c *ClearTmpFile) IsCanHandler(ctx *cli.Context, dl *Downloader) bool {
	return true
}
