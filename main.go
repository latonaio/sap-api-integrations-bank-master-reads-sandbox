package main

import (
	sap_api_caller "sap-api-integrations-bank-master-reads/SAP_API_Caller"
	"sap-api-integrations-bank-master-reads/sap_api_input_reader"

	"github.com/latonaio/golang-logging-library/logger"
)

func main() {
	l := logger.NewLogger()
	fr := sap_api_input_reader.NewFileReader()
	inoutSDC := fr.ReadSDC("./Inputs//SDC_Bank_Master_Bank_sample.json")
	caller := sap_api_caller.NewSAPAPICaller(
		"https://sandbox.api.sap.com/s4hanacloud/sap/opu/odata4/sap/", l,
	)

	accepter := inoutSDC.Accepter
	if len(accepter) == 0 || accepter[0] == "All" {
		accepter = []string{
			"Bank",
		}
	}

	caller.AsyncGetBankMaster(
		inoutSDC.Bank.BankCountry,
		inoutSDC.Bank.Bank,
		accepter,
	)
}
