package dto

type GetPaymentsResponse struct {
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
