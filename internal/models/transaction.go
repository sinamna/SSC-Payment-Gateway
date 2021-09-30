package models

import (
	"crypto/sha1"
	"encoding/base64"
	"fmt"
	"math/rand"
	"time"
)

type CreateTransactionReq struct {
	OrderID  string `json:"order_id"`
	Amount   int32  `json:"amount"` // between 1000 to 500,000,000 rials
	Name     string `json:"name"`
	Phone    int32  `json:"phone"` //9382198592 or 09382198592 or 989382198592
	Mail     string `json:"mail"`
	Desc     string `json:"desc"`     // transaction description 255char
	Callback string `json:"callback"` // redirected after payment
}

func (ctr *CreateTransactionReq) ToString()string{
	return fmt.Sprintf("%+v",ctr)
}


func (ctr *CreateTransactionReq) GenerateID(){
	rand.Seed(time.Now().UnixNano())
	hasher := sha1.New()
	hasher.Write([]byte(ctr.ToString()))
	hash := base64.URLEncoding.EncodeToString(hasher.Sum(nil))
	ctr.OrderID = hash
}

type CreateTransactionResp struct {
	TrxID string `json:"id"`
	Link  string `json:"link"`
}

type IDPayErr struct {
	ErrorCode    int    `json:"error_code"`
	ErrorMessage string `json:"error_message"`
}
