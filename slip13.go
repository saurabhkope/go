package main
import "fmt"

func main(){
	even:=0
	odd:=0
	for i := 0; i <100; i++{
		if i%2==0 {
			even+=i
		}else{
			odd+=i
		}
	}
	fmt.Println("Sum of all even in 100: ",even);
	fmt.Println("Sum of all odd in 100: ",odd);
}