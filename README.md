## Threaded Timed Implementations for Factorial Calculation in Go

This repository provides three times and threaded different implementations for calculating the factorial of a number in Go:

* **Iterative:** Calculates the factorial using a loop, efficient for smaller numbers.
* **Recursive:** Calculates the factorial using recursion, elegant but less efficient for large numbers.
* **Tree-Based (Illustrative):** Simulates a tree-based approach for demonstration purposes. Replace it with an actual tree implementation for real-world use.

### Features

* Each implementation runs in a separate thread for performance comparison.
* Timing measurements are included to show execution time differences.
* User input allows for calculating factorials of various numbers.

### How to Use

1. Clone this repository.
2. Run `go run FactorialInThreads.go`.
3. Enter a non-negative integer when prompted.
4. Observe the calculated factorial and execution time for each approach.

### Code Structure

* `main.go`: Entry point, handles user input, starts threads, and prints results.
* `factorial_iterative.go`: Implements iterative factorial calculation.
* `factorial_recursive.go`: Implements recursive factorial calculation.
* `factorial_tree.go`: Simulates a tree-based approach (replace with actual implementation).

### Additional Notes

* This code is for demonstration purposes and may not be optimized for large factorials.
* Consider using `int64` or `math/big.Int` for handling larger numbers and preventing overflow.
* Replace the `factorialTree` function with an actual tree structure for the tree-based approach.

### Contributing

Pull requests and suggestions are welcome. Please ensure your contributions follow Go conventions and coding style.

## License

This project is licensed under the MIT License. See the [`LICENSE`](LICENSE.MD) file for details.

## Exampe with user Interaction

<pre>
➜  FactorialInThreads go run FactorialInThreads.go 
Enter a number: 10
Iterative factorial: 3628800
Recursive factorial: 3628800
Iterative calculation time: 52.709µs
Recursive calculation time: 49.607µs
Tree-based factorial calculation in progress...
Tree-based factorial: 3628800
Tree-based calculation time: 1.001115727s
Total execution time: 1.001193495s
➜  FactorialInThreads 
</re>