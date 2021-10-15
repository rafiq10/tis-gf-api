package main

import (
	"context"
	"log"
	"net/http"
	"os"
	"os/signal"
	"time"

	"tis-gf-api/api/config"
	ha "tis-gf-api/api/handlers/accounting"

	hu "tis-gf-api/api/handlers/fileupload"
	hr "tis-gf-api/api/handlers/reports"

	// hu "tis-gf-api/api/fileupload"

	"github.com/gorilla/mux"
)

func main() {
	l := log.New(os.Stdout, "tis-gf-api", log.LstdFlags)
	hh := ha.NewAccounting(l)
	myReps := hr.NewReports(l)
	repsAging := hr.NewReportsAgingCLI(l)
	repsDCC := hr.NewReportsDCC(l)
	fileUp := hu.NewFileUploader(l)

	sm := mux.NewRouter()

	getRouter := sm.Methods(http.MethodGet).Subrouter()
	getRouter.HandleFunc(config.API_VERSION+"/accounting/financial-statements", hh.GetAllFS)
	getRouter.HandleFunc(config.API_VERSION+"/reports/mgt/reports-loaded-resume", myReps.GetReportsLoadedResume)
	getRouter.HandleFunc(config.API_VERSION+"/reports/mgt/reports-aging-cli", repsAging.GetAgingCLI)
	getRouter.HandleFunc(config.API_VERSION+"/reports/mgt/reports-dcc", repsDCC.GetDCC)

	postFormRouter := sm.Methods(http.MethodPost).Subrouter()
	postFormRouter.HandleFunc(config.API_VERSION+"/upload", fileUp.FileUpload)
	postFormRouter.HandleFunc(config.API_VERSION+"/accounting/financial-statements", hh.AddFinancialStatement)

	s := &http.Server{
		Addr:         ":8099",
		Handler:      sm,
		IdleTimeout:  120 * time.Second,
		ReadTimeout:  60 * time.Second,
		WriteTimeout: 60 * time.Second,
	}
	gracefulShutdown(s, l)

}

func gracefulShutdown(s *http.Server, l *log.Logger) {
	go func() {
		err := s.ListenAndServe()
		if err != nil {
			l.Fatal(err)
		}
	}()

	sigChan := make(chan os.Signal)
	signal.Notify(sigChan, os.Interrupt)
	signal.Notify(sigChan, os.Kill)
	sig := <-sigChan
	l.Println("Received terminate shutdown", sig)

	tc, _ := context.WithTimeout(context.Background(), 30*time.Second)
	s.Shutdown(tc)
}
