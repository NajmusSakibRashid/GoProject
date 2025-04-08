package main	
import "fmt"

// Function to reverse a string
func reverseString(s string) string {
	runes := []rune(s)
	for i, j := 0, len(runes)-1; i < j; i, j = i+1, j-1 {
		runes[i], runes[j] = runes[j], runes[i]
	}
	return string(runes)
}

func main() {
	// Example usage of the function
	original := "Hello, world!"
	reversed := reverseString(original)
	fmt.Println(reversed) // Output: !dlrow ,olleH
}