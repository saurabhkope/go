package main
import "fmt"
func main(){
	var num1,num2,n float32;
	fmt.Println("------Arithmatic Operation------");
	fmt.Print("Enter first Number: ")
	fmt.Scan(&num1)
	fmt.Print("Enter Second Number: ")
	fmt.Scan(&num2)
	fmt.Println("1.Add");
	fmt.Println("2.Sub");
	fmt.Println("3.Mul");
	fmt.Println("4.Div");
	fmt.Print("Enter Your Choice(1,2,3,4) : ");
	fmt.Scan(&n)
	switch(n){
		case 1:fmt.Println("Addtion is: ",num1+num2); 
			break;
		case 2:fmt.Println("Substraction is: ",num1-num2); 
			break;
		case 3:fmt.Println("Multiplication is: ",num1*num2); 
			break;
		case 4:fmt.Println("Division is: ",num1/num2); 
			break;
		default:fmt.Println("Please enter valid choice"); 
			break;
			
	}
	
	
}