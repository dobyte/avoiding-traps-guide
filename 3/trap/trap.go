/**
 * @Author: fuxiao
 * @Email: 576101059@qq.com
 * @Date: 2022/3/17 5:14 下午
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
		go func(i int) {
			wg.Add(1)
			defer wg.Done()
			fmt.Printf("index: %d\n", i)
		}(i)
	}

	wg.Wait()

	fmt.Println("over")
}
