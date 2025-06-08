# Go PromptPay Generator

A Go library for generating PromptPay payment strings for Thailand's PromptPay system.

## Installation

```bash
go get github.com/naruebaet/go-promptpay
```

## Usage

```go
package main

import (
    "fmt"
    "github.com/naruebaet/go-promptpay/pp"
)

func main() {
    // Generate PromptPay string for phone number
    phoneStr, err := pp.GenPromptpay(pp.AccountTypePhone, "0901234567")
    if err != nil {
        fmt.Println("Error:", err)
        return
    }
    fmt.Println("Basic PromptPay:", phoneStr)

    // Generate PromptPay string with amount
    phoneWithAmount, err := pp.GenPromptpayWithAmount(pp.AccountTypePhone, "0901234567", 1000.50)
    if err != nil {
        fmt.Println("Error:", err)
        return
    }
    fmt.Println("PromptPay with amount:", phoneWithAmount)

    // Generate PromptPay string for national ID with amount
    idWithAmount, err := pp.GenPromptpayWithAmount(pp.AccountTypeID, "1234567890123", 500.75)
    if err != nil {
        fmt.Println("Error:", err)
        return
    }
    fmt.Println("ID with amount:", idWithAmount)
}
```

## Features

- Generate PromptPay strings for phone numbers and national IDs
- Support for payment amounts
- EMV Co QR Code specification compliant
- Simple and easy to use API

## Account Types

The library supports two types of PromptPay identifiers:

- `pp.AccountTypePhone`: For Thai mobile phone numbers (10 digits)
- `pp.AccountTypeID`: For Thai national ID numbers (13 digits)

## Amount Support

The library supports generating PromptPay QR codes with specified amounts:

- Use `GenPromptpayWithAmount` to include a payment amount
- Amounts are specified in Thai Baht (THB)
- Supports decimal values up to 2 decimal places
- Generate fixed-amount QR codes for specific payments

## Notes

- Phone numbers should be in the format "0XXXXXXXXX" (10 digits)
- National IDs must be 13 digits
- The generated string is EMVCo QR Code compliant and can be used with any QR code generator

## License

MIT License
