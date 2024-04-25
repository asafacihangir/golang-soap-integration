package main

import (
	"fmt"
	"github.com/hooklift/gowsdl/soap"
	"golang-soap-integration/gen/country"
	"log"
)

func main() {

	countryServiceClient := soap.NewClient("http://webservices.oorsprong.org/websamples.countryinfo/CountryInfoService.wso")
	countryService := country.NewCountryInfoServiceSoapType(countryServiceClient)

	fullCountryInfoRequest := &country.CountriesUsingCurrency{SISOCurrencyCode: "USD"}
	fullCountryInfoResponse, err := countryService.CountriesUsingCurrency(fullCountryInfoRequest)
	if err != nil {
		log.Fatalf("Error calling FullCountryInfo: %v", err)
	}

	for i := range fullCountryInfoResponse.CountriesUsingCurrencyResult.TCountryCodeAndName {
		fmt.Printf("Country Code: %s, Country Name: %s\n",
			fullCountryInfoResponse.CountriesUsingCurrencyResult.TCountryCodeAndName[i].SISOCode,
			fullCountryInfoResponse.CountriesUsingCurrencyResult.TCountryCodeAndName[i].SName)
	}

}
