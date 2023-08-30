package classification

type UrlPath struct {
	QueryUrl string
	Pages    int
	Remark   string
}

var Paths = map[string][]UrlPath{
	"小学-三年级-习近平新时代中国特色社会主义思想学生读本": {
		{
			QueryUrl: "https://book.pep.com.cn/1291001103221/files/mobile/%d.jpg?%d",
			Pages:    2,
		},
	},
	"小学-五年级-习近平新时代中国特色社会主义思想学生读本": {
		{
			QueryUrl: "https://book.pep.com.cn/1291001403221/files/mobile/%d.jpg?%d",
			Pages:    102,
		},
	},
	"小学-一年级-道德与法治": {
		{
			QueryUrl: "https://book.pep.com.cn/1284001101161/files/mobile/%d.jpg?%d",
			Pages:    78,
		},
		{
			QueryUrl: "https://book.pep.com.cn/1284001102161/files/mobile/%d.jpg?%d",
			Pages:    80,
		},
	},
	"小学-一年级-语文": {
		{
			QueryUrl: "https://book.pep.com.cn/1211001101161/files/mobile/%d.jpg?%d",
			Pages:    128,
		},
		{
			QueryUrl: "https://book.pep.com.cn/1211001102161/files/mobile/%d.jpg?%d",
			Pages:    128,
		},
	},
	"小学-一年级-数学": {
		{
			QueryUrl: "https://book.pep.com.cn/1221001101121/files/mobile/%d.jpg?%d",
			Pages:    120,
		},
		{
			QueryUrl: "https://book.pep.com.cn/1221001102121/files/mobile/%d.jpg?%d",
			Pages:    112,
		},
	},
	"小学-一年级-英语": {
		{
			QueryUrl: "https://book.pep.com.cn/1212001101123/files/mobile/%d.jpg?%d",
			Pages:    78,
		},
		{
			QueryUrl: "https://book.pep.com.cn/1212001102123/files/mobile/%d.jpg?%d",
			Pages:    78,
		},
	},
	"小学-一年级-科学": {
		{
			QueryUrl: "https://book.pep.com.cn/1244001101171/files/mobile/%d.jpg?%d",
			Pages:    52,
		},
		{
			QueryUrl: "https://book.pep.com.cn/1244001102172/files/mobile/%d.jpg?%d",
			Pages:    20,
		},
		{
			QueryUrl: "https://book.pep.com.cn/1244001102171/files/mobile/%d.jpg?%d",
			Pages:    52,
		},
	},
	"小学-一年级-音乐": {
		{
			QueryUrl: "https://book.pep.com.cn/1262001101122/files/mobile/%d.jpg?%d",
			Pages:    68,
		},
		{
			QueryUrl: "https://book.pep.com.cn/1262001101121/files/mobile/%d.jpg?%d",
			Pages:    68,
		},
		{
			QueryUrl: "https://book.pep.com.cn/1262001102122/files/mobile/%d.jpg?%d",
			Pages:    68,
		},
		{
			QueryUrl: "https://book.pep.com.cn/1262001102121/files/mobile/%d.jpg?%d",
			Pages:    68,
		},
	},
	"小学-一年级-体育与健康": {
		{
			QueryUrl: "https://book.pep.com.cn/1272001103221/files/mobile/%d.jpg?%d",
			Pages:    344,
		},
	},
	"小学-一年级-美术": {
		{
			QueryUrl: "https://book.pep.com.cn/1263001101121/files/mobile/%d.jpg?%d",
			Pages:    52,
		},
		{
			QueryUrl: "https://book.pep.com.cn/1263001102121/files/mobile/%d.jpg?%d",
			Pages:    52,
		},
	},

	"初中-七年级-数学": {
		{
			QueryUrl: "https://book.pep.com.cn/1321001101121/files/mobile/%d.jpg?%d",
			Pages:    162,
			Remark:   "上册",
		},
		{
			QueryUrl: "https://book.pep.com.cn/1321001102121/files/mobile/%d.jpg?%d",
			Pages:    176,
			Remark:   "下册",
		},
	},
	"初中-八年级-数学": {
		{
			QueryUrl: "https://book.pep.com.cn/1321001201131/files/mobile/%d.jpg?%d",
			Pages:    170,
			Remark:   "上册",
		},
		{
			QueryUrl: "https://book.pep.com.cn/1321001202131/files/mobile/%d.jpg?%d",
			Pages:    148,
			Remark:   "下册",
		},
	},
	"初中-九年级-数学": {
		{
			QueryUrl: "https://book.pep.com.cn/1321001301141/files/mobile/%d.jpg?%d",
			Pages:    164,
			Remark:   "上册",
		},
		{
			QueryUrl: "https://book.pep.com.cn/1321001302141/files/mobile/%d.jpg?%d",
			Pages:    124,
			Remark:   "下册",
		},
	},
	"高中-必修-数学-第一册": {
		{
			QueryUrl: "https://book.pep.com.cn/1421001121201/files/mobile/%d.jpg?%d",
			Pages:    152,
			Remark:   "B版",
		},
		{
			QueryUrl: "https://book.pep.com.cn/1421001121191/files/mobile/%d.jpg?%d",
			Pages:    270,
			Remark:   "A版",
		},
	},
	"高中-选择性必修-数学-第一册": {
		{
			QueryUrl: "https://book.pep.com.cn/1421001127202/files/mobile/%d.jpg?%d",
			Pages:    192,
			Remark:   "B版",
		},
		{
			QueryUrl: "https://book.pep.com.cn/1421001127201/files/mobile/%d.jpg?%d",
			Pages:    156,
			Remark:   "A版",
		},
	},
	"高中-选择性必修-数学-第二册": {
		{
			QueryUrl: "https://book.pep.com.cn/1421001128202/files/mobile/%d.jpg?%d",
			Pages:    140,
			Remark:   "B版",
		},
		{
			QueryUrl: "https://book.pep.com.cn/1421001128201/files/mobile/%d.jpg?%d",
			Pages:    114,
			Remark:   "A版",
		},
	},
	"高中-必修-数学-第二册": {
		{
			QueryUrl: "https://book.pep.com.cn/1421001122201/files/mobile/%d.jpg?%d",
			Pages:    192,
			Remark:   "B版",
		},
		{
			QueryUrl: "https://book.pep.com.cn/1421001122191/files/mobile/%d.jpg?%d",
			Pages:    280,
			Remark:   "A版",
		},
	},
	"高中-必修-数学-第三册": {
		{
			QueryUrl: "https://book.pep.com.cn/1421001123201/files/mobile/%d.jpg?%d",
			Pages:    126,
			Remark:   "B版",
		},
	},
	"高中-必修-数学-第四册": {
		{
			QueryUrl: "https://book.pep.com.cn/1421001137201/files/mobile/%d.jpg?%d",
			Pages:    142,
			Remark:   "B版",
		},
	},
}
