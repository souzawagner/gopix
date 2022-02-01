package main

import (
	"fmt"
	qr "github.com/skip2/go-qrcode"
	"os"
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

func (pix Pix) GetAmount() string {
	return utils.NumberFormat(pix.Amount, 2, ".", "")
}

func (pix Pix) String() string {
	result := fmt.Sprintf(
		"%s%s%s%s%s%s%s%s%s",
		utils.FormatValue(PayloadFormatIndicator, DefaultPayloadFormatIndicator),

		// Merchant
		pix.GetMerchantAccountInformation(),

		utils.FormatValue(MerchantCategoryCode, DefaultMerchantCategoryCode),

		// Transaction
		utils.FormatValue(TransactionCurrency, BrazilianRealCurrency),
		utils.FormatValue(TransactionAmount, pix.GetAmount()),
		utils.FormatValue(CountryCode, BrazilianCountryCode),

		// Merchant
		utils.FormatValue(MerchantName, pix.Merchant.Name),
		utils.FormatValue(MerchantCity, pix.Merchant.City),

		pix.GetAdditionalDataFieldTemplate(),
	)
	return fmt.Sprintf("%s%s", result, pix.GetCheckSum(result))
}

func (pix Pix) QRCode(level qr.RecoveryLevel, size int) ([]byte, error) {
	qrcode, err := qr.Encode(pix.String(), level, size)
	return qrcode, err
}

func (pix Pix) File(path string) error {
	var qrcode, _ = pix.QRCode(qr.Medium, 256)
	file, err := os.Create(path)
	if err != nil {
		return err
	}
	_, err = file.Write(qrcode)
	if err != nil {
		return err
	}
	err = file.Sync()
	if err != nil {
		return err
	}
	return nil
}

func (pix Pix) GetMerchantAccountInformation() string {
	gui := utils.FormatValue(MerchantAccountInformationGui, DefaultMerchantAccountInformationGui)
	key := utils.FormatValue(MerchantAccountInformationKey, pix.Key)
	description := utils.FormatValue(MerchantAccountInformationDescription, pix.Description)

	return utils.FormatValue(
		MerchantAccountInformation,
		fmt.Sprintf("%s%s%s", gui, key, description),
	)
}

func (pix Pix) GetAdditionalDataFieldTemplate() string {
	txId := utils.FormatValue(AdditionalDataFieldTemplateTxid, pix.TxId)
	return utils.FormatValue(AdditionalDataFieldTemplate, txId)
}

func (_ Pix) GetCheckSum(pix string) string {
	pix = fmt.Sprintf("%s%s%s", pix, CRC16, "04")
	return utils.FormatValue(
		CRC16,
		utils.GetCRC16Checksum(pix),
	)
}
