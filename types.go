package flexepin

type Flexepin struct {
	SiteKey string
	Secret string
	RootUrl string
	TerminalId string
	TransactionIdGenerator func() string
}

type Customer struct {
	Id string `json:"customer_id"`
	Ip string `json:"customer_ip"`
	Data string `json:"customer_ip_data"` // undocumented. Your guess is as good as mine...
}

type StatusRes struct {
	Status string `json:"status"`
}

// Some of these values are not always present, depending on the result.
type ValidateRes struct {
	Result string `json:"result"`
	Msg string `json:"result_description"`
	TransactionId string `json:"transaction_id"`
	TransactionNumber string `json:"trans_no"`
	Serial string `json:"serial"`
	Value int `json:"value"`
	Cost float64 `json:"cost"`
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
	Cost float64 `json:"cost"`
	ResidualValue int `json:"residual_value"`
	Status string `json:"status"`
	Currency string `json:"currency"`
	Ean string `json:"ean"`
	Description string `json:"description"`
}

type Transaction struct {
	Result string `json:"result"`
	Msg string `json:"result_description"`
	TransactionNumber string `json:"trans_no"`
	TransactionId string `json:"transaction_id"`
	Serial string `json:"serial"`
	Value int `json:"value"`
	Cost float64 `json:"cost"`
	Status string `json:"status"`
	Currency string `json:"currency"`
	Description string `json:"description"`
	Timestamp string `json:"timestamp"`
}

// Annoyingly, querying transactions can either return a singular Transaction struct, or a list of Transaction structs.
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
