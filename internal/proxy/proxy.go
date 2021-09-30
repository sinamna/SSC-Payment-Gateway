package proxy

import (
	"github.com/gin-gonic/gin"
	paymentHandler2 "github.com/sinamna/PaymentGateway/internal/paymentHandler"
	"log"
)

var paymentHandler *paymentHandler2.PaymentHandler

func StartServer() {

	paymentHandler = paymentHandler2.NewPaymentHandler()

	router := gin.Default()

	router.POST("/transaction/new", NewTransactionHandler)
	router.POST("/transaction/verify", VerifyTransactionHandler)
	router.POST("/transaction/state", GetTransactionState)

	log.Fatalln(router.Run(":8080"))
}
