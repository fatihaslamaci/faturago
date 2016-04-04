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
	invoice2.DespatchDocumentReference = DocumentReferenceTypeListMap(V.DespatchDocumentReference)
	invoice2.ReceiptDocumentReference = DocumentReferenceTypeListMap(V.ReceiptDocumentReference)
	invoice2.OriginatorDocumentReference = DocumentReferenceTypeListMap(V.OriginatorDocumentReference)
	invoice2.ContractDocumentReference = DocumentReferenceTypeListMap(V.ContractDocumentReference)
	invoice2.AdditionalDocumentReference = DocumentReferenceTypeListMap(V.AdditionalDocumentReference)
	invoice2.Signature = SignatureTypeListMap(V.Signature)
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

func SignatureTypeListMap(v []SignatureType) []serialized.SignatureType {
	if v == nil {
		return nil
	}
	var R []serialized.SignatureType

	for i := range v {
		R = append(R, *SignatureTypeMap(&v[i]))
	}
	return R
}

func SignatureTypeMap(v *SignatureType) *serialized.SignatureType {
	if v == nil {
		return nil
	}

	R := &serialized.SignatureType{}
	R.ID = IdentifierType1Map(v.ID)
	//TODO yapılacak
	//R.DigitalSignatureAttachment = AttachmentTypeMap(v.DigitalSignatureAttachment)
	R.SignatoryParty = PartyTypeMap(v.SignatoryParty)
	return R
}

func PartyTypeMap(v *PartyType) *serialized.PartyType {
	if v == nil {
		return nil
	}

	R := &serialized.PartyType{}

	R.WebsiteURI = IdentifierType1Map(v.WebsiteURI)
	R.EndpointID = IdentifierType1Map(v.EndpointID)
	R.IndustryClassificationCode = CodeTypeMap(v.IndustryClassificationCode)
	R.PartyIdentification = PartyIdentificationTypeListMap(v.PartyIdentification)
	R.PartyName = PartyNameTypeMap(v.PartyName)
	R.PostalAddress = AddressTypeMap(v.PostalAddress)
	//R.PartyTaxScheme            =PartyTaxSchemeTypeMap(v.PartyTaxScheme)
	//R.PartyLegalEntity          =PartyLegalEntityTypeListMap(v.PartyLegalEntity)
	//R.Contact                   =ContactTypeMap(v.Contact)
	//R.Person                    =PersonTypeMap(v.Person)
	//R.AgentParty                =PartyTypeMap(v.AgentParty)

	return R
}

func AddressTypeMap(v *AddressType) *serialized.AddressType {
	if v == nil {
		return nil
	}
	R := &serialized.AddressType{}
	R.ID = IdentifierType1Map(v.ID)
	R.Postbox = TextTypeMap(v.Postbox)
	R.Room = TextTypeMap(v.Room)
	R.StreetName = TextTypeMap(v.StreetName)
	R.BlockName = TextTypeMap(v.BlockName)
	R.BuildingName = TextTypeMap(v.BuildingName)
	R.BuildingNumber = TextTypeMap(v.BuildingNumber)
	R.CitySubdivisionName = TextTypeMap(v.CitySubdivisionName)
	R.CityName = TextTypeMap(v.CityName)
	R.PostalZone = TextTypeMap(v.PostalZone)
	R.Region = TextTypeMap(v.Region)
	R.District = TextTypeMap(v.District)
	R.Country = CountryTypeMap(v.Country)

	return R
}

func CountryTypeMap(v *CountryType) *serialized.CountryType {
	if v == nil {
		return nil
	}
	R := &serialized.CountryType{}
	R.IdentificationCode = CodeTypeMap(v.IdentificationCode)
	R.Name = TextTypeMap(v.Name)

	return R
}

func PartyNameTypeMap(v *PartyNameType) *serialized.PartyNameType {
	if v == nil {
		return nil
	}
	R := &serialized.PartyNameType{}
	R.Name = TextTypeMap(v.Name)

	return R
}

func PartyIdentificationTypeListMap(v []PartyIdentificationType) []serialized.PartyIdentificationType {
	if v == nil {
		return nil
	}

	var R []serialized.PartyIdentificationType

	for i := range v {
		R = append(R, *PartyIdentificationTypeMap(&v[i]))
	}

	return R
}

func PartyIdentificationTypeMap(v *PartyIdentificationType) *serialized.PartyIdentificationType {
	if v == nil {
		return nil
	}
	R := &serialized.PartyIdentificationType{}
	R.ID = IdentifierType1Map(v.ID)
	return R
}

func DocumentReferenceTypeListMap(v []DocumentReferenceType) []serialized.DocumentReferenceType {
	if v == nil {
		return nil
	}
	var R []serialized.DocumentReferenceType
	for i := range v {
		R = append(R, *DocumentReferenceTypeMap(&v[i]))
	}
	return R
}

func DocumentReferenceTypeMap(v *DocumentReferenceType) *serialized.DocumentReferenceType {
	if v == nil {
		return nil
	}

	R := &serialized.DocumentReferenceType{}
	//TODO yapılacak
	//R.Attachment = AttachmentTypeMap(v.Attachment)
	//R.IssuerParty = PartyTypeMap(v.IssuerParty)

	R.DocumentDescription = TextTypeListMap(v.DocumentDescription)
	R.DocumentType = TextTypeMap(v.DocumentType)
	R.DocumentTypeCode = CodeTypeMap(v.DocumentTypeCode)

	R.ValidityPeriod = PeriodTypeMap(v.ValidityPeriod)
	R.IssueDate = DateTypeMap(v.IssueDate)
	R.ID = IdentifierType1Map(v.ID)

	return R
}

func BillingReferenceTypeMapList(v []BillingReferenceType) []serialized.BillingReferenceType {
	if v == nil {
		return nil
	}

	var R []serialized.BillingReferenceType

	for i := range v {
		R = append(R, *BillingReferenceTypeMap(&v[i]))
	}

	return R
}

func BillingReferenceTypeMap(v *BillingReferenceType) *serialized.BillingReferenceType {
	if v == nil {
		return nil
	}

	R := &serialized.BillingReferenceType{}
	//TODO Map işlemi yapılacak

	return R
}

func PeriodTypeMap(v *PeriodType) *serialized.PeriodType {
	if v == nil {
		return nil
	}
	R := &serialized.PeriodType{}
	R.Description = TextTypeMap(v.Description)
	R.DurationMeasure = MeasureTypeMap(v.DurationMeasure)
	R.EndDate = DateTypeMap(v.EndDate)
	R.EndTime = TimeTypeMap(v.EndTime)
	R.StartDate = DateTypeMap(v.StartDate)
	R.StartTime = TimeTypeMap(v.StartTime)

	return R
}

func MeasureTypeMap(v *MeasureType) *serialized.MeasureType {
	if v == nil {
		return nil
	}
	R := &serialized.MeasureType{}
	R.Value = v.Value
	R.UnitCode = v.unitCode
	R.UnitCodeListVersionID = v.unitCodeListVersionID

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

	var R []serialized.TextType

	for i := range v {
		R = append(R, *TextTypeMap(&v[i]))
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
