package main
import "fmt"
func perOperation(a,b int ) (int , int ,int ,int ){
	add:=a+b
	sub:=a-b
	mul:=a*b
	div:=a/b
	return add,sub,mul,div
}
func main() {	
	a:=5
	b:=2
	add,sub,mul,div := perOperation(a,b);
	fmt.Println("Add: ",add);
	fmt.Println("Sub: ",sub);
	fmt.Println("Mul: ",mul);
	fmt.Println("Div: ",div);
}
