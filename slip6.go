package main

import "fmt"

// CopyArray copies all elements of source array into target array
func CopyArray(source []int, target []int) {
    // Ensure the target array has enough capacity to hold all elements
    if len(target) < len(source) {
        fmt.Println("Target array does not have enough capacity.")
        return
    }

    // Copy elements from source to target array
    for i := 0; i < len(source); i++ {
        target[i] = source[i]
    }
}

func main() {
    // Define source and target arrays
    source := []int{1, 2, 3, 4, 5}
    target := make([]int, len(source))

    // Call the CopyArray method to copy elements
    CopyArray(source, target)

    // Print the target array to verify the copy
    fmt.Println("Source array:", source)
    fmt.Println("Target array:", target)
}
