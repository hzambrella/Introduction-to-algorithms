// 堆排序的使用
// 优先级序列


package main
import(
"heapsort"
"fmt"
"tree"
)

func main(){
array:=[]int{10,20,5,43,13,45,23,17,14,13}
// 堆排序
fmt.Println("before:",array)
fmt.Println("now:",heapsort.MainHeadsort(array))


// 下面的优先级队列的操作都是针对最大堆的
// 返回s中具有最大关键字的元素
fmt.Println("返回s中具有最大关键字的元素:\n",hmaxnum(array))

// 去掉并返回S中的具有最大关键字的元素
l,m:=hextractmax(array)
fmt.Println("去掉并返回S中的具有最大关键字的元素:\n",l,m)

// 将元素x的关键字的值增加到k
fmt.Println("将元素x的关键字的值增加到k:\n",hincreseKey(heapsort.MainHeadsort(array),5,100))

// 把元素x插入到集合s
fmt.Println("把元素x插入到集合s:\n",hinsert(heapsort.MainHeadsort(array),100))
}

// 返回a中具有最大关键字的元素
func hmaxnum(a []int)int{
x:=heapsort.MainHeadsort(a)
return x[0]
}


// 去掉并返回a中的具有最大关键字的元素，并保持最大堆的性质
func hextractmax(a []int)(int,[]int){
x:=heapsort.MainHeadsort(a)
max:=x[0]
tree.Swap(&x[0],&x[len(x)-1])
y:=[]int{}
num:=len(x)
y=x[:num-1]
heapsort.MaxHeapify(y,1)
return max,y
}


// 将元素x的关键字的值增加到key,并保持a最大堆的性质
// i是第几个元素
func hincreseKey(a []int,i int ,key int)[]int{
if key<a[i-1]{
fmt.Println("new key is smaller than current key")
}

a[i-1]=key
// 这里采用插入循环的方式，循环不变式

for{

if i<=1{
break
}

if a[i-1]<a[tree.Parent(i)-1]{
break
}

tree.Swap(&a[i-1],&a[tree.Parent(i)-1])
i=tree.Parent(i)

// fmt.Println(a)
}
return a
}
// 插入元素k到集合a当中
func hinsert(a []int,k int)[]int{
// golang的切片看上去很简单
// 这里做的是树heap_size+1，并且让a[heap_size[a]]趋近负无穷大
a=append(a,k)

// 这里是为了保证最大堆性质
hincreseKey(a,len(a),k)

return a
}

