# GO-Smarshal

Smarshal (SMARt + unMARSHAL)

(i'm just putting it out there beacuse i make use of it in few of my personal projects, contributions are welcome)

<a href="https://www.buymeacoffee.com/jakoboo" target="_blank"><img src="https://cdn.buymeacoffee.com/buttons/lato-orange.png" alt="Buy Me A Coffee" style="height: 51px !important;width: 217px !important; border-radius: 5px !important;" ></a>

# Use case and examples

## Marshal

```go
type User struct {
	ID       uint   `json:"id"`
	Email    string `json:"email"`
	Username string `json:"username"`
}

type Meta struct {
	RequestID string `json:"requestID"`
}
```

```go
user := &User{
	ID:      1,
	Email:    "mail@example.com",
	Username: "example",

}
meta := &Meta{
	RequestID: "be700d399b9761c5f021a8774b304175",
}

b, err := smarshal.Marshal(user, meta)
if err != nil {
	...
}
```

Marshal will return combined struct like below:

```json
{
  "id": 1,
  "email": "mail@example.com",
  "username": "example",
  "requestID": "be700d399b9761c5f021a8774b304175"
}
```

## Unmarshal

Consider this possible API responses:

```json
{
  "status": "failed",
  "error": "Something went really bad! :c",
  "requestID": "be700d399b9761c5f021a8774b304175"
}
```

```json
{
  "id": 1,
  "email": "mail@example.com",
  "username": "example",
  "requestID": "be700d399b9761c5f021a8774b304175"
}
```

Normally you could create a struct containing all possible fields preffarably as pointers so you could check if they are nil and act accordingly. Using this package you can create single CustomError for your API that is universal accross different responses.

```go
type CustomError struct {
	Status string `json:"status"`
	Error  string `json:"error"`
}

type User struct {
	ID       uint   `json:"id"`
	Email    string `json:"email"`
	Username string `json:"username"`
	HasPaid  bool   `json:"hasPaid"`
}

type Meta struct {
	RequestID string `json:"requestID"`
}
```

```go
customErr := &CustomError{}
user := &User{}
meta := &Meta{}

err := smarshal.Unmarshal(apiResponse, &user, &meta, &customErr)
if err != nil {
	...
}

if customErr != nil {
	...
}

if user != nil {
	...
}

if meta != nil {
	...
}
```

Notice **double pointer** used here to be able to check for `nil` when populated struct was equal to its zero value (Unmarshal if provided single pointer will work as expected but will return zero value struct)
