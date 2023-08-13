package upgrade

import (
	"os/exec"
	"runtime"

	"github.com/maogou/pep-ebook/internal/constant"
	"github.com/urfave/cli/v2"
)

func (u *Upgrade) Execute(ctx *cli.Context) error {
	u.PrintLog("Upgrade-Execute", "开始执行升级,当前版本为:", constant.Version)

	arg := "GO111MODULE=" + constant.GO111MODULE + "GOPROXY=" + constant.GOPROXY + " go install " + constant.GithubRepo
	execCmd := exec.Command("sh", "-c", arg)

	if runtime.GOOS == constant.Windows {
		arg = "set GOPROXY=" + constant.GOPROXYWIn + " && go install " + constant.GithubRepo
		execCmd = exec.Command("cmd.exe", "/c", arg)

	}

	if err := execCmd.Run(); err != nil {
		u.PrintLog("upgrade", "升级错误err=", err)
		u.ZLog.Warn().Msg("😒😒😒升级失败,请检查是否可以正常访问github.com后重试!")
		return err
	}

	u.ZLog.Info().Msg("🚀🚀🚀恭喜你升级成功!")
	return nil
}
