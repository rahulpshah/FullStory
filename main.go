package main

func main() {

}

// multiply multiples a and b
func multiply(a, b int) int {
	if b == 0 {
		return 0
	}
	return a + multiply(a, b-1)
}
