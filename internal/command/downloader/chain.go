package downloader

import (
	"github.com/urfave/cli/v2"
)

type Chain interface {
	SetNext(handler Chain)
	HandlerRequest(ctx *cli.Context, dl *Downloader)
	IsCanHandler(ctx *cli.Context, dl *Downloader) bool
}

type BaseHandler struct {
	NextHandler Chain
}

func (b *BaseHandler) SetNext(handler Chain) {
	b.NextHandler = handler
}
