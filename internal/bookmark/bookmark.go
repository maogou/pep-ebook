package bookmark

import (
	"github.com/maogou/pep-ebook/internal/bookmark/high_school/compulsory/math_compulsory"
	"github.com/maogou/pep-ebook/internal/bookmark/high_school/selective_compulsory/math_selective_compulsory"
	"github.com/maogou/pep-ebook/internal/bookmark/middle_school/eighth_grade/math_eighth"
	"github.com/maogou/pep-ebook/internal/bookmark/middle_school/ninth_grade/math_ninth"
	"github.com/maogou/pep-ebook/internal/bookmark/middle_school/seventh_grade/math_seventh"
	"github.com/pdfcpu/pdfcpu/pkg/pdfcpu"
)

var Bookmark = map[string][]pdfcpu.Bookmark{
	"小学/三年级/习近平新时代中国特色社会主义思想学生读本_0": math_seventh.Bms7_0,
	"初中/七年级/数学/上册":    math_seventh.Bms7_0,
	"初中/七年级/数学/下册":    math_seventh.Bms7_1,
	"初中/八年级/数学/上册":    math_eighth.Bms8_0,
	"初中/八年级/数学/下册":    math_eighth.Bms8_1,
	"初中/九年级/数学/上册":    math_ninth.Math9_0,
	"初中/九年级/数学/下册":    math_ninth.Math9_1,
	"高中/必修/数学/第一册/B版": math_compulsory.MathCompulsoryFirstB,
	"高中/必修/数学/第一册/A版": math_compulsory.MathCompulsoryFirstA,
	"高中/必修/数学/第二册/B版": math_compulsory.MathCompulsorySecondB,
	"高中/必修/数学/第二册/A版": math_compulsory.MathCompulsorySecondA,
	"高中/必修/数学/第三册/B版": math_compulsory.MathCompulsoryThirdB,
	"高中/必修/数学/第四册/B版": math_compulsory.MathCompulsoryFourB,

	"高中/选择性必修/数学/第一册/B版": math_selective_compulsory.MathSelectiveCompulsoryFirstB,
	"高中/选择性必修/数学/第一册/A版": math_selective_compulsory.MathSelectiveCompulsoryFirstA,
	"高中/选择性必修/数学/第二册/B版": math_selective_compulsory.MathSelectiveCompulsorySecondB,
}
