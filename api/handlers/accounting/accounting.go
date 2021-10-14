package handlers

import (
	"errors"
	"log"
	"net/http"

	"tis-gf-api/api/accounting"
	"tis-gf-api/models"
	"tis-gf-api/mydb"
)

type Accounting struct {
	l *log.Logger
	// afs *models.AccountingFinancialStatement
}

func NewAccounting(l *log.Logger) *Accounting {
	return &Accounting{l}
}

func (acc *Accounting) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	var err error

	if r.Method == http.MethodGet {
		country := r.FormValue("country")

		if country == "" {
			err = errors.New("empty country")
		} else {
			err = accounting.GetAllFinancialStatements(w, r)
		}
		//
	}

	if r.Method == http.MethodPost {
		// fs := &models.AccountingFinancialStatement{}
		err = accounting.AddFinancialStatement(w, r)

	}
	if r.Method == http.MethodPut {
		// expect parameters in the URI

		mydb.GetTableFields("")
		s := models.AccountingFinancialStatement{}
		mydb.GetTableNameByStruct(s)
		// err = accounting.UpdateFinancialStatement(w, r)
	}

	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}
}

func (a *Accounting) addFS(r *http.Request, w http.ResponseWriter) {
	a.l.Println("Handle POST Financial Statemens")
}
