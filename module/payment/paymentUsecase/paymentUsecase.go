package paymentUsecase

import "github.com/bunkai888/hello-microservice/module/payment/paymentRepository"

type (
	PaymentUsecaseService interface{}

	paymentUsecase struct {
		paymentRepository paymentRepository.PaymentRepositoryService
	}
)

func NewPaymentUsecase(paymentRepository paymentRepository.PaymentRepositoryService) PaymentUsecaseService {
	return &paymentUsecase{
		paymentRepository: paymentRepository,
	}
}
