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
	UBLVersionID    IdentifierType
	CustomizationID IdentifierType

	ProfileID            IdentifierType
	ID                   IdentifierType
	CopyIndicator        CopyIndicatorType
	UUID                 IdentifierType
	IssueDate            IssueDateType
	IssueTime            IssueTimeType
	InvoiceTypeCode      CodeType
	Note                 []TextType
	DocumentCurrencyCode CodeType
	//private TaxCurrencyCodeType taxCurrencyCodeField;
	//private PricingCurrencyCodeType pricingCurrencyCodeField;
	//private PaymentCurrencyCodeType paymentCurrencyCodeField;
	//private PaymentAlternativeCurrencyCodeType paymentAlternativeCurrencyCodeField;
	//private AccountingCostType accountingCostField;

	LineCountNumeric        NumericType
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
	ID IdentifierType

	SignatoryParty PartyType `xml:",omitempty"`

	DigitalSignatureAttachment AttachmentType
}

type AttachmentType struct {
	EmbeddedDocumentBinaryObject BinaryObjectType
	ExternalReference            ExternalReferenceType
}

type ExternalReferenceType struct {
	URI IdentifierType
}

type BinaryObjectType struct {
	Format           string `xml:"format,attr,omitempty"`
	MimeCode         string `xml:"mimeCode,attr,omitempty"`
	EncodingCode     string `xml:"encodingCode,attr,omitempty"`
	CharacterSetCode string `xml:"characterSetCode,attr,omitempty"`
	Uri              string `xml:"uri,attr,omitempty"`
	Filename         string `xml:"filename,attr,omitempty"`
	Value            []byte `xml:",chardata"`
}

type PartyType struct {
	WebsiteURI string `xml:",omitempty"`

	//private EndpointIDType endpointIDField;
	//private IndustryClassificationCodeType industryClassificationCodeField;
	PartyIdentification []PartyIdentificationType

	//private PartyNameType partyNameField;

	PostalAddress AddressType

	//private PartyTaxSchemeType partyTaxSchemeField;
	//private PartyLegalEntityType[] partyLegalEntityField;
	//private ContactType contactField;
	//private PersonType personField;
	//private PartyType agentPartyField;

}

type AddressType struct {
	ID IdentifierType

	//PostboxType postboxField;
	//RoomType roomField;
	StreetName          TextType
	BlockName           TextType
	BuildingName        TextType
	BuildingNumber      TextType
	CitySubdivisionName TextType
	CityName            TextType
	PostalZone          TextType
	RegionType          TextType
	District            TextType
	Country             CountryType
}

type CountryType struct {
	IdentificationCode CodeType

	Name TextType
}

type PartyIdentificationType struct {
	ID IdentifierType
}

type SupplierPartyType struct {
	Party PartyType `xml:",omitempty"`
}
type CustomerPartyType struct {
	Party PartyType `xml:",omitempty"`
}

type CopyIndicatorType struct {
	IndicatorType
}

type IndicatorType struct {
	Value bool `xml:",chardata"`
}

type CodeType struct {
	ListID         string `xml:"listID,attr,omitempty"`
	ListAgencyID   string `xml:"listAgencyID,attr,omitempty"`
	ListAgencyName string `xml:"listAgencyName,attr,omitempty"`
	ListName       string `xml:"listName,attr,omitempty"`
	ListVersionID  string `xml:"listVersionID,attr,omitempty"`
	Name           string `xml:"name,attr,omitempty"`
	LanguageID     string `xml:"languageID,attr,omitempty"`
	ListURI        string `xml:"listURI,attr,omitempty"`
	ListSchemeURI  string `xml:"listSchemeURI,attr,omitempty"`
	Value          string `xml:",chardata"`
}

type TextType struct {
	LanguageID       string `xml:"languageID,attr,omitempty"`
	LanguageLocaleID string `xml:"languageLocaleID,attr,omitempty"`
	Value            string `xml:",chardata"`
}

type NumericType struct {
	Format string  `xml:"Format,attr,omitempty"`
	Value  float64 `xml:",chardata"`
}
