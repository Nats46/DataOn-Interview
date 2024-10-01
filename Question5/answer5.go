package answer5

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

func LenString() {
	reader := bufio.NewReader(os.Stdin)
	fmt.Print("please input string: ")
	input, _ := reader.ReadString('\n')
	input = strings.TrimSpace(input)
	fmt.Println(len(input))
}