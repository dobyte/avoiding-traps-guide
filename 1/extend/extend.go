/**
 * @Author: fuxiao
 * @Email: 576101059@qq.com
 * @Date: 2022/3/17 2:58 下午
 * @Desc: TODO
 */

package main

import "fmt"

func main() {
	var s *string
	
	var i interface{} = s
	
	if i == nil {
		fmt.Println("it's empty")
	} else {
		fmt.Println("it's non-empty")
	}
}
