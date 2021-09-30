package paymentHandler

import (
	"bytes"
	"encoding/json"
	"errors"
	"fmt"
	"github.com/sinamna/PaymentGateway/internal/models"
	"github.com/valyala/fasthttp"
	"io/ioutil"
	"net/http"
	"os"
)

var ACCESSTOKEN = os.Getenv("IDPAY_ACCESS_TOKEN")


type PaymentHandler struct{
	Server fasthttp.Server
}

func NewPaymentHandler() *PaymentHandler {
	return &PaymentHandler{}
}

func (ph *PaymentHandler) MakeTransaction(trx *models.CreateTransactionReq) (*models.CreateTransactionResp,*models.IDPayErr,error){
	if trx.Amount ==0 || trx.Callback == "" || trx.OrderID == "" {
		return nil, nil, errors.New("missing required fields")
	}

	url := "https://api.idpay.ir/v1.1/payment"
	payload, _:= json.Marshal(trx)

	req, _ := http.NewRequest("POST", url, bytes.NewBuffer(payload))

	req.Header.Set("Content-Type", "application/json")
	req.Header.Set("X-API-KEY", ACCESSTOKEN)
	req.Header.Set("X-SANDBOX", "true")

	resp, err := http.DefaultClient.Do(req)
	if err != nil{
		fmt.Println(err)
		return nil, nil, err
	}
	defer resp.Body.Close()
	body, _ := ioutil.ReadAll(resp.Body)

	var transactionResp models.CreateTransactionResp
	var idPayErr models.IDPayErr
	if resp.Status == "201"{
		_ = json.Unmarshal(body,&transactionResp)
	}else{
		_ = json.Unmarshal(body,&idPayErr)
	}
	return &transactionResp, &idPayErr, nil
}

func (ph *PaymentHandler) VerifyTransaction (verifyReq *models.TransactionReq)([]byte, int, error){

	if verifyReq.TrxID == "" || verifyReq.OrderID == "" {
		return nil, http.StatusBadRequest, errors.New("missing required fields")
	}

	url := "https://api.idpay.ir/v1.1/payment/verify"
	payload, _ := json.Marshal(verifyReq)

	req, _ := http.NewRequest("POST", url, bytes.NewBuffer(payload))

	req.Header.Set("Content-Type", "application/json")
	req.Header.Set("X-API-KEY", ACCESSTOKEN)
	req.Header.Set("X-SANDBOX", "true")

	resp, _ := http.DefaultClient.Do(req)

	defer resp.Body.Close()
	body, _ := ioutil.ReadAll(resp.Body)

	return body, resp.StatusCode , nil
}

func (ph *PaymentHandler) GetTransactionStat (getTrxReq *models.TransactionReq)([]byte, int, error){
	if getTrxReq.TrxID == "" || getTrxReq.OrderID == "" {
		return nil, http.StatusBadRequest, errors.New("missing required fields")
	}
	url := "https://api.idpay.ir/v1.1/payment/inquiry"
	payload, _ := json.Marshal(getTrxReq)

	req, _ := http.NewRequest("POST", url, bytes.NewBuffer(payload))

	req.Header.Set("Content-Type", "application/json")
	req.Header.Set("X-API-KEY", ACCESSTOKEN)
	req.Header.Set("X-SANDBOX",	"true")

	resp, _ := http.DefaultClient.Do(req)

	defer resp.Body.Close()
	body, _ := ioutil.ReadAll(resp.Body)

	return body, resp.StatusCode,nil
}