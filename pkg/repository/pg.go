package repository

import (
	"database/sql"
	"github.com/sinamna/PaymentGateway/internal/models"
)

type PgDatabase struct{
	client *sql.DB
}

func(pg *PgDatabase)SaveTransaction(req *models.CreateTransactionReq)error{

	return nil
}
