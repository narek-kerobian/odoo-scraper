package service

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

// Prompts for cli action
func CliPrompt(label string, defValue string) string {
    fmt.Print(label)
	reader := bufio.NewReader(os.Stdin)

	// ReadString will block until the delimiter is entered
	input, err := reader.ReadString('\n')
	if err != nil {
		panic(err)
	}

	// remove the delimeter from the string
	input = strings.TrimSuffix(input, "\n")
    if input == "" {
        input = defValue
    }

    return input
}

func CliSelect(label string, options []string, defIndex int) (key int) {
    fmt.Println(label)

    for i, v := range(options) {
        fmt.Printf("%d. %s \n", i+1, v)
    }

    _, err := fmt.Scanf("%d", &key)
    if err != nil {
        panic(err)
    }

    return
}

