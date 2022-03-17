/**
 * @Author: fuxiao
 * @Email: 576101059@qq.com
 * @Date: 2022/3/17 10:38 上午
 * @Desc: TODO
 */

package main

import (
	"fmt"
	"reflect"
)

func test(s interface{}) {
	if reflect.ValueOf(s).IsNil() {
		fmt.Println("it's empty")
	} else {
		fmt.Println("it's non-empty")
	}
}

func main() {
	var s *string

	test(s)
}
