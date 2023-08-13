package downloader

import (
	"os"
	"sort"
	"time"

	"github.com/urfave/cli/v2"
)

type SortImage struct {
	BaseHandler
}

type fileInfo struct {
	Name    string    // 文件名
	ModTime time.Time // 文件修改时间
}

var _ Chain = (*SortImage)(nil)

func (s *SortImage) HandlerRequest(ctx *cli.Context, dl *Downloader) {
	dl.PrintLog("SortImage-HandlerRequest", "开始处理")
	if s.IsCanHandler(ctx, dl) {
		for index, images := range dl.images {
			if dl.err != nil {
				break
			}

			var sortImage []fileInfo
			for _, image := range images {
				var tmp fileInfo
				tmp.Name = image

				if info, err := os.Stat(image); err != nil {
					dl.err = err
					break
				} else {
					tmp.ModTime = info.ModTime()
				}

				sortImage = append(sortImage, tmp)
			}

			if len(sortImage) > 0 {
				sort.Slice(
					sortImage, func(i, j int) bool {
						return sortImage[i].ModTime.Before(sortImage[j].ModTime)
					},
				)

				var names []string
				for _, name := range sortImage {
					names = append(names, name.Name)
				}

				dl.images[index] = names
			}
		}
	}

	s.NextHandler.HandlerRequest(ctx, dl)
}

func (s *SortImage) IsCanHandler(ctx *cli.Context, dl *Downloader) bool {
	dl.PrintLog("SortImage-IsCanHandler", "判断接受处理dl.err=", dl.err, "dl.images=", len(dl.images))

	if dl.err != nil || len(dl.images) == 0 {

		return false
	}
	return true
}
