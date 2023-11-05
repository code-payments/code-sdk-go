# Code Wallet Golang SDK

The Code Wallet Golang SDK is a module that allows Go developers to integrate Code into their applications. Seamlessly start accepting payments with minimal setup and just a few lines of code.

See the [documentation](https://code-wallet.github.io/code-sdk/docs/guide/introduction.html) for more details.

## What is Code?

[Code](https://getcode.com) is a mobile wallet app leveraging self-custodial blockchain technology to provide an instant, global, and private payments experience.

## Installation

You can install the Code Wallet Golang SDK using to Go toolset:

```bash
go get github.com/code-wallet/code-sdk-go
```

## Usage
Here's a simple example showcasing how to create a payment intent using the Golang SDK:

```go
package main

import (
	"context"
	"log"

	codesdk "github.com/code-wallet/code-sdk-go"
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
		codesdk.USD,
		0.05,
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

If you have any questions or need help integrating Code into your website or application, please reach out to us on [Discord](https://discord.gg/DunN9aNS) or [Twitter](https://twitter.com/getcode).

##  Contributing

For now the best way to contribute is to share feedback on [Discord](https://discord.gg/DunN9aNS). This will evolve as we continue to build out the platform and open up more ways to contribute.
