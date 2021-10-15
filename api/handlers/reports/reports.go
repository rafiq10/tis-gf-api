package reports

import (
	"log"
	"net/http"

	"tis-gf-api/mydb"
)

type Reports struct {
	l *log.Logger
	// afs *models.AccountingFinancialStatement
}

type ReportsAging struct {
	l *log.Logger
	// afs *models.AccountingFinancialStatement
}

type ReportsDCC struct {
	l *log.Logger
	// afs *models.AccountingFinancialStatement
}

func NewReports(l *log.Logger) *Reports {
	return &Reports{l}
}
func NewReportsAgingCLI(l *log.Logger) *ReportsAging {
	return &ReportsAging{l}
}

func NewReportsDCC(l *log.Logger) *ReportsDCC {
	return &ReportsDCC{l}
}

func (acc *Reports) GetReportsLoadedResume(w http.ResponseWriter, r *http.Request) {
	myYear := r.FormValue("year")
	myMonth := r.FormValue("month")

	err := mydb.ExecuteSQL("select country,funnel_numOfertas,CPY_numPY,RfCartera_numPY,CLI_numPY,PRO_numPY,TESO_numConceptos from REPORTS_Get_Reports_Loaded_Resume("+myYear+","+myMonth+") order by country", w)
	if err != nil {
		log.Fatal(err)
	}

}

func (rep *ReportsAging) GetAgingCLI(w http.ResponseWriter, r *http.Request) {
	country := r.FormValue("country")
	pdcYear := r.FormValue("pdcyear")
	pdcMonth := r.FormValue("pdcmonth")
	reportYear := r.FormValue("reportyear")
	reportMonth := r.FormValue("reportmonth")

	err := mydb.ExecuteSQL(" sp_REPORTS_Get_CLI_Aging '"+country+"',"+pdcYear+","+pdcMonth+","+reportYear+","+reportMonth, w)
	if err != nil {
		log.Fatal(err)
	}
}

func (rep *ReportsDCC) GetDCC(w http.ResponseWriter, r *http.Request) {
	country := r.FormValue("country")
	pdcYear := r.FormValue("pdcyear")
	pdcMonth := r.FormValue("pdcmonth")
	reportYear := r.FormValue("reportyear")
	reportMonth := r.FormValue("reportmonth")

	mysql := "select * from [TESO_get_DCC]('" + country + "'," + reportYear + "," + reportMonth + "," + pdcYear + "," + pdcMonth + ")"
	err := mydb.ExecuteSQL(mysql, w)
	if err != nil {
		log.Fatal(err, mysql)
	}

}

// func (acc *Reports) ServeHTTP(w http.ResponseWriter, r *http.Request) {
// 	var err error
// 	if r.Method == http.MethodGet {
// 		err = reports.GetReportsLoadedResume(w, r)
// 	}
// 	if err != nil {
// 		http.Error(w, err.Error(), http.StatusBadRequest)
// 		return
// 	}
// }
// func (rep *ReportsAging) ServeHTTP(w http.ResponseWriter, r *http.Request) {
// 	var err error
// 	if r.Method == http.MethodGet {
// 		err = reports.GetAgingCLI(w, r)
// 	}
// 	if err != nil {
// 		http.Error(w, err.Error(), http.StatusBadRequest)
// 		return
// 	}
// }

// func (rep *ReportsDCC) ServeHTTP(w http.ResponseWriter, r *http.Request) {
// 	var err error
// 	if r.Method == http.MethodGet {
// 		err = reports.GetDCC(w, r)
// 	}

// 	if err != nil {
// 		http.Error(w, err.Error(), http.StatusBadRequest)
// 		// http.Error(w, err.Error(), http.StatusBadRequest)
// 		return
// 	}
// }
