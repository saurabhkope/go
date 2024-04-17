package main
import "fmt"
func perAddSub(a, b int) (int , int ){
	sum :=a + b
	dif :=a - b
	return sum, dif
}
func main(){
	num1 := 15
	num2 := 5
	add,sub:=perAddSub(num1,num2)
	fmt.Println("Addition: ",add)
	fmt.Println("Subition: ",sub)
}