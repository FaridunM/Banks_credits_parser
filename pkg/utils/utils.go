package utils

import (
	"github.com/FaridunM/Banks_credits_parser/pkg/config"
	. "github.com/FaridunM/Banks_credits_parser/pkg/logger"
	"github.com/FaridunM/Banks_credits_parser/pkg/structs"
	"encoding/json"
	"fmt"
	"net/http"
	"strings"

	"github.com/spf13/viper"
)

// if the condition is true, then return a, else return b
func TernarOp[T any](condition bool, a, b T) T {
	if condition {
		return a
	}

	return b
}

// Getting url by credit type
func GetCreditUrls() (creditsUrl config.Credits, err error) {
	viper.SetConfigFile("../credits.json")
	if err = viper.ReadInConfig(); err != nil {
		Log.Error("error in read in json-file:", err)
		return config.Credits{}, err
	}

	if err = viper.Unmarshal(&creditsUrl); err != nil {
		Log.Error("error in unmarshaling:", err)
		return config.Credits{}, err
	}

	return
}

func ModifyCreditText(credit structs.Credit) structs.Credit {
	for key, value := range credit.Collateral {
		value = strings.TrimSpace(value)
		credit.Collateral[key] = value
	}

	for key, value := range credit.Amount {
		value = strings.TrimSpace(value)
		credit.Amount[key] = value
	}

	for key, value := range credit.LoanTerms {
		value = strings.TrimSpace(value)
		credit.LoanTerms[key] = value
	}

	for key, value := range credit.NeedDocuments {
		value = strings.TrimSpace(value)
		credit.NeedDocuments[key] = value
	}

	for key, value := range credit.InterestRate {
		value = strings.TrimSpace(value)
		credit.NeedDocuments[key] = value
	}

	return credit
}

// WriteJson write structure or base type to http.ResponseWriter
func WriteJson(w http.ResponseWriter, structure any) error {
	defer PanicCatcher()

	body, err := json.Marshal(structure)
	if err != nil {
		return err
	}

	fmt.Fprint(w, string(body))

	return nil
}

// Cather panic and write to log. Need it to keep the server from crashing
func PanicCatcher() {
	if err := recover(); err != nil {
		Log.Panic(err)
		// log.Println("panic:", err)

		var w http.ResponseWriter
		w.WriteHeader(http.StatusInternalServerError)
		WriteJson(w, "Internal server error")
	}
}
