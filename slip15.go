package main
import "fmt"
func opration(a,b float64)(float64,float64,float64){
	add:=a+b
	sub:=a-b
	mul:=a*b
	return add,sub,mul
}
func main() {
	a:=10.0;
	b:=5.0;
	add,sub,mul := opration(a,b);
	fmt.Println("Add : ",add);
	fmt.Println("Sub : ",sub);
	fmt.Println("Mul : ",mul);
}