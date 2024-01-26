package main

import (
	"fmt"
	"sync"
	"time"
)

// Iterative factorial function
func factorialIterative(n int64) int64 {
	result := int64(1)
	for i := int64(2); i <= n; i++ {
		result *= i
	}
	return result
}

// Recursive factorial function
func factorialRecursive(n int64) int64 {
	if n == 0 {
		return 1
	}
	return n * factorialRecursive(n-1)
}

// Tree-based factorial calculation (illustrative approach)
func factorialTree(n int64) int64 {
	// Simulate a tree-like structure for demonstration
	// Actual implementation would involve tree data structures
	time.Sleep(time.Second) // Simulate tree processing time
	fmt.Println("Tree-based factorial calculation in progress...")
	result := int64(1)
	for i := int64(2); i <= n; i++ {
		result *= i
	}
	return result
}

func main() {
	var number int64
	fmt.Print("Enter a number: ")
	fmt.Scanf("%d", &number)

	var wg sync.WaitGroup
	wg.Add(3)

	start := time.Now()
	go func() {
		defer wg.Done()
		startTime := time.Now()
		fmt.Println("Iterative factorial:", factorialIterative(number))
		fmt.Println("Iterative calculation time:", time.Since(startTime))
	}()

	go func() {
		defer wg.Done()
		startTime := time.Now()
		fmt.Println("Recursive factorial:", factorialRecursive(number))
		fmt.Println("Recursive calculation time:", time.Since(startTime))
	}()

	go func() {
		defer wg.Done()
		startTime := time.Now()
		fmt.Println("Tree-based factorial:", factorialTree(number))
		fmt.Println("Tree-based calculation time:", time.Since(startTime))
	}()

	wg.Wait()
	fmt.Println("Total execution time:", time.Since(start))
}
