/*
 * @Author: GG
 * @Date: 2022-12-05 09:39:10
 * @LastEditTime: 2022-12-05 15:43:41
 * @LastEditors: GG
 * @Description:
 * @FilePath: \golang-utils\main.go
 *
 */
package main

import (
	"fmt"
	"golang-utils/digit"
)

func main() {
	fmt.Println("first commit")
	d := digit.NewDigit(123456)
	fmt.Printf("digit: %v\n", d)

	fmt.Printf("digit.PerfectNumber(): %v\n", d.PerfectNumber())

	fmt.Printf("SliceFew(12345, 2, 3): %v\n", digit.SliceFew(12345, 2, 3))

	fmt.Printf("digit.LastFew(2): %v\n", d.LastFew(2))
	fmt.Printf("digit.FirstFew(4): %v\n", d.FirstFew(4))
	fmt.Printf("digit.GetNumber(3): %v\n", d.GetNumber(3))
}
