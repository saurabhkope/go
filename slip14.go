package main

import "fmt"

func main() {
    // Creating a slice
    slice := []int{1, 2, 3, 4, 5}

    // Appending elements to the slice
    slice = append(slice, 6)
    fmt.Println("After appending 6:", slice)

    // Removing an element from the slice
    indexToRemove := 2
    slice = append(slice[:indexToRemove], slice[indexToRemove+1:]...)
    fmt.Println("After removing element at index", indexToRemove, ":", slice)

    // Copying a slice
    copiedSlice := make([]int, len(slice))
    copy(copiedSlice, slice)
    fmt.Println("Copied slice:", copiedSlice)
}
