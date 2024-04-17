package main
import "fmt"
func table(a int){
	fmt.Println("Multiplication Table");
	for i := 1; i <= 10; i++{
		fmt.Println(a," * ", i," = ",i*a);
	}
}
func main() {
	a:=5
	table(a);
}