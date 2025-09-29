package tochka

import (
	"context"
	"fmt"
	"testing"

	"github.com/stretchr/testify/require"
)

func TestGetPayments(t *testing.T) {

	c := NewClient("https://enter.tochka.com/sandbox/v2", "v1.0", "sandbox.jwt.token")

	cases := []struct {
		name       string
		ctx        context.Context
		customerID string
	}{
		{
			name:       "Positive GetPayments",
			ctx:        context.Background(),
			customerID: "1234567ab",
		},
	}

	for _, tCase := range cases {
		t.Run(tCase.name, func(t *testing.T) {
			payments, err := c.Acquiring.GetPayments(tCase.ctx, tCase.customerID)
			for _, p := range payments.Data.Operation {
				fmt.Printf("Payment - %+v\n", p)
			}
			require.NoError(t, err)
		})
	}

}

func TestGetOperationInfo(t *testing.T) {

	c := NewClient("https://enter.tochka.com/sandbox/v2", "v1.0", "sandbox.jwt.token")

	cases := []struct {
		name        string
		ctx         context.Context
		operationID string
	}{
		{
			name:        "Positive GetOperationInfo",
			ctx:         context.Background(),
			operationID: "48232c9a-ce82-1593-3cb6-5c85a1ffef8f",
		},
		{
			name:        "Positive GetOperationInfo 2",
			ctx:         context.Background(),
			operationID: "74a21bae-53ca-4103-a6a8-5ca1d1fcf100",
		},
	}

	for _, tCase := range cases {
		t.Run(tCase.name, func(t *testing.T) {
			opInfo, err := c.Acquiring.GetOperationInfo(tCase.ctx, tCase.operationID)
			fmt.Printf("Operation info - %+v\n", opInfo)
			require.NoError(t, err)
		})
	}

}

func TestGetPaymentRegistry(t *testing.T) {

	c := NewClient("https://enter.tochka.com/sandbox/v2", "v1.0", "sandbox.jwt.token")

	cases := []struct {
		name       string
		ctx        context.Context
		customerID string
		merchantID string
		paymentID  string
		date       string
	}{
		{
			name:       "Positive GetPaymentRegistry no paymentID",
			ctx:        context.Background(),
			customerID: "1234567ab",
			merchantID: "200000000001097",
			paymentID:  "",
			date:       "2025-05-10",
		},
	}

	for _, tCase := range cases {
		t.Run(tCase.name, func(t *testing.T) {
			registry, err := c.Acquiring.GetPaymentRegistry(tCase.ctx, tCase.customerID, tCase.merchantID, tCase.paymentID, tCase.date)
			fmt.Printf("Payment Registry - %+v\n", registry)
			require.NoError(t, err)
		})
	}
}

func TestGetRetailers(t *testing.T) {

	c := NewClient("https://enter.tochka.com/sandbox/v2", "v1.0", "sandbox.jwt.token")

	cases := []struct {
		name       string
		ctx        context.Context
		customerID string
	}{
		{
			name:       "Positive GetRetailers",
			ctx:        context.Background(),
			customerID: "1234567ab",
		},
	}

	for _, tCase := range cases {
		t.Run(tCase.name, func(t *testing.T) {
			registry, err := c.Acquiring.GetRetailers(tCase.ctx, tCase.customerID)
			fmt.Printf("Payment Registry - %+v\n", registry)
			require.NoError(t, err)
		})
	}
}

func TestCreatePaymentOperation(t *testing.T) {

	c := NewClient("https://enter.tochka.com/sandbox/v2", "v1.0", "sandbox.jwt.token")

	cases := []struct {
		name          string
		ctx           context.Context
		customerCode  string
		merchantID    string
		purpose       string
		paymentModes  []string
		amountOfMoney float64
	}{
		{
			name:          "Positive Create no paymentID",
			ctx:           context.Background(),
			customerCode:  "1234567ab",
			merchantID:    "200000000001097",
			purpose:       "test",
			paymentModes:  []string{"sbp", "card"},
			amountOfMoney: 100.0,
		},
	}

	for _, tCase := range cases {
		t.Run(tCase.name, func(t *testing.T) {
			op, err := c.Acquiring.CreatePaymentOperation(tCase.ctx, tCase.customerCode, tCase.merchantID, tCase.purpose, tCase.paymentModes, tCase.amountOfMoney)
			require.NoError(t, err)
			fmt.Printf("Created Payment Operation - %+v\n", op)
		})
	}
}
