package server

import (
	"github.com/bunkai888/hello-microservice/module/payment/paymentHandler"
	"github.com/bunkai888/hello-microservice/module/payment/paymentRepository"
	"github.com/bunkai888/hello-microservice/module/payment/paymentUsecase"
)

func (s *server) PaymentService() {
	repo := paymentRepository.NewPaymentRepository(s.db)
	usecase := paymentUsecase.NewPaymentUsecase(repo)
	httpHandler := paymentHandler.NewPaymentHttpHandler(s.cfg, usecase)
	queueHandler := paymentHandler.NewPaymentQueueHandler(s.cfg, usecase)

	_ = httpHandler
	_ = queueHandler

	payment := s.app.Group("/payment_v1")

	payment.GET("", s.healthCheckService)
}
