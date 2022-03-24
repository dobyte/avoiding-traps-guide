/**
 * @Author: fuxiao
 * @Email: 576101059@qq.com
 * @Date: 2022/3/24 4:44 下午
 * @Desc: TODO
 */

package main

import (
	"fmt"
	"io/ioutil"
	"regexp"
)

func find(filename string) []byte {
	b, _ := ioutil.ReadFile(filename)

	loc := regexp.MustCompile("1[1-9][0-9]{9}").FindIndex(b)

	if loc == nil {
		return nil
	}

	return b[loc[0]:loc[1]]
}

func main() {
	telephone := find("telephone.txt")

	fmt.Printf("%s\n", telephone)
}
