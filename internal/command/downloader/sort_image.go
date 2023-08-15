package downloader

import (
	"github.com/spf13/cast"
	"github.com/urfave/cli/v2"
	"sort"
	"strings"
)

type SortImage struct {
	BaseHandler
}

type fileInfo struct {
	Name string // 文件名
	Page int    //页数
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
				if len(image) < 2 {
					continue
				}
				rImage := strings.ReplaceAll(image, "\\", "/")
				imageSlice := strings.Split(rImage, "/")
				page := strings.Split(imageSlice[len(imageSlice)-1], ".")[0]

				sortImage = append(sortImage, fileInfo{Name: image, Page: cast.ToInt(page)})
			}

			if len(sortImage) > 0 {
				sort.Slice(
					sortImage, func(i, j int) bool {
						return sortImage[i].Page < sortImage[j].Page
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
