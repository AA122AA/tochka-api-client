package dto

type GetOperationsResponse struct {
	Data  Data  `json:"Data"`
	Links Links `json:"Links"`
	Meta  Meta  `json:"Meta"`
}

type Data struct {
	Operation []Operation `json:"Operation"`
}

type Operation struct {
	CustomerCode     string   `json:"customerCode"`
	TaxSystemCode    string   `json:"taxSystemCode"`
	PaymentType      string   `json:"paymentType"`
	PaymentId        string   `json:"paymentId"`
	TransactionId    string   `json:"transactionId"`
	CreatedAt        string   `json:"createdAt"`
	PaymentMode      []string `json:"paymentMode"`
	RedirectUrl      string   `json:"redirectUrl"`
	FailRedirectUrl  string   `json:"failRedirectUrl"`
	Client           Client   `json:"Client"`
	Items            []Item   `json:"Items"`
	Purpose          string   `json:"purpose"`
	Amount           float64  `json:"amount"`
	Status           string   `json:"status"`
	OperationId      string   `json:"operationId"`
	PaymentLink      string   `json:"paymentLink"`
	MerchantId       string   `json:"merchantId"`
	ConsumerId       string   `json:"consumerId"`
	Order            []Order  `json:"Order"`
	Supplier         Supplier `json:"Supplier"`
	PreAuthorization bool     `json:"preAuthorization"`
	PaidAt           string   `json:"paidAt"`
	PaymentLinkId    string   `json:"paymentLinkId"`
}

type Client struct {
	Name  string `json:"name"`
	Email string `json:"email"`
	Phone string `json:"phone"`
}

type Item struct {
	VatType       string   `json:"vatType"`
	Name          string   `json:"name"`
	Amount        float64  `json:"amount"`
	Quantity      float64  `json:"quantity"`
	PaymentMethod string   `json:"paymentMethod"`
	PaymentObject string   `json:"paymentObject"`
	Measure       string   `json:"measure"`
	Supplier      Supplier `json:"Supplier"`
}

type Supplier struct {
	Phone   string `json:"phone"`
	Name    string `json:"name"`
	TaxCode string `json:"taxCode"`
}

type Order struct {
	OrderId string  `json:"orderId"`
	Type    string  `json:"type"`
	Amount  float64 `json:"amount"`
	Time    string  `json:"time"`
}

type Links struct {
	Self  string `json:"self"`
	First string `json:"first"`
	Prev  string `json:"prev"`
	Next  string `json:"next"`
	Last  string `json:"last"`
}

type Meta struct {
	TotalPages int `json:"totalPages"`
}

type GetPaymentRegistryResponse struct {
	Data  RegistryData `json:"Data"`
	Links Links        `json:"Links"`
	Meta  Meta         `json:"Meta"`
}

type RegistryData struct {
	Registry []Registry `json:"Registry"`
}

type Registry struct {
	PaymentType string    `json:"paymentType"`
	TotalAmount float64   `json:"totalAmount"`
	PaymentId   string    `json:"paymentId"`
	Payments    []Payment `json:"payments"`
}

type Payment struct {
	Purpose          string  `json:"purpose"`
	Status           string  `json:"status"`
	Amount           float64 `json:"amount"`
	OperationId      string  `json:"operationId"`
	PaymentLink      string  `json:"paymentLink"`
	Time             string  `json:"time"`
	Number           string  `json:"number"`
	Commission       float64 `json:"commission"`
	EnrollmentAmount float64 `json:"enrollmentAmount"`
}

type GetRetailersResponse struct {
	Data  RetailerData `json:"Data"`
	Links Links        `json:"Links"`
	Meta  Meta         `json:"Meta"`
}

type RetailerData struct {
	Retailer []Retailer `json:"Retailer"`
}

type Retailer struct {
	Status       string   `json:"status"`
	IsActive     bool     `json:"isActive"`
	Mcc          string   `json:"mcc"`
	Rate         float64  `json:"rate"`
	Name         string   `json:"name"`
	Url          string   `json:"url"`
	MerchantId   string   `json:"merchantId"`
	TerminalId   string   `json:"terminalId"`
	PaymentModes []string `json:"paymentModes"`
	Cashbox      string   `json:"cashbox"`
}

type CreatePaymentOperationData struct {
	Data PaymentData `json:"Data"`
}

type PaymentData struct {
	CustomerCode     string   `json:"customerCode"`
	Amount           float64  `json:"amount"`
	Purpose          string   `json:"purpose"`
	RedirectURL      string   `json:"redirectUrl,omitempty"`
	FailRedirectURL  string   `json:"failRedirectUrl,omitempty"`
	PaymentMode      []string `json:"paymentMode"`
	SaveCard         bool     `json:"saveCard,omitempty"`
	ConsumerID       string   `json:"consumerId,omitempty"`
	MerchantID       string   `json:"merchantId,requiered"`
	PreAuthorization bool     `json:"preAuthorization"`
	TTL              int      `json:"ttl,omitempty"`
	PaymentLinkID    string   `json:"paymentLinkId,omitempty"`
}
