package heapsort
import (
"testing"
)

func TestMainHeadsort(t *testing.T){
array :=[]int{1,32,41,56,123,78,90,32,10}
num:=MainHeadsort(array)
t.Log(num)
}