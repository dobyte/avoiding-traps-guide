/**
 * @Author: fuxiao
 * @Email: 576101059@qq.com
 * @Date: 2022/3/17 10:28 上午
 * @Desc: TODO
 */

package main

import "fmt"

func test(s interface{}) {
	if s == nil {
		fmt.Println("it's empty")
	} else {
		fmt.Println("it's non-empty")
	}
}

func main() {
	var s *string
	
	test(s)
}
