package handler

import (
	"github.com/FaridunM/Banks_credits_parser/pkg/gateway"
	"github.com/FaridunM/Banks_credits_parser/pkg/structs"
	"github.com/FaridunM/Banks_credits_parser/pkg/utils"
	"fmt"
	"net/http"
	"regexp"

	"github.com/gorilla/mux"
)

func (h *Handler) Pong(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("Pong"))
}

func (h *Handler) Handle404(w http.ResponseWriter, r *http.Request) {
	w.WriteHeader(http.StatusNotFound)
	fmt.Fprintln(w, "Sorry, page not found.")
}

func (h *Handler) Handle405(w http.ResponseWriter, r *http.Request) {
	w.WriteHeader(http.StatusMethodNotAllowed)
	fmt.Fprintln(w, "Sorry, method not allowed.")
}

// func (h *Handler) GetCredits(w http.ResponseWriter, r *http.Request) {
// 	log.Println("Eskhata Credits:", h.credits.Eskhata.CreditsTitles)
// 	log.Println("Humo Credits:", h.credits.Humo.CreditsTitles)
// 	w.Write([]byte("GetCredits"))
// }

func (h *Handler) GetCreditsBy(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	bank := vars["bank"]
	var (
		titles  string
		credits []string
	)

	switch {
	case regexp.MustCompile(`^H|humo`).MatchString(bank):
		credits = h.credits.Humo.CreditsTitles
	case regexp.MustCompile(`^E|es(h|kh)ata`).MatchString(bank):
		credits = h.credits.Eskhata.CreditsTitles
	default:
		w.Write([]byte(fmt.Sprintf("%q", "Bank not found")))
		return
	}

	for _, title := range credits {
		titles += `"` + title + `", `
	}
	w.Write([]byte(fmt.Sprintf("{%v}", titles)))
}

func (h *Handler) GetCredit(w http.ResponseWriter, r *http.Request) {
	// fmt.Println("h.credits.Humo -", h.credits.Humo)
	vars := mux.Vars(r)
	bank := vars["bank"]
	creditType := vars["credit_type"]

	var (
		humoCredit    structs.Credit
		eskhataCredit structs.Credit
	)

	switch {
	case regexp.MustCompile(`^H|humo`).MatchString(bank):
		url, err := gateway.GetHumoUrlBy(creditType)
		if err != nil {
			h.logger.Error(err)
		}

		switch creditType {
		case "consumer", "start_business":
			humoCredit = gateway.GetHumoStandartCredit(creditType, url, h.credits.Humo.ConsumerAndBusinessPrompts)
		case "mortgage":
			humoCredit = gateway.GetHumoStandartCredit(creditType, url, h.credits.Humo.MortgagePrompts)
		case "education":
			humoCredit = gateway.GetHumoStandartCredit(creditType, url, h.credits.Humo.EducationPrompts)
		case "agricultural", "livestock":
			humoCredit = gateway.GetHumoStandartCredit(creditType, url, h.credits.Humo.AgriculturalAndLivestockPrompts)
		case "orzu":
			humoCredit = gateway.GetHumoOrzuCredit(url)
		default:
			w.Write([]byte(fmt.Sprintf("%q", "Credit not found")))
			return
		}

		utils.WriteJson(w, humoCredit)
	case regexp.MustCompile(`^E|es(h|kh)ata`).MatchString(bank):
		url, anchor, err := gateway.GetEskhataUrlBy(creditType)
		if err != nil {
			h.logger.Error(err)
		}

		switch creditType {
		case "multi_purpose_consumer":
			eskhataCredit = gateway.GetEskhataCredit(url, anchor, h.credits.Eskhata.MultiPurposeConsumerPrompts)
		case "express", "hunarkhoi_mardumi", "car":
			eskhataCredit = gateway.GetEskhataCredit2(url, anchor, h.credits.Eskhata.PromptsForExpressHunarkhoiMardumiCar)
		case "manzili":
			eskhataCredit = gateway.GetEskhataCredit3(url, anchor, h.credits.Eskhata.ManziliPrompts)
		default:
			w.Write([]byte(fmt.Sprintf("%q", "Credit not found")))
			return
		}

		utils.WriteJson(w, eskhataCredit)
	default:
		w.Write([]byte(fmt.Sprintf("%q", "Bank not found")))
		return
	}
}
