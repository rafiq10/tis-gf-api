package reports_mgt

import (
	"log"
	"net/http"

	"tis-gf-api/mydb"
)

func GetReportsLoadedResume(w http.ResponseWriter, r *http.Request) error {
	myYear := r.FormValue("year")
	myMonth := r.FormValue("month")

	err := mydb.ExecuteSQL("select country,funnel_numOfertas,CPY_numPY,RfCartera_numPY,CLI_numPY,PRO_numPY,TESO_numConceptos from REPORTS_Get_Reports_Loaded_Resume("+myYear+","+myMonth+") order by country", w)
	if err != nil {
		log.Fatal(err)
	}
	return err
}

func GetAgingCLI(w http.ResponseWriter, r *http.Request) error {
	country := r.FormValue("country")
	pdcYear := r.FormValue("pdcyear")
	pdcMonth := r.FormValue("pdcmonth")
	reportYear := r.FormValue("reportyear")
	reportMonth := r.FormValue("reportmonth")

	err := mydb.ExecuteSQL(" sp_REPORTS_Get_CLI_Aging '"+country+"',"+pdcYear+","+pdcMonth+","+reportYear+","+reportMonth, w)
	if err != nil {
		log.Fatal(err)
	}
	return err
}

func GetDCC(w http.ResponseWriter, r *http.Request) error {
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
	return err
}
