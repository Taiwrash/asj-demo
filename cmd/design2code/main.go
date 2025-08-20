package main

import (
	"fmt"

	"github.com/taiwrash/asj-demo/pkg/bedrock"
)

func main() {
	client := bedrock.New("us-east-1")

	design := `
	We need a VPC with 2 private subnets, 
	an RDS PostgreSQL instance, and an ALB.
	`

	prompt := fmt.Sprintf(`Convert this design into Terraform HCL.
	Only return valid Terraform code:

	%s`, design)

	fmt.Println(client.Ask(prompt, 400))
}
