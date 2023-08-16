package gateway

import (
	. "github.com/FaridunM/Banks_credits_parser/pkg/logger"
	"github.com/FaridunM/Banks_credits_parser/pkg/utils"
	"fmt"
)

// Getting url by credit type
func GetHumoUrlBy(creditType string) (url string, err error) {
	creditUrls, err := utils.GetCreditUrls()
	if err != nil {
		Log.Error("Error in getting credit urls")
	}

	switch creditType {
	case "consumer":
		url = creditUrls.Humo.Consumer
	case "start_business":
		url = creditUrls.Humo.Start_Business
	case "mortgage":
		url = creditUrls.Humo.Mortgage
	case "education":
		url = creditUrls.Humo.Education
	case "agricultural":
		url = creditUrls.Humo.Agricultural
	case "livestock":
		url = creditUrls.Humo.Livestock
	case "orzu":
		url = creditUrls.Humo.Orzu
	default:
		url = ""
		Log.Errorf("creditType %s not found", creditType)
		err = fmt.Errorf("creditType %s not found", creditType)
	}

	return
}

// Getting url and anchor by credit type
func GetEskhataUrlBy(creditType string) (url, anchor string, err error) {
	creditUrls, err := utils.GetCreditUrls()
	if err != nil {
		Log.Println("Error in getting credit urls")
	}
	url = creditUrls.Eskhata.CommonUrl

	switch creditType {
	case "multi_purpose_consumer":
		anchor = creditUrls.Eskhata.Multi_Purpose_Consumer
	case "express":
		anchor = creditUrls.Eskhata.Express
	case "hunarkhoi_mardumi":
		anchor = creditUrls.Eskhata.Hunarkhoi_Mardumi
	case "car":
		anchor = creditUrls.Eskhata.Car
	case "manzili":
		anchor = creditUrls.Eskhata.Manzili
	default:
		anchor = ""
		Log.Errorf("creditType %s not found", creditType)
		err = fmt.Errorf("creditType %s not found", creditType)
	}

	return
}
