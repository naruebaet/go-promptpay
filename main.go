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
}
