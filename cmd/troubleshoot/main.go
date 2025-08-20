package main

import (
	"fmt"

	"github.com/taiwrash/asj-demo/pkg/bedrock"
)

func main() {
	client := bedrock.New("us-east-1")

	logData := `
	2025-08-17 12:34:56 ERROR Failed to connect to database at db-prod-1:5432
	timeout expired
	`

	prompt := fmt.Sprintf(`You are a DevOps assistant. 
	Analyze the following log and suggest possible causes and fixes:

	%s`, logData)

	fmt.Println(client.Ask(prompt, 300))
}
