/*
 * @Author: GG
 * @Date: 2022-12-05 09:35:44
 * @LastEditTime: 2022-12-05 15:50:07
 * @LastEditors: GG
 * @Description: 数字工具类
 * @FilePath: \golang-utils\digit\digit.go
 *
 */
package digit

import (
	"math"
	"strconv"
)

type slice map[int]int

type Digit struct {
	Num   int   // 值
	Len   int   // 长度
	Slice slice // 元素拆解后的map
}

/**
 * @description: 实例化一个数字结构体
 * @param {int} n
 * @return {*}
 */
func NewDigit(n int) *Digit {
	return &Digit{
		Num:   n,
		Len:   Len(n),
		Slice: TakeApart(n),
	}
}

/**
 * @description: 获取数字长度
 * @param {int} n
 * @return {*}
 */
func Len(n int) int {
	return len(ToString(n))
}

/**
 * @description: 将int转成字符串
 * @param {int} n
 * @return {*}
 */
func ToString(n int) string {
	return strconv.Itoa(n)
}

/**
 * @description: 将string转成int
 * @param {string} n
 * @return {*}
 */
func ToInt(n string) (int, error) {
	return strconv.Atoi(n)
}

/**
 * @description: 将一个数字每个元素拆解成一个map
 * @param {int} n
 * @return {*}
 */
func TakeApart(n int) slice {
	// n = 1234
	len := Len(n)                    // 4
	currentSlice := make(slice, len) // 数字拆解后的map
	k := 1                           // 下标

	// for i := len ; i > 0; i-- {
	// 	// i = 4, 1234 % 10000 / 1000 = 1234/1000 = 1
	// 	// i = 3, 1234 % 1000 / 100 = 234/100 = 2
	// 	// i = 2, 1234 % 100 / 10 = 34/10 = 3
	// 	// i = 1, 1234 % 10 / 1 = 4/1 = 4
	// 	e := n % int(math.Pow10(i)) / int(math.Pow10(i-1))
	// 	currentSlice[k] = e
	// 	k++
	// }

	// 算法2
	for i := len; i > 0; i-- {
		// i=4,1234/1000%10 = 1
		// i=3,1234/100%10 = 2
		// i=2,1234/10%10 = 3
		// i=1,1234/1%10 = 4
		e := n / int(math.Pow10(i-1)) % 10
		currentSlice[k] = e
		k++
	}
	return currentSlice
}

/**
 * @description: 以字符方式截取
 * @param {int} i 数值
 * @param {int} m 下标
 * @param {int} n 下标
 * @return {*}
 */
func SliceFew(i, m, n int) string {
	s := ToString(i)
	digit := s[m:n]
	return digit
}

/**
 * @description: 获取对应位数值,填充0
 * @param {int} i
 * @return {*}
 */
func (d *Digit) GetNumber(i int) int {
	// d.Num = 1234, len = 4
	// i = 1, 1 * 10^(4-1) = 1 * 10*10*10 = 1000
	// i = 2, 2 * 10^(4-2) = 2 * 10*10 = 200
	// i = 3, 3 * 10^(4-3) = 3 * 10 = 30
	// i = 4, 4 * 10^(4-4) = 4 * 1 = 4
	return d.Slice[i] * int(math.Pow10(d.Len-i))
}

/**
 * @description: 返回后几位数
 * @param {int} n
 * @return {*}
 */
func (d *Digit) LastFew(n int) int {
	m := 0
	l := d.Len
	if n >= l {
		return d.Num
	}
	for i := 0; i < n; i++ {
		m += d.GetNumber(l - i)
	}
	return m
}

/**
 * @description: 返回前几位数，填充0
 * @param {int} n
 * @return {*}
 */
func (d *Digit) FirstFew(n int) int {
	m := 0
	l := d.Len
	if n >= l {
		return d.Num
	}
	for i := 1; i <= n; i++ {
		m += d.GetNumber(i)
	}
	return m
}

/**
 * @description: 完全数
 * @return {*}
 */
func (d *Digit) PerfectNumber() bool {
	m := d.Num
	for i := 1; i < d.Len; i++ {
		if d.Num%i == 0 {
			m -= i
		}
	}

	return m == 0
}
