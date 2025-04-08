package downloader

import (
	"github.com/urfave/cli/v2"
)

type SuccessPrint struct {
	BaseHandler
}

func (c *SuccessPrint) HandlerRequest(ctx *cli.Context, dl *Downloader) {
	dl.ZLog.Info().Msg("💯🏆下载完成,让学习成为一种习惯🚀")
}

func (c *SuccessPrint) IsCanHandler(ctx *cli.Context, dl *Downloader) bool {
	return true
}
