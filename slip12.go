package main
import "fmt"
func swap(a, b *int){
	
	temp:=*a
	*a=*b
	*b=temp
}
func main(){
	a:=10;
	b:=20;
	fmt.Println("Before swap");
	fmt.Println("a : ",a)
	fmt.Println("b : ",b)

	swap(&a, &b);

	fmt.Println("After swap")
	fmt.Println("a : ",a)
	fmt.Println("b : ",b)
}