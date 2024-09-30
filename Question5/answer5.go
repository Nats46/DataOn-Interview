package answer5

import (
	"bufio"
	"fmt"
	"os"
)

func LenString() {
	reader := bufio.NewReader(os.Stdin)
	fmt.Print("please input string: ")
	input, _ := reader.ReadString('\n')
	fmt.Println(len(input))
}