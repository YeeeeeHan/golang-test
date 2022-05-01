package utils

import (
	"TechnicalAssignment/pkg/constants"
	"fmt"
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestParseInput_Commands(t *testing.T) {
	registerInput := "register abc def"
	loginInput := "login abc"
	depositInput := "deposit 123"
	withdrawInput := "withdraw 123"
	sendInput := "send abc 123"
	balanceInput := "balance"
	logoutInput := "logout"
	accountsInput := "accounts"

	loginArgs := []string{"abc"}
	depositArgs := []string{"123"}
	withdrawArgs := []string{"123"}
	sendArgs := []string{"abc", "123"}
	emptyArgs := []string{}

	command0, args0 := ParseInput(registerInput)
	assert.Equal(t, constants.Register, command0, fmt.Sprintf("Parsed command is not %s", constants.Register))
	assert.Equal(t, loginArgs, args0, fmt.Sprintf("Parsed args is not %+v", loginArgs))

	command1, args1 := ParseInput(loginInput)
	assert.Equal(t, constants.Login, command1, fmt.Sprintf("Parsed command is not %s", constants.Login))
	assert.Equal(t, loginArgs, args1, fmt.Sprintf("Parsed args is not %+v", loginArgs))

	command2, args2 := ParseInput(depositInput)
	assert.Equal(t, constants.Deposit, command2, fmt.Sprintf("Parsed command is not %s", constants.Deposit))
	assert.Equal(t, depositArgs, args2, fmt.Sprintf("Parsed args is not %+v", depositArgs))

	command3, args3 := ParseInput(withdrawInput)
	assert.Equal(t, constants.Withdraw, command3, fmt.Sprintf("Parsed command is not %s", constants.Withdraw))
	assert.Equal(t, withdrawArgs, args3, fmt.Sprintf("Parsed args is not %+v", withdrawArgs))

	command4, args4 := ParseInput(sendInput)
	assert.Equal(t, constants.Send, command4, fmt.Sprintf("Parsed command is not %s", constants.Send))
	assert.Equal(t, sendArgs, args4, fmt.Sprintf("Parsed args is not %+v", sendArgs))

	command5, args5 := ParseInput(balanceInput)
	assert.Equal(t, constants.Balance, command5, fmt.Sprintf("Parsed command is not %s", constants.Balance))
	assert.Equal(t, emptyArgs, args5, fmt.Sprintf("Parsed args is not %+v", emptyArgs))

	command6, args6 := ParseInput(logoutInput)
	assert.Equal(t, constants.Logout, command6, fmt.Sprintf("Parsed command is not %s", constants.Logout))
	assert.Equal(t, emptyArgs, args6, fmt.Sprintf("Parsed args is not %+v", emptyArgs))

	command7, args7 := ParseInput(accountsInput)
	assert.Equal(t, constants.Accounts, command7, fmt.Sprintf("Parsed command is not %s", constants.Accounts))
	assert.Equal(t, emptyArgs, args7, fmt.Sprintf("Parsed args is not %+v", emptyArgs))
}

func TestParseInput_MultipleSpaces(t *testing.T) {
	withdrawInput := "withdraw   123"
	sendInput := " send abc   123"
	balanceInput := "   balance  "

	withdrawArgs := []string{"123"}
	sendArgs := []string{"abc", "123"}
	emptyArgs := []string{}

	command3, args3 := ParseInput(withdrawInput)
	assert.Equal(t, constants.Withdraw, command3, fmt.Sprintf("Parsed command is not %s", constants.Withdraw))
	assert.Equal(t, withdrawArgs, args3, fmt.Sprintf("Parsed args is not %+v", withdrawArgs))

	command4, args4 := ParseInput(sendInput)
	assert.Equal(t, constants.Send, command4, fmt.Sprintf("Parsed command is not %s", constants.Send))
	assert.Equal(t, sendArgs, args4, fmt.Sprintf("Parsed args is not %+v", sendArgs))

	command5, args5 := ParseInput(balanceInput)
	assert.Equal(t, constants.Balance, command5, fmt.Sprintf("Parsed command is not %s", constants.Balance))
	assert.Equal(t, emptyArgs, args5, fmt.Sprintf("Parsed args is not %+v", emptyArgs))
}

//func TestParseInput_InvalidCommands(t *testing.T) {
//	withdrawInput := "withdra   123"
//	sendInput := " sen abc   123"
//	balanceInput := "   balanc  "
//
//
//	command3, _ := ParseInput(withdrawInput)
//	assert.Equal(t, constants.InvalidCommandMsg, command3, fmt.Sprintf("Parsed response is not %s", constants.Withdraw))
//
//	command4, _ := ParseInput(sendInput)
//	assert.Equal(t, constants.Send, command4, fmt.Sprintf("Parsed command is not %s", constants.Send))
//
//	command5, _ := ParseInput(balanceInput)
//	assert.Equal(t, constants.Balance, command5, fmt.Sprintf("Parsed command is not %s", constants.Balance))
//}
