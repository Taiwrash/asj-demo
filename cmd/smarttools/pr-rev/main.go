// cmd/smarttools/prreview/main.go
package main

import (
	"fmt"

	"github.com/taiwrash/asj-demo/pkg/bedrock"
)

func main() {
	client := bedrock.New("us-east-1")

	prDesc := `
	Adds new S3 bucket for storing user uploads.
	Bucket is public for testing.
	`

	prompt := fmt.Sprintf(`Review this PR for DevOps best practices:
	%s

	Highlight:
	- Security concerns
	- IaC quality
	- Suggested improvements`, prDesc)

	fmt.Println(client.Ask(prompt, 400))
}
