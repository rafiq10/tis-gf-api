package accounting

import (
	"context"
	"database/sql"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"reflect"

	"github.com/gorilla/mux"
	"github.com/volatiletech/sqlboiler/v4/boil"
	"tis-gf-api/models"
	"tis-gf-api/mydb"
	"tis-gf-api/utils"
)

func GetAllFinancialStatements(w http.ResponseWriter, r *http.Request) error {
	db, err := mydb.GetDb()
	if err != nil {
		log.Fatal(err)
	}

	fs, err := models.AccountingFinancialStatements().All(context.Background(), db)
	if err != nil {
		log.Fatal(err)
	}
	return utils.ToJSON(w, fs)
}

func AddFinancialStatement(w http.ResponseWriter, r *http.Request) error {
	// API.SetApiHeaders(w)
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
		log.Fatal(err)
	}
	err = fs.Insert(context.Background(), db, boil.Infer())
	if err != nil {
		log.Fatal(err)
	}
	return err
}

func UpdateFinancialStatement(w http.ResponseWriter, r *http.Request) error {
	// var fs models.AccountingFinancialStatement
	// var err error

	// db, err = mydb.GetDb()
	// x := models.AccountingFinancialStatementColumns
	x := models.AccountingFinancialStatementColumns
	s := reflect.ValueOf(&x).Elem().Type()

	// values := make([]interface{}, v.NumField())
	for i := 0; i < s.NumField(); i++ {
		log.Println("v.Field(i).Type():" + s.Field(i).Name)

	}
	return nil

}

func getBodyString(r *http.Request) (bodyString string) {
	var bodyBytes []byte
	var err error
	bodyBytes, err = ioutil.ReadAll(r.Body)
	if err != nil {
		log.Println("Error en getBodyString")
		return ""
	}
	bodyString = string(bodyBytes)
	log.Println("getBodyString before return: " + bodyString)
	return bodyString

}

// func SetApiHeaders(w http.ResponseWriter) {
// 	w.Header().Set("Content-Type", "application/json")
// }
func GetFinancialStatementsByCountry(w http.ResponseWriter, r *http.Request) {

	fmt.Println("Endpoint hit")
	// API.SetApiHeaders(w)
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
	// db.Close()
}
