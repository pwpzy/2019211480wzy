package main

import (
	"fmt"
)

var (
	myres = make(map[int]int, 20)
	//mu    sync.Mutex
)


func factorial(n int) {
	var res = 1
	//var b int
	ch := make(chan int,1)
	for i := 1; i <= n; i++ {
		res *= i
	}
	//mu.Lock()
	ch <- res
	myres[n] = <- ch
	//myres[n] =res

	//mu.Unlock()
}

func main() {

	for i := 1; i <= 20; i++ {
		go factorial(i)
		//time.Sleep(100 * time.Microsecond)
	}
	//mu.Lock()
	//	b := make(chan int, 20)

	//	for i := 0; i <= 20; i++ {
	//		b <- myres[i]
	//		fmt.Println(<-b)
	//	}
	//}
	for i, v := range myres {

		fmt.Printf("myres[%d] = %d\n", i, v)
	}
	//mu.Unlock()

}