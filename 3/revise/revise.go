/**
 * @Author: fuxiao
 * @Email: 576101059@qq.com
 * @Date: 2022/3/17 5:24 下午
 * @Desc: TODO
 */

package main

import (
	"fmt"
	"sync"
)

func main() {
	var wg sync.WaitGroup

	for i := 0; i < 5; i++ {
		wg.Add(1)
		go func(i int) {
			defer wg.Done()
			fmt.Printf("index: %d\n", i)
		}(i)
	}

	wg.Wait()

	fmt.Println("over")
}
