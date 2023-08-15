package bookmark

import (
	"github.com/maogou/pep-ebook/internal/bookmark/primary_school/eighth_grade/math_eighth"
	"github.com/maogou/pep-ebook/internal/bookmark/primary_school/seventh_grade/math_seventh"
	"github.com/pdfcpu/pdfcpu/pkg/pdfcpu"
)

var Bookmark = map[string][]pdfcpu.Bookmark{
	"小学/三年级/习近平新时代中国特色社会主义思想学生读本_0": math_seventh.Bms7_0,
	"初中/七年级/数学/上册": math_seventh.Bms7_0,
	"初中/七年级/数学/下册": math_seventh.Bms7_1,
	"初中/八年级/数学/上册": math_eighth.Bms8_0,
}
