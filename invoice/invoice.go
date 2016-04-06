package invoice

import (
	"encoding/xml"
	//"time"
)

//func (c *DateType) UnmarshalXML(d *xml.Decoder, start xml.StartElement) error {
//	const shortForm = "2006-01-02" // yyyy-mm-dd date format
//	var v string
//	d.DecodeElement(&v, &start)
//	parse, err := time.Parse(shortForm, v)
//	if err != nil {
//		return err
//	}
//	*c = DateType{parse}
//	return nil
//}

//func (c *TimeType) UnmarshalXML(d *xml.Decoder, start xml.StartElement) error {
//	const shortForm = "15:04:05" // hh-MM-ss time format
//	var v string
//	d.DecodeElement(&v, &start)
//	parse, err := time.Parse(shortForm, v)
//	if err != nil {
//		return err
//	}
//	*c = TimeType{parse}
//	return nil
//}

type DateType struct {
	//Normalde time.time olması gerekirken geçi cözüm olarak string yapılmıştır
	//time.Time
	XMLName xml.Name
	Value   string `xml:",chardata"`
}

type TimeType struct {
	//Normalde time.time olması gerekirken geçi cözüm olarak string yapılmıştır
	//time.Time
	XMLName xml.Name
	Value   string `xml:",chardata"`
}

type InvoiceType struct {
	XMLName                        xml.Name           `xml:"Invoice"`
	UBLExtensions                  []UBLExtensionType `xml:"UBLExtensions,omitempty"`
	UBLVersionID                   *IdentifierType1
	CustomizationID                *IdentifierType1
	ProfileID                      *IdentifierType1
	ID                             *IdentifierType1
	CopyIndicator                  *IndicatorType   `xml:"CopyIndicator,omitempty"`
	UUID                           *IdentifierType1 `xml:"UUID,omitempty"`
	IssueDate                      *DateType
	IssueTime                      *TimeType
	InvoiceTypeCode                *CodeType
	Note                           []TextType
	DocumentCurrencyCode           *CodeType
	TaxCurrencyCode                *CodeType
	PricingCurrencyCode            *CodeType
	PaymentCurrencyCode            *CodeType
	PaymentAlternativeCurrencyCode *CodeType
	AccountingCost                 *TextType
	LineCountNumeric               *NumericType
	InvoicePeriod                  *PeriodType
	OrderReference                 *OrderReferenceType
	BillingReference               []BillingReferenceType
	DespatchDocumentReference      []DocumentReferenceType
	ReceiptDocumentReference       []DocumentReferenceType
	OriginatorDocumentReference    []DocumentReferenceType
	ContractDocumentReference      []DocumentReferenceType
	AdditionalDocumentReference    []DocumentReferenceType
	Signature                      []SignatureType
	AccountingSupplierParty        *PartyType `xml:"AccountingSupplierParty>Party,omitempty"`
	AccountingCustomerParty        *PartyType `xml:"AccountingCustomerParty>Party,omitempty"`
	BuyerCustomerParty             *PartyType `xml:"BuyerCustomerParty>Party,omitempty"`
	SellerSupplierParty            *PartyType `xml:"SellerSupplierParty>Party,omitempty"`
	TaxRepresentativeParty         *PartyType `xml:"TaxRepresentativeParty>Party,omitempty"`
	Delivery                       []DeliveryType
	PaymentMeans                   []PaymentMeansType
	PaymentTerms                   *PaymentTermsType
	AllowanceCharge                []AllowanceChargeType
	TaxExchangeRate                *ExchangeRateType
	PricingExchangeRate            *ExchangeRateType
	PaymentExchangeRate            *ExchangeRateType
	PaymentAlternativeExchangeRate *ExchangeRateType
	TaxTotal                       []TaxTotalType
	WithholdingTaxTotal            []TaxTotalType
	LegalMonetaryTotal             *MonetaryTotalType
	InvoiceLine                    []InvoiceLineType
}

type UBLExtensionType struct {
	XMLName          xml.Name
	ExtensionContent string `xml:",innerxml"` //System.Xml.XmlElement
}

type PriceType struct {
	PriceAmountType
}
type PriceAmountType struct {
	PriceAmount AmountType
}

type AmountType struct {
	CurrencyID                string  `xml:"currencyID,attr,omitempty"`
	CurrencyCodeListVersionID string  `xml:"currencyCodeListVersionID,attr,omitempty"`
	Value                     float64 `xml:",chardata"`
}

type LineReferenceType struct {
	LineID            *LineIDType
	LineStatusCode    *CodeType
	DocumentReference *DocumentReferenceType
}
type LineIDType struct {
	IdentifierType1
}
type IdentifierType1 struct {
	XMLName xml.Name
	IdentifierType
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

type DocumentReferenceType struct {
	ID                  *IdentifierType1
	IssueDate           *DateType
	DocumentTypeCode    *CodeType
	DocumentType        *TextType
	DocumentDescription []TextType
	Attachment          *AttachmentType
	ValidityPeriod      *PeriodType
	IssuerParty         *PartyType
}

type TextType struct {
	LanguageID       string `xml:"languageID,attr,omitempty"`
	LanguageLocaleID string `xml:"languageLocaleID,attr,omitempty"`
	Value            string `xml:",chardata"`
}

type AttachmentType struct {
	EmbeddedDocumentBinaryObject *BinaryObjectType
	ExternalReference            *ExternalReferenceType
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

type ExternalReferenceType struct {
	URI *IdentifierType1
}
type PeriodType struct {
	StartDate       *DateType
	StartTime       *TimeType
	EndDate         *DateType
	EndTime         *TimeType
	DurationMeasure *MeasureType
	Description     *TextType
}

type MeasureType struct {
	unitCode              string  `xml:"unitCode,attr,omitempty"`
	unitCodeListVersionID string  `xml:"unitCodeListVersionID,attr,omitempty"`
	Value                 float64 `xml:",chardata"`
}

type PartyType struct {
	WebsiteURI                 *IdentifierType1
	EndpointID                 *IdentifierType1
	IndustryClassificationCode *CodeType
	PartyIdentification        []PartyIdentificationType
	PartyName                  *PartyNameType
	PostalAddress              *AddressType
	PartyTaxScheme             *PartyTaxSchemeType
	PartyLegalEntity           []PartyLegalEntityType
	Contact                    *ContactType
	Person                     *PersonType
	AgentParty                 *PartyType
}
type PartyIdentificationType struct {
	ID *IdentifierType1
}
type PartyNameType struct {
	Name *TextType
}
type AddressType struct {
	ID                  *IdentifierType1
	Postbox             *TextType
	Room                *TextType
	StreetName          *TextType
	BlockName           *TextType
	BuildingName        *TextType
	BuildingNumber      *TextType
	CitySubdivisionName *TextType
	CityName            *TextType
	PostalZone          *TextType
	Region              *TextType
	District            *TextType
	Country             *CountryType
}
type CountryType struct {
	IdentificationCode *CodeType
	Name               *TextType
}
type PartyTaxSchemeType struct {
	TaxScheme *TaxSchemeType
}
type TaxSchemeType struct {
	Name        *TextType
	TaxTypeCode *CodeType
}
type PartyLegalEntityType struct {
	RegistrationName            *TextType
	CompanyID                   *IdentifierType1
	RegistrationDate            *DateType
	SoleProprietorshipIndicator *IndicatorType
	CorporateStockAmount        *AmountType
	FullyPaidSharesIndicator    *IndicatorType
	CorporateRegistrationScheme *CorporateRegistrationSchemeType
	HeadOfficeParty             *PartyType
}

type IndicatorType struct {
	Value bool `xml:",chardata"`
}
type CorporateRegistrationSchemeType struct {
	ID                            *IdentifierType1
	Name                          *TextType
	CorporateRegistrationTypeCode *CodeType
	JurisdictionRegionAddress     []AddressType
}
type ContactType struct {
	Telephone          *TextType
	Telefax            *TextType
	ElectronicMail     *TextType
	Note               *TextType
	OtherCommunication []CommunicationType
}
type CommunicationType struct {
	ChannelCode *CodeType
	Channel     *TextType
	Value       *TextType
}
type PersonType struct {
	FirstName                 *TextType
	FamilyName                *TextType
	Title                     *TextType
	MiddleName                *TextType
	NameSuffix                *TextType
	NationalityID             *IdentifierType1
	FinancialAccount          *FinancialAccountType
	IdentityDocumentReference *DocumentReferenceType
}
type FinancialAccountType struct {
	ID                         *IdentifierType1
	CurrencyCode               *CodeType
	PaymentNote                *TextType
	FinancialInstitutionBranch *BranchType
}
type BranchType struct {
	Name                 *TextType
	FinancialInstitution *FinancialInstitutionType
}
type FinancialInstitutionType struct {
	Name *TextType
}
type OrderLineReferenceType struct {
	LineID           *LineIDType
	SalesOrderLineID *IdentifierType1
	UUID             *IdentifierType1
	LineStatusCode   *CodeType
	OrderReference   *OrderReferenceType
}
type OrderReferenceType struct {
	ID                *IdentifierType1
	SalesOrderID      *IdentifierType1
	IssueDate         *DateType
	OrderTypeCode     *CodeType
	DocumentReference *DocumentReferenceType
}
type InvoiceLineType struct {
	ID                    *IdentifierType1
	Note                  []TextType
	InvoicedQuantity      *QuantityType1
	LineExtensionAmount   *AmountType
	OrderLineReference    []OrderLineReferenceType
	DespatchLineReference []LineReferenceType
	ReceiptLineReference  []LineReferenceType
	Delivery              []DeliveryType
	AllowanceCharge       []AllowanceChargeType
	TaxTotal              *TaxTotalType
	WithholdingTaxTotal   []TaxTotalType
	Item                  *ItemType
	Price                 *PriceType
	SubInvoiceLine        []InvoiceLineType
}

type QuantityType1 struct {
	QuantityType
}
type QuantityType struct {
	UnitCode               string  `xml:"unitCode,attr,omitempty"`
	UnitCodeListID         string  `xml:"unitCodeListID,attr,omitempty"`
	UnitCodeListAgencyID   string  `xml:"unitCodeListAgencyID,attr,omitempty"`
	UnitCodeListAgencyName string  `xml:"unitCodeListAgencyName,attr,omitempty"`
	Value                  float64 `xml:",chardata"`
}

type QuantityType2 struct {
	QuantityType1
}

type DeliveryType struct {
	ID                          *IdentifierType1
	Quantity                    *QuantityType2
	ActualDeliveryDate          *DateType
	ActualDeliveryTime          *TimeType
	LatestDeliveryDate          *DateType
	LatestDeliveryTime          *TimeType
	TrackingID                  *IdentifierType1
	DeliveryAddress             *AddressType
	AlternativeDeliveryLocation *LocationType1
	EstimatedDeliveryPeriod     *PeriodType
	CarrierParty                *PartyType
	DeliveryParty               *PartyType
	Despatch                    *DespatchType
	DeliveryTerms               []DeliveryTermsType
	Shipment                    *ShipmentType
}
type LocationType1 struct {
	ID      *IdentifierType1
	Address *AddressType
}
type DespatchType struct {
	ID                      *IdentifierType1
	ActualDespatchDate      *DateType
	ActualDespatchTime      *TimeType
	Instructions            *TextType
	DespatchAddress         *AddressType
	DespatchParty           *PartyType
	Contact                 *ContactType
	EstimatedDespatchPeriod *PeriodType
}
type DeliveryTermsType struct {
	ID           *IdentifierType1
	SpecialTerms *TextType
	Amount       *AmountType
}
type ShipmentType struct {
	ID                                 *IdentifierType1
	HandlingCode                       *CodeType
	HandlingInstructions               *TextType
	GrossWeightMeasure                 *MeasureType
	NetWeightMeasure                   *MeasureType
	GrossVolumeMeasure                 *MeasureType
	NetVolumeMeasure                   *MeasureType
	TotalGoodsItemQuantity             *QuantityType1
	TotalTransportHandlingUnitQuantity *QuantityType1
	InsuranceValueAmount               *AmountType
	DeclaredCustomsValueAmount         *AmountType
	DeclaredForCarriageValueAmount     *AmountType
	DeclaredStatisticsValueAmount      *AmountType
	FreeOnBoardValueAmount             *AmountType
	SpecialInstructions                []TextType
	TransportHandlingUnit              []TransportHandlingUnitType
	ReturnAddress                      *AddressType
	FirstArrivalPortLocation           *LocationType1
	LastExitPortLocation               *LocationType1
}
type TransportHandlingUnitType struct {
	ID                              *IdentifierType1
	TransportHandlingUnitTypeCode   *CodeType
	HandlingCode                    *CodeType
	HandlingInstructions            *TextType
	HazardousRiskIndicator          *IndicatorType
	TotalGoodsItemQuantity          *QuantityType1
	TotalPackageQuantity            *QuantityType1
	DamageRemarks                   []TextType
	TraceID                         *IdentifierType1
	ActualPackage                   []PackageType
	TransportEquipment              []TransportEquipmentType
	TransportMeans                  []TransportMeansType
	HazardousGoodsTransit           []HazardousGoodsTransitType
	MeasurementDimension            []DimensionType
	MinimumTemperature              *TemperatureType
	MaximumTemperature              *TemperatureType
	FloorSpaceMeasurementDimension  *DimensionType
	PalletSpaceMeasurementDimension *DimensionType
	ShipmentDocumentReference       []DocumentReferenceType
	CustomsDeclaration              []CustomsDeclarationType
}
type PackageType struct {
	CustomsDeclaration          *IdentifierType1
	Quantity                    *QuantityType2
	ReturnableMaterialIndicator *IndicatorType
	PackageLevelCode            *CodeType
	PackagingTypeCode           *CodeType
	PackingMaterial             []TextType
	ContainedPackage            []PackageType
	GoodsItem                   []GoodsItemType
	MeasurementDimension        []DimensionType
}
type GoodsItemType struct {
	ID                               *IdentifierType1
	Description                      []TextType
	HazardousRiskIndicator           *IndicatorType
	DeclaredCustomsValueAmount       *AmountType
	DeclaredForCarriageValueAmount   *AmountType
	DeclaredStatisticsValueAmount    *AmountType
	FreeOnBoardValueAmount           *AmountType
	InsuranceValueAmount             *AmountType
	ValueAmount                      *AmountType
	GrossWeightMeasure               *MeasureType
	NetWeightMeasure                 *MeasureType
	ChargeableWeightMeasure          *MeasureType
	GrossVolumeMeasure               *MeasureType
	NetVolumeMeasure                 *MeasureType
	Quantity                         *QuantityType2
	RequiredCustomsID                *IdentifierType1
	CustomsStatusCode                *CodeType
	CustomsTariffQuantity            *QuantityType1
	CustomsImportClassifiedIndicator *IndicatorType
	ChargeableQuantity               *QuantityType1
	ReturnableQuantity               *QuantityType1
	TraceID                          *IdentifierType1
	Item                             []ItemType
	FreightAllowanceCharge           []AllowanceChargeType
	Temperature                      []TemperatureType
	OriginAddress                    *AddressType
	MeasurementDimension             []DimensionType
}
type ItemType struct {
	Description                     *TextType
	Name                            *TextType
	Keyword                         *TextType
	BrandName                       *TextType
	ModelName                       *TextType
	BuyersItemIdentification        *ItemIdentificationType
	SellersItemIdentification       *ItemIdentificationType
	ManufacturersItemIdentification *ItemIdentificationType
	AdditionalItemIdentification    []ItemIdentificationType
	CommodityClassification         []CommodityClassificationType
	ItemInstance                    []ItemInstanceType
}
type ItemIdentificationType struct {
	ID *IdentifierType1
}
type CommodityClassificationType struct {
	ItemClassificationCode *CodeType
}
type ItemInstanceType struct {
	ProductTraceID         *IdentifierType1
	ManufactureDate        *DateType
	ManufactureTime        *TimeType
	BestBeforeDate         *DateType
	RegistrationID         *IdentifierType1
	SerialID               *IdentifierType1
	AdditionalItemProperty []ItemPropertyType
	LotIdentification      *LotIdentificationType
}
type ItemPropertyType struct {
	ID                *IdentifierType1
	Name              *TextType
	NameCode          *CodeType
	TestMethod        *TextType
	Value             *TextType
	ValueQuantity     *QuantityType1
	ValueQualifier    []TextType
	ImportanceCode    *CodeType
	ListValue         []TextType
	UsabilityPeriod   *PeriodType
	ItemPropertyGroup []ItemPropertyGroupType
	RangeDimension    *DimensionType
	ItemPropertyRange *ItemPropertyRangeType
}
type ItemPropertyGroupType struct {
	ID             *IdentifierType1
	Name           *TextType
	ImportanceCode *CodeType
}
type DimensionType struct {
	AttributeID    *IdentifierType1
	Measure        *MeasureType
	Description    []TextType
	MinimumMeasure *MeasureType
	MaximumMeasure *MeasureType
}
type ItemPropertyRangeType struct {
	MinimumValue *TextType
	MaximumValue *TextType
}
type LotIdentificationType struct {
	LotNumberID            *IdentifierType1
	ExpiryDate             *DateType
	AdditionalItemProperty []ItemPropertyType
}
type AllowanceChargeType struct {
	ChargeIndicator         *IndicatorType
	AllowanceChargeReason   *TextType
	MultiplierFactorNumeric *NumericType
	SequenceNumeric         *NumericType
	Amount                  *AmountType
	BaseAmount              *AmountType
	PerUnitAmount           *AmountType
}

type NumericType struct {
	Format string  `xml:"Format,attr,omitempty"`
	Value  float64 `xml:",chardata"`
}

type TemperatureType struct {
	AttributeID *IdentifierType1
	Measure     *MeasureType
	Description []TextType
}
type TransportEquipmentType struct {
	ID                         *IdentifierType1
	TransportEquipmentTypeCode *CodeType
	Description                *TextType
}
type TransportMeansType struct {
	JourneyID                 *IdentifierType1
	RegistrationNationalityID *IdentifierType1
	RegistrationNationality   []TextType
	DirectionCode             *CodeType
	TransportMeansTypeCode    *CodeType
	TradeServiceCode          *CodeType
	Stowage                   *StowageType
	AirTransport              *AirTransportType
	RoadTransport             *RoadTransportType
	RailTransport             *RailTransportType
	MaritimeTransport         *MaritimeTransportType
	OwnerParty                *PartyType
	MeasurementDimension      []DimensionType
}
type StowageType struct {
	LocationID           *IdentifierType1
	Location             []TextType
	MeasurementDimension []DimensionType
}
type AirTransportType struct {
	AircraftID *IdentifierType1
}
type RoadTransportType struct {
	LicensePlateID *IdentifierType1
}
type RailTransportType struct {
	TrainID   *IdentifierType1
	RailCarID *IdentifierType1
}
type MaritimeTransportType struct {
	VesselID                             *IdentifierType1
	VesselName                           *TextType
	RadioCallSignID                      *IdentifierType1
	ShipsRequirements                    []TextType
	GrossTonnageMeasure                  *MeasureType
	NetTonnageMeasure                    *MeasureType
	RegistryCertificateDocumentReference *DocumentReferenceType
	RegistryPortLocation                 *LocationType1
}
type HazardousGoodsTransitType struct {
	TransportEmergencyCardCode *CodeType
	PackingCriteriaCode        *CodeType
	HazardousRegulationCode    *CodeType
	InhalationToxicityZoneCode *CodeType
	TransportAuthorizationCode *CodeType
	MaximumTemperature         *TemperatureType
	MinimumTemperature         *TemperatureType
}
type CustomsDeclarationType struct {
	ID          *IdentifierType1
	IssuerParty *PartyType
}
type TaxTotalType struct {
	TaxAmount   *AmountType
	TaxSubtotal []TaxSubtotalType
}
type TaxSubtotalType struct {
	TaxableAmount                *AmountType
	TaxAmount                    *AmountType
	CalculationSequenceNumeric   *NumericType
	TransactionCurrencyTaxAmount *AmountType
	Percent                      *NumericType
	BaseUnitMeasure              *MeasureType
	PerUnitAmount                *AmountType
	TaxCategory                  *TaxCategoryType
}
type TaxCategoryType struct {
	Name                   *TextType
	TaxExemptionReasonCode *CodeType
	TaxExemptionReason     *TextType
	TaxScheme              *TaxSchemeType
}
type MonetaryTotalType struct {
	TaxExclusiveAmount    *AmountType
	TaxInclusiveAmount    *AmountType
	LineExtensionAmount   *AmountType
	AllowanceTotalAmount  *AmountType
	ChargeTotalAmount     *AmountType
	PayableRoundingAmount *AmountType
	PayableAmount         *AmountType
}
type ExchangeRateType struct {
	SourceCurrencyCode *CodeType
	TargetCurrencyCode *CodeType
	CalculationRate    *NumericType
	Date               *DateType
}
type PaymentTermsType struct {
	Note                    *TextType
	PenaltySurchargePercent *NumericType
	Amount                  *AmountType
	PenaltyAmount           *AmountType
	PaymentDueDate          *DateType
	SettlementPeriod        *PeriodType
}
type PaymentMeansType struct {
	PaymentMeansCode      *CodeType
	PaymentDueDate        *DateType
	PaymentChannelCode    *CodeType
	InstructionNote       *TextType
	PayerFinancialAccount *FinancialAccountType
	PayeeFinancialAccount *FinancialAccountType
}

type SignatureType struct {
	ID                         *IdentifierType1
	SignatoryParty             *PartyType
	DigitalSignatureAttachment *AttachmentType
}
type BillingReferenceLineType struct {
	ID              *IdentifierType1
	Amount          *AmountType
	AllowanceCharge []AllowanceChargeType
}
type BillingReferenceType struct {
	InvoiceDocumentReference              *DocumentReferenceType
	SelfBilledInvoiceDocumentReference    *DocumentReferenceType
	CreditNoteDocumentReference           *DocumentReferenceType
	SelfBilledCreditNoteDocumentReference *DocumentReferenceType
	DebitNoteDocumentReference            *DocumentReferenceType
	ReminderDocumentReference             *DocumentReferenceType
	AdditionalDocumentReference           *DocumentReferenceType
	BillingReferenceLine                  []BillingReferenceLineType
}
type SignatureType1 struct {
	SignedInfo     *SignedInfoType
	SignatureValue *SignatureValueType
	KeyInfo        *KeyInfoType
	Object         []ObjectType
	Id             string
}
type SignedInfoType struct {
	CanonicalizationMethod *CanonicalizationMethodType1
	SignatureMethod        *SignatureMethodType1
	Reference              []ReferenceType1
	Id                     string
}
type CanonicalizationMethodType1 struct {
	Any       []string //string//System.Xml.XmlNode
	Algorithm string
}
type SignatureMethodType1 struct {
	HMACOutputLength string
	Any              []string //string//System.Xml.XmlNode
	Algorithm        string
}
type ReferenceType1 struct {
	Transforms   []TransformType
	DigestMethod *DigestMethodType
	DigestValue  []byte
	Id           string
	URI          string
	Type         string
}
type TransformType struct {
	Items     []string //object
	Text      []string
	Algorithm string
}
type DigestMethodType struct {
	Any       []string //string//System.Xml.XmlNode
	Algorithm string
}
type SignatureValueType struct {
	Id    string
	Value []byte
}
type KeyInfoType struct {
	Items []string //object
	//Burayı kontrol et
	ItemsElementName []string //[]ItemsChoiceType2
	Text             []string
	Id               string
}

/*
enum ItemsChoiceType2 {
        Item,
        KeyName,
        KeyValue,
        MgmtData,
        PGPData,
        RetrievalMethod,
        SPKIData,
        X509Data,
    }
*/

type KeyValueType struct {
	Item string //object
	Text []string
}
type DSAKeyValueType struct {
	P           []byte
	Q           []byte
	G           []byte
	Y           []byte
	J           []byte
	Seed        []byte
	PgenCounter []byte
}
type RSAKeyValueType struct {
	Modulus  []byte
	Exponent []byte
}
type PGPDataType struct {
	Items []string //object
	//Burayı kontrol et
	ItemsElementName []string //[]ItemsChoiceType1
}

//ItemsChoiceType1 enum
//Item,
//PGPKeyID,
//PGPKeyPacket,
//}
type RetrievalMethodType struct {
	Transforms []TransformType
	URI        string
	Type       string
}
type SPKIDataType struct {
	SPKISexp []byte
	Any      string `xml:",innerxml"` //System.Xml.XmlElement
}
type X509DataType struct {
	Items []string //object
	//Burayı kontrol et
	ItemsElementName []string //[]ItemsChoiceType
}
type X509IssuerSerialType struct {
	X509IssuerName   string
	X509SerialNumber string
}

//ItemsChoiceType enum
//Item,
//X509CRL,
//X509Certificate,
//X509IssuerSerial,
//X509SKI,
//X509SubjectName,
//}
//ItemsChoiceType2 enum
//Item,
//KeyName,
//KeyValue,
//MgmtData,
//PGPData,
//RetrievalMethod,
//SPKIData,
//X509Data,
//}
type ObjectType struct {
	Any      []string //string//System.Xml.XmlNode
	Id       string
	MimeType string
	Encoding string
}
