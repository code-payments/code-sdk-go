![Code Golang SDK](https://repository-images.githubusercontent.com/714496320/d9fde93e-cefa-4276-8542-34befdeaa983)

# Code SDK - Go

[![Release](https://img.shields.io/github/v/release/code-payments/code-sdk-go.svg)](https://github.com/code-payments/code-sdk-go/releases/latest)
[![PkgGoDev](https://pkg.go.dev/badge/github.com/code-payments/code-sdk-go)](https://pkg.go.dev/github.com/code-payments/code-sdk-go/sdk)
[![Tests](https://github.com/code-payments/code-sdk-go/actions/workflows/test.yml/badge.svg)](https://github.com/code-payments/code-sdk-go/actions/workflows/test.yml)
[![GitHub License](https://img.shields.io/badge/license-MIT-lightgrey.svg?style=flat)](https://github.com/code-payments/code-sdk-go/blob/main/LICENSE.md)

The Code Golang SDK is a module that allows Go developers to integrate Code into their applications. Seamlessly start accepting payments with minimal setup and just a few lines of code.

See the [documentation](https://code-payments.github.io/code-sdk/docs/guide/introduction.html) for more details.

## What is Code?

[Code](https://getcode.com) is a mobile wallet app leveraging self-custodial blockchain technology to provide an instant, global, and private payments experience.

## Installation

You can install the Code Golang SDK using to Go toolset:

```bash
go get github.com/code-payments/code-sdk-go
```

## Usage
Here's a simple example showcasing how to create a payment intent using the Golang SDK:

```go
package main

import (
	"context"
	"log"

	codesdk "github.com/code-payments/code-sdk-go/sdk"
)

func check(err error) {
	if err != nil {
		log.Fatal(err)
	}
}

func main() {
	ctx := context.Background()

	// Setup the Code web client
	client := codesdk.NewWebClient()

	// Specify payment request details
	intent, err := codesdk.NewPaymentRequestIntent(
		// Or the string "usd"
		codesdk.USD,
		// Minimum amount is $0.05 USD
		0.05,
		// Code Deposit Address or any Kin token account
		"E8otxw1CVX9bfyddKu3ZB3BVLa4VVF9J7CTPdnUwT9jR",
	)
	check(err)

	// Create a payment request intent
	_, err = client.CreatePaymentRequest(ctx, intent)
	check(err)

	// Check the intent status
	_, err = client.GetIntentStatus(ctx, intent.GetIntentId())
	check(err)
}
```

## Getting Help

If you have any questions or need help integrating Code into your website or application, please reach out to us on [Discord](https://discord.gg/T8Tpj8DBFp) or [Twitter](https://twitter.com/getcode).

##  Contributing

For now the best way to contribute is to share feedback on [Discord](https://discord.gg/T8Tpj8DBFp). This will evolve as we continue to build out the platform and open up more ways to contribute.
