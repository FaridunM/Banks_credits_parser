package config

import (
	"github.com/spf13/viper"
	"go.uber.org/fx"
)

type CreditsPrompts struct {
	Title                     string `json:"title"`
	Description               string `json:"description,omitempty"`
	Purpose                   string `json:"purpose,omitempty"`
	LandingMethodology        string `json:"landingMethodology,omitempty"`
	Collateral                string `json:"collateral,omitempty"`
	Currency                  string `json:"currency,omitempty"`
	Amount                    string `json:"amount"`
	InterestRate              string `json:"interestRate,omitempty"`
	LoanPeriod                string `json:"loanPeriod"`
	ItinialPayment            string `json:"itinialPayment,omitempty"`
	GracePeriod               string `json:"gracePeriod,omitempty"`
	LoanTerms                 string `json:"loanTerms"`
	NeedDocuments             string `json:"needDocuments"`
	PurposeAmountLoanPeriod   string `json:"purposeAmountLoanPeriod,omitempty"`
	InterestRateNeedDocuments string `json:"interestRateNeedDocuments,omitempty"`
	AdditionalField           string `json:"additionalField,omitempty"`
}

type HumoCreditsUrl struct {
	CreditsTitles                   []string       `json:"creditsTitles"`
	Consumer                        string         `json:"consumer"`
	Start_Business                  string         `json:"start_business"`
	ConsumerAndBusinessPrompts      CreditsPrompts `json:"consumerAndBusinessPrompts"`
	Mortgage                        string         `json:"mortgage"`
	MortgagePrompts                 CreditsPrompts `json:"mortgagePrompts"`
	Education                       string         `json:"education"`
	EducationPrompts                CreditsPrompts `json:"educationPrompts"`
	Agricultural                    string         `json:"agricultural"`
	Livestock                       string         `json:"livestock"`
	AgriculturalAndLivestockPrompts CreditsPrompts `json:"agriculturalAndLivestockPrompts"`
	Orzu                            string         `json:"orzu"`
}

type EskhataCreditsUrlAnchors struct {
	CreditsTitles                        []string       `json:"creditsTitles"`
	CommonUrl                            string         `json:"commonUrl"`
	Multi_Purpose_Consumer               string         `json:"multi_purpose_consumer"`
	MultiPurposeConsumerPrompts          CreditsPrompts `json:"multiPurposeConsumerPrompts"`
	Express                              string         `json:"express"`
	Hunarkhoi_Mardumi                    string         `json:"hunarkhoi_mardumi"`
	Car                                  string         `json:"car"`
	PromptsForExpressHunarkhoiMardumiCar CreditsPrompts `json:"promptsForExpressHunarkhoiMardumiCar"`
	Manzili                              string         `json:"manzili"`
	ManziliPrompts                       CreditsPrompts `json:"manziliPrompts"`
}

type Credits struct {
	Humo    HumoCreditsUrl           `json:"humo"`
	Eskhata EskhataCreditsUrlAnchors `json:"eskhata"`
}

func ImportCreditsUrl() *Credits {
	viper.SetConfigFile("../credits.json")

	err := viper.ReadInConfig()
	if err != nil {
		panic(err)
	}

	var credits *Credits

	err = viper.Unmarshal(&credits)

	if err != nil {
		panic(err)
	}

	return credits
}

var CreditModule = fx.Provide(ImportCreditsUrl)
