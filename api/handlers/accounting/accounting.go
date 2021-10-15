package handlers

import (
	"context"
	"database/sql"
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"tis-gf-api/models"
	"tis-gf-api/mydb"
	"tis-gf-api/utils"

	"github.com/gorilla/mux"
	"github.com/volatiletech/sqlboiler/v4/boil"
)

type Accounting struct {
	l *log.Logger
	// afs *models.AccountingFinancialStatement
}

func NewAccounting(l *log.Logger) *Accounting {
	return &Accounting{l}
}

func (a *Accounting) GetAllFS(w http.ResponseWriter, r *http.Request) {
	db, err := mydb.GetDb()
	if err != nil {
		log.Fatal(err)
	}

	fs, err := models.AccountingFinancialStatements().All(context.Background(), db)
	if err != nil {
		http.Error(w, "Unable to get Financial Statement", http.StatusBadRequest)
	}
	utils.ToJSON(w, fs)
}

func (a *Accounting) GetFinancialStatementsByCountry(w http.ResponseWriter, r *http.Request) {

	fmt.Println("Endpoint hit")
	db, err := mydb.GetDb()
	params := mux.Vars(r)
	myCountry := params["country"]

	fModel := &models.AccountingFinancialStatement{Country: myCountry}
	fs, err := models.FindAccountingFinancialStatement(context.Background(), db,
		myCountry, fModel.ReportType, fModel.ReportYear, fModel.ReportMonth, fModel.AccNum)
	utils.DieIf(err)

	d, err := json.Marshal(fs)
	utils.DieIf((err))
	w.Write(d)
}

func (a *Accounting) AddFinancialStatement(w http.ResponseWriter, r *http.Request) {
	var db *sql.DB
	var fs models.AccountingFinancialStatement
	var err error
	var bodyString string

	err = utils.FromJSON(r.Body, &fs)
	if err != nil {
		http.Error(w, "unable to unmarshall POST request: "+bodyString, http.StatusBadRequest)
	}

	db, err = mydb.GetDb()
	if err != nil {
		http.Error(w, "Error connecting to db: "+err.Error(), http.StatusBadRequest)
	}
	err = fs.Insert(context.Background(), db, boil.Infer())
	if err != nil {
		http.Error(w, "Error inserting into db: "+err.Error(), http.StatusBadRequest)
	}

}

type KeyFS struct{}

func (a *Accounting) MiddlewareFinancialStatementValidator(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		fs := &models.AccountingFinancialStatement{}
		err := utils.FromJSON(r.Body, fs)
		if err != nil {
			http.Error(w, "unable to unmarshall POST request", http.StatusBadRequest)
			return
		}

		ctx := context.WithValue(r.Context(), KeyFS{}, fs)
		r = r.WithContext(ctx)

		next.ServeHTTP(w, r)
	})
}
