/*
 * @Author: your name
 * @Date: 2021-01-05 13:21:23
 * @LastEditTime: 2021-01-05 13:27:37
 * @LastEditors: Please set LastEditors
 * @Description: In User Settings Edit
 * @FilePath: /demo/calc/computer.go
 */
package calc

// Sum 加
func Sum(a, b int) (c int) {
	c = a + b
	return
}

// Reduce 加
func Reduce(a *int, b int) {
	*a -= b
}

// Mul 乘
func Mul(a, b float32) float32 {
	return a * b
}
