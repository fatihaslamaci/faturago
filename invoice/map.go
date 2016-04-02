package invoice

import (
	"faturago/invoice/serialized"
)

func InvoiceMap(V *InvoiceType) serialized.InvoiceType {
	var invoice2 serialized.InvoiceType

	invoice2.UBLVersionID = IdentifierType1Map(V.UBLVersionID)
	invoice2.CustomizationID = IdentifierType1Map(V.CustomizationID)
	invoice2.ProfileID = IdentifierType1Map(V.ProfileID)
	invoice2.ID = IdentifierType1Map(V.ID)
	invoice2.CopyIndicator = IndicatorTypeMap(V.CopyIndicator)

	invoice2.UUID = IdentifierType1Map(V.UUID)
	invoice2.IssueDate = DateTypeMap(V.IssueDate)
	invoice2.IssueTime = TimeTypeMap(V.IssueTime)
	invoice2.InvoiceTypeCode = CodeTypeMap(V.InvoiceTypeCode)
	invoice2.Note = TextTypeListMap(V.Note)
	invoice2.DocumentCurrencyCode = CodeTypeMap(V.DocumentCurrencyCode)
	invoice2.TaxCurrencyCode = CodeTypeMap(V.TaxCurrencyCode)
	invoice2.PricingCurrencyCode = CodeTypeMap(V.PricingCurrencyCode)
	invoice2.PaymentCurrencyCode = CodeTypeMap(V.PaymentCurrencyCode)
	invoice2.PaymentAlternativeCurrencyCode = CodeTypeMap(V.PaymentAlternativeCurrencyCode)
	invoice2.AccountingCost = TextTypeMap(V.AccountingCost)
	invoice2.LineCountNumeric = NumericTypeMap(V.LineCountNumeric)
	invoice2.InvoicePeriod = PeriodTypeMap(V.InvoicePeriod)
	invoice2.BillingReference = BillingReferenceTypeMapList(V.BillingReference)
	//invoice2.DespatchDocumentReference = []DocumentReferenceTypeMap(V.DespatchDocumentReference)
	//invoice2.ReceiptDocumentReference = []DocumentReferenceTypeMap(V.ReceiptDocumentReference)
	//invoice2.OriginatorDocumentReference = []DocumentReferenceTypeMap(V.OriginatorDocumentReference)
	//invoice2.ContractDocumentReference = []DocumentReferenceTypeMap(V.ContractDocumentReference)
	//invoice2.AdditionalDocumentReference = []DocumentReferenceTypeMap(V.AdditionalDocumentReference)
	//invoice2.Signature = []SignatureTypeMap(V.Signature)
	//invoice2.AccountingSupplierParty = *SupplierPartyTypeMap(V.AccountingSupplierParty)
	//invoice2.AccountingCustomerParty = *CustomerPartyTypeMap(V.AccountingCustomerParty)
	//invoice2.BuyerCustomerParty = *CustomerPartyTypeMap(V.BuyerCustomerParty)
	//invoice2.SellerSupplierParty = *SupplierPartyTypeMap(V.SellerSupplierParty)
	//invoice2.TaxRepresentativeParty = *PartyTypeMap(V.TaxRepresentativeParty)
	//invoice2.Delivery = []DeliveryTypeMap(V.Delivery)
	//invoice2.PaymentMeans = []PaymentMeansTypeMap(V.PaymentMeans)
	//invoice2.PaymentTerms = *PaymentTermsTypeMap(V.PaymentTerms)
	//invoice2.AllowanceCharge = []AllowanceChargeTypeMap(V.AllowanceCharge)
	//invoice2.TaxExchangeRate = *ExchangeRateTypeMap(V.TaxExchangeRate)
	//invoice2.PricingExchangeRate = *ExchangeRateTypeMap(V.PricingExchangeRate)
	//invoice2.PaymentExchangeRate = *ExchangeRateTypeMap(V.PaymentExchangeRate)
	//invoice2.PaymentAlternativeExchangeRate = *ExchangeRateTypeMap(V.PaymentAlternativeExchangeRate)
	//invoice2.TaxTotal = []TaxTotalTypeMap(V.TaxTotal)
	//invoice2.WithholdingTaxTotal = []TaxTotalTypeMap(V.WithholdingTaxTotal)
	//invoice2.LegalMonetaryTotal = *MonetaryTotalTypeMap(V.LegalMonetaryTotal)
	//invoice2.InvoiceLine = []InvoiceLineTypeMap(V.InvoiceLine)

	return invoice2
}

func BillingReferenceTypeMapList(v []BillingReferenceType) []serialized.BillingReferenceType {
	if v == nil {
		return nil
	}
	//TODO map yapılacak
	return nil
}

func PeriodTypeMap(v *PeriodType) *serialized.PeriodType {
	if v == nil {
		return nil
	}
	R := &serialized.PeriodType{}
	R.Description = TextTypeMap(v.Description)
	//TODO DurationMeasure yapılmadı
	//R.DurationMeasure = PeriodTypeMap(v.DurationMeasure)
	R.EndDate = DateTypeMap(v.EndDate)
	R.EndTime = TimeTypeMap(v.EndTime)
	R.StartDate = DateTypeMap(v.StartDate)
	R.StartTime = TimeTypeMap(v.StartTime)

	return R
}

func NumericTypeMap(v *NumericType) *serialized.NumericType {
	if v == nil {
		return nil
	}
	R := &serialized.NumericType{}
	R.Format = v.Format
	R.Value = v.Value

	return R
}

func TextTypeListMap(v []TextType) []serialized.TextType {
	if v == nil {
		return nil
	}

	var j = len(v)

	R := make([]serialized.TextType, j)

	for i := range v {
		R[i].Value = v[i].Value
		R[i].LanguageLocaleID = v[i].LanguageLocaleID
		R[i].LanguageID = v[i].LanguageID
	}

	return R
}

func TextTypeMap(v *TextType) *serialized.TextType {
	if v == nil {
		return nil
	}
	R := &serialized.TextType{}
	R.LanguageID = v.LanguageID
	R.LanguageLocaleID = v.LanguageLocaleID
	R.Value = v.Value

	return R
}

func CodeTypeMap(v *CodeType) *serialized.CodeType {
	if v == nil {
		return nil
	}
	R := &serialized.CodeType{}
	R.Value = v.Value
	R.LanguageID = v.LanguageID
	R.ListAgencyID = v.ListAgencyID
	R.ListID = v.ListID
	R.ListName = v.ListName
	R.ListSchemeURI = v.ListSchemeURI
	R.ListURI = v.ListURI
	R.ListVersionID = v.ListVersionID
	R.Name = v.Name

	return R
}

func IdentifierType1Map(v *IdentifierType1) *serialized.IdentifierType1 {

	if v == nil {
		return nil
	}

	R := &serialized.IdentifierType1{}

	R.Value = v.Value
	R.SchemeAgencyID = v.SchemeAgencyID
	R.SchemeAgencyName = v.SchemeAgencyName
	R.SchemeDataURI = v.SchemeDataURI
	R.SchemeID = v.SchemeID
	R.SchemeName = v.SchemeName
	R.SchemeURI = v.SchemeURI
	R.SchemeVersionID = v.SchemeVersionID

	return R
}

func IndicatorTypeMap(v *IndicatorType) *serialized.IndicatorType {
	if v == nil {
		return nil
	}
	R := &serialized.IndicatorType{}
	R.Value = v.Value
	return R
}
func DateTypeMap(v *DateType) *serialized.DateType {
	if v == nil {
		return nil
	}
	R := &serialized.DateType{}
	R.Value = v.Value
	return R
}
func TimeTypeMap(v *TimeType) *serialized.TimeType {
	if v == nil {
		return nil
	}
	R := &serialized.TimeType{}
	R.Value = v.Value
	return R
}
