// 堆排序算法

package heapsort

// import "fmt"
import "os"
import "strconv"
import "tree"


// 1.exe 1 3 4 6 0 8 4 31

func MainHeadsort (array []int)[]int{
if len(array)==0{
args:=os.Args[1:]
for _,v:=range args{
x,_:=strconv.Atoi(v)
array=append(array,x)
}
}
// fmt.Println("begin:",array)

sort:=Headsort(array)
// fmt.Println("end:",sort)
return sort
}

// 保证堆的性质,i为算法中的序号，比如array中第一个元素序号为1
func MaxHeapify(array []int, i int){
var largest int
l:=tree.Left(i)
r:=tree.Right(i)

// fmt.Printf("l=%d   r=%d\n",l,r)

// go 的数组中第一个的序号为0，而算法中第一个元素的序号为1
if l-1<len(array){

if array[l-1]>array[i-1]{
largest=l
}else{
largest=i
}

}else{
largest=-1
}

// fmt.Println("left largest:",largest)

if r-1<len(array){
if array[r-1]>array[largest-1]{
largest=r
}
}else{
if largest==-1{
largest=-1
}
}

// fmt.Println("right largest:",largest)

if largest!=i&&largest!=-1{

tree.Swap(&array[i-1],&array[largest-1])
//c:=array[i-1]
//array[i-1]=array[largest-1]
//array[largest-1]=c

// fmt.Printf("保证堆的性质%d  %v\n：",i,array)
// 递归
MaxHeapify(array,largest)
}
return
}

// 建立最大堆
func BuildMaxHeap(array []int){
for i:=len(array)/2;i>0;i--{
// fmt.Println(i)
MaxHeapify(array,i)
}
// fmt.Println(array)
}



// 排序
func Headsort(array []int)[]int{
var sort []int
BuildMaxHeap(array)
t:=len(array)
for i:=t;i>1;i--{
sort =append(sort,array[0])
tree.Swap(&array[0],&array[i-1])
array=array[0:len(array)-1]
MaxHeapify(array,1)

// fmt.Println("now:",array)

}
sort =append(sort,array[0])
return sort
}
