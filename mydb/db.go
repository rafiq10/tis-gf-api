package mydb

import (
	"database/sql"
	"errors"
	"io/ioutil"
	"log"
	"net/http"
	"os"

	"tis-gf-api/models"
	"tis-gf-api/utils"

	_ "github.com/denisenkom/go-mssqldb"
	"github.com/iancoleman/orderedmap"
	_ "github.com/iancoleman/orderedmap"

	"tis-gf-api/secrets"
)

var tableTypes = make(map[string]interface{})
var tableColumns = make(map[string]struct{})

func init() {
	tableTypes["AccountingCMBFinStatementType"] = &models.AccountingCMBFinStatementType{}
	tableTypes["AccountingFinancialStatement"] = &models.AccountingFinancialStatement{}
	tableTypes["Sysdiagram"] = &models.Sysdiagram{}
	tableTypes["TESOAccountsMapping"] = &models.TESOAccountsMapping{}
	tableTypes["TESOAging"] = &models.TESOAging{}
	tableTypes["TESOAsientoType"] = &models.TESOAsientoType{}
	tableTypes["TESOCMBBalanceSide"] = &models.TESOCMBBalanceSide{}
	tableTypes["TESOCMBCat1"] = &models.TESOCMBCat1{}
	tableTypes["TESOCMBCat2"] = &models.TESOCMBCat2{}
	tableTypes["TESOCompany"] = &models.TESOCompany{}
	tableTypes["TESOControl"] = &models.TESOControl{}
	tableTypes["TESOCreditorToAccNumMAP"] = &models.TESOCreditorToAccNumMAP{}
	tableTypes["TESOTEAActual"] = &models.TESOTEAActual{}
	tableTypes["TESOTEAInitialBalance"] = &models.TESOTEAInitialBalance{}
	tableTypes["TESOTEARF"] = &models.TESOTEARF{}
	tableTypes["TblAccBook"] = &models.TblAccBook{}
	tableTypes["TblAccBooksMgt"] = &models.TblAccBooksMgt{}
	tableTypes["TblAccBooksMgtCOPY"] = &models.TblAccBooksMgtCOPY{}
	tableTypes["TblAccBooksShadow"] = &models.TblAccBooksShadow{}
	tableTypes["TblAccountingMapping"] = &models.TblAccountingMapping{}
	tableTypes["TblCompany"] = &models.TblCompany{}
	tableTypes["TblCostComparisonConfig"] = &models.TblCostComparisonConfig{}
	tableTypes["TblCountry"] = &models.TblCountry{}
	tableTypes["TblCountriesEsp"] = &models.TblCountriesEsp{}
	tableTypes["TblCountryLimit"] = &models.TblCountryLimit{}
	tableTypes["TblCurrency"] = &models.TblCurrency{}
	tableTypes["TblExchangeRate"] = &models.TblExchangeRate{}
	tableTypes["TblFunnelOfertas"] = &models.TblFunnelOfertas{}
	tableTypes["TblFunnelOfertasEstimacion"] = &models.TblFunnelOfertasEstimacion{}
	tableTypes["TblFunnelOfertasEstimacionShadow"] = &models.TblFunnelOfertasEstimacionShadow{}
	tableTypes["TblFunnelRF"] = &models.TblFunnelRF{}
	tableTypes["TblFunnelRFShadow"] = &models.TblFunnelRFShadow{}
	tableTypes["TblFunnelSalesObjective"] = &models.TblFunnelSalesObjective{}
	tableTypes["TblLineaDeNegocio"] = &models.TblLineaDeNegocio{}
	tableTypes["TblLnToPlan2020Map"] = &models.TblLnToPlan2020Map{}
	tableTypes["TblMercadoCliente"] = &models.TblMercadoCliente{}
	tableTypes["TblMonth"] = &models.TblMonth{}
	tableTypes["TblNegocio2020"] = &models.TblNegocio2020{}
	tableTypes["TblOfferIsRecurrent"] = &models.TblOfferIsRecurrent{}
	tableTypes["TblOfferLineaNegocio"] = &models.TblOfferLineaNegocio{}
	tableTypes["TblOfferNegocio2020"] = &models.TblOfferNegocio2020{}
	tableTypes["TblOfferPhase"] = &models.TblOfferPhase{}
	tableTypes["TblOfferStateType"] = &models.TblOfferStateType{}
	tableTypes["TblOfferSubcategory"] = &models.TblOfferSubcategory{}
	tableTypes["TblPlan2020ToSubcategoryMap"] = &models.TblPlan2020ToSubcategoryMap{}
	tableTypes["TblProjectBudget"] = &models.TblProjectBudget{}
	tableTypes["TblProjectBudgetsMgt"] = &models.TblProjectBudgetsMgt{}
	tableTypes["TblProjectBudgetsMgtCOPY"] = &models.TblProjectBudgetsMgtCOPY{}
	tableTypes["TblProjectBudgetsShadow"] = &models.TblProjectBudgetsShadow{}
	tableTypes["TblProjectCalcType"] = &models.TblProjectCalcType{}
	tableTypes["TblProjectPhase"] = &models.TblProjectPhase{}
	tableTypes["TblProjectType"] = &models.TblProjectType{}
	tableTypes["TblProjectType"] = &models.TblProjectType{}
	tableTypes["TblProject"] = &models.TblProject{}
	tableTypes["TblProjectsMgt"] = &models.TblProjectsMgt{}
	tableTypes["TblProv"] = &models.TblProv{}
	tableTypes["TblRollingForecast"] = &models.TblRollingForecast{}
	tableTypes["TblRollingForecastMgt"] = &models.TblRollingForecastMgt{}
	tableTypes["TblSendMailBackup"] = &models.TblSendMailBackup{}
	tableTypes["TblSendMailConfig"] = &models.TblSendMailConfig{}
	tableTypes["TblUser"] = &models.TblUser{}
	tableTypes["TblUsers2"] = &models.TblUsers2{}
	tableTypes["TblWebStat"] = &models.TblWebStat{}
	tableTypes["TblYear"] = &models.TblYear{}
	tableTypes["TempTblAging"] = &models.TempTblAging{}

}

type DB interface {
	ExecuteSQL(queryStr string, w http.ResponseWriter) error
}

func GetDb() (db *sql.DB, err error) {
	db, err = sql.Open("mssql", secrets.SQL_CONN_STR)
	return
}

func GetTestDb() (db *sql.DB, err error) {
	db, err = sql.Open("mssql", secrets.SQL_TEST_CONN_STR)
	return
}

// func GetTableNameByStruct(myStruct interface{}) string {
// 	fullName := reflect.TypeOf(myStruct).String()
// 	tblName := strings.ReplaceAll(fullName, "*models.", "")
// 	return tblName
// }

// func ExistsTable(tblName string) bool {
// 	_, ok := tableTypes[tblName]
// 	return ok

// }
// func GetTableFields(myStruct interface{}) []string {
// 	tblName := GetTableNameByStruct(myStruct)
// 	tp := tableColumns[tblName]

// 	fmt.Println("tp field 1 name:")

// 	// tp := models.AccountingFinancialStatementColumns
// 	tpReflected := reflect.ValueOf(&tp).Elem().Type()
// 	// fmt.Printf(tpReflected.Name())

// 	numOfFields := tpReflected.NumField()

// 	fieldNames := make([]string, numOfFields)
// 	for i := 0; i < numOfFields; i++ {
// 		fieldNames[i] = tpReflected.Field(i).Name
// 	}

// 	// fmt.Println(fieldNames)
// 	return fieldNames

// }

func ExecuteSQL(queryStr string, w http.ResponseWriter) error {
	var conn *sql.DB
	var err error
	isTesting := os.Getenv("TESTING")
	if isTesting == "TRUE" {
		conn, err = GetTestDb()
	} else {
		conn, err = GetDb()
	}

	if err != nil {
		log.Fatal("Error while opening database connection:", err.Error())
	}
	defer conn.Close()

	rows, err := conn.Query(queryStr)
	if err != nil {
		// log.Fatal("query: " + queryStr)
		log.Fatal("Query failed:", err.Error())
	}
	defer rows.Close()

	columns, _ := rows.Columns()

	count := len(columns)

	var v []interface{} // `json:"data"`

	for rows.Next() {
		values := make([]interface{}, count)
		valuePtrs := make([]interface{}, count)
		o := orderedmap.New()

		for i := range columns {
			valuePtrs[i] = &values[i]
		}
		if err := rows.Scan(valuePtrs...); err != nil {
			log.Fatal(err)
		}

		for i := range columns {
			o.Set(columns[i], values[i])
		}
		v = append(v, o)
	}
	r := utils.ToJSON(w, v)
	return r
}

func ExecSqlFromFile(fullPath string) (err error) {

	c, ioErr := ioutil.ReadFile(fullPath)
	if ioErr != nil {
		// log.Fatalf("Error: %v", ioErr.Error())
		return errors.New("ExecSqlFromFile --> ioutil.ReadFile(fullPath): " + ioErr.Error())
	}
	sql := string(c)

	conn, err := GetTestDb()
	if err != nil {
		// log.Fatal("Error while opening database connection:", err.Error())
		return errors.New("ExecSqlFromFile --> mydb.GetTestDb(): " + err.Error())
	}
	defer conn.Close()
	_, err = conn.Exec(sql)

	if err != nil {
		// log.Fatalf("Error en ejecutar la consulta: %v", err.Error())
		return errors.New("ExecSqlFromFile --> conn.Exec(sql) " + err.Error())
	}
	return nil
}
