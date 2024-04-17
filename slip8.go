package main
import "fmt"

func main(){
	var n int
	fmt.Println("Enter a number of books: ")
	fmt.Scanln(&n)
	var bookid= make([]int,n)
	var title= make([]string,n)
	var authorName= make([]string,n)
	var price= make([]float64,n)

	for i := 0; i<n; i++ {
		fmt.Println("Enter Book Details :", i+1);
		fmt.Println("Enter BookID: ");
		fmt.Scanln(&bookid[i])

		fmt.Println("Enter BookTitle: ");
		fmt.Scanln(&title[i])

		fmt.Println("Enter AuthorName: ");
		fmt.Scanln(&authorName[i])

		fmt.Println("Enter BookPrice: ");
		fmt.Scanln(&price[i])
	}
	fmt.Println("\n----------Book Details---------");

	for i := 0;i<n;i++ {
		fmt.Println("Book Details: ",i+1);
		fmt.Println("Book ID: ",bookid[i]);
		fmt.Println("Book Title: ",title[i]);
		fmt.Println("Book Author: ",authorName[i]);
		fmt.Println("Book Price: ",price[i]);
	}
}
