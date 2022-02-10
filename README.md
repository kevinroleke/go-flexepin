# go-flexepin

A Golang wrapper for the Flexepin Merchant API. 

# Usage

```go
import "log"
import "github.com/kevinroleke/go-flexepin"

func main() {
	siteKey := "ABC45FJFDE" // Provided by Flexepin
	secret := "ABCDAPWPFJEDFV" // Provided by Flexepin
	terminalId := "1234" // Indentifies this script/server.
	rootUrl := "https://testrest.flexepin.com" // API root, minus trailing slash.
	txidGenerator := flexepin.GetNonce // Function that generates a transactio ID, up to 32 characters. By default flexepin.GetNonce uses hashed microtime.
	fl := flexepin.New(siteKey, secret, terminalId, rootUrl, txidGenerator)

	ok, err := fl.Status()
	if err != nil {
		log.Fatal(err)
	}

	if ok {
		log.Println("API is active!")
	} else {
		log.Fatal("API is offline!")
	}
}
```

# Methods

### Validate

```go
Validate(pin string) (ValidateRes, error)
```

### Redeem

```go
Redeem(pin string, customer Customer) (RedeemRes, error)
```

### Get Transactions

```go
GetTransactionById(id string) (TransactionRes, error)
GetTransactionByNumber(no string) (TransactionRes, error)
GetTransactionsByDateRange(startDate time.Time, endDate time.Time) (TransactionsRes, error)
```

### Get Stores

```go
GetStoresByCountry(country string) (StoreRes, error)
GetStoresByRange(lng string, lat string, kmRange string) (StoreRes, error) // kmRange = radius in kilometers 
```

# Types

### Flexepin

```go
type Flexepin struct {
	SiteKey string
	Secret string
	RootUrl string
	TerminalId string
	TransactionIdGenerator func() string
}
```

### Customer

```go
type Customer struct {
	Id string `json:"customer_id"`
	Ip string `json:"customer_ip"`
	Data string `json:"customer_ip_data"` // undocumented. Your guess is as good as mine...
}
```

### Pin Responses

```go
type ValidateRes struct {
	Result string `json:"result"`
	Msg string `json:"result_description"`
	TransactionId string `json:"transaction_id"`
	TransactionNumber string `json:"trans_no"`
	Serial string `json:"serial"`
	Value int `json:"value"`
	Cost int `json:"cost"`
	ResidualValue int `json:"residual_value"`
	Status string `json:"status"`
	Currency string `json:"currency"`
	Ean string `json:"ean"`
	Description string `json:"description"`
}

type RedeemRes struct {
	Result string `json:"result"`
	Msg string `json:"result_description"`
	TransactionId string `json:"transaction_id"`
	TransactionNumber string `json:"trans_no"`
	Serial string `json:"serial"`
	Value int `json:"value"`
	Cost int `json:"cost"`
	ResidualValue int `json:"residual_value"`
	Status string `json:"status"`
	Currency string `json:"currency"`
	Ean string `json:"ean"`
	Description string `json:"description"`
}
```

### Transactions

```go
type Transaction struct {
	Result string `json:"result"`
	Msg string `json:"result_description"`
	TransactionNumber string `json:"trans_no"`
	TransactionId string `json:"transaction_id"`
	Serial string `json:"serial"`
	Value int `json:"value"`
	Cost int `json:"cost"`
	Status string `json:"status"`
	Currency string `json:"currency"`
	Description string `json:"description"`
	Timestamp string `json:"timestamp"`
}

type TransactionRes struct {
	TransactionNumber string `json:"trans_no"`
	TransactionId string `json:"transaction_id"`
	Timestamp string `json:"timestamp"`
	Result string `json:"result"`
	Msg string `json:"result_description"`
	Transaction Transaction `json:"transaction"`
}

type TransactionsRes struct {
	TransactionNumber string `json:"trans_no"`
	TransactionId string `json:"transaction_id"`
	Timestamp string `json:"timestamp"`
	Result string `json:"result"`
	Msg string `json:"result_description"`
	Transactions []Transaction `json:"transactions"`
}
```

### Stores

```go
type Store struct {
	Name string `json:"name"`
	Address string `json:"address1"`
	City string `json:"city"`
	State string `json:"state"`
	Country string `json:"country"`
	Latitude string `json:"latitude"`
	Longitude string `json:"longitude"`
	Category string `json:"category"`
	Distance string `json:"distance"`
}

type StoreRes struct {
	TransactionNumber string `json:"trans_no"`
	Result int `json:"result"` // Its an integer here, but a string in the other methods
	Msg string `json:"result_description"`
	Stores []Store
}
```
