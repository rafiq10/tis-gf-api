package reports

import (
	"log"
	"net/http"

	"tis-gf-api/mydb"
	"tis-gf-api/utils"
)

type Reports struct {
	l *log.Logger
	// afs *models.AccountingFinancialStatement
}

type errMsg struct {
	errKey string
	errTxt string
}

type ReportsAging struct {
	l *log.Logger
	// afs *models.AccountingFinancialStatement
}

type ReportsDCC struct {
	l           *log.Logger
	country     string
	reportYear  string
	reportMonth string
	pdcYear     string
	pdcMonth    string
	// afs *models.AccountingFinancialStatement
}

func NewReports(l *log.Logger) *Reports {
	return &Reports{l}
}
func NewReportsAgingCLI(l *log.Logger) *ReportsAging {
	return &ReportsAging{l}
}

func NewReportsDCC(l *log.Logger, ctry, rYear, rMonth, pdcYear, pdcMonth string) *ReportsDCC {
	return &ReportsDCC{l, ctry, rYear, rMonth, pdcYear, pdcMonth}
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
	w.Header().Set("content-type", "application/json")

	var errMsg []utils.ErrMsg

	if r.FormValue("country") == "" {
		errMsg = append(errMsg, utils.ErrMsg{ErrTxt: "Empty field: country"})
	}
	if r.FormValue("pdcyear") == "" {
		errMsg = append(errMsg, utils.ErrMsg{ErrTxt: "Empty field: pdcyear"})
	}
	if r.FormValue("pdcmonth") == "" {
		errMsg = append(errMsg, utils.ErrMsg{ErrTxt: "Empty field: pdcmonth"})
	}
	if r.FormValue("reportyear") == "" {
		errMsg = append(errMsg, utils.ErrMsg{ErrTxt: "Empty field: reportyear"})
	}
	if r.FormValue("reportmonth") == "" {
		errMsg = append(errMsg, utils.ErrMsg{ErrTxt: "Empty field: reportmonth"})
	}

	if len(errMsg) != 0 {
		// rep.l.Println("len(errMsg) != 0" + fmt.Sprintf("%v", errMsg))
		utils.RespondWithError(w, http.StatusBadRequest, errMsg)
		return
	}
	country := r.FormValue("country")
	pdcYear := r.FormValue("pdcyear")
	pdcMonth := r.FormValue("pdcmonth")
	reportYear := r.FormValue("reportyear")
	reportMonth := r.FormValue("reportmonth")

	rep.country = country
	rep.pdcYear = pdcYear
	rep.pdcMonth = pdcMonth
	rep.reportYear = reportYear
	rep.reportMonth = reportMonth
	mysql := "select * from [TESO_get_DCC]('" + country + "'," + reportYear + "," + reportMonth + "," + pdcYear + "," + pdcMonth + ")"
	err := mydb.ExecuteSQL(mysql, w)
	if err != nil {
		log.Fatal(err, mysql)
	}

}
