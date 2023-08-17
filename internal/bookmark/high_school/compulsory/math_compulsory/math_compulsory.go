package math_compulsory

import "github.com/pdfcpu/pdfcpu/pkg/pdfcpu"

var MathCompulsoryFirstB = []pdfcpu.Bookmark{
	{PageFrom: 1, Title: "封面"},
	{PageFrom: 5, Title: "前言"},
	{PageFrom: 7, Title: "目录"},
	{PageFrom: 9, Title: "第一章: 集合与常用逻辑用语", Children: []pdfcpu.Bookmark{
		{PageFrom: 10, Title: "1.0: 本章导语"},
		{PageFrom: 11, Title: "1.1: 集合", Children: []pdfcpu.Bookmark{
			{PageFrom: 11, Title: "1.1.1 集合及其表示方法"},
			{PageFrom: 18, Title: "1.1.2 集合的基本关系"},
			{PageFrom: 23, Title: "1.1.3 集合的基本运算"},
		}},
		{PageFrom: 31, Title: "1.2: 常用逻辑用语", Children: []pdfcpu.Bookmark{
			{PageFrom: 31, Title: "1.2.1 命题与量词"},
			{PageFrom: 36, Title: "1.2.2 全称量词命题与存在量词命题的否定"},
			{PageFrom: 39, Title: "1.2.3 充分条件,必要条件"},
		}},
		{PageFrom: 47, Title: "本章小结1"},
	}},
	{PageFrom: 51, Title: "第二章: 等式与不等式", Children: []pdfcpu.Bookmark{
		{PageFrom: 52, Title: "2.0: 本章导语"},

		{PageFrom: 53, Title: "2.1: 等式", Children: []pdfcpu.Bookmark{
			{PageFrom: 53, Title: "2.1.1 等式的性质与方程解集"},
			{PageFrom: 57, Title: "2.1.2 一元二次方程的解集及其根与系数的关系"},
			{PageFrom: 62, Title: "2.1.3 方程组的解集"},
		}},
		{PageFrom: 69, Title: "2.2: 不等式", Children: []pdfcpu.Bookmark{
			{PageFrom: 69, Title: "2.2.1 不等式及其性质"},
			{PageFrom: 75, Title: "2.2.2 不等式的解集"},
			{PageFrom: 79, Title: "2.2.3 一元二次不等式的解法"},
			{PageFrom: 84, Title: "2.2.4 均值不等式及其应用"},
		}},
		{PageFrom: 91, Title: "本章小结2"},
	}},
	{PageFrom: 95, Title: "第三章: 函数", Children: []pdfcpu.Bookmark{
		{PageFrom: 96, Title: "3.0: 本章导语"},
		{PageFrom: 97, Title: "3.1: 函数的概念及其性质", Children: []pdfcpu.Bookmark{
			{PageFrom: 97, Title: "3.1.1 函数及其表示方法"},
			{PageFrom: 107, Title: "3.1.2 函数的单调性"},
			{PageFrom: 117, Title: "3.1.3 函数的奇偶性"},
		}},
		{PageFrom: 126, Title: "3.2 函数与方程,不等式之间的关系"},
		{PageFrom: 136, Title: "3.3 函数的应用(一)"},
		{PageFrom: 140, Title: "3.4 数学建模活动:决定苹果的最佳出售时间点"},
		{PageFrom: 146, Title: "本章小结3"},
	}},
	{PageFrom: 150, Title: "后记", Bold: true, Italic: true},
	{PageFrom: 152, Title: "尾部封面", Bold: true, Italic: true},
}
