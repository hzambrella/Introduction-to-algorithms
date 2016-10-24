package tree

import (
	"math/rand"
	// "fmt"
)

// 这里的i都是array的第几个元素。使用的话-1即可
// 例子 tree.Swap(&a[i-1],$a[tree.Parent(i)])

// 左儿子
func Left(i int) int {
	return 2 * i
}

// 右儿子
func Right(i int) int {
	return 2*i + 1
}

func Parent(i int) int {
	return i / 2
}

// 交换
func Swap(a *int, b *int) {
	c := *b
	*b = *a
	*a = c
}

// 随机数字
func Randnum(min int, max int) int {

	i := rand.Intn(max)
	// fmt.Println("i", i)
	if i < min {
		return Randnum(min, max)
	}
	return i
}
