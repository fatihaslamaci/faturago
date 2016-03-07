package main

//
import (
	"encoding/xml"
	"time"
)

type IssueDateType struct {
	DateType
}

type DateType struct {
	time.Time
}

type TimeType struct {
	time.Time
}
type IssueTimeType struct {
	TimeType
}

func (c *DateType) UnmarshalXML(d *xml.Decoder, start xml.StartElement) error {
	const shortForm = "2006-01-02" // yyyy-mm-dd date format
	var v string
	d.DecodeElement(&v, &start)
	parse, err := time.Parse(shortForm, v)
	if err != nil {
		return err
	}
	*c = DateType{parse}
	return nil
}

func (c *TimeType) UnmarshalXML(d *xml.Decoder, start xml.StartElement) error {
	const shortForm = "15:04:05" // hh-MM-ss time format
	var v string
	d.DecodeElement(&v, &start)
	parse, err := time.Parse(shortForm, v)
	if err != nil {
		return err
	}
	*c = TimeType{parse}
	return nil
}

type Invoice struct {
	XMLName xml.Name `xml:"Invoice"`
	//private UBLExtensionType[] uBLExtensionsField;
	UBLVersionID    UBLVersionIDType
	CustomizationID CustomizationIDType

	ProfileID       ProfileIDType
	ID              IDType
	CopyIndicator   CopyIndicatorType
	UUID            UUIDType
	IssueDate       IssueDateType
	IssueTime       IssueTimeType
	InvoiceTypeCode string
	//private NoteType[] noteField;
	DocumentCurrencyCode string
	//private TaxCurrencyCodeType taxCurrencyCodeField;
	//private PricingCurrencyCodeType pricingCurrencyCodeField;
	//private PaymentCurrencyCodeType paymentCurrencyCodeField;
	//private PaymentAlternativeCurrencyCodeType paymentAlternativeCurrencyCodeField;
	//private AccountingCostType accountingCostField;

	LineCountNumeric        int
	Signature               []SignatureType   `xml:",omitempty"`
	AccountingSupplierParty SupplierPartyType `xml:",omitempty"`

	AccountingCustomerParty CustomerPartyType `xml:",omitempty"`
	BuyerCustomerParty      CustomerPartyType `xml:",omitempty"`
	SellerSupplierParty     SupplierPartyType `xml:",omitempty"`
	TaxRepresentativeParty  PartyType         `xml:",omitempty"`
	//delivery  DeliveryType[]
	//paymentMeans PaymentMeansType[]
	//paymentTerms PaymentTermsType
	//allowanceCharge AllowanceChargeType[]
	//taxExchangeRate ExchangeRateType
	//pricingExchangeRate ExchangeRateType
	//paymentExchangeRate ExchangeRateType
	//paymentAlternativeExchangeRate ExchangeRateType
	//taxTotal TaxTotalType[]
	//withholdingTaxTotal TaxTotalType[]
	//legalMonetaryTotal MonetaryTotalType
	//invoiceLine InvoiceLineType[]

}

type IdentifierType struct {
	SchemeID         string `xml:"schemeID,attr,omitempty"`
	SchemeName       string `xml:"schemeName,attr,omitempty"`
	SchemeAgencyID   string `xml:"schemeAgencyID,attr,omitempty"`
	SchemeAgencyName string `xml:"schemeAgencyName,attr,omitempty"`
	SchemeVersionID  string `xml:"schemeVersionID,attr,omitempty"`
	SchemeDataURI    string `xml:"schemeDataURI,attr,omitempty"`
	SchemeURI        string `xml:"schemeURI,attr,omitempty"`
	Value            string `xml:",chardata"`
}

type SignatureType struct {
	ID IDType

	SignatoryParty PartyType `xml:",omitempty"`

	//private AttachmentType digitalSignatureAttachmentField;
}

type PartyType struct {
	WebsiteURI string `xml:",omitempty"`

	//private EndpointIDType endpointIDField;
	//private IndustryClassificationCodeType industryClassificationCodeField;
	//PartyIdentification []PartyIdentificationType
	//private PartyNameType partyNameField;
	//private AddressType postalAddressField;
	//private PartyTaxSchemeType partyTaxSchemeField;
	//private PartyLegalEntityType[] partyLegalEntityField;
	//private ContactType contactField;
	//private PersonType personField;
	//private PartyType agentPartyField;

}

type SupplierPartyType struct {
	Party PartyType `xml:",omitempty"`
}
type CustomerPartyType struct {
	Party PartyType `xml:",omitempty"`
}

type IdentifierType1 struct {
	IdentifierType
}

type UBLVersionIDType struct {
	IdentifierType1
}
type CustomizationIDType struct {
	IdentifierType1
}

type ProfileIDType struct {
	IdentifierType1
}

type IDType struct {
	IdentifierType1
}
type UUIDType struct {
	IdentifierType1
}

type CopyIndicatorType struct {
	IndicatorType
}

type IndicatorType struct {
	Value bool `xml:",chardata"`
}
