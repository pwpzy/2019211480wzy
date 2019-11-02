package main

import (
	"fmt"
	"time"
)

func sushu(a int) {

	/*for i := 2; i < a; i++ {

		if (a%i == 0) {

			break
		} else {
			fmt.Println(a)
			break
		}
	}*/
	//for i:=2;i<=10000;i++{
		flag:=1
		for j:=2;j<=a/2;j++{
			if(a%j==0){
				flag=0
				break
			}

		}
		if(flag!=0){
			fmt.Println(a)

		}
	//}
}

func main() {
	for j := 2; j <= 10000; j++ {

		time.Sleep(1)
		go  sushu(j)
	}
}
