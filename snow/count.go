/*
 * @Author: your name
 * @Date: 2021-01-05 13:29:21
 * @LastEditTime: 2021-01-05 13:57:18
 * @LastEditors: Please set LastEditors
 * @Description: In User Settings Edit
 * @FilePath: /demo/snow/count.go
 */
package snow

import (
	"fmt"

	"pkg.deepin.com/web/service/demo/calc"
)

func initew() {
	num := calc.Sum(2, 3)
	fmt.Printf("num: %d\n", num)

	var (
		it1 = int(10)
		it2 = int(2)
		fl1 = float32(12)
		fl2 = float32(1.24)
		fl  float32
	)
	calc.Reduce(&it1, it2)
	fmt.Printf("it1: %d\n", it1)
	fl = calc.Mul(fl1, fl2)
	fmt.Printf("fl: %.1f\n", fl)

}
