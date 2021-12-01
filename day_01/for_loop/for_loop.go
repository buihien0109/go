package for_loop

import "fmt"

// For loop
func forLoopExample(arr []int) (sum int) {
	for i := 0; i < len(arr); i++ {
		sum += arr[i]
	}
	return sum
}

// For ... range
func forRangeExample(arr []int) (sum int) {
	for _, value := range arr {
		sum += value
	}

	return sum
}

// While
func whileExample(arr []int) (sum int) {
	i := 0
	for i < len(arr) {
		sum += arr[i]
		i++
	}
	return sum
}

// Do ... while
func doWhileExample(arr []int) (sum int) {
	i := 0
	for {
		if i == len(arr) {
			break
		}
		sum += arr[i]
		i++
	}

	return sum
}

func DemoFor() {
	numbers := []int{1,2,3,4,5}

	result := forLoopExample(numbers)
	fmt.Println(result)

	result1 := forRangeExample(numbers)
	fmt.Println(result1)

	result2 := whileExample(numbers)
	fmt.Println(result2)

	result3 := doWhileExample(numbers)
	fmt.Println(result3)
}
