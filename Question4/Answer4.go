package answer4

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

func SplitInput() {
	var results []int
	reader := bufio.NewReader(os.Stdin)

	fmt.Print("Please input number: ")
	input, err := reader.ReadString('\n')
	if err != nil {
		fmt.Println("Error reading input:", err)
		return
	}
	delcomma := strings.Replace(input, ",", "", -1)
	delcomma = strings.TrimSpace(delcomma)
	number, err := strconv.Atoi(delcomma)
	if err != nil {
		fmt.Println("Error converting to number:", err)
		return
	}
	place := 1
	for number > 0 {
		digit := number % 10
		if digit != 0 {
			results = append(results, digit*place) 
		}
		number /= 10
		place *= 10 
	}
	for i := len(results) - 1; i >= 0; i-- {
		fmt.Println(results[i])
	}
}
