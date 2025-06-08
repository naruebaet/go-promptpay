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
    fmt.Println(phoneStr)

    // Generate PromptPay string for national ID
    idStr, err := pp.GenPromptpay(pp.AccountTypeID, "1234567890123")
    if err != nil {
        fmt.Println("Error:", err)
        return
    }
    fmt.Println(idStr)
}
```

## Account Types

The library supports two types of PromptPay identifiers:

- `pp.AccountTypePhone`: For Thai mobile phone numbers (10 digits)
- `pp.AccountTypeID`: For Thai national ID numbers (13 digits)

## Features

- Generate PromptPay strings for phone numbers and national IDs
- EMV Co QR Code specification compliant
- Simple and easy to use API

## Notes

- Phone numbers should be in the format "0XXXXXXXXX" (10 digits)
- National IDs must be 13 digits
- The generated string is EMVCo QR Code compliant and can be used with any QR code generator

## License

MIT License
