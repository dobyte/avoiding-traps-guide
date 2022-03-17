/**
 * @Author: fuxiao
 * @Email: 576101059@qq.com
 * @Date: 2022/3/17 4:32 下午
 * @Desc: TODO
 */

package main

import (
	"fmt"
	"time"
)

type People struct {
	Name string
	Age  int
}

func main() {
	peoples := []People{
		{
			Name: "Tom",
			Age:  18,
		},
		{
			Name: "lucy",
			Age:  15,
		},
		{
			Name: "lily",
			Age:  20,
		},
	}

	for _, people := range peoples {
		go func(people People) {
			fmt.Printf("My name is %s, I am %d years old.\n", people.Name, people.Age)
		}(people)
	}

	time.Sleep(time.Second)
}
