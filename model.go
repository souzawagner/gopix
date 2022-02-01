package main

import (
	"fmt"
	"pix/v1/utils"
)

type Pix struct {
	Key         string
	Description string
	TxId        string
	Amount      float64
	Merchant    Merchant
}

type Merchant struct {
	Name string
	City string
}

func (payload Pix) GetAmount() string {
	return utils.NumberFormat(payload.Amount, 2, ".", "")
}

func (payload Pix) GetPayload() string {
	var result = fmt.Sprintf(
		"%s%s%s%s%s%s%s%s%s",
		utils.FormatValue(PayloadFormatIndicator, DefaultPayloadFormatIndicator),

		// Merchant
		payload.GetMerchantAccountInformation(),

		utils.FormatValue(MerchantCategoryCode, DefaultMerchantCategoryCode),

		// Transaction
		utils.FormatValue(TransactionCurrency, BrazilianRealCurrency),
		utils.FormatValue(TransactionAmount, payload.GetAmount()),
		utils.FormatValue(CountryCode, BrazilianCountryCode),

		// Merchant
		utils.FormatValue(MerchantName, payload.Merchant.Name),
		utils.FormatValue(MerchantCity, payload.Merchant.City),

		payload.GetAdditionalDataFieldTemplate(),
	)
	return fmt.Sprintf("%s%s", result, payload.GetCheckSum(result))
}

func (payload Pix) GetMerchantAccountInformation() string {
	var gui = utils.FormatValue(MerchantAccountInformationGui, DefaultMerchantAccountInformationGui)
	var key = utils.FormatValue(MerchantAccountInformationKey, payload.Key)
	var description = utils.FormatValue(MerchantAccountInformationDescription, payload.Description)

	return utils.FormatValue(
		MerchantAccountInformation,
		fmt.Sprintf("%s%s%s", gui, key, description),
	)
}

func (payload Pix) GetAdditionalDataFieldTemplate() string {
	var txId = utils.FormatValue(AdditionalDataFieldTemplateTxid, payload.TxId)
	return utils.FormatValue(AdditionalDataFieldTemplate, txId)
}

func (_ Pix) GetCheckSum(payload string) string {
	payload = fmt.Sprintf("%s%s%s", payload, CRC16, "04")
	return utils.FormatValue(
		CRC16,
		utils.GetCRC16Checksum(payload),
	)
}
