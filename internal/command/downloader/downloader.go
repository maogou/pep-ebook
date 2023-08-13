package downloader

import (
	"github.com/maogou/pep-ebook/internal/classification"
	"github.com/maogou/pep-ebook/internal/command"
	"github.com/maogou/pep-ebook/internal/constant"
	"github.com/rs/xid"
	"github.com/rs/zerolog"
	"github.com/rs/zerolog/log"
	"github.com/spf13/viper"
	"github.com/urfave/cli/v2"
)

type Downloader struct {
	command.DebugLog
	enableLog    bool
	err          error
	ZLog         zerolog.Logger
	period       string //学段
	grade        string //年级
	subject      string //学科
	paths        []classification.UrlPath
	imagesTmpDir string
	pdfDir       string
	pathKey      string
	images       map[string][]string
	pdfBookmark  map[string]string
	success      map[string]string
	fail         map[string]string
}

func newDownloader() *Downloader {
	return &Downloader{
		imagesTmpDir: constant.ImageTmpDir,
		pdfDir:       constant.SavePdfDir,
		images:       make(map[string][]string),
		pdfBookmark:  make(map[string]string),
		success:      make(map[string]string),
		fail:         make(map[string]string),
		enableLog:    viper.GetBool("debug"),
		ZLog:         log.Logger.With().Str("qid", xid.New().String()).Logger(),
	}
}

func (d *Downloader) Execute(ctx *cli.Context) error {
	d.PrintLog("Downloader-Execute", "开始执行")
	if err := d.prepareSelect(); err != nil {
		return err
	}

	downloadHandler := &Download{}
	sortImageHandler := &SortImage{}
	createPdfHandler := &CreatePdf{}
	addBookmarkHandler := &AddBookmark{}
	printFinishTipsHandler := &PrintFinishTipsHandler{}
	clearTmpFileHandler := &ClearTmpFile{}

	downloadHandler.SetNext(sortImageHandler)
	sortImageHandler.SetNext(createPdfHandler)
	createPdfHandler.SetNext(addBookmarkHandler)
	addBookmarkHandler.SetNext(printFinishTipsHandler)
	printFinishTipsHandler.SetNext(clearTmpFileHandler)

	downloadHandler.HandlerRequest(ctx, d)

	return d.err
}
func (d *Downloader) EnableDebug() bool {
	return d.enableLog
}

func (d *Downloader) PrintLog(key string, value ...any) {
	if d.EnableDebug() {
		d.ZLog.Info().Any(key, value).Send()
	}
}

var _ command.DebugLog = (*Downloader)(nil)

func DownloaderCommand() *cli.Command {
	cmd := &cli.Command{
		Name:        "download",
		Usage:       "自动下载具体学科带书签的电子书",
		Description: "使用示例: ebook download",
		Action: func(cCtx *cli.Context) error {
			downloader := newDownloader()
			return downloader.Execute(cCtx)
		},
	}

	return cmd
}
