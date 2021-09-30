package proxy

import (
	"github.com/gin-gonic/gin"
	"github.com/sinamna/PaymentGateway/internal/models"
	"net/http"
)

func NewTransactionHandler(ctx *gin.Context){
	var newCreateTransactionReq models.CreateTransactionReq
	if err := ctx.ShouldBindJSON(&newCreateTransactionReq); err!= nil{
		ctx.JSON(http.StatusBadRequest,gin.H{"error":err})
	}

	newCreateTransactionReq.GenerateID()
	//saving to db
	trxResponse, idPayErr, err := paymentHandler.MakeTransaction(&newCreateTransactionReq)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError,gin.H{"error": err})
		return
	}
	if idPayErr != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": idPayErr.ErrorMessage})
		return
	}

	ctx.JSON(http.StatusCreated,gin.H{"transaction":trxResponse, "order_id": newCreateTransactionReq.OrderID})
}


func VerifyTransactionHandler(ctx *gin.Context){
	var verifyReq models.TransactionReq
	if err := ctx.ShouldBindJSON(&verifyReq); err!= nil{
		ctx.JSON(http.StatusBadRequest,gin.H{"error":err})
	}
	body, statusCode, err := paymentHandler.VerifyTransaction(&verifyReq)
	if statusCode == http.StatusBadRequest {
		ctx.JSON(400, gin.H{
			"error": err,
		})
	}
	//saving to db
	ctx.JSON(statusCode,body)
}

func GetTransactionState(ctx *gin.Context){
	var getTrx models.TransactionReq
	if err := ctx.ShouldBindJSON(&getTrx); err!= nil{
		ctx.JSON(http.StatusBadRequest,gin.H{"error":err})
	}
	body, statusCode, err := paymentHandler.VerifyTransaction(&getTrx)
	if statusCode == http.StatusBadRequest {
		ctx.JSON(500, gin.H{
			"error": err,
		})
	}
	//saving to db
	ctx.JSON(statusCode,body)
}