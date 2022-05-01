package custError

import (
	"TechnicalAssignment/pkg/constants"
	"errors"
)

var InternalDBError = errors.New(constants.InternalDBErrMsg)
var InvalidNumArguments = errors.New(constants.InvalidNumArgumentsMsg)
var AccountExistsError = errors.New(constants.AccountExistsMsg)
