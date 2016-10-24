package quicksort

import (
	// "fmt"
	"tree"
)

// 快速排序(从小到大)
func Quicksort(a []int, p int, r int) []int {
	if p < r {
		q := RandomPartition(a, p, r)
		Quicksort(a, p, q-1)
		Quicksort(a, q+1, r)
	}
	return a
}

func RandomPartition(a []int, p int, r int) int {
	x := tree.Randnum(p, r) //随机化
	tree.Swap(&a[x-1], &a[r-1])
	return Partition(a, p, r)
}

// 对子数组分解
// golang中，第一个元素对应的为a[0]
// x是主元，i标记第一部分尾部，初始值为0，如果第j个元素小于主元，把第j个元素加入第一部分
// 这样第一部分每个元素小于x且小于第二部分的每个元素
func Partition(a []int, p int, r int) int {

	x := a[r-1]
	i := p - 1
	for j := p; j <= r-1; j++ {
		if a[j-1] <= x {
			//if a[j-1] >= x {
			tree.Swap(&a[j-1], &a[(i-1)+1])
			i++
			//fmt.Println("swap:", i)
		}
		//fmt.Printf("partition%d:%v\n", r, a)
	}
	tree.Swap(&a[(i-1)+1], &a[r-1])
	//fmt.Println("---------ok-----------:", a, i+1)
	return i + 1
}
