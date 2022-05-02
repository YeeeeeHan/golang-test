package custError

import (
	"TechnicalAssignment/pkg/constants"
	"errors"
)

var InternalDBError = errors.New(constants.InternalDBErrMsg)
var InvalidNumArguments = errors.New(constants.InvalidNumArgumentsMsg)
var InvalidArguments = errors.New(constants.InvalidArgumentsMsg)
var AccountAlreadyExistsError = errors.New(constants.AccountExistsMsg)
var AccountsDoesNotExistError = errors.New(constants.AccountsDoesNotExistMsg)
var WrongCredentialsError = errors.New(constants.WrongCredentialsMsg)
var PermissionError = errors.New(constants.PermissionErrorMsg)
var InsufficientFunds = errors.New(constants.InsufficientFundsMsg)
var NegativeValueError = errors.New(constants.NegativeValueMsg)
