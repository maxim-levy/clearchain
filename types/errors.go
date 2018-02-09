package types

import (
	"fmt"

	sdk "github.com/cosmos/cosmos-sdk/types"
)

// List of error codes generated by the application.
// Base SDK reserves 0 ~ 99, app's internal error codes
// start from 1000.
const (
	CodeInvalidAmount      sdk.CodeType = 1000
	CodeInvalidAddress     sdk.CodeType = 1001
	CodeInvalidPubKey      sdk.CodeType = 1002
	CodeInvalidAccount     sdk.CodeType = 1003
	CodeWrongSigner        sdk.CodeType = 1004
	CodeWrongMessageFormat sdk.CodeType = 1005
)

func ErrInvalidAmount(typ string) sdk.Error {
	return sdk.NewError(CodeInvalidAmount, fmt.Sprintf("invalid amount: %s", typ))
}

func ErrInvalidAddress(typ string) sdk.Error {
	return sdk.NewError(CodeInvalidAddress, fmt.Sprintf("invalid address: %s", typ))
}

func ErrInvalidPubKey(typ string) sdk.Error {
	return sdk.NewError(CodeInvalidPubKey, fmt.Sprintf("invalid pub key: %s", typ))
}

func ErrInvalidAccount(typ string) sdk.Error {
	return sdk.NewError(CodeInvalidAccount, fmt.Sprintf("invalid account: %s", typ))
}

func ErrWrongSigner(typ string) sdk.Error {
	return sdk.NewError(CodeWrongSigner, fmt.Sprintf("wrong signer: %s", typ))
}

func ErrWrongMsgFormat(typ string) sdk.Error {
	return sdk.NewError(CodeWrongMessageFormat, fmt.Sprintf("wrong message format: %s", typ))
}

func ErrUnauthorized(typ string) sdk.Error {
	return sdk.NewError(sdk.CodeUnauthorized, fmt.Sprintf("unauthorized: %s", typ))
}
