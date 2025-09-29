package tochka

import (
	"context"
	"encoding/json"
	"errors"
	"log"
	"net/http"
	"net/url"

	"github.com/AA122AA/tochka-api-client/tochka/dto"
)

type AcquiringClient struct {
	client   *Client
	basePath string
}

// Поддержать все необязательные параметры: fromDate, toDate, page, perPage, status
func (c *AcquiringClient) GetPayments(ctx context.Context, customerCode string) (*dto.GetOperationsResponse, error) {
	if customerCode == "" {
		return nil, errors.New("customerCode is empty")
	}

	u, err := c.client.buildURL(c.basePath, "/payments")
	if err != nil {
		return nil, err
	}

	params := url.Values{}
	params.Add("customerCode", customerCode)

	u.RawQuery = params.Encode()

	req, err := c.client.newRequest(ctx, http.MethodGet, u.String(), nil)
	if err != nil {
		return nil, err
	}

	respBody, err := c.client.do(req)
	if err != nil {
		return nil, err
	}

	var data dto.GetOperationsResponse
	err = json.Unmarshal(respBody, &data)
	if err != nil {
		return nil, err
	}

	return &data, nil
}

func (c *AcquiringClient) GetOperationInfo(ctx context.Context, operationID string) (*dto.GetOperationsResponse, error) {
	if operationID == "" {
		return nil, errors.New("operationID is empty")
	}

	path := "/payments/" + operationID

	u, err := c.client.buildURL(c.basePath, path)
	if err != nil {
		return nil, err
	}

	req, err := c.client.newRequest(ctx, http.MethodGet, u.String(), nil)
	if err != nil {
		return nil, err
	}

	respBody, err := c.client.do(req)
	if err != nil {
		return nil, err
	}

	var data dto.GetOperationsResponse
	err = json.Unmarshal(respBody, &data)
	if err != nil {
		return nil, err
	}

	return &data, nil
}

// Поддержать все необязательные параметры: fromDate, toDate, page, perPage, status
// Переделать date с string на datetime
func (c *AcquiringClient) GetPaymentRegistry(ctx context.Context, customerCode, merchantID, paymentID, date string) (*dto.GetPaymentRegistryResponse, error) {
	switch {
	case customerCode == "":
		return nil, errors.New("customerCode is empty")
	case merchantID == "":
		return nil, errors.New("merchantID is empty")
	case date == "":
		return nil, errors.New("date is empty")
	}

	u, err := c.client.buildURL(c.basePath, "/registry")
	if err != nil {
		return nil, err
	}

	params := url.Values{}
	params.Add("customerCode", customerCode)
	params.Add("merchantId", merchantID)
	if paymentID != "" {
		params.Add("paymentId", paymentID)
	}
	params.Add("date", date)

	u.RawQuery = params.Encode()

	req, err := c.client.newRequest(ctx, http.MethodGet, u.String(), nil)
	if err != nil {
		return nil, err
	}

	respBody, err := c.client.do(req)
	if err != nil {
		return nil, err
	}

	var data dto.GetPaymentRegistryResponse
	err = json.Unmarshal(respBody, &data)
	if err != nil {
		return nil, err
	}

	return &data, nil
}

func (c *AcquiringClient) GetRetailers(ctx context.Context, customerCode string) (*dto.GetRetailersResponse, error) {
	if customerCode == "" {
		return nil, errors.New("customerCode is empty")
	}

	u, err := c.client.buildURL(c.basePath, "/retailers")
	if err != nil {
		return nil, err
	}

	params := url.Values{}
	params.Add("customerCode", customerCode)

	u.RawQuery = params.Encode()

	req, err := c.client.newRequest(ctx, http.MethodGet, u.String(), nil)
	if err != nil {
		return nil, err
	}

	log.Printf("Request - %v", req)

	respBody, err := c.client.do(req)
	if err != nil {
		return nil, err
	}

	var data dto.GetRetailersResponse
	err = json.Unmarshal(respBody, &data)
	if err != nil {
		return nil, err
	}

	return &data, nil
}

// TODO: add optional params: redirectUrl,failRedirectUrl, saveCard, consumerId, preAthorization, ttl, paymentLinkId
func (c *AcquiringClient) CreatePaymentOperation(
	ctx context.Context,
	customerCode, merchantID, purpose string,
	paymentModes []string,
	amountOfMoney float64,
	preAuth bool,
) (*dto.CreatePaymentOperationDataResponse, error) {
	u, err := c.client.buildURL(c.basePath, "/payments")
	if err != nil {
		return nil, err
	}

	op := dto.CreatePaymentOperationData{
		Data: dto.PaymentData{
			CustomerCode:     customerCode,
			Amount:           amountOfMoney,
			Purpose:          purpose,
			PaymentMode:      paymentModes,
			MerchantID:       merchantID,
			PreAuthorization: preAuth,
		},
	}

	body, err := json.Marshal(op)
	if err != nil {
		return nil, err
	}

	req, err := c.client.newRequest(ctx, http.MethodPost, u.String(), body)
	if err != nil {
		return nil, err
	}

	log.Printf("Request - %v", req)

	respBody, err := c.client.do(req)
	if err != nil {
		return nil, err
	}

	var data dto.CreatePaymentOperationDataResponse
	err = json.Unmarshal(respBody, &data)
	if err != nil {
		return nil, err
	}

	return &data, nil
}

func (c *AcquiringClient) CapturePayment(ctx context.Context, operationID string) (*dto.CapturePaymentResponse, error) {
	path := "/payments/" + operationID + "/capture"
	u, err := c.client.buildURL(c.basePath, path)
	if err != nil {
		return nil, err
	}

	req, err := c.client.newRequest(ctx, http.MethodPost, u.String(), nil)
	if err != nil {
		return nil, err
	}

	log.Printf("Request - %v", req)

	respBody, err := c.client.do(req)
	if err != nil {
		return nil, err
	}

	var data dto.CapturePaymentResponse
	err = json.Unmarshal(respBody, &data)
	if err != nil {
		return nil, err
	}

	return &data, nil
}

func (c *AcquiringClient) RefundPaymentOperation(
	ctx context.Context, operationID string, amountOfMoney float64,
) (*dto.RefundPaymentOperationResponse, error) {
	path := "/payments/" + operationID + "/refund"
	u, err := c.client.buildURL(c.basePath, path)
	if err != nil {
		return nil, err
	}

	op := dto.RefundPaymentOperationRequest{
		Data: dto.RefundData{
			Amount: amountOfMoney,
		},
	}

	body, err := json.Marshal(op)
	if err != nil {
		return nil, err
	}

	req, err := c.client.newRequest(ctx, http.MethodPost, u.String(), body)
	if err != nil {
		return nil, err
	}

	log.Printf("Request - %v", req)

	respBody, err := c.client.do(req)
	if err != nil {
		return nil, err
	}

	var data dto.RefundPaymentOperationResponse
	err = json.Unmarshal(respBody, &data)
	if err != nil {
		return nil, err
	}

	return &data, nil
}
