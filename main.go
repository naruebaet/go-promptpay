package main

import (
	"fmt"

	"github.com/naruebaet/go-promptpay/pp"
)

func main() {
	// for phone number
	str, _ := pp.GenPromptpay(pp.AccountTypePhone, "990844901")
	fmt.Println(str)

	// for ID
	str, _ = pp.GenPromptpay(pp.AccountTypeID, "1234567890123")
	fmt.Println(str)

	// for phone number with amount
	str, _ = pp.GenPromptpayWithAmount(pp.AccountTypePhone, "990844901", 100.00)
	fmt.Println(str)

	// for ID with amount
	str, _ = pp.GenPromptpayWithAmount(pp.AccountTypeID, "1234567890123", 5.12)
	fmt.Println(str)
}
