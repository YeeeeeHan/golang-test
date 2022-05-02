package constants

// Commands
const Register = "register"
const Login = "login"
const Deposit = "deposit"
const Withdraw = "withdraw"
const Send = "send"
const Balance = "balance"
const Logout = "logout"
const Accounts = "accounts"

// Error Messages
const InternalDBErrMsg = "internal error with database"
const InvalidCommandMsg = "command not recognised"
const InvalidNumArgumentsMsg = "invalid number of arguments"
const InvalidArgumentsMsg = "invalid argument format"
const AccountExistsMsg = "username already exists"
const AccountsDoesNotExistMsg = "username does not exist"
const WrongCredentialsMsg = "wrong username or password"
const NotLoggedInMsg = "please log in first"
const PermissionErrorMsg = "you do not have the permissions"
const InsufficientFundsMsg = "account has insufficient funds"
const NegativeValueMsg = "value must be positive"

// User
const Unregistered = "NIL"

// Files
const UsernameFile = "cmd/db/username.txt"
const BalanceFile = "cmd/db/balance.db"
const PasswordFile = "cmd/db/password.db"

const UsernameFileTest = "../../cmd/db/username_test.txt"
const BalanceFileTest = "../../cmd/db/balance_test.db"
const PasswordFileTest = "../../cmd/db/password_test.db"
