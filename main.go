package PriorQueue

import (
	"fmt"
	"strconv"
)

func main() {

	
	
	sl := []string{}
	for i := 0; i < 10; i++ {
		sl = append(sl,strconv.Itoa(i))
	}
	qu := NewQueueWithSlice(sl)
//	fmt.Println(qu.Size())
fmt.Println(qu.DequeueMin())
qu.DequeueMin()
qu.Enqueue("55")
	for qu.Size() > 0 {
		fmt.Println(qu.DequeueMin())


	}

}
