package main

const (
	PayloadFormatIndicator                = "00"
	MerchantAccountInformation            = "26"
	MerchantAccountInformationGui         = "00"
	MerchantAccountInformationKey         = "01"
	MerchantAccountInformationDescription = "02"
	MerchantCategoryCode                  = "52"
	TransactionCurrency                   = "53"
	TransactionAmount                     = "54"
	CountryCode                           = "58"
	MerchantName                          = "59"
	MerchantCity                          = "60"
	AdditionalDataFieldTemplate           = "62"
	AdditionalDataFieldTemplateTxid       = "05"
	CRC16                                 = "63"
)

const (
	BrazilianRealCurrency = "986"
	BrazilianCountryCode  = "BR"
)

const (
	DefaultPayloadFormatIndicator        = "01"
	DefaultMerchantCategoryCode          = "0000"
	DefaultMerchantAccountInformationGui = "br.gov.bcb.pix"
)
