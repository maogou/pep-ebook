package doc

import (
	"fmt"
	"io"
	"net/http"
	"os"
	"sort"
	"time"

	"github.com/pdfcpu/pdfcpu/pkg/api"
	"github.com/pdfcpu/pdfcpu/pkg/pdfcpu"

	"github.com/pdfcpu/pdfcpu/pkg/pdfcpu/model"
	"github.com/pdfcpu/pdfcpu/pkg/pdfcpu/types"
)

func demo() {
	var queryUrl string = "https://book.pep.com.cn/1321001101121/files/mobile/%d.jpg?%d"

	for i := 1; i < 163; i++ {
		bookUrl := fmt.Sprintf(queryUrl, i, time.Now().Unix())
		resp, err := http.Get(bookUrl)

		if err != nil {
			fmt.Println("获取电子书失败:", err)
			return
		}

		defer resp.Body.Close()

		contents, err := io.ReadAll(resp.Body)

		if err != nil {
			fmt.Println("读取内容失败:", err)
			continue
		}

		name := fmt.Sprintf("./book/数学/%d.jpg", i)
		err = os.WriteFile(name, contents, 0666)

		if err != nil {
			fmt.Println("读取到的内容写入失败", err)
		}

		fmt.Println("写入成功", i)
		time.Sleep(2 * time.Second)
	}

	toPdf()
}

type fileInfo struct {
	Name    string    // 文件名
	ModTime time.Time // 文件修改时间
}

func toPdf() {

	books, err := model.ImageFileNames("./book/数学", types.MB)

	if err != nil {
		fmt.Println("打开文件目录失败", err)
		return
	}

	sortBooks := []fileInfo{}

	for _, book := range books {
		var tmp fileInfo
		tmp.Name = book

		info, err := os.Stat(book)
		if err != nil {
			fmt.Println(book, "获取对应的系统信息失败", err)
			return
		}

		tmp.ModTime = info.ModTime()

		sortBooks = append(sortBooks, tmp)
	}

	sort.Slice(
		sortBooks, func(i, j int) bool {
			return sortBooks[i].ModTime.Before(sortBooks[j].ModTime)
		},
	)

	names := []string{}
	for _, name := range sortBooks {
		names = append(names, name.Name)
	}

	imp, err := api.Import("pos:full", types.POINTS)
	if err != nil {
		fmt.Println("导入失败api.Import", err)
		return
	}
	err = api.ImportImagesFile(names, "./数学.pdf", imp, nil)

	if err != nil {
		fmt.Println("创建pdf失败", err)
	}

	bms := []pdfcpu.Bookmark{
		{PageFrom: 1, Title: "封面"},
		{PageFrom: 5, Title: "主编的话"},
		{PageFrom: 7, Title: "目录"},
		{PageFrom: 9, Title: "第一章: 有理数", Children: []pdfcpu.Bookmark{
			{PageFrom: 10, Title: "1.1: 整数和负数"},
			{PageFrom: 14, Title: "1.2: 有理数"},
			{PageFrom: 24, Title: "1.3: 有理数的加减法",
				Children: []pdfcpu.Bookmark{
					{PageFrom: 29, Title: "实验与探究 填幻方"},
					{PageFrom: 35, Title: "阅读与思考 中国人最先使用负数"},
				}},
			{PageFrom: 36, Title: "1.4: 有理数的乘除法", Children: []pdfcpu.Bookmark{
				{PageFrom: 48, Title: "观察与猜想 翻牌游戏中的数学道理"},
			}},
			{PageFrom: 49, Title: "1.5: 有理数的乘方"},
			{PageFrom: 57, Title: "数学活动1"},
			{PageFrom: 58, Title: "小结1"},
			{PageFrom: 59, Title: "复习题1"},
		}},
		{PageFrom: 61, Title: "第二章: 整式的加法", Children: []pdfcpu.Bookmark{
			{PageFrom: 62, Title: "2.1 整式", Children: []pdfcpu.Bookmark{
				{PageFrom: 69, Title: "阅读与思考 数字1与字母X的对话"},
			}},
			{PageFrom: 70, Title: "2.2 整式的加减", Children: []pdfcpu.Bookmark{
				{PageFrom: 79, Title: "信息技术应用 电子表格与数据计算"},
			}},
			{PageFrom: 80, Title: "数学活动2"},
			{PageFrom: 82, Title: "小结2"},
			{PageFrom: 83, Title: "复习题2"},
		}},

		{PageFrom: 85, Title: "第三章: 一元一次方程", Italic: true, Bold: true, Children: []pdfcpu.Bookmark{
			{PageFrom: 86, Title: "3.1 从算式到方程", Children: []pdfcpu.Bookmark{
				{PageFrom: 92, Title: "阅读与思考 “方程” 史话"},
			}},
			{PageFrom: 94, Title: "3.2 解一元一次方程(一)", Children: []pdfcpu.Bookmark{
				{PageFrom: 94, Title: "----合并同类项与移项"},
				{PageFrom: 100, Title: "实验与探究 无线循环小数化分数"},
			}},
			{PageFrom: 101, Title: "3.2 解一元一次方程(二)", Children: []pdfcpu.Bookmark{
				{PageFrom: 101, Title: "----去括号与去分母"},
			}},
			{PageFrom: 108, Title: "3.4 实际问题与一元一次方程"},
			{PageFrom: 117, Title: "数学活动3"},
			{PageFrom: 118, Title: "小结3"},
			{PageFrom: 119, Title: "复习题3"},
		}},
		{PageFrom: 121, Title: "第四章: 几何图形初步", Bold: true, Italic: true, Children: []pdfcpu.Bookmark{
			{PageFrom: 122, Title: "4.1 几何图形", Children: []pdfcpu.Bookmark{
				{PageFrom: 132, Title: "阅读与思考 几何学的起源"},
			}},
			{PageFrom: 133, Title: "4.2 直线 射线 线段", Children: []pdfcpu.Bookmark{
				{PageFrom: 138, Title: "阅读与思考 长度的测量"},
			}},
			{PageFrom: 139, Title: "4.3 角"},
			{PageFrom: 149, Title: "4.4 课题学习 设计制作长方体形状的包装纸盒"},
			{PageFrom: 151, Title: "数学活动4"},
			{PageFrom: 153, Title: "小结4"},
			{PageFrom: 154, Title: "复习题4"},
		}},
		{PageFrom: 159, Title: "部分中英文词汇索引", Bold: true, Italic: true},
		{PageFrom: 161, Title: "后记", Bold: true, Italic: true},
		{PageFrom: 162, Title: "尾部封面", Bold: true, Italic: true},
	}

	err = api.AddBookmarksFile("./数学.pdf", "./数学_bookMark.pdf", bms, true, nil)
	if err != nil {
		fmt.Println("添加书签错误", err)
		return
	}

	fmt.Println("处理pdf完成")
}
