package upgrade

import (
	"github.com/maogou/pep-ebook/internal/command"
	"github.com/rs/xid"
	"github.com/rs/zerolog"
	"github.com/rs/zerolog/log"
	"github.com/spf13/viper"
	"github.com/urfave/cli/v2"
)

type Upgrade struct {
	command.DebugLog
	enableLog bool
	ZLog      zerolog.Logger
}

func newUpgrade() *Upgrade {
	return &Upgrade{
		enableLog: viper.GetBool("debug"),
		ZLog:      log.Logger.With().Str("qid", xid.New().String()).Logger(),
	}
}

func (u *Upgrade) EnableDebug() bool {
	return u.enableLog
}

func (u *Upgrade) PrintLog(key string, value ...any) {
	if u.EnableDebug() {
		u.ZLog.Info().Any(key, value).Send()
	}
}

var _ command.DebugLog = (*Upgrade)(nil)

func UpgradeCommand() *cli.Command {
	cmd := &cli.Command{
		Name:        "upgrade",
		Usage:       "升级pep-ebook到最新版本",
		Description: "使用示例: ebook upgrade",

		Action: func(cCtx *cli.Context) error {
			upgrade := newUpgrade()
			return upgrade.Execute(cCtx)
		},
	}

	return cmd
}
