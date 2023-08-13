package downloader

import "github.com/urfave/cli/v2"

type PrintFinishTipsHandler struct {
	BaseHandler
}

func (p *PrintFinishTipsHandler) HandlerRequest(ctx *cli.Context, dl *Downloader) {
	dl.PrintLog(
		"PrintFinishTipsHandler-HandlerRequest", "开始处理dl.err=", dl.err, "dl.success=", dl.success, "dl.fail=",
		dl.fail,
	)

	if p.IsCanHandler(ctx, dl) {
		for _, value := range dl.success {
			dl.ZLog.Info().Msg("🏆🏆🏆文件保存在:" + value)
		}

		for filename := range dl.fail {
			dl.ZLog.Info().Msg(filename + "处理失败🥴🥴🥴")
		}
	}

	p.NextHandler.HandlerRequest(ctx, dl)

}

func (p *PrintFinishTipsHandler) IsCanHandler(ctx *cli.Context, dl *Downloader) bool {
	dl.PrintLog("PrintFinishTipsHandler-IsCanHandler", "开始判断")

	if dl.err != nil || len(dl.success) == 0 && len(dl.fail) == 0 {
		return false
	}

	return true
}
