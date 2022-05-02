# Coding Test

In this coding test, I have opted to use Golang as my preferred language.


1. [Time Allocation]()
2. [Project Structure]()
3. [Project Walk Through]()
4. [Architecture]()
5. [Project Design Considerations]()
6. [Reflection and Areas of Improvement]()


# Time Allocation*

1. Planning --- 40 mins
2. Draft out starter code --- 2 hours
3. Implementation --- 2 hours
4. Testing and writing tests --- 3 hours
5. Documentation  --- 2 hours 30 min

*Time spent is not calculated in 1 seating, and is a very close approximation via stopwatch

# Project Structure

```bash
├── README.md
├── TechnicalAssignment
├── cmd
 	├── db
 	└── server
├── main.go
└── pkg
	├── constants
	├── custError
	├── service
	└── utils
```

Explanation of packages:
- **main.go**: Main entry point for the application
- **cmd**: Contains packages that are used internally in this application
- **db**: Contains logic related to CRUD operations to the database, as well as DB files
- **server**: Contains blocking logic to listen for user input
- **pkg**: Contains packages with standalone logic
- **constants**: Contains all the constants in the project, from error messages to file paths
- **custError**: Custom errors that are used throughout the project
- **service**: Logic to handle individual user commands
- **utils**: Contains utility logic that can be reused throughout the package

# Project Walk Through

1. Run  `./TechnicalAssignment `. A list of command instructions and their description will be shown. The bottom left corner shows a username corresponding to the current user session, which is currently `NIL`.
   ![landing page](https://i.ibb.co/23ZdG3K/Screen-Shot-2022-05-02-at-4-32-02-PM.png)

2. Enter `register testreviewer password123`. This creates a new account with username `testreviewer` and password `password123`, subsequently logging into the newly created account.
   ![landing page](https://i.ibb.co/7VMvYCf/Screen-Shot-2022-05-02-at-4-45-11-PM.png)
3. Enter `deposit 12` to deposit money into the account.
   ![landing page](https://i.ibb.co/SJGRbSP/Screen-Shot-2022-05-02-at-4-47-42-PM.png)
4. Enter `withdraw 2` to withdraw money from the account.
   ![landing page](https://i.ibb.co/JrPnmMD/Screen-Shot-2022-05-02-at-4-49-36-PM.png)
5. Enter `balance` to view the current balance in the account.
   ![landing page](https://i.ibb.co/Swh31ng/Screen-Shot-2022-05-02-at-4-50-44-PM.png)
6. Enter `send admin 7` to send money to another account, in this case `admin`.
   ![landing page](https://i.ibb.co/5FRqyBF/Screen-Shot-2022-05-02-at-4-52-02-PM.png)
7. Enter `logout` to log out from the account. Notice the username shown is now `NIL`
   ![landing page](https://i.ibb.co/hd2HTLv/Screen-Shot-2022-05-02-at-4-53-30-PM.png)
8. Enter `login admin password123`. The admin account is created upon project compilation and has special permissions.
   ![landing page](https://i.ibb.co/MpcW4Nc/Screen-Shot-2022-05-02-at-4-54-55-PM.png)

9. Enter `accounts` to keep track of the balances in all the accounts.
   ![landing page](https://i.ibb.co/zJxtNJY/Screen-Shot-2022-05-02-at-4-55-46-PM.png )

# Architecture
![](https://i.ibb.co/YchnzSb/Screen-Shot-2022-05-02-at-5-10-43-PM.png )

1. Server listens to user input from `stdin` via a `for` loop, in `server.ListenAndServe()`
2. User's input is parsed with `utils.ParseInput()` to determine the command
3. According to the parsed command, handlers in `service` would be called to service the user's request
4. The handlers might make CRUD operations to 3 database files, `username.txt`, `balance.db`, `password.db`


# Project Design Considerations

### Server
A server implemented with a for-loop is used because I felt that the optimal user flow would be a synchronous exchange of instructions, where the user inputs a command, waits for a reply, before executing the next command. It is also the most straight-forward implementation I could think of, mimicking the behaviour of always-on hosted servers continuously listening on a connection.

User input is read from `os.Stdin` and it is parsed with `utils.ParseInput()`. Command parsed is then fed into a `switch` case to be handled by the `service` package.

![](https://i.ibb.co/D1sjCgd/Screen-Shot-2022-05-02-at-5-51-36-PM.png )


### Database
I was debating between implementing a persistent datastore versus an in-memory datastore, because I felt that a persistent datastore would make the application too complicated. However,  without a persistent datastore, the reviewer will have to keep `registering` accounts to test the logic and thus I decided to use a simple key-value persistent datastore.

I used a public package `github.com/rapidloop/skv` to implement the persistent datastore. It writes an encoded key-value pair a `.db` file, namely  `balance.db` and `password.db`. This `skv` package only support `get` and `put` operations, and in order to `get all usernames`  in the `Accounts()` handler, a plaintext file of `username.txt` is used to store all the usernames.

### User Authentication
In order to effectively emulate a wallet application where a unified interface serves all users, user authentication must be implemented. This is handled by the `passsord.db` table and the `Register()` and `Login()` handlers.

The `password.db` table stores a mapping of usernames and passwords that are created upon calling `Register()` command. User authentication happens upon calling the `Login()` command, which checks the password the user passed against those found in the `password.db` table.

### User Sessions
To effectively keep track of which user is operating the application, user sessions are used. A user session is a `Wallet` struct that has a `Username` field.
![](https://i.ibb.co/ftstYQg/Screen-Shot-2022-05-02-at-5-43-02-PM.png )

The `sessionUser` is instantiated in `main.go` and its reference is passed into `server.ListenAndServe()`.

![](https://i.ibb.co/1LtLJyM/Screen-Shot-2022-05-02-at-5-44-30-PM.png )

To log in in or log out, pointer to the `sessionUser` is modified by assigning a `username`  or an empty struct respectively. If the `Username` field is not set, `GetUsername()` returns `NIL`, which is displayed on the terminal display.

Once `Username` is set, subsequent commands will be effective upon the binded `Username` in the `sessionUser(aka wallet)` pointer.
![](https://i.ibb.co/XStJz47/Screen-Shot-2022-05-02-at-5-46-17-PM.png )
![](https://i.ibb.co/FnZ1WR4/Screen-Shot-2022-05-02-at-5-56-09-PM.png )

### Parsing user input
The simplest way for user to input instructions via terminal would be to enter a white-space separated string, according to instructions displayed upon starting the programme. Hence, `utils.ParseInput()` is used to parse the input.

It does the following:
1. return if input is an empty string
2. splits input around each instance of a word with `string.Fields()`
3. de-capitalise input for standardisation
4.  returns (command, arguments)

### User Logical Flow
![](https://i.ibb.co/Js0gqz9/Screen-Shot-2022-05-02-at-6-16-42-PM.png )

1. If command is `Register()`, register user and log in  --- Logging in user when they register is a logical userflow.
2. If no user currently logged in `and` command is not `Login()`, prompt user to log in --- This is to ensure users login first and foremost before any commands.
3. Else, serve command normally for logged-in user

### Handlers
Here are the main considerations for handlers in  `service.go`:
1. Take in pointer to current `sessionUser` for user details (wallet)
2. Take in `args` if necessary
3. Return any errors, whose messages are displayed to the user via `fmt.Println()`
   ![](https://i.ibb.co/ZT3pkPS/Screen-Shot-2022-05-02-at-6-25-25-PM.png )

_Register()_
1.  Checks if `len(args) == 2`, to ensure username and password is entered
2.  Calls `db.CreateUser` to persist information to DB

_Login()_
1.  Checks if `len(args) == 2`, to ensure username and password is entered
2. Queries `password.db` table to determine is username exists
3. If so, checks the password returned matches user input
4. Logs user in

_Deposit()_
1.  Checks if `len(args) == 1`, to ensure deposit value is entered
2. Calls `topUp()` helper function  _(To be elaborated below)_

_Withdraw()_
1.  Checks if `len(args) == 1`, to ensure withdraw value is entered
2. Calls `drawDown()` helper function  _(To be elaborated below)_

_Send()_
1.  Checks if `len(args) == 2`, to ensure destination account and value is entered
2. Calls `drawDown()` on current user, and `topUp()` on destination account

_Balance()_
1. Sets userSession pointer to point to an empty `Wallet` struct.

_Accounts()_
1. Checks if account is `admin` account
2. Reads `username.txt` file to obtain list of usernames
3. Queries `balance.db` for each username and print it out

### service.topUp() and service.drawDown()
These couple of helper functions in the `service` package exists to extract the repeated logic of:
1. Adding or subtracting `x` amount from an account
2. Performing sanity checks --- e.g. Ensure no negative values, ensure funds > amount to be subtracted

This prevents repeated code, more specifically in the `Send()`, `Withdraw()`, and `Deposit()` function, because sending money from A to B is essentially withdrawing money from A and depositing money to B.

### Constants
`pkg/constants` contains all constants that are used in the project. This provides a source of truth for all constants that are depended upon throughout the project, and one only needs to edit them in this file for the changes to be propagated throughout.

These are the main type of constants:
1. Commands --- e.g. `"register"`, `"withdraw"`
2. Error messages --- e.g.  `"username already exists"`,  `account has insufficient funds`
3. FilePaths

### Error Handling
`pkg/custError` contains all the custom errors that are unique to this project. There are 3 benefits to creating custom errors:
1. Custom errors can be re-used throughout the project, following the Don't Repeat Yourself (DRY) philosophy.
2. Custom errors can have custom messages that are directly shown to users, and no extra parsing of errors need to be done on the server side.
3. Custom errors make testing way easier and cleaner by providing an exhaustive list of negative outcomes that can happen.


### Testing
Tests in this project follow 2 structures:

1. `TestMain()` provides more high order functionality such a connecting to a test database and creating files before the tests, and deleting test database and files after the tests.

   ![](https://i.ibb.co/KDpcB6t/Screen-Shot-2022-05-02-at-6-52-08-PM.png )
2. `TestExample()` is the meat of the testing logic, consisting of a `testTable` slice of custom test objects. Each test object specifies the test name, required inputs, and desired outputs. A for loop is used to loop over the range of `testTable` slice, and calls a specific function to be tested (`functionToBeTest` in this case), asserting that the input and output fields of each `testTable` object (`tt`) are equal.
   ![](https://i.ibb.co/XJtD2Nf/Screen-Shot-2022-05-02-at-6-50-40-PM.png )

There are 3 main areas that needs to be tested:
1. Parsing of user input --- `utils.ParseInput()`
2. Individual handlers in `service.go` --- `Login()`, `Register()`, etc
3. User flow --- e.g. User can only withdraw after logging in

_parse_test.go()_
1. Ensure empty inputs are handled
2. Ensure erratic spaces are handles --- e.g. `"  login   abc   def  "`

_service_test.go()_
1. Ensure all handlers are working fine with valid inputs
2. Ensure all handlers output the accurate error for invalid inputs --- e.g. "$" sign in inputs, wrong number of arguments, negative inputs
3. Ensure `register` does not create duplicate accounts
4. Ensure `login` only happens with correct credentials
5. Ensure `withdraw` and `send` does not happen when insufficient funds
6. Ensure `send` deducts from source account and credits destination account
7. Ensure `send` argument's destination account exists

_user flow_
User flow is tested manually.


# Reflection and Areas of Improvement

Throughout the course of this task, I had 2 main goals:
1 ) Keep the code/project as simple and clean as possible
2) Make the user flow and experience as foolproof and bulletproof as possible

In terms of my first goal, I felt that I was off to a good start initially as I had taken a pretty long time to plan the overall concept of the application, as well as a clear and differentiated project structure. However, the persistent data store added a bit of complexity to the project because I had to balance between creating a simple implementation and ensuring its persistence was reliable.

As a result, I am aware that the implementation of files as tables might not be a very clean solution compared to using a separate database like Postgres. However, I felt that I had navigated this trade-off to the best of my abilities.

Regarding the code cleanliness, I tried to extract repeated logic and constants out as much as possible, and attempted to provide clear comments in the code.

In terms of my second goal, I am confident that I have considered all cases of user flow, and handled all errors gracefully with concise error messages. I have spent the largest proportion of my time on testing - manual user flow testing and writing test. However, I feel that I could have included flow testing as integration tests, to ensure certain user flows are forbidden and handled.

Lastly, I hope that the user experience is as clear as possible, where all information shown or required is minimal and necessary, and users will not need much prompts to understand how to operation the application. 

