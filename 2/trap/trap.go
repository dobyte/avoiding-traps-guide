/**
 * @Author: fuxiao
 * @Email: 576101059@qq.com
 * @Date: 2022/3/17 3:43 下午
 * @Desc: TODO
 */

package main

import "fmt"

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

	students := make([]*People, 0, len(peoples))

	for _, people := range peoples {
		students = append(students, &people)
	}

	for _, student := range students {
		fmt.Printf("My name is %s, I am %d years old.\n", student.Name, student.Age)
	}
}
