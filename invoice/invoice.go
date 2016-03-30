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
	Note                           []TextType1
	DocumentCurrencyCode           *CodeType
	TaxCurrencyCode                *CodeType
	PricingCurrencyCode            *CodeType
	PaymentCurrencyCode            *CodeType
	PaymentAlternativeCurrencyCode *CodeType
	AccountingCost                 *TextType1
	LineCountNumeric               *NumericType1
	InvoicePeriod                  *PeriodType
	OrderReference                 *OrderReferenceType
	BillingReference               []BillingReferenceType
	DespatchDocumentReference      []DocumentReferenceType
	ReceiptDocumentReference       []DocumentReferenceType
	OriginatorDocumentReference    []DocumentReferenceType
	ContractDocumentReference      []DocumentReferenceType
	AdditionalDocumentReference    []DocumentReferenceType
	Signature                      []SignatureType
	AccountingSupplierParty        *SupplierPartyType
	AccountingCustomerParty        *CustomerPartyType
	BuyerCustomerParty             *CustomerPartyType
	SellerSupplierParty            *SupplierPartyType
	TaxRepresentativeParty         *PartyType
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

//type UBLExtensionType struct {
//	UBLExtension UBLExtensionType2
//}
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
	DocumentType        *DocumentTypeType
	DocumentDescription []TextType1
	Attachment          *AttachmentType
	ValidityPeriod      *PeriodType
	IssuerParty         *PartyType
}

type DocumentTypeType struct {
	TextType1
}
type TextType1 struct {
	TextType
}
type TextType struct {
	LanguageID       string `xml:"languageID,attr,omitempty"`
	LanguageLocaleID string `xml:"languageLocaleID,attr,omitempty"`
	Value            string `xml:",chardata"`
}
type NameType struct {
	TextType
}
type NameType1 struct {
	NameType
}

type TextType2 struct {
	TextType1
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
	StartTime       *StartTimeType
	EndDate         *DateType
	EndTime         *TimeType
	DurationMeasure *DurationMeasureType
	Description     *TextType1
}
type StartTimeType struct {
	TimeType
}

type DurationMeasureType struct {
	MeasureType1
}
type MeasureType1 struct {
	MeasureType
}
type MeasureType struct {
	unitCode              string  `xml:"unitCode,attr,omitempty"`
	unitCodeListVersionID string  `xml:"unitCodeListVersionID,attr,omitempty"`
	Value                 float64 `xml:",chardata"`
}

type MeasureType2 struct {
	MeasureType1
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
	ID IdentifierType1
}
type PartyNameType struct {
	Name NameType1
}
type AddressType struct {
	ID                  *IdentifierType1
	Postbox             *TextType1
	Room                *TextType1
	StreetName          *NameType
	BlockName           *NameType
	BuildingName        *NameType
	BuildingNumber      *TextType1
	CitySubdivisionName *NameType
	CityName            *NameType
	PostalZone          *TextType1
	Region              *TextType1
	District            *TextType1
	Country             *CountryType
}
type CountryType struct {
	IdentificationCode *CodeType
	Name               *NameType1
}
type PartyTaxSchemeType struct {
	TaxScheme *TaxSchemeType
}
type TaxSchemeType struct {
	Name        *NameType1
	TaxTypeCode *CodeType
}
type PartyLegalEntityType struct {
	RegistrationName            *NameType
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
	Name                          *NameType1
	CorporateRegistrationTypeCode *CodeType
	JurisdictionRegionAddress     []AddressType
}
type ContactType struct {
	Telephone          *TextType1
	Telefax            *TextType1
	ElectronicMail     *TextType1
	Note               *TextType1
	OtherCommunication []CommunicationType
}
type CommunicationType struct {
	ChannelCode *CodeType
	Channel     *TextType1
	Value       *TextType1
}
type PersonType struct {
	FirstName                 *NameType
	FamilyName                *NameType
	Title                     *TextType1
	MiddleName                *NameType
	NameSuffix                *TextType1
	NationalityID             *IdentifierType1
	FinancialAccount          *FinancialAccountType
	IdentityDocumentReference *DocumentReferenceType
}
type FinancialAccountType struct {
	ID                         *IdentifierType1
	CurrencyCode               *CodeType
	PaymentNote                *TextType1
	FinancialInstitutionBranch *BranchType
}
type BranchType struct {
	Name                 *NameType1
	FinancialInstitution *FinancialInstitutionType
}
type FinancialInstitutionType struct {
	Name *NameType1
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
	Note                  []TextType1
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
	Instructions            *TextType1
	DespatchAddress         *AddressType
	DespatchParty           *PartyType
	Contact                 *ContactType
	EstimatedDespatchPeriod *PeriodType
}
type DeliveryTermsType struct {
	ID           *IdentifierType1
	SpecialTerms *TextType1
	Amount       *AmountType
}
type ShipmentType struct {
	ID                                 *IdentifierType1
	HandlingCode                       *CodeType
	HandlingInstructions               *TextType1
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
	SpecialInstructions                []TextType1
	TransportHandlingUnit              []TransportHandlingUnitType
	ReturnAddress                      *AddressType
	FirstArrivalPortLocation           *LocationType1
	LastExitPortLocation               *LocationType1
}
type TransportHandlingUnitType struct {
	ID                              *IdentifierType1
	TransportHandlingUnitTypeCode   *CodeType
	HandlingCode                    *CodeType
	HandlingInstructions            *TextType1
	HazardousRiskIndicator          *IndicatorType
	TotalGoodsItemQuantity          *QuantityType1
	TotalPackageQuantity            *QuantityType1
	DamageRemarks                   []TextType1
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
	PackingMaterial             []TextType1
	ContainedPackage            []PackageType
	GoodsItem                   []GoodsItemType
	MeasurementDimension        []DimensionType
}
type GoodsItemType struct {
	ID                               *IdentifierType1
	Description                      []TextType1
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
	Description                     *TextType1
	Name                            *NameType1
	Keyword                         *TextType1
	BrandName                       *NameType
	ModelName                       *NameType
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
	Name              *NameType1
	NameCode          *CodeType
	TestMethod        *TextType1
	Value             *TextType1
	ValueQuantity     *QuantityType1
	ValueQualifier    []TextType1
	ImportanceCode    *CodeType
	ListValue         []TextType1
	UsabilityPeriod   *PeriodType
	ItemPropertyGroup []ItemPropertyGroupType
	RangeDimension    *DimensionType
	ItemPropertyRange *ItemPropertyRangeType
}
type ItemPropertyGroupType struct {
	ID             *IdentifierType1
	Name           *NameType1
	ImportanceCode *CodeType
}
type DimensionType struct {
	AttributeID    *IdentifierType1
	Measure        *MeasureType2
	Description    []TextType1
	MinimumMeasure *MeasureType
	MaximumMeasure *MeasureType
}
type ItemPropertyRangeType struct {
	MinimumValue *TextType1
	MaximumValue *TextType1
}
type LotIdentificationType struct {
	LotNumberID            *IdentifierType1
	ExpiryDate             *DateType
	AdditionalItemProperty []ItemPropertyType
}
type AllowanceChargeType struct {
	ChargeIndicator         *IndicatorType
	AllowanceChargeReason   *TextType1
	MultiplierFactorNumeric *NumericType1
	SequenceNumeric         *NumericType1
	Amount                  *AmountType
	BaseAmount              *AmountType
	PerUnitAmount           *AmountType
}

type NumericType1 struct {
	NumericType
}
type NumericType struct {
	Format string  `xml:"Format,attr,omitempty"`
	Value  float64 `xml:",chardata"`
}
type RateType struct {
	NumericType
}

type RateType1 struct {
	RateType
}

type PercentType struct {
	NumericType
}

type PercentType1 struct {
	PercentType
}

type ValueType1 struct {
	NumericType
}
type TemperatureType struct {
	AttributeID *IdentifierType1
	Measure     *MeasureType2
	Description []TextType1
}
type TransportEquipmentType struct {
	ID                         *IdentifierType1
	TransportEquipmentTypeCode *CodeType
	Description                *TextType1
}
type TransportMeansType struct {
	JourneyID                 *IdentifierType1
	RegistrationNationalityID *IdentifierType1
	RegistrationNationality   []TextType1
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
	Location             []TextType1
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
	VesselName                           *NameType
	RadioCallSignID                      *IdentifierType1
	ShipsRequirements                    []TextType1
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
	CalculationSequenceNumeric   *NumericType1
	TransactionCurrencyTaxAmount *AmountType
	Percent                      *PercentType1
	BaseUnitMeasure              *MeasureType
	PerUnitAmount                *AmountType
	TaxCategory                  *TaxCategoryType
}
type TaxCategoryType struct {
	Name                   *NameType1
	TaxExemptionReasonCode *CodeType
	TaxExemptionReason     *TextType1
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
	CalculationRate    *RateType
	Date               *DateType
}
type PaymentTermsType struct {
	Note                    *TextType1
	PenaltySurchargePercent *PercentType
	Amount                  *AmountType
	PenaltyAmount           *AmountType
	PaymentDueDate          *DateType
	SettlementPeriod        *PeriodType
}
type PaymentMeansType struct {
	PaymentMeansCode      *CodeType
	PaymentDueDate        *DateType
	PaymentChannelCode    *CodeType
	InstructionNote       *TextType1
	PayerFinancialAccount *FinancialAccountType
	PayeeFinancialAccount *FinancialAccountType
}
type CustomerPartyType struct {
	Party *PartyType
}
type SupplierPartyType struct {
	Party *PartyType
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
