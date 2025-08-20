# AWS Bedrock DevOps Demo (AWS Summit Johannesburg)

This demo showcases using AWS Bedrock with Claude to analyze infrastructure code and provide DevOps best practices recommendations.

- Session slides: [Go + Amazon Bedrock for DevOps](https://docs.google.com/presentation/d/1l0V1Oa5rZkj7KMWENFTSVjmnA_6wAxFMU45fTxwvEjQ/edit?usp=sharing)

## Prerequisites

- Go 1.21+
- AWS Account with Bedrock access
- AWS CLI v2
- macOS (for this guide)

## Quick Start

```bash
# Clone the repository
git clone https://github.com/taiwrash/asj-demo.git
cd asj-demo

# Set up AWS credentials
aws configure

# Install dependencies
go mod tidy

# Run the demo
go run cmd/smarttools/pr-rev/main.go -name "S3 Config" -desc "Adding S3 bucket with public access"
```

## Setup Guide

### 1. AWS Configuration

Configure your AWS credentials:

```bash
aws configure
```

Enter your details:
```
AWS Access Key ID: [Your Access Key]
AWS Secret Access Key: [Your Secret Key]
Default region name: us-east-1
Default output format: json
```

### 2. Enable Bedrock Model Access

1. Open [AWS Bedrock Console](https://console.aws.amazon.com/bedrock)
2. Go to "Model access"
3. Enable "Anthropic - Claude"
4. Click "Save changes"

### 3. Environment Variables

```bash
export AWS_REGION=us-east-1
export AWS_ACCESS_KEY_ID=your_access_key
export AWS_SECRET_ACCESS_KEY=your_secret_key
```

## Usage

Review infrastructure changes:

```bash
# Basic review
go run cmd/smarttools/pr-rev/main.go -name "VPC Setup" 

# Review with custom description
go run cmd/smarttools/pr-rev/main.go -name "S3 Config" -desc "Adding secure S3 bucket"
```

## Project Structure

```
asj-demo/
├── cmd/
│   └── smarttools/
│       └── pr-rev/
│           └── main.go      # CLI entry point
├── pkg/
│   └── bedrock/
│       └── client.go        # Bedrock client implementation
└── README.md
```

## Common Issues

### Model Access Error
```
ValidationException: You don't have access to the model
```
**Solution**: 
- Verify model access in Bedrock console
- Check AWS region configuration
- Validate IAM permissions

### Authentication Errors
```
InvalidSignatureException
```
**Solution**:
- Run `aws configure list`
- Check environment variables
- Verify AWS credentials file

## Development

```bash
# Run tests
go test ./...

# Build binary
go build -o pr-review cmd/smarttools/pr-rev/main.go

# Run binary
./pr-review -name "Test PR" -desc "Test description"
```

## License

MIT License - See LICENSE file

## Author

Rasheed Taiwo (@taiwrash)

---

**Note**: Replace placeholder credentials and values with your actual configuration