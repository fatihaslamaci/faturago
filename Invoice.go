package main

import (
	"encoding/xml"
	"time"
)

type customDate struct {
	time.Time
}

type customTime struct {
	time.Time
}

func (c *customDate) UnmarshalXML(d *xml.Decoder, start xml.StartElement) error {
	const shortForm = "2006-01-02" // yyyymmdd date format
	var v string
	d.DecodeElement(&v, &start)
	parse, err := time.Parse(shortForm, v)
	if err != nil {
		return err
	}
	*c = customDate{parse}
	return nil
}

func (c *customTime) UnmarshalXML(d *xml.Decoder, start xml.StartElement) error {
	const shortForm = "15:04:05" // yyyymmdd date format
	var v string
	d.DecodeElement(&v, &start)
	parse, err := time.Parse(shortForm, v)
	if err != nil {
		return err
	}
	*c = customTime{parse}
	return nil
}

type Invoice struct {
	XMLName         xml.Name `xml:"Invoice"`
	UBLVersionID    string
	CustomizationID string

	ProfileID            string
	ID                   IdType
	CopyIndicator        string
	UUID                 string
	IssueDate            customDate
	IssueTime            customTime
	InvoiceTypeCode      string
	DocumentCurrencyCode string
	LineCountNumeric     int
	Signature            SignatureType
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
	ID IdType

	SignatoryParty PartyType

	//private AttachmentType digitalSignatureAttachmentField;
}

type PartyType struct {
	WebsiteURI string

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

type IdentifierType1 struct {
	IdentifierType
}

type IdType struct {
	IdentifierType1
}
