package internal

import (
	"os"
	"time"

	"github.com/maogou/pep-ebook/internal/command/upgrade"

	"github.com/maogou/pep-ebook/internal/command/downloader"

	"github.com/maogou/pep-ebook/internal/constant"
	"github.com/spf13/viper"
	"github.com/urfave/cli/v2"

	"github.com/rs/zerolog"
	"github.com/rs/zerolog/log"
)

func init() {
	log.Logger = log.Output(zerolog.ConsoleWriter{Out: os.Stdout, TimeFormat: "15:04:05.000"})
}

func initConfig() error {
	viper.SetConfigType("yaml")

	viper.SetConfigFile(constant.DefaultConfigPath)

	if err := viper.ReadInConfig(); err != nil {
		return err
	}

	return nil
}

func NewPepEbAppCommand() *cli.App {
	ebook := &cli.App{
		Name:                  "pep-ebook",
		Usage:                 "自动下载带书签(人民教育出版社)的电子书",
		Version:               constant.Version,
		Copyright:             "(c) 2023 Wang Xingyuan",
		Compiled:              time.Now(),
		CustomAppHelpTemplate: constant.Logo + cli.AppHelpTemplate,
		Authors: []*cli.Author{
			{
				Name:  "Wang Xingyuan",
				Email: "kinyou_xy@foxmail.com",
			},
		},
		Before: func(cCtx *cli.Context) error {
			_ = initConfig()
			return nil
		},
		Commands: []*cli.Command{
			downloader.DownloaderCommand(),
			upgrade.UpgradeCommand(),
		},
	}

	return ebook

}
