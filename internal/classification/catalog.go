package classification

type UrlPath struct {
	BookID string
	Pages  int
	Remark string
}

var Paths = map[string][]UrlPath{
	"小学-三年级-习近平新时代中国特色社会主义思想学生读本": {
		{
			BookID: "1291001103221",
			Pages:  2,
		},
	},
	"小学-五年级-习近平新时代中国特色社会主义思想学生读本": {
		{
			BookID: "1291001403221",
			Pages:  102,
		},
	},
	"小学-一年级-道德与法治": {
		{
			BookID: "1284001101161",
			Pages:  78,
			Remark: "上册",
		},
		{
			BookID: "1284001102161",
			Pages:  80,
			Remark: "下册",
		},
	},
	"小学-一年级-语文": {
		{
			BookID: "1211001101161",
			Pages:  128,
			Remark: "上册",
		},
		{
			BookID: "1211001102161",
			Pages:  128,
			Remark: "下册",
		},
	},
	"小学-一年级-数学": {
		{
			BookID: "1221001101121",
			Pages:  120,
			Remark: "上册",
		},
		{
			BookID: "1221001102121",
			Pages:  112,
			Remark: "下册",
		},
	},
	"小学-一年级-英语": {
		{
			BookID: "1212001101123",
			Pages:  78,
			Remark: "上册",
		},
		{
			BookID: "1212001102123",
			Pages:  78,
			Remark: "下册",
		},
	},
	"小学-三年级-英语": {
		{
			BookID: "1212001301244",
			Pages:  106,
			Remark: "上册",
		},
		{
			BookID: "1212001302244",
			Pages:  102,
			Remark: "下册",
		},
	},
	"小学-一年级-科学": {
		{
			BookID: "1244001101171",
			Pages:  52,
			Remark: "一",
		},
		{
			BookID: "1244001102172",
			Pages:  20,
			Remark: "二",
		},
		{
			BookID: "1244001102171",
			Pages:  52,
			Remark: "三",
		},
	},
	"小学-一年级-音乐": {
		{
			BookID: "1262001101122",
			Pages:  68,
			Remark: "一",
		},
		{
			BookID: "1262001101121",
			Pages:  68,
			Remark: "二",
		},
		{
			BookID: "1262001102122",
			Pages:  68,
			Remark: "三",
		},
		{
			BookID: "1262001102121",
			Pages:  68,
			Remark: "四",
		},
	},
	"小学-一年级-体育与健康": {
		{
			BookID: "1272001103221",
			Pages:  344,
		},
	},
	"小学-一年级-美术": {
		{
			BookID: "1263001101121",
			Pages:  52,
			Remark: "上册",
		},
		{
			BookID: "1263001102121",
			Pages:  52,
			Remark: "下册",
		},
	},

	"初中-七年级-数学": {
		{
			BookID: "1321001101121",
			Pages:  162,
			Remark: "上册",
		},
		{
			BookID: "1321001102121",
			Pages:  176,
			Remark: "下册",
		},
	},
	"初中-八年级-数学": {
		{
			BookID: "1321001201131",
			Pages:  170,
			Remark: "上册",
		},
		{
			BookID: "1321001202131",
			Pages:  148,
			Remark: "下册",
		},
	},
	"初中-九年级-数学": {
		{
			BookID: "1321001301141",
			Pages:  164,
			Remark: "上册",
		},
		{
			BookID: "1321001302141",
			Pages:  124,
			Remark: "下册",
		},
	},
	"高中-必修-数学-第一册": {
		{
			BookID: "1421001121201",
			Pages:  152,
			Remark: "B版",
		},
		{
			BookID: "1421001121191",
			Pages:  270,
			Remark: "A版",
		},
	},
	"高中-选择性必修-数学-第一册": {
		{
			BookID: "1421001127202",
			Pages:  192,
			Remark: "B版",
		},
		{
			BookID: "1421001127201",
			Pages:  156,
			Remark: "A版",
		},
	},
	"高中-选择性必修-数学-第二册": {
		{
			BookID: "1421001128202",
			Pages:  140,
			Remark: "B版",
		},
		{
			BookID: "1421001128201",
			Pages:  114,
			Remark: "A版",
		},
	},
	"高中-选择性必修-数学-第三册": {
		{
			BookID: "1421001129202",
			Pages:  126,
			Remark: "B版",
		},
		{
			BookID: "1421001129201",
			Pages:  156,
			Remark: "A版",
		},
	},
	"高中-必修-数学-第二册": {
		{
			BookID: "1421001122201",
			Pages:  192,
			Remark: "B版",
		},
		{
			BookID: "1421001122191",
			Pages:  280,
			Remark: "A版",
		},
	},
	"高中-必修-数学-第三册": {
		{
			BookID: "1421001123201",
			Pages:  126,
			Remark: "B版",
		},
	},
	"高中-必修-数学-第四册": {
		{
			BookID: "1421001137201",
			Pages:  142,
			Remark: "B版",
		},
	},
}
