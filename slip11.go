package main
import "fmt"

func main() {
	var n int 
	fmt.Println("Enter a number: ")
	fmt.Scanln(&n)
	
	if n>=10 && n<=99{
		fmt.Println(n," is two digit number");
	}else{
		fmt.Println(n,"is not two digit number");
	}

}