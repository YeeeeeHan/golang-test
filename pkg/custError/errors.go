package custError

import (
	"TechnicalAssignment/pkg/constants"
	"errors"
)

var InternalDBError = errors.New(constants.InternalDBErrMsg)
var InvalidNumArguments = errors.New(constants.InvalidNumArgumentsMsg)
var AccountExistsError = errors.New(constants.AccountExistsMsg)
var WrongCredentialsError = errors.New(constants.WrongCredentialsMsg)
var PermissionError = errors.New(constants.PermissionErrorMsg)
