package services

import (
	"context"
	"github.com/newage-saint/insuretech-go-sdk/pkg/models"
	"strings"
)

// PaymentService handles payment-related API calls
type PaymentService struct {
	Client Client
}

// GetPayment Get payment
func (s *PaymentService) GetPayment(ctx context.Context, paymentId string) (*models.PaymentRetrievalResponse, error) {
	path := "/v1/payments/{payment_id}"
	path = strings.ReplaceAll(path, "{payment_id}", paymentId)
	var result models.PaymentRetrievalResponse
	err := s.Client.DoRequest(ctx, "GET", path, nil, &result)
	if err != nil {
		return nil, err
	}
	return &result, nil
}

// ListPaymentMethods Payment Methods
func (s *PaymentService) ListPaymentMethods(ctx context.Context, userId string) (*models.PaymentMethodsListingResponse, error) {
	path := "/v1/users/{user_id}/payment-methods"
	path = strings.ReplaceAll(path, "{user_id}", userId)
	var result models.PaymentMethodsListingResponse
	err := s.Client.DoRequest(ctx, "GET", path, nil, &result)
	if err != nil {
		return nil, err
	}
	return &result, nil
}

// AddPaymentMethod Add payment method
func (s *PaymentService) AddPaymentMethod(ctx context.Context, userId string, req *models.AddPaymentMethodRequest) (*models.AddPaymentMethodResponse, error) {
	path := "/v1/users/{user_id}/payment-methods"
	path = strings.ReplaceAll(path, "{user_id}", userId)
	var result models.AddPaymentMethodResponse
	err := s.Client.DoRequest(ctx, "POST", path, req, &result)
	if err != nil {
		return nil, err
	}
	return &result, nil
}

// VerifyPayment Verify payment
func (s *PaymentService) VerifyPayment(ctx context.Context, paymentId string, req *models.PaymentVerificationRequest) (*models.PaymentVerificationResponse, error) {
	path := "/v1/payments/{payment_id}:verify"
	path = strings.ReplaceAll(path, "{payment_id}", paymentId)
	var result models.PaymentVerificationResponse
	err := s.Client.DoRequest(ctx, "POST", path, req, &result)
	if err != nil {
		return nil, err
	}
	return &result, nil
}

// InitiateRefund Refund Management
func (s *PaymentService) InitiateRefund(ctx context.Context, paymentId string, req *models.InitiateRefundRequest) (*models.InitiateRefundResponse, error) {
	path := "/v1/payments/{payment_id}/refunds"
	path = strings.ReplaceAll(path, "{payment_id}", paymentId)
	var result models.InitiateRefundResponse
	err := s.Client.DoRequest(ctx, "POST", path, req, &result)
	if err != nil {
		return nil, err
	}
	return &result, nil
}

// GetRefundStatus Get refund status
func (s *PaymentService) GetRefundStatus(ctx context.Context, refundId string) (*models.RefundStatusRetrievalResponse, error) {
	path := "/v1/refunds/{refund_id}"
	path = strings.ReplaceAll(path, "{refund_id}", refundId)
	var result models.RefundStatusRetrievalResponse
	err := s.Client.DoRequest(ctx, "GET", path, nil, &result)
	if err != nil {
		return nil, err
	}
	return &result, nil
}

// InitiatePayment Payment Processing
func (s *PaymentService) InitiatePayment(ctx context.Context, req *models.PaymentInitiatePaymentRequest) (*models.PaymentInitiatePaymentResponse, error) {
	path := "/v1/payments"
	var result models.PaymentInitiatePaymentResponse
	err := s.Client.DoRequest(ctx, "POST", path, req, &result)
	if err != nil {
		return nil, err
	}
	return &result, nil
}

// ListPayments List payments
func (s *PaymentService) ListPayments(ctx context.Context) (*models.PaymentsListingResponse, error) {
	path := "/v1/payments"
	var result models.PaymentsListingResponse
	err := s.Client.DoRequest(ctx, "GET", path, nil, &result)
	if err != nil {
		return nil, err
	}
	return &result, nil
}

// ReconcilePayments Reconciliation
func (s *PaymentService) ReconcilePayments(ctx context.Context, req *models.ReconcilePaymentsRequest) (*models.ReconcilePaymentsResponse, error) {
	path := "/v1/payments:reconcile"
	var result models.ReconcilePaymentsResponse
	err := s.Client.DoRequest(ctx, "POST", path, req, &result)
	if err != nil {
		return nil, err
	}
	return &result, nil
}
