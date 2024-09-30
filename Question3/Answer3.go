package answer3

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

func EmailString(){
	reader := bufio.NewReader(os.Stdin)
	fmt.Print("please input email: ")
	email, _ := reader.ReadString('\n')
	change := strings.Replace(email, "@", " ", -1)
	fmt.Println(change)
}
