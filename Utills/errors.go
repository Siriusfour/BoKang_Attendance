package Utills

import (
	"fmt"
	"github.com/pkg/errors"
)

func AppendErr(existing error, newErr error) error {
	if existing == nil {
		return newErr
	}
	return fmt.Errorf("%v; %w", existing, newErr)
}

const (
	Binding_Data_is_Failed = "binding data is failed:"
	Insect_is_Failed       = "insect is failed"
	Access_Token_is_failed = "access token is failed"
	Query_is_failed        = "query userinfo is failed"

	AccessTokenIsInvalid = 10001
	RefreshTokenIsValid  = 10002
	QueryIsFailed        = 10003
)

var (
	ErrIsATokenIsInvalid = errors.New("access token is invalid")
	ErrIsRTokenIsInvalid = errors.New("refresh token is invalid")
)
