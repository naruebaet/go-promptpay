package pp

import (
	"fmt"
	"strings"
)

// AccountType represents the type of PromptPay account
type AccountType string

// Error types for PromptPay
type Error string

func (e Error) Error() string { return string(e) }

const (
	// EMV QR Code version
	Version = "000201"
	// QR Type (11 = Static)
	QRType = "010211"
	// Merchant Account Information with length
	MerchantAccountInfo = "2937"
	// Application ID for PromptPay
	MerchantAccountInfoAID = "0016A000000677010111"
	// Merchant Account Info prefix for phone numbers (includes Thai country code)
	MerchantAccountInfoForPhone = "01130066"
	// Merchant Account Info prefix for national ID
	MerchantAccountInfoForID = "0213"
	// Country code (TH)
	Country = "5802TH"
	// Currency code (764 for THB)
	Currency = "5303764"
	// Currency symbol (THB)
	CurrencySymbol = "54"
	// Checksum template
	CheckSum = "6304"

	// AccountTypePhone represents a phone number account
	AccountTypePhone AccountType = "phone"
	// AccountTypeID represents a national ID account
	AccountTypeID AccountType = "id"

	// Error constants
	ErrInvalidAccountType Error = "invalid account type, use 'phone' or 'id'"
	ErrInvalidIDLength    Error = "invalid ID length, must be 13 digits"
	ErrInvalidIDFormat    Error = "invalid ID format, must be digits only"
	ErrInvalidIDStart     Error = "invalid ID, cannot start with 0"
)

// GenPromptpay generates a QR code string for PromptPay based on the provided parameters.
// It accepts an AccountType and an account number (phone number or national ID).
// Returns the generated QR code string or an error if validation fails.
func GenPromptpay(accountType AccountType, accountNumber string) (string, error) {
	accountInfo, err := validateAndFormatAccount(accountType, accountNumber)
	if err != nil {
		return "", err
	}

	raw := Version + QRType + MerchantAccountInfo + MerchantAccountInfoAID + accountInfo + Country + Currency + CheckSum
	return raw + CRC16XMODEM(raw), nil
}

func GenPromptpayWithAmount(accountType AccountType, accountNumber string, amount float64) (string, error) {
	accountInfo, err := validateAndFormatAccount(accountType, accountNumber)
	if err != nil {
		return "", err
	}

	raw := Version + QRType + MerchantAccountInfo + MerchantAccountInfoAID + accountInfo + Country + Currency +
		formatAmount(amount) + CheckSum
	return raw + CRC16XMODEM(raw), nil
}

// validateAndFormatAccount handles common account validation and formatting
func validateAndFormatAccount(accountType AccountType, accountNumber string) (string, error) {
	if accountType != AccountTypePhone && accountType != AccountTypeID {
		return "", ErrInvalidAccountType
	}

	if accountType == AccountTypePhone {
		return MerchantAccountInfoForPhone + normalizePhoneNumber(accountNumber), nil
	}

	id, err := validateThaiID(accountNumber)
	if err != nil {
		return "", err
	}
	return MerchantAccountInfoForID + id, nil
}

// formatAmount formats the amount for QR code generation
func formatAmount(amount float64) string {
	amountStr := fmt.Sprintf("%.2f", amount)
	amountLength := fmt.Sprintf("%02d", len(amountStr))
	return CurrencySymbol + amountLength + amountStr
}

// normalizePhoneNumber removes country code prefixes and leading zeros from phone numbers
func normalizePhoneNumber(phone string) string {
	phone = strings.TrimPrefix(phone, "+66")
	phone = strings.TrimPrefix(phone, "66")
	phone = strings.TrimPrefix(phone, "+660")
	phone = strings.TrimPrefix(phone, "660")
	return strings.TrimPrefix(phone, "0")
}

// validateThaiID validates and formats Thai national ID numbers
func validateThaiID(id string) (string, error) {
	if len(id) != 13 {
		return "", ErrInvalidIDLength
	}

	if id[0] == '0' {
		return "", ErrInvalidIDStart
	}

	for _, char := range id {
		if char < '0' || char > '9' {
			return "", ErrInvalidIDFormat
		}
	}

	return id, nil
}

// Find the checksum for the QR code using CRC-16 with static polynomial 0x1021 (XMODEM) and initial value 0xFFFF
// CRC16XMODEM calculates the CRC-16/XMODEM checksum for the given string.
func CRC16XMODEM(data string) string {
	crc := uint16(0xFFFF)
	for _, c := range data {
		crc ^= uint16(c) << 8
		for i := 0; i < 8; i++ {
			if crc&0x8000 != 0 {
				crc = (crc << 1) ^ 0x1021
			} else {
				crc <<= 1
			}
		}
	}
	return strings.ToUpper(fmt.Sprintf("%04X", crc&0xFFFF))
}
