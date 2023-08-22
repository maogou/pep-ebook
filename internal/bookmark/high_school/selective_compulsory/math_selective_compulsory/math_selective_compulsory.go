package math_selective_compulsory

import "github.com/pdfcpu/pdfcpu/pkg/pdfcpu"

var MathSelectiveCompulsoryFirstB = []pdfcpu.Bookmark{
	{PageFrom: 1, Title: "封面"},
	{PageFrom: 5, Title: "前言"},
	{PageFrom: 7, Title: "目录"},
	{PageFrom: 9, Title: "第一章: 空间向量与立体几何", Children: []pdfcpu.Bookmark{
		{PageFrom: 10, Title: "本章导语"},
		{PageFrom: 11, Title: "1.1 空间向量及其运算", Children: []pdfcpu.Bookmark{
			{PageFrom: 11, Title: "1.1.1 空间向量及其运算"},
			{PageFrom: 21, Title: "1.1.2 空间向量基本定理"},
			{PageFrom: 26, Title: "1.1.3 空间向量的坐标与空间直角坐标系"},
		}},
		{PageFrom: 38, Title: "1.2 空间向量在立体几何中的应用", Children: []pdfcpu.Bookmark{
			{PageFrom: 38, Title: "1.2.1 空间中的点,直线与空间向量"},
			{PageFrom: 46, Title: "1.2.2 空间中的平面与空间向量"},
			{PageFrom: 52, Title: "1.2.3 直线与平面的夹角"},
			{PageFrom: 57, Title: "1.2.4 二面角"},
			{PageFrom: 62, Title: "1.2.5 空间中的距离"},
		}},
		{PageFrom: 72, Title: "本章小结1"},
	}},
	{PageFrom: 77, Title: "第二章: 平面解析几何", Children: []pdfcpu.Bookmark{
		{PageFrom: 78, Title: "本章导语"},
		{PageFrom: 79, Title: "2.1 坐标法"},
		{PageFrom: 83, Title: "2.2 直线及其方程", Children: []pdfcpu.Bookmark{
			{PageFrom: 83, Title: "2.2.1 直线的倾斜角与斜率"},
			{PageFrom: 91, Title: "2.2.2 直线的方程"},
			{PageFrom: 99, Title: "2.2.3 两条直线的位置关系"},
			{PageFrom: 105, Title: "2.2.4 点到直线的距离"},
		}},
		{PageFrom: 111, Title: "2.3 圆及其方程", Children: []pdfcpu.Bookmark{
			{PageFrom: 111, Title: "2.3.1 圆的标准方程"},
			{PageFrom: 115, Title: "2.3.2 圆的一般方程"},
			{PageFrom: 118, Title: "2.3.3 直线与圆的位置关系"},
			{PageFrom: 124, Title: "2.3.4 圆与圆的位置关系"},
		}},
		{PageFrom: 131, Title: "2.4 曲线与方程"},
		{PageFrom: 137, Title: "2.5 椭圆及其方程", Children: []pdfcpu.Bookmark{
			{PageFrom: 137, Title: "2.5.1 椭圆的标准方程"},
			{PageFrom: 143, Title: "2.5.2 椭圆的几何性质"},
		}},
		{PageFrom: 152, Title: "2.6 双曲线及其方程", Children: []pdfcpu.Bookmark{
			{PageFrom: 152, Title: "2.6.1 双曲线的标准方程"},
			{PageFrom: 157, Title: "2.6.2 双曲线的几何性质"},
		}},
		{PageFrom: 166, Title: "2.7 抛物线及其方程", Children: []pdfcpu.Bookmark{
			{PageFrom: 166, Title: "2.7.1 抛物线的标准方程"},
			{PageFrom: 170, Title: "2.7.2 抛物线的几何性质"},
		}},
		{PageFrom: 176, Title: "2.8 直线与圆锥曲线的位置关系"},
		{PageFrom: 183, Title: "本章小结2"},
	}},

	{PageFrom: 189, Title: "后记", Bold: true, Italic: true},
	{PageFrom: 192, Title: "尾部封面", Bold: true, Italic: true},
}
