package downloader

import (
	"errors"
	"strings"

	"github.com/maogou/pep-ebook/internal/classification"
	"github.com/spf13/viper"

	"github.com/AlecAivazis/survey/v2"
)

func (d *Downloader) prepareSelect() error {

	var (
		period  string
		grade   string
		subject string
		curl    string
	)

	prompt := &survey.Select{
		Message: "请选择用书学段",
		Options: classification.Periods,
	}

	if err := survey.AskOne(prompt, &period); err != nil {
		return errors.New("你中断了选择用书学段")
	}

	prompt = &survey.Select{
		Message: "请选择学生年级",
		Options: classification.Grades[period],
	}

	if err := survey.AskOne(prompt, &grade); err != nil {
		return errors.New("你中断了选择学生年级")
	}

	prompt = &survey.Select{
		Message: "请选择学科",
		Options: classification.Subjects[period],
	}

	if err := survey.AskOne(prompt, &subject); err != nil {
		return errors.New("你中断了选择学科")
	}

	mustOrNotSelect := ""
	if period == "高中" {
		//          数学            必修
		highKey := subject + "-" + grade

		prompt = &survey.Select{
			Message: "请选择必修项:",
			Options: classification.HighMustOrNot[highKey],
		}

		if err := survey.AskOne(prompt, &mustOrNotSelect); err != nil {
			return errors.New("你中断了选择必修项")
		}
	}

	key := period + "-" + grade + "-" + subject

	if mustOrNotSelect != "" {
		key = period + "-" + grade + "-" + subject + "-" + mustOrNotSelect
	}

	if _, ok := classification.Paths[key]; !ok {
		return errors.New("你选择的学段+年级+学科不存在(维护人员正在努力更新中....),请重新选择")
	}

	if len(classification.Paths[key]) == 0 {
		return errors.New("你选择的学段+年级+学科不存在(维护人员正在努力更新中....),请重新选择")
	}

	if curl = viper.GetString("authenticated_curl"); curl == "" {
		textPrompt := &survey.Editor{
			Message: "请访问 " + MakeBookURL(classification.Paths[key][0].BookID) + " 按照 README 输入认证请求",
			Help:    "输入完成后点击 ESC 键, 再次输入 :wq 保存并关闭编辑器",
		}
		if err := survey.AskOne(textPrompt, &curl); err != nil {
			return errors.New("你中断了输入认证请求")
		}
	}

	d.period = period
	d.grade = grade
	d.subject = subject
	d.authenticatedCurl = curl
	d.paths = classification.Paths[key]
	d.pathKey = strings.ReplaceAll(key, "-", "/")

	return nil
}
