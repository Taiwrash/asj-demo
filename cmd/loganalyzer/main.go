package main

import (
	"fmt"

	"github.com/taiwrash/asj-demo/pkg/bedrock"
)

func main() {
	client := bedrock.New("us-east-1")

	logs := `
	[INFO] serviceA started
	[WARN] serviceB high memory
	[ERROR] serviceC panic: nil pointer
	[INFO] serviceA completed
	`

	prompt := fmt.Sprintf(`Summarize the following logs into:
	1. Key issues
	2. Error counts
	3. Recommended next actions

	%s`, logs)

	fmt.Println(client.Ask(prompt, 400))
}
