package main

import (
	"encoding/xml"
	"time"
)

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

type InvoiceType struct {
	XMLName                        xml.Name `xml:"Invoice"`
	UBLExtensions                  []UBLExtensionType
	UBLVersionID                   *IdentifierType1
	CustomizationID                *IdentifierType1
	ProfileID                      *IdentifierType1
	ID                             *IDType
	CopyIndicator                  *CopyIndicatorType `xml:",omitempty"`
	UUID                           *UUIDType
	IssueDate                      *IssueDateType
	IssueTime                      *IssueTimeType
	InvoiceTypeCode                *CodeType1
	Note                           []NoteType
	DocumentCurrencyCode           *CodeType1
	TaxCurrencyCode                *CodeType1
	PricingCurrencyCode            *CodeType1
	PaymentCurrencyCode            *CodeType1
	PaymentAlternativeCurrencyCode *CodeType1
	AccountingCost                 *AccountingCostType
	LineCountNumeric               *LineCountNumericType
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
type UBLExtensionType struct {
	ExtensionContent string //System.Xml.XmlElement
}
type PriceType struct {
	PriceAmount PriceAmountType
}
type PriceAmountType struct {
	PriceAmount AmountType1
}

type AmountType struct {
	CurrencyID                string  `xml:"currencyID,attr,omitempty"`
	CurrencyCodeListVersionID string  `xml:"currencyCodeListVersionID,attr,omitempty"`
	Value                     float64 `xml:",chardata"`
}

type AmountType2 struct {
	AmountType1
}
type AmountType1 struct {
	AmountType
}

type LineReferenceType struct {
	LineID            *LineIDType
	LineStatusCode    *LineStatusCodeType
	DocumentReference *DocumentReferenceType
}
type LineIDType struct {
	IdentifierType1
}
type IdentifierType1 struct {
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
type IDType struct {
	IdentifierType1
}

type UUIDType struct {
	IdentifierType1
}

type LineStatusCodeType struct {
	CodeType1
}
type CodeType1 struct {
	CodeType
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
	ID                  *IDType
	IssueDate           *IssueDateType
	DocumentTypeCode    *CodeType1
	DocumentType        *DocumentTypeType
	DocumentDescription []DocumentDescriptionType
	Attachment          *AttachmentType
	ValidityPeriod      *PeriodType
	IssuerParty         *PartyType
}
type IssueDateType struct {
	DateType
}

//type DateType struct {
//	Value time.Time//System.DateTime
//}
type DateType struct {
	time.Time
}

type ValidityStartDateType struct {
	DateType
}
type ValidationDateType struct {
	DateType
}
type TaxPointDateType struct {
	DateType
}
type SubmissionDueDateType struct {
	DateType
}
type SubmissionDateType struct {
	DateType
}
type StartDateType struct {
	DateType
}
type SourceForecastIssueDateType struct {
	DateType
}
type RevisionDateType struct {
	DateType
}
type ResponseDateType struct {
	DateType
}
type ResolutionDateType struct {
	DateType
}
type RequiredDeliveryDateType struct {
	DateType
}
type RequestedPublicationDateType struct {
	DateType
}
type RequestedDespatchDateType struct {
	DateType
}
type RequestedDeliveryDateType struct {
	DateType
}
type RegistrationExpirationDateType struct {
	DateType
}
type RegistrationDateType struct {
	DateType
}
type RegisteredDateType struct {
	DateType
}
type ReferenceDateType struct {
	DateType
}
type ReceivedDateType struct {
	DateType
}
type PreviousMeterReadingDateType struct {
	DateType
}
type PlannedDateType struct {
	DateType
}
type PaymentDueDateType struct {
	DateType
}
type PaidDateType struct {
	DateType
}
type OccurrenceDateType struct {
	DateType
}
type NominationDateType struct {
	DateType
}
type ManufactureDateType struct {
	DateType
}
type LatestSecurityClearanceDateType struct {
	DateType
}
type LatestProposalAcceptanceDateType struct {
	DateType
}
type LatestPickupDateType struct {
	DateType
}
type LatestMeterReadingDateType struct {
	DateType
}
type LatestDeliveryDateType struct {
	DateType
}
type LastRevisionDateType struct {
	DateType
}
type InstallmentDueDateType struct {
	DateType
}
type GuaranteedDespatchDateType struct {
	DateType
}
type FirstShipmentAvailibilityDateType struct {
	DateType
}
type ExpiryDateType struct {
	DateType
}
type EstimatedDespatchDateType struct {
	DateType
}
type EstimatedDeliveryDateType struct {
	DateType
}
type EndDateType struct {
	DateType
}
type EffectiveDateType struct {
	DateType
}
type EarliestPickupDateType struct {
	DateType
}
type DueDateType struct {
	DateType
}
type DateType1 struct {
	DateType
}
type ComparisonForecastIssueDateType struct {
	DateType
}
type CallDateType struct {
	DateType
}
type BirthDateType struct {
	DateType
}
type BestBeforeDateType struct {
	DateType
}
type AwardDateType struct {
	DateType
}
type AvailabilityDateType struct {
	DateType
}
type ApprovalDateType struct {
	DateType
}
type ActualPickupDateType struct {
	DateType
}
type ActualDespatchDateType struct {
	DateType
}
type ActualDeliveryDateType struct {
	DateType
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
type VesselNameType struct {
	NameType
}
type TechnicalNameType struct {
	NameType
}
type StreetNameType struct {
	NameType
}
type ServiceNameType struct {
	NameType
}
type RoamingPartnerNameType struct {
	NameType
}
type RetailEventNameType struct {
	NameType
}
type RegistrationNameType struct {
	NameType
}
type OtherNameType struct {
	NameType
}
type NameType1 struct {
	NameType
}
type ModelNameType struct {
	NameType
}
type MiddleNameType struct {
	NameType
}
type HolderNameType struct {
	NameType
}
type FirstNameType struct {
	NameType
}
type FileNameType struct {
	NameType
}
type FamilyNameType struct {
	NameType
}
type CitySubdivisionNameType struct {
	NameType
}
type CityNameType struct {
	NameType
}
type CategoryNameType struct {
	NameType
}
type BuildingNameType struct {
	NameType
}
type BrandNameType struct {
	NameType
}
type BlockNameType struct {
	NameType
}
type AliasNameType struct {
	NameType
}
type AdditionalStreetNameType struct {
	NameType
}
type ExtensionReasonType struct {
	TextType1
}
type ExtensionAgencyNameType struct {
	TextType1
}
type XPathType struct {
	TextType1
}
type WorkPhaseType struct {
	TextType1
}
type WeightType struct {
	TextType1
}
type WarrantyInformationType struct {
	TextType1
}
type ValueType struct {
	TextType1
}
type ValueQualifierType struct {
	TextType1
}
type ValidateToolVersionType struct {
	TextType1
}
type ValidateToolType struct {
	TextType1
}
type ValidateProcessType struct {
	TextType1
}
type TransportationServiceDescriptionType struct {
	TextType1
}
type TransportUserSpecialTermsType struct {
	TextType1
}
type TransportUserRemarksType struct {
	TextType1
}
type TransportServiceProviderSpecialTermsType struct {
	TextType1
}
type TransportServiceProviderRemarksType struct {
	TextType1
}
type TradingRestrictionsType struct {
	TextType1
}
type TitleType struct {
	TextType1
}
type TimingComplaintType struct {
	TextType1
}
type TimezoneOffsetType struct {
	TextType1
}
type TimeAmountType struct {
	TextType1
}
type TierRangeType struct {
	TextType1
}
type TextType2 struct {
	TextType1
}
type TestMethodType struct {
	TextType1
}
type TelephoneType struct {
	TextType1
}
type TelefaxType struct {
	TextType1
}
type TelecommunicationsSupplyTypeType struct {
	TextType1
}
type TelecommunicationsServiceCategoryType struct {
	TextType1
}
type TelecommunicationsServiceCallType struct {
	TextType1
}
type TechnicalCommitteeDescriptionType struct {
	TextType1
}
type TaxExemptionReasonType struct {
	TextType1
}
type TariffDescriptionType struct {
	TextType1
}
type SummaryDescriptionType struct {
	TextType1
}
type SubscriberTypeType struct {
	TextType1
}
type StatusReasonType struct {
	TextType1
}
type SpecialTransportRequirementsType struct {
	TextType1
}
type SpecialTermsType struct {
	TextType1
}
type SpecialServiceInstructionsType struct {
	TextType1
}
type SpecialInstructionsType struct {
	TextType1
}
type SignatureMethodType struct {
	TextType1
}
type ShipsRequirementsType struct {
	TextType1
}
type ShippingMarksType struct {
	TextType1
}
type ServiceTypeType struct {
	TextType1
}
type ServiceNumberCalledType struct {
	TextType1
}
type SealingPartyTypeType struct {
	TextType1
}
type RoomType struct {
	TextType1
}
type RoleDescriptionType struct {
	TextType1
}
type ResolutionType struct {
	TextType1
}
type ResidenceTypeType struct {
	TextType1
}
type ReplenishmentOwnerDescriptionType struct {
	TextType1
}
type RemarksType struct {
	TextType1
}
type RejectionNoteType struct {
	TextType1
}
type RejectReasonType struct {
	TextType1
}
type RegulatoryDomainType struct {
	TextType1
}
type RegistrationNationalityType struct {
	TextType1
}
type RegionType struct {
	TextType1
}
type ReferenceType struct {
}
type RankType struct {
	TextType1
}
type PurposeType struct {
	TextType1
}
type ProcessReasonType struct {
	TextType1
}
type ProcessDescriptionType struct {
	TextType1
}
type PrizeDescriptionType struct {
	TextType1
}
type PriorityType struct {
	TextType1
}
type PrintQualifierType struct {
	TextType1
}
type PriceTypeType struct {
	TextType1
}
type PriceRevisionFormulaDescriptionType struct {
	TextType1
}
type PriceChangeReasonType struct {
	TextType1
}
type PreviousMeterReadingMethodType struct {
	TextType1
}
type PostboxType struct {
	TextType1
}
type PostalZoneType struct {
	TextType1
}
type PlotIdentificationType struct {
	TextType1
}
type PlacardNotationType struct {
	TextType1
}
type PlacardEndorsementType struct {
	TextType1
}
type PhoneNumberType struct {
}
type PersonalSituationType struct {
	TextType1
}
type PaymentOrderReferenceType struct {
	TextType1
}
type PaymentNoteType struct {
	TextType1
}
type PaymentDescriptionType struct {
	TextType1
}
type PayerReferenceType struct {
	TextType1
}
type PayPerViewType struct {
	TextType1
}
type PasswordType struct {
	TextType1
}
type PartyTypeType struct {
	TextType1
}
type PackingMaterialType struct {
	TextType1
}
type OutstandingReasonType struct {
}
type OtherInstructionType struct {
	TextType1
}
type OrganizationDepartmentType struct {
	TextType1
}
type OrderableUnitType struct {
	TextType1
}
type OptionsDescriptionType struct {
	TextType1
}
type OneTimeChargeTypeType struct {
	TextType1
}
type NoteType struct {
	TextType1
}
type NegotiationDescriptionType struct {
	TextType1
}
type NameSuffixType struct {
	TextType1
}
type MovieTitleType struct {
	TextType1
}
type MonetaryScopeType struct {
	TextType1
}
type MinimumValueType struct {
	TextType1
}
type MinimumImprovementBidType struct {
	TextType1
}
type MeterReadingTypeType struct {
	TextType1
}
type MeterReadingCommentsType struct {
	TextType1
}
type MeterNumberType struct {
	TextType1
}
type MeterNameType struct {
	TextType1
}
type MeterConstantType struct {
	TextType1
}
type MaximumValueType struct {
	TextType1
}
type MarkCareType struct {
	TextType1
}
type MarkAttentionType struct {
	TextType1
}
type LowTendersDescriptionType struct {
	TextType1
}
type LossRiskType struct {
	TextType1
}
type LoginType struct {
	TextType1
}
type LocationType struct {
	TextType1
}
type ListValueType struct {
	TextType1
}
type LineType struct {
	TextType1
}
type LimitationDescriptionType struct {
	TextType1
}
type LegalReferenceType struct {
	TextType1
}
type LatestMeterReadingMethodType struct {
	TextType1
}
type KeywordType struct {
	TextType1
}
type JustificationType struct {
	TextType1
}
type JustificationDescriptionType struct {
	TextType1
}
type JobTitleType struct {
	TextType1
}
type InvoicingPartyReferenceType struct {
	TextType1
}
type InstructionsType struct {
	TextType1
}
type InstructionNoteType struct {
	TextType1
}
type InhouseMailType struct {
	TextType1
}
type InformationType struct {
	TextType1
}
type HeatingTypeType struct {
	TextType1
}
type HaulageInstructionsType struct {
	TextType1
}
type HashAlgorithmMethodType struct {
	TextType1
}
type HandlingInstructionsType struct {
	TextType1
}
type FundingProgramType struct {
	TextType1
}
type FrequencyType struct {
	TextType1
}
type ForwarderServiceInstructionsType struct {
	TextType1
}
type FloorType struct {
	TextType1
}
type FeeDescriptionType struct {
	TextType1
}
type ExtensionType struct {
	TextType1
}
type ExpressionType struct {
	TextType1
}
type ExemptionReasonType struct {
	TextType1
}
type ExclusionReasonType struct {
	TextType1
}
type ElectronicMailType struct {
	TextType1
}
type ElectronicDeviceDescriptionType struct {
	TextType1
}
type DutyType struct {
	TextType1
}
type DocumentStatusReasonDescriptionType struct {
	TextType1
}
type DocumentHashType struct {
	TextType1
}
type DocumentDescriptionType struct {
	TextType1
}
type DistrictType struct {
	TextType1
}
type DescriptionType struct {
	TextType1
}
type DepartmentType struct {
	TextType1
}
type DemurrageInstructionsType struct {
	TextType1
}
type DeliveryInstructionsType struct {
	TextType1
}
type DataSendingCapabilityType struct {
	TextType1
}
type DamageRemarksType struct {
	TextType1
}
type CustomsClearanceServiceInstructionsType struct {
	TextType1
}
type CustomerReferenceType struct {
	TextType1
}
type CurrentChargeTypeType struct {
	TextType1
}
type CountrySubentityType struct {
	TextType1
}
type CorrectionTypeType struct {
	TextType1
}
type ContractTypeType struct {
	TextType1
}
type ContractSubdivisionType struct {
	TextType1
}
type ContractNameType struct {
	TextType1
}
type ContentType struct {
	TextType1
}
type ConsumptionTypeType struct {
	TextType1
}
type ConsumptionLevelType struct {
	TextType1
}
type ConsumersEnergyLevelType struct {
	TextType1
}
type ConditionsType struct {
	TextType1
}
type ConditionsDescriptionType struct {
	TextType1
}
type ConditionType struct {
	TextType1
}
type CompanyLegalFormType struct {
	TextType1
}
type CommentType struct {
	TextType1
}
type CodeValueType struct {
	TextType1
}
type CharacteristicsType struct {
	TextType1
}
type ChannelType struct {
	TextType1
}
type ChangeConditionsType struct {
	TextType1
}
type CertificateTypeType struct {
	TextType1
}
type CarrierServiceInstructionsType struct {
	TextType1
}
type CanonicalizationMethodType struct {
	TextType1
}
type CandidateStatementType struct {
	TextType1
}
type CancellationNoteType struct {
	TextType1
}
type CalculationExpressionType struct {
	TextType1
}
type BuyerReferenceType struct {
	TextType1
}
type BuildingNumberType struct {
	TextType1
}
type BirthplaceNameType struct {
	TextType1
}
type BackorderReasonType struct {
	TextType1
}
type AwardingCriterionDescriptionType struct {
	TextType1
}
type ApprovalStatusType struct {
	TextType1
}
type AllowanceChargeReasonType struct {
	TextType1
}
type AgencyNameType struct {
	TextType1
}
type AdditionalInformationType struct {
	TextType1
}
type AdditionalConditionsType struct {
	TextType1
}
type ActivityTypeType struct {
	TextType1
}
type AccountingCostType struct {
	TextType1
}
type AcceptedVariantsDescriptionType struct {
	TextType1
}
type AttachmentType struct {
	EmbeddedDocumentBinaryObject *EmbeddedDocumentBinaryObjectType
	ExternalReference            *ExternalReferenceType
}
type EmbeddedDocumentBinaryObjectType struct {
	BinaryObjectType1
}
type BinaryObjectType1 struct {
	BinaryObjectType
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
type VideoType struct {
	BinaryObjectType
}
type SoundType struct {
}
type PictureType struct {
	BinaryObjectType
}
type GraphicType struct {
	BinaryObjectType
}
type ExternalReferenceType struct {
	URI *IdentifierType1
}
type PeriodType struct {
	StartDate       *StartDateType
	StartTime       *StartTimeType
	EndDate         *EndDateType
	EndTime         *EndTimeType
	DurationMeasure *DurationMeasureType
	Description     *DescriptionType
}
type StartTimeType struct {
	TimeType
}

//type TimeType struct {
//	Value time.Time//System.DateTime
//}
type TimeType struct {
	time.Time
}
type ValidationTimeType struct {
	TimeType
}
type SourceForecastIssueTimeType struct {
	TimeType
}
type RevisionTimeType struct {
	TimeType
}
type ResponseTimeType struct {
	TimeType
}
type ResolutionTimeType struct {
	TimeType
}
type RequiredDeliveryTimeType struct {
	TimeType
}
type RequestedDespatchTimeType struct {
	TimeType
}
type RegisteredTimeType struct {
	TimeType
}
type ReferenceTimeType struct {
	TimeType
}
type PaidTimeType struct {
	TimeType
}
type OccurrenceTimeType struct {
	TimeType
}
type NominationTimeType struct {
	TimeType
}
type ManufactureTimeType struct {
	TimeType
}
type LatestPickupTimeType struct {
	TimeType
}
type LatestDeliveryTimeType struct {
	TimeType
}
type LastRevisionTimeType struct {
	TimeType
}
type IssueTimeType struct {
	TimeType
}
type GuaranteedDespatchTimeType struct {
	TimeType
}
type ExpiryTimeType struct {
	TimeType
}
type EstimatedDespatchTimeType struct {
	TimeType
}
type EstimatedDeliveryTimeType struct {
	TimeType
}
type EndTimeType struct {
	TimeType
}
type EffectiveTimeType struct {
	TimeType
}
type EarliestPickupTimeType struct {
	TimeType
}
type ComparisonForecastIssueTimeType struct {
	TimeType
}
type CallTimeType struct {
	TimeType
}
type AwardTimeType struct {
	TimeType
}
type ActualPickupTimeType struct {
	TimeType
}
type ActualDespatchTimeType struct {
	TimeType
}
type ActualDeliveryTimeType struct {
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
type ValueMeasureType struct {
	MeasureType1
}
type TareWeightMeasureType struct {
	MeasureType1
}
type SourceValueMeasureType struct {
	MeasureType1
}
type PreEventNotificationDurationMeasureType struct {
	MeasureType1
}
type PostEventNotificationDurationMeasureType struct {
	MeasureType1
}
type NetWeightMeasureType struct {
	MeasureType1
}
type NetVolumeMeasureType struct {
	MeasureType1
}
type NetTonnageMeasureType struct {
	MeasureType1
}
type NetNetWeightMeasureType struct {
	MeasureType1
}
type MinimumMeasureType struct {
	MeasureType1
}
type MeasureType2 struct {
	MeasureType1
}
type MaximumMeasureType struct {
	MeasureType1
}
type LongitudeMinutesMeasureType struct {
	MeasureType1
}
type LongitudeDegreesMeasureType struct {
	MeasureType1
}
type LoadingLengthMeasureType struct {
	MeasureType1
}
type LeadTimeMeasureType struct {
	MeasureType1
}
type LatitudeMinutesMeasureType struct {
	MeasureType1
}
type LatitudeDegreesMeasureType struct {
	MeasureType1
}
type GrossWeightMeasureType struct {
	MeasureType1
}
type GrossVolumeMeasureType struct {
	MeasureType1
}
type GrossTonnageMeasureType struct {
	MeasureType1
}
type ComparedValueMeasureType struct {
	MeasureType1
}
type ChargeableWeightMeasureType struct {
	MeasureType1
}
type BaseUnitMeasureType struct {
	MeasureType1
}
type AltitudeMeasureType struct {
	MeasureType1
}
type PartyType struct {
	WebsiteURI                 *IdentifierType1
	EndpointID                 *IdentifierType1
	IndustryClassificationCode *CodeType1
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
	ID IDType
}
type PartyNameType struct {
	Name NameType1
}
type AddressType struct {
	ID                  *IDType
	Postbox             *PostboxType
	Room                *RoomType
	StreetName          *StreetNameType
	BlockName           *BlockNameType
	BuildingName        *BuildingNameType
	BuildingNumber      *BuildingNumberType
	CitySubdivisionName *CitySubdivisionNameType
	CityName            *CityNameType
	PostalZone          *PostalZoneType
	Region              *RegionType
	District            *DistrictType
	Country             *CountryType
}
type CountryType struct {
	IdentificationCode *CodeType1
	Name               *NameType1
}
type PartyTaxSchemeType struct {
	TaxScheme *TaxSchemeType
}
type TaxSchemeType struct {
	Name        *NameType1
	TaxTypeCode *CodeType1
}
type PartyLegalEntityType struct {
	RegistrationName            *RegistrationNameType
	CompanyID                   *IdentifierType1
	RegistrationDate            *RegistrationDateType
	SoleProprietorshipIndicator *SoleProprietorshipIndicatorType
	CorporateStockAmount        *AmountType1
	FullyPaidSharesIndicator    *FullyPaidSharesIndicatorType
	CorporateRegistrationScheme *CorporateRegistrationSchemeType
	HeadOfficeParty             *PartyType
}
type SoleProprietorshipIndicatorType struct {
	IndicatorType
}
type IndicatorType struct {
	Value bool `xml:",chardata"`
}
type VariantConstraintIndicatorType struct {
	IndicatorType
}
type UnknownPriceIndicatorType struct {
	IndicatorType
}
type ToOrderIndicatorType struct {
	IndicatorType
}
type ThirdPartyPayerIndicatorType struct {
	IndicatorType
}
type TaxIncludedIndicatorType struct {
	IndicatorType
}
type TaxEvidenceIndicatorType struct {
	IndicatorType
}
type StatusAvailableIndicatorType struct {
	IndicatorType
}
type SplitConsignmentIndicatorType struct {
	IndicatorType
}
type SpecialSecurityIndicatorType struct {
	IndicatorType
}
type ReturnableMaterialIndicatorType struct {
	IndicatorType
}
type ReturnabilityIndicatorType struct {
	IndicatorType
}
type RequiredCurriculaIndicatorType struct {
	IndicatorType
}
type RefrigerationOnIndicatorType struct {
	IndicatorType
}
type RefrigeratedIndicatorType struct {
	IndicatorType
}
type PublishAwardIndicatorType struct {
	IndicatorType
}
type PrizeIndicatorType struct {
	IndicatorType
}
type PricingUpdateRequestIndicatorType struct {
	IndicatorType
}
type PrepaidIndicatorType struct {
	IndicatorType
}
type PreCarriageIndicatorType struct {
	IndicatorType
}
type PowerIndicatorType struct {
	IndicatorType
}
type PartialDeliveryIndicatorType struct {
	IndicatorType
}
type OtherConditionsIndicatorType struct {
	IndicatorType
}
type OrderableIndicatorType struct {
	IndicatorType
}
type OptionalLineItemIndicatorType struct {
	IndicatorType
}
type OnCarriageIndicatorType struct {
	IndicatorType
}
type MarkCareIndicatorType struct {
	IndicatorType
}
type MarkAttentionIndicatorType struct {
	IndicatorType
}
type LivestockIndicatorType struct {
	IndicatorType
}
type LegalStatusIndicatorType struct {
	IndicatorType
}
type ItemUpdateRequestIndicatorType struct {
	IndicatorType
}
type IndicationIndicatorType struct {
	IndicatorType
}
type HumanFoodIndicatorType struct {
	IndicatorType
}
type HumanFoodApprovedIndicatorType struct {
	IndicatorType
}
type HazardousRiskIndicatorType struct {
	IndicatorType
}
type GovernmentAgreementConstraintIndicatorType struct {
	IndicatorType
}
type GeneralCargoIndicatorType struct {
	IndicatorType
}
type FullyPaidSharesIndicatorType struct {
	IndicatorType
}
type FrozenDocumentIndicatorType struct {
	IndicatorType
}
type FreeOfChargeIndicatorType struct {
	IndicatorType
}
type FollowupContractIndicatorType struct {
	IndicatorType
}
type DangerousGoodsApprovedIndicatorType struct {
	IndicatorType
}
type CustomsImportClassifiedIndicatorType struct {
	IndicatorType
}
type CopyIndicatorType struct {
	IndicatorType
}
type ContainerizedIndicatorType struct {
	IndicatorType
}
type ConsolidatableIndicatorType struct {
	IndicatorType
}
type CompletionIndicatorType struct {
	IndicatorType
}
type ChargeIndicatorType struct {
	IndicatorType
}
type CatalogueIndicatorType struct {
	IndicatorType
}
type CandidateReductionConstraintIndicatorType struct {
	IndicatorType
}
type BulkCargoIndicatorType struct {
	IndicatorType
}
type BindingOnBuyerIndicatorType struct {
	IndicatorType
}
type BasedOnConsensusIndicatorType struct {
	IndicatorType
}
type BalanceBroughtForwardIndicatorType struct {
	IndicatorType
}
type BackOrderAllowedIndicatorType struct {
	IndicatorType
}
type AuctionConstraintIndicatorType struct {
	IndicatorType
}
type AnimalFoodIndicatorType struct {
	IndicatorType
}
type AnimalFoodApprovedIndicatorType struct {
	IndicatorType
}
type AdValoremIndicatorType struct {
	IndicatorType
}
type AcceptedIndicatorType struct {
	IndicatorType
}
type CorporateRegistrationSchemeType struct {
	ID                            *IDType
	Name                          *NameType1
	CorporateRegistrationTypeCode *CodeType1
	JurisdictionRegionAddress     []AddressType
}
type ContactType struct {
	Telephone          *TelephoneType
	Telefax            *TelefaxType
	ElectronicMail     *ElectronicMailType
	Note               *NoteType
	OtherCommunication []CommunicationType
}
type CommunicationType struct {
	ChannelCode *CodeType1
	Channel     *ChannelType
	Value       *ValueType
}
type PersonType struct {
	FirstName                 *FirstNameType
	FamilyName                *FamilyNameType
	Title                     *TitleType
	MiddleName                *MiddleNameType
	NameSuffix                *NameSuffixType
	NationalityID             *IdentifierType1
	FinancialAccount          *FinancialAccountType
	IdentityDocumentReference *DocumentReferenceType
}
type FinancialAccountType struct {
	ID                         *IDType
	CurrencyCode               *CodeType1
	PaymentNote                *PaymentNoteType
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
	UUID             *UUIDType
	LineStatusCode   *LineStatusCodeType
	OrderReference   *OrderReferenceType
}
type OrderReferenceType struct {
	ID                *IDType
	SalesOrderID      *IdentifierType1
	IssueDate         *IssueDateType
	OrderTypeCode     *CodeType1
	DocumentReference *DocumentReferenceType
}
type InvoiceLineType struct {
	ID                    IDType
	Note                  []NoteType
	InvoicedQuantity      *InvoicedQuantityType
	LineExtensionAmount   *AmountType1
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
type InvoicedQuantityType struct {
	QuantityType1
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
type VarianceQuantityType struct {
	QuantityType1
}
type ValueQuantityType struct {
	QuantityType1
}
type TotalTransportHandlingUnitQuantityType struct {
	QuantityType1
}
type TotalPackagesQuantityType struct {
	QuantityType1
}
type TotalPackageQuantityType struct {
	QuantityType1
}
type TotalMeteredQuantityType struct {
	QuantityType1
}
type TotalGoodsItemQuantityType struct {
	QuantityType1
}
type TotalDeliveredQuantityType struct {
	QuantityType1
}
type TotalConsumedQuantityType struct {
	QuantityType1
}
type TimeDeltaDaysQuantityType struct {
	QuantityType1
}
type ThresholdQuantityType struct {
	QuantityType1
}
type TargetInventoryQuantityType struct {
	QuantityType1
}
type ShortQuantityType struct {
	QuantityType1
}
type SharesNumberQuantityType struct {
	QuantityType1
}
type ReturnableQuantityType struct {
	QuantityType1
}
type RejectedQuantityType struct {
	QuantityType1
}
type ReceivedTenderQuantityType struct {
	QuantityType1
}
type ReceivedQuantityType struct {
	QuantityType1
}
type ReceivedForeignTenderQuantityType struct {
	QuantityType1
}
type ReceivedElectronicTenderQuantityType struct {
	QuantityType1
}
type QuantityType2 struct {
	QuantityType1
}
type PreviousMeterQuantityType struct {
	QuantityType1
}
type PerformanceValueQuantityType struct {
	QuantityType1
}
type PassengerQuantityType struct {
	QuantityType1
}
type PackQuantityType struct {
	QuantityType1
}
type OversupplyQuantityType struct {
	QuantityType1
}
type OutstandingQuantityType struct {
	QuantityType1
}
type OperatingYearsQuantityType struct {
	QuantityType1
}
type NormalTemperatureReductionQuantityType struct {
	QuantityType1
}
type MultipleOrderQuantityType struct {
	QuantityType1
}
type MinimumQuantityType struct {
	QuantityType1
}
type MinimumOrderQuantityType struct {
	QuantityType1
}
type MinimumInventoryQuantityType struct {
	QuantityType1
}
type MinimumBackorderQuantityType struct {
	QuantityType1
}
type MaximumVariantQuantityType struct {
	QuantityType1
}
type MaximumQuantityType struct {
	QuantityType1
}
type MaximumOrderQuantityType struct {
	QuantityType1
}
type MaximumOperatorQuantityType struct {
	QuantityType1
}
type MaximumBackorderQuantityType struct {
	QuantityType1
}
type LatestMeterQuantityType struct {
	QuantityType1
}
type GasPressureQuantityType struct {
	QuantityType1
}
type ExpectedQuantityType struct {
	QuantityType1
}
type ExpectedOperatorQuantityType struct {
	QuantityType1
}
type EstimatedOverallContractQuantityType struct {
	QuantityType1
}
type EstimatedConsumedQuantityType struct {
	QuantityType1
}
type EmployeeQuantityType struct {
	QuantityType1
}
type DifferenceTemperatureReductionQuantityType struct {
	QuantityType1
}
type DeliveredQuantityType struct {
	QuantityType1
}
type DebitedQuantityType struct {
	QuantityType1
}
type CustomsTariffQuantityType struct {
	QuantityType1
}
type CrewQuantityType struct {
	QuantityType1
}
type CreditedQuantityType struct {
	QuantityType1
}
type ContentUnitQuantityType struct {
	QuantityType1
}
type ConsumptionWaterQuantityType struct {
	QuantityType1
}
type ConsumptionEnergyQuantityType struct {
	QuantityType1
}
type ConsumerUnitQuantityType struct {
	QuantityType1
}
type ConsignmentQuantityType struct {
	QuantityType1
}
type ChildConsignmentQuantityType struct {
	QuantityType1
}
type ChargeableQuantityType struct {
	QuantityType1
}
type BatchQuantityType struct {
	QuantityType1
}
type BasicConsumedQuantityType struct {
	QuantityType1
}
type BaseQuantityType struct {
	QuantityType1
}
type BackorderQuantityType struct {
	QuantityType1
}
type ActualTemperatureReductionQuantityType struct {
	QuantityType1
}
type DeliveryType struct {
	ID                          *IDType
	Quantity                    *QuantityType2
	ActualDeliveryDate          *ActualDeliveryDateType
	ActualDeliveryTime          *ActualDeliveryTimeType
	LatestDeliveryDate          *LatestDeliveryDateType
	LatestDeliveryTime          *LatestDeliveryTimeType
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
	ID      *IDType
	Address *AddressType
}
type DespatchType struct {
	ID                      *IDType
	ActualDespatchDate      *ActualDespatchDateType
	ActualDespatchTime      *ActualDespatchTimeType
	Instructions            *InstructionsType
	DespatchAddress         *AddressType
	DespatchParty           *PartyType
	Contact                 *ContactType
	EstimatedDespatchPeriod *PeriodType
}
type DeliveryTermsType struct {
	ID           *IDType
	SpecialTerms *SpecialTermsType
	Amount       *AmountType2
}
type ShipmentType struct {
	ID                                 *IDType
	HandlingCode                       *CodeType1
	HandlingInstructions               *HandlingInstructionsType
	GrossWeightMeasure                 *GrossWeightMeasureType
	NetWeightMeasure                   *NetWeightMeasureType
	GrossVolumeMeasure                 *GrossVolumeMeasureType
	NetVolumeMeasure                   *NetVolumeMeasureType
	TotalGoodsItemQuantity             *TotalGoodsItemQuantityType
	TotalTransportHandlingUnitQuantity *TotalTransportHandlingUnitQuantityType
	InsuranceValueAmount               *AmountType1
	DeclaredCustomsValueAmount         *AmountType1
	DeclaredForCarriageValueAmount     *AmountType1
	DeclaredStatisticsValueAmount      *AmountType1
	FreeOnBoardValueAmount             *AmountType1
	SpecialInstructions                []SpecialInstructionsType
	TransportHandlingUnit              []TransportHandlingUnitType
	ReturnAddress                      *AddressType
	FirstArrivalPortLocation           *LocationType1
	LastExitPortLocation               *LocationType1
}
type TransportHandlingUnitType struct {
	ID                              *IDType
	TransportHandlingUnitTypeCode   *CodeType1
	HandlingCode                    *CodeType1
	HandlingInstructions            *HandlingInstructionsType
	HazardousRiskIndicator          *HazardousRiskIndicatorType
	TotalGoodsItemQuantity          *TotalGoodsItemQuantityType
	TotalPackageQuantity            *TotalPackageQuantityType
	DamageRemarks                   []DamageRemarksType
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
	CustomsDeclaration          *IDType
	Quantity                    *QuantityType2
	ReturnableMaterialIndicator *ReturnableMaterialIndicatorType
	PackageLevelCode            *CodeType1
	PackagingTypeCode           *CodeType1
	PackingMaterial             []PackingMaterialType
	ContainedPackage            []PackageType
	GoodsItem                   []GoodsItemType
	MeasurementDimension        []DimensionType
}
type GoodsItemType struct {
	ID                               *IDType
	Description                      []DescriptionType
	HazardousRiskIndicator           *HazardousRiskIndicatorType
	DeclaredCustomsValueAmount       *AmountType1
	DeclaredForCarriageValueAmount   *AmountType1
	DeclaredStatisticsValueAmount    *AmountType1
	FreeOnBoardValueAmount           *AmountType1
	InsuranceValueAmount             *AmountType1
	ValueAmount                      *AmountType1
	GrossWeightMeasure               *GrossWeightMeasureType
	NetWeightMeasure                 *NetWeightMeasureType
	ChargeableWeightMeasure          *ChargeableWeightMeasureType
	GrossVolumeMeasure               *GrossVolumeMeasureType
	NetVolumeMeasure                 *NetVolumeMeasureType
	Quantity                         *QuantityType2
	RequiredCustomsID                *IdentifierType1
	CustomsStatusCode                *CodeType1
	CustomsTariffQuantity            *CustomsTariffQuantityType
	CustomsImportClassifiedIndicator *CustomsImportClassifiedIndicatorType
	ChargeableQuantity               *ChargeableQuantityType
	ReturnableQuantity               *ReturnableQuantityType
	TraceID                          *IdentifierType1
	Item                             []ItemType
	FreightAllowanceCharge           []AllowanceChargeType
	Temperature                      []TemperatureType
	OriginAddress                    *AddressType
	MeasurementDimension             []DimensionType
}
type ItemType struct {
	Description                     *DescriptionType
	Name                            *NameType1
	Keyword                         *KeywordType
	BrandName                       *BrandNameType
	ModelName                       *ModelNameType
	BuyersItemIdentification        *ItemIdentificationType
	SellersItemIdentification       *ItemIdentificationType
	ManufacturersItemIdentification *ItemIdentificationType
	AdditionalItemIdentification    []ItemIdentificationType
	CommodityClassification         []CommodityClassificationType
	ItemInstance                    []ItemInstanceType
}
type ItemIdentificationType struct {
	ID *IDType
}
type CommodityClassificationType struct {
	ItemClassificationCode *CodeType1
}
type ItemInstanceType struct {
	ProductTraceID         *IdentifierType1
	ManufactureDate        *ManufactureDateType
	ManufactureTime        *ManufactureTimeType
	BestBeforeDate         *BestBeforeDateType
	RegistrationID         *IdentifierType1
	SerialID               *IdentifierType1
	AdditionalItemProperty []ItemPropertyType
	LotIdentification      *LotIdentificationType
}
type ItemPropertyType struct {
	ID                *IDType
	Name              *NameType1
	NameCode          *CodeType1
	TestMethod        *TestMethodType
	Value             *ValueType
	ValueQuantity     *ValueQuantityType
	ValueQualifier    []ValueQualifierType
	ImportanceCode    *CodeType1
	ListValue         []ListValueType
	UsabilityPeriod   *PeriodType
	ItemPropertyGroup []ItemPropertyGroupType
	RangeDimension    *DimensionType
	ItemPropertyRange *ItemPropertyRangeType
}
type ItemPropertyGroupType struct {
	ID             *IDType
	Name           *NameType1
	ImportanceCode *CodeType1
}
type DimensionType struct {
	AttributeID    *IdentifierType1
	Measure        *MeasureType2
	Description    []DescriptionType
	MinimumMeasure *MinimumMeasureType
	MaximumMeasure *MaximumMeasureType
}
type ItemPropertyRangeType struct {
	MinimumValue *MinimumValueType
	MaximumValue *MaximumValueType
}
type LotIdentificationType struct {
	LotNumberID            *IdentifierType1
	ExpiryDate             *ExpiryDateType
	AdditionalItemProperty []ItemPropertyType
}
type AllowanceChargeType struct {
	ChargeIndicator         *ChargeIndicatorType
	AllowanceChargeReason   *AllowanceChargeReasonType
	MultiplierFactorNumeric *MultiplierFactorNumericType
	SequenceNumeric         *SequenceNumericType
	Amount                  *AmountType2
	BaseAmount              *AmountType1
	PerUnitAmount           *AmountType1
}
type MultiplierFactorNumericType struct {
	NumericType1
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
type TargetCurrencyBaseRateType struct {
	RateType
}
type SourceCurrencyBaseRateType struct {
	RateType
}
type RateType1 struct {
	RateType
}
type OrderableUnitFactorRateType struct {
	RateType
}
type CalculationRateType struct {
	RateType
}
type AmountRateType struct {
	RateType
}
type PercentType struct {
	NumericType
}
type TierRatePercentType struct {
	PercentType
}
type TargetServicePercentType struct {
	PercentType
}
type SettlementDiscountPercentType struct {
	PercentType
}
type ReliabilityPercentType struct {
	PercentType
}
type ProgressPercentType struct {
	PercentType
}
type PercentType1 struct {
	PercentType
}
type PenaltySurchargePercentType struct {
	PercentType
}
type PaymentPercentType struct {
	PercentType
}
type ParticipationPercentType struct {
	PercentType
}
type PartecipationPercentType struct {
	PercentType
}
type MinimumPercentType struct {
	PercentType
}
type MaximumPercentType struct {
	PercentType
}
type HumidityPercentType struct {
	PercentType
}
type AirFlowPercentType struct {
	PercentType
}
type ValueType1 struct {
	NumericType
}
type WeightNumericType struct {
	NumericType1
}
type SequenceNumericType struct {
	NumericType1
}
type ResidentOccupantsNumericType struct {
	NumericType1
}
type ReminderSequenceNumericType struct {
	NumericType1
}
type PackSizeNumericType struct {
	NumericType1
}
type OrderQuantityIncrementNumericType struct {
	NumericType1
}
type OrderIntervalDaysNumericType struct {
	NumericType1
}
type MinimumNumberNumericType struct {
	NumericType1
}
type MaximumPaymentInstructionsNumericType struct {
	NumericType1
}
type MaximumNumberNumericType struct {
	NumericType1
}
type MaximumCopiesNumericType struct {
	NumericType1
}
type LineNumberNumericType struct {
	NumericType1
}
type LineCountNumericType struct {
	NumericType1
}
type FrozenPeriodDaysNumericType struct {
	NumericType1
}
type CalculationSequenceNumericType struct {
	NumericType1
}
type BudgetYearNumericType struct {
	NumericType1
}
type TemperatureType struct {
	AttributeID *IdentifierType1
	Measure     *MeasureType2
	Description []DescriptionType
}
type TransportEquipmentType struct {
	ID                         *IDType
	TransportEquipmentTypeCode *CodeType1
	Description                *DescriptionType
}
type TransportMeansType struct {
	JourneyID                 *IdentifierType1
	RegistrationNationalityID *IdentifierType1
	RegistrationNationality   []RegistrationNationalityType
	DirectionCode             *CodeType1
	TransportMeansTypeCode    *CodeType1
	TradeServiceCode          *CodeType1
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
	Location             []LocationType
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
	VesselName                           *VesselNameType
	RadioCallSignID                      *IdentifierType1
	ShipsRequirements                    []ShipsRequirementsType
	GrossTonnageMeasure                  *GrossTonnageMeasureType
	NetTonnageMeasure                    *NetTonnageMeasureType
	RegistryCertificateDocumentReference *DocumentReferenceType
	RegistryPortLocation                 *LocationType1
}
type HazardousGoodsTransitType struct {
	TransportEmergencyCardCode *CodeType1
	PackingCriteriaCode        *CodeType1
	HazardousRegulationCode    *CodeType1
	InhalationToxicityZoneCode *CodeType1
	TransportAuthorizationCode *CodeType1
	MaximumTemperature         *TemperatureType
	MinimumTemperature         *TemperatureType
}
type CustomsDeclarationType struct {
	ID          *IDType
	IssuerParty *PartyType
}
type TaxTotalType struct {
	TaxAmount   *AmountType1
	TaxSubtotal []TaxSubtotalType
}
type TaxSubtotalType struct {
	TaxableAmount                *AmountType1
	TaxAmount                    *AmountType1
	CalculationSequenceNumeric   *CalculationSequenceNumericType
	TransactionCurrencyTaxAmount *AmountType1
	Percent                      *PercentType1
	BaseUnitMeasure              *BaseUnitMeasureType
	PerUnitAmount                *AmountType1
	TaxCategory                  *TaxCategoryType
}
type TaxCategoryType struct {
	Name                   *NameType1
	TaxExemptionReasonCode *CodeType1
	TaxExemptionReason     *TaxExemptionReasonType
	TaxScheme              *TaxSchemeType
}
type MonetaryTotalType struct {
	LineExtensionAmount   *AmountType1
	TaxExclusiveAmount    *AmountType1
	TaxInclusiveAmount    *AmountType1
	AllowanceTotalAmount  *AmountType1
	ChargeTotalAmount     *AmountType1
	PayableRoundingAmount *AmountType1
	PayableAmount         *AmountType1
}
type ExchangeRateType struct {
	SourceCurrencyCode *CodeType1
	TargetCurrencyCode *CodeType1
	CalculationRate    *CalculationRateType
	Date               *DateType1
}
type PaymentTermsType struct {
	Note                    *NoteType
	PenaltySurchargePercent *PenaltySurchargePercentType
	Amount                  *AmountType2
	PenaltyAmount           *AmountType1
	PaymentDueDate          *PaymentDueDateType
	SettlementPeriod        *PeriodType
}
type PaymentMeansType struct {
	PaymentMeansCode      *CodeType1
	PaymentDueDate        *PaymentDueDateType
	PaymentChannelCode    *CodeType1
	InstructionNote       *InstructionNoteType
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
	ID                         *IDType
	SignatoryParty             *PartyType
	DigitalSignatureAttachment *AttachmentType
}
type BillingReferenceLineType struct {
	ID              *IDType
	Amount          *AmountType2
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
type TransportEventType struct {
	IdentificationID       *IdentifierType1
	OccurrenceDate         *OccurrenceDateType
	OccurrenceTime         *OccurrenceTimeType
	TransportEventTypeCode *CodeType1
	Description            []DescriptionType
	CompletionIndicator    *CompletionIndicatorType
	ReportedShipment       *ShipmentType
	CurrentStatus          []StatusType
	Contact                []ContactType
	Location               *LocationType1
	Signature              *SignatureType
	Period                 []PeriodType
}
type StatusType struct {
	ConditionCode       *CodeType1
	ReferenceDate       *ReferenceDateType
	ReferenceTime       *ReferenceTimeType
	Description         []DescriptionType
	StatusReasonCode    *CodeType1
	StatusReason        []StatusReasonType
	SequenceID          *IdentifierType1
	Text                []TextType2
	IndicationIndicator *IndicationIndicatorType
	Percent             *PercentType1
	ReliabilityPercent  *ReliabilityPercentType
	Condition           []ConditionType1
}
type ConditionType1 struct {
	AttributeID    *IdentifierType1
	Measure        *MeasureType2
	Description    []DescriptionType
	MinimumMeasure *MinimumMeasureType
	MaximumMeasure *MaximumMeasureType
}
type RelatedItemType struct {
	ID          *IDType
	Quantity    *QuantityType2
	Description []DescriptionType
}
type ActivityDataLineType struct {
	ID                          *IDType
	SupplyChainActivityTypeCode *CodeType1
	BuyerCustomerParty          *CustomerPartyType
	SellerSupplierParty         *SupplierPartyType
	ActivityPeriod              *PeriodType
	ActivityOriginLocation      *LocationType1
	ActivityFinalLocation       *LocationType1
	SalesItem                   []SalesItemType
}
type SalesItemType struct {
	Quantity          *QuantityType2
	ActivityProperty  []ActivityPropertyType
	TaxExclusivePrice []PriceType
	TaxInclusivePrice []PriceType
	Item              *ItemType
}
type ActivityPropertyType struct {
	Name  *NameType1
	Value *ValueType
}
type DocumentResponseType struct {
	Response          *ResponseType
	DocumentReference *DocumentReferenceType
	LineResponse      []LineResponseType
}
type ResponseType struct {
	ReferenceID  *IdentifierType1
	ResponseCode *CodeType1
	Description  []DescriptionType
}
type LineResponseType struct {
	LineReference *LineReferenceType
	Response      []ResponseType
}
type QualifyingPartyType struct {
	ParticipationPercent             *ParticipationPercentType
	PersonalSituation                []PersonalSituationType
	OperatingYearsQuantity           *OperatingYearsQuantityType
	EmployeeQuantity                 *EmployeeQuantityType
	BusinessClassificationEvidenceID *IdentifierType1
	BusinessIdentityEvidenceID       *IdentifierType1
	TendererRoleCode                 *CodeType1
	BusinessClassificationScheme     *ClassificationSchemeType
	TechnicalCapability              []CapabilityType
	FinancialCapability              []CapabilityType
	CompletedTask                    []CompletedTaskType
	Declaration                      []DeclarationType
	Party                            *PartyType
	EconomicOperatorRole             *EconomicOperatorRoleType
}
type ClassificationSchemeType struct {
	ID                     *IDType
	UUID                   *UUIDType
	LastRevisionDate       *LastRevisionDateType
	LastRevisionTime       *LastRevisionTimeType
	Note                   []NoteType
	Name                   *NameType1
	Description            []DescriptionType
	AgencyID               *IdentifierType1
	AgencyName             *AgencyNameType
	VersionID              *IdentifierType1
	URI                    *IdentifierType1
	SchemeURI              *IdentifierType1
	LanguageID             *IdentifierType1
	ClassificationCategory []ClassificationCategoryType
}
type ClassificationCategoryType struct {
	Name                              *NameType1
	CodeValue                         *CodeValueType
	Description                       []DescriptionType
	CategorizesClassificationCategory []ClassificationCategoryType
}
type CapabilityType struct {
	CapabilityTypeCode *CodeType1
	Description        []DescriptionType
	ValueAmount        *AmountType1
	ValueQuantity      *ValueQuantityType
	EvidenceSupplied   []EvidenceSuppliedType
	ValidityPeriod     *PeriodType
}
type EvidenceSuppliedType struct {
	ID *IDType
}
type CompletedTaskType struct {
	AnnualAverageAmount *AmountType1
	TotalTaskAmount     *AmountType1
	//Burayı bir kontrol et
	//PartyCapacityAmountType
	Description            []DescriptionType
	EvidenceSupplied       []EvidenceSuppliedType
	Period                 *PeriodType
	RecipientCustomerParty *CustomerPartyType
}
type DeclarationType struct {
	Name                []NameType1
	DeclarationTypeCode *CodeType1
	Description         []DescriptionType
	EvidenceSupplied    []EvidenceSuppliedType
}
type EconomicOperatorRoleType struct {
	RoleCode        *CodeType1
	RoleDescription []RoleDescriptionType
}
type TransportationServiceType struct {
	TransportServiceCode *CodeType1

	TariffClassCode                          *CodeType1
	Priority                                 *PriorityType
	FreightRateClassCode                     *CodeType1
	TransportationServiceDescription         []TransportationServiceDescriptionType
	TransportationServiceDetailsURI          *IdentifierType1
	NominationDate                           *NominationDateType
	NominationTime                           *NominationTimeType
	Name                                     *NameType1
	SequenceNumeric                          *SequenceNumericType
	TransportEquipment                       []TransportEquipmentType
	SupportedTransportEquipment              []TransportEquipmentType
	UnsupportedTransportEquipment            []TransportEquipmentType
	CommodityClassification                  []CommodityClassificationType
	SupportedCommodityClassification         []CommodityClassificationType
	UnsupportedCommodityClassification       []CommodityClassificationType
	TotalCapacityDimension                   *DimensionType
	ShipmentStage                            []ShipmentStageType
	TransportEvent                           []TransportEventType
	ResponsibleTransportServiceProviderParty *PartyType
	EnvironmentalEmission                    []EnvironmentalEmissionType
	EstimatedDurationPeriod                  *PeriodType
	ScheduledServiceFrequency                []ServiceFrequencyType
}
type ShipmentStageType struct {
	ID                               *IDType
	TransportModeCode                *CodeType1
	TransportMeansTypeCode           *CodeType1
	TransitDirectionCode             *CodeType1
	PreCarriageIndicator             *PreCarriageIndicatorType
	OnCarriageIndicator              *OnCarriageIndicatorType
	EstimatedDeliveryDate            *EstimatedDeliveryDateType
	EstimatedDeliveryTime            *EstimatedDeliveryTimeType
	RequiredDeliveryDate             *RequiredDeliveryDateType
	RequiredDeliveryTime             *RequiredDeliveryTimeType
	LoadingSequenceID                *IdentifierType1
	SuccessiveSequenceID             *IdentifierType1
	Instructions                     []InstructionsType
	DemurrageInstructions            []DemurrageInstructionsType
	CrewQuantity                     *CrewQuantityType
	PassengerQuantity                *PassengerQuantityType
	TransitPeriod                    *PeriodType
	CarrierParty                     []PartyType
	TransportMeans                   *TransportMeansType
	LoadingPortLocation              *LocationType1
	UnloadingPortLocation            *LocationType1
	TransshipPortLocation            *LocationType1
	LoadingTransportEvent            *TransportEventType
	ExaminationTransportEvent        *TransportEventType
	AvailabilityTransportEvent       *TransportEventType
	ExportationTransportEvent        *TransportEventType
	DischargeTransportEvent          *TransportEventType
	WarehousingTransportEvent        *TransportEventType
	TakeoverTransportEvent           *TransportEventType
	OptionalTakeoverTransportEvent   *TransportEventType
	DropoffTransportEvent            *TransportEventType
	ActualPickupTransportEvent       *TransportEventType
	DeliveryTransportEvent           *TransportEventType
	ReceiptTransportEvent            *TransportEventType
	StorageTransportEvent            *TransportEventType
	AcceptanceTransportEvent         *TransportEventType
	TerminalOperatorParty            *PartyType
	CustomsAgentParty                *PartyType
	EstimatedTransitPeriod           *PeriodType
	FreightAllowanceCharge           []AllowanceChargeType
	FreightChargeLocation            *LocationType1
	DetentionTransportEvent          []TransportEventType
	RequestedDepartureTransportEvent *TransportEventType
	RequestedArrivalTransportEvent   *TransportEventType
	RequestedWaypointTransportEvent  []TransportEventType
	PlannedDepartureTransportEvent   *TransportEventType
	PlannedArrivalTransportEvent     *TransportEventType
	PlannedWaypointTransportEvent    []TransportEventType
	ActualDepartureTransportEvent    *TransportEventType
	ActualWaypointTransportEvent     *TransportEventType
	ActualArrivalTransportEvent      *TransportEventType
	TransportEvent                   []TransportEventType
	EstimatedDepartureTransportEvent *TransportEventType
	EstimatedArrivalTransportEvent   *TransportEventType
	PassengerPerson                  []PersonType
	DriverPerson                     []PersonType
	ReportingPerson                  *PersonType
	CrewMemberPerson                 []PersonType
	SecurityOfficerPerson            *PersonType
	MasterPerson                     *PersonType
	ShipsSurgeonPerson               *PersonType
}
type EnvironmentalEmissionType struct {
	EnvironmentalEmissionTypeCode *CodeType1
	ValueMeasure                  *ValueMeasureType
	Description                   []DescriptionType
	EmissionCalculationMethod     []EmissionCalculationMethodType
}
type EmissionCalculationMethodType struct {
	CalculationMethodCode   *CodeType1
	FullnessIndicationCode  *CodeType1
	MeasurementFromLocation *LocationType1
	MeasurementToLocation   *LocationType1
}
type ServiceFrequencyType struct {
	WeekDayCode *CodeType1
}
type AddressLineType struct {
	Line *LineType
}
type SubcontractTermsType struct {
	Rate                         *RateType1
	UnknownPriceIndicator        *UnknownPriceIndicatorType
	Description                  []DescriptionType
	Amount                       *AmountType2
	SubcontractingConditionsCode *CodeType1
	MaximumPercent               *MaximumPercentType
	MinimumPercent               *MinimumPercentType
}
type LineItemType struct {
	ID                        *IDType
	SalesOrderID              *IdentifierType1
	UUID                      *UUIDType
	Note                      []NoteType
	LineStatusCode            *LineStatusCodeType
	Quantity                  *QuantityType2
	LineExtensionAmount       *AmountType1
	TotalTaxAmount            *AmountType1
	MinimumQuantity           *MinimumQuantityType
	MaximumQuantity           *MaximumQuantityType
	MinimumBackorderQuantity  *MinimumBackorderQuantityType
	MaximumBackorderQuantity  *MaximumBackorderQuantityType
	InspectionMethodCode      *CodeType1
	PartialDeliveryIndicator  *PartialDeliveryIndicatorType
	BackOrderAllowedIndicator *BackOrderAllowedIndicatorType
	AccountingCostCode        *CodeType1
	AccountingCost            *AccountingCostType
	WarrantyInformation       []WarrantyInformationType
	Delivery                  []DeliveryType
	DeliveryTerms             *DeliveryTermsType
	OriginatorParty           *PartyType
	OrderedShipment           []OrderedShipmentType
	PricingReference          *PricingReferenceType
	AllowanceCharge           []AllowanceChargeType
	Price                     *PriceType
	Item                      *ItemType
	SubLineItem               []LineItemType
	WarrantyValidityPeriod    *PeriodType
	WarrantyParty             *PartyType
	TaxTotal                  []TaxTotalType
	ItemPriceExtension        *PriceExtensionType
	LineReference             []LineReferenceType
}
type OrderedShipmentType struct {
	Shipment *ShipmentType
	Package  []PackageType
}
type PricingReferenceType struct {
	OriginalItemLocationQuantity *ItemLocationQuantityType
	AlternativeConditionPrice    []PriceType
}
type ItemLocationQuantityType struct {
	LeadTimeMeasure            *LeadTimeMeasureType
	MinimumQuantity            *MinimumQuantityType
	MaximumQuantity            *MaximumQuantityType
	HazardousRiskIndicator     *HazardousRiskIndicatorType
	TradingRestrictions        []TradingRestrictionsType
	ApplicableTerritoryAddress []AddressType
	Price                      *PriceType
	DeliveryUnit               []DeliveryUnitType
	ApplicableTaxCategory      []TaxCategoryType
	Package                    *PackageType
	AllowanceCharge            []AllowanceChargeType
	DependentPriceReference    *DependentPriceReferenceType
}
type DeliveryUnitType struct {
	BatchQuantity          *BatchQuantityType
	ConsumerUnitQuantity   *ConsumerUnitQuantityType
	HazardousRiskIndicator *HazardousRiskIndicatorType
}
type DependentPriceReferenceType struct {
	Percent                *PercentType1
	LocationAddress        *AddressType
	DependentLineReference *LineReferenceType
}
type PriceExtensionType struct {
	Amount   *AmountType2
	TaxTotal []TaxTotalType
}
type AppealTermsType struct {
	Description            []DescriptionType
	PresentationPeriod     *PeriodType
	AppealInformationParty *PartyType
	AppealReceiverParty    *PartyType
	MediationParty         *PartyType
}
type RegulationType struct {
	Name           *NameType1
	LegalReference *LegalReferenceType
	OntologyURI    *IdentifierType1
}
type AuctionTermsType struct {
	AuctionConstraintIndicator  *AuctionConstraintIndicatorType
	JustificationDescription    []JustificationDescriptionType
	Description                 []DescriptionType
	ProcessDescription          []ProcessDescriptionType
	ConditionsDescription       []ConditionsDescriptionType
	ElectronicDeviceDescription []ElectronicDeviceDescriptionType
	AuctionURI                  *IdentifierType1
}
type TenderedProjectType struct {
	VariantID                 *IdentifierType1
	FeeAmount                 *AmountType1
	FeeDescription            []FeeDescriptionType
	TenderEnvelopeID          *IdentifierType1
	TenderEnvelopeTypeCode    *CodeType1
	ProcurementProjectLot     *ProcurementProjectLotType
	EvidenceDocumentReference []DocumentReferenceType
	TaxTotal                  []TaxTotalType
	LegalMonetaryTotal        *MonetaryTotalType
	TenderLine                []TenderLineType
	AwardingCriterionResponse []AwardingCriterionResponseType
}
type ProcurementProjectLotType struct {
	ID                 *IDType
	TenderingTerms     *TenderingTermsType
	ProcurementProject *ProcurementProjectType
}
type TenderingTermsType struct {
	AwardingMethodTypeCode                    *CodeType1
	PriceEvaluationCode                       *CodeType1
	MaximumVariantQuantity                    *MaximumVariantQuantityType
	VariantConstraintIndicator                *VariantConstraintIndicatorType
	AcceptedVariantsDescription               []AcceptedVariantsDescriptionType
	PriceRevisionFormulaDescription           []PriceRevisionFormulaDescriptionType
	FundingProgramCode                        *CodeType1
	FundingProgram                            []FundingProgramType
	MaximumAdvertisementAmount                *AmountType1
	Note                                      []NoteType
	PaymentFrequencyCode                      *CodeType1
	EconomicOperatorRegistryURI               *IdentifierType1
	RequiredCurriculaIndicator                *RequiredCurriculaIndicatorType
	OtherConditionsIndicator                  *OtherConditionsIndicatorType
	AdditionalConditions                      []AdditionalConditionsType
	LatestSecurityClearanceDate               *LatestSecurityClearanceDateType
	DocumentationFeeAmount                    *AmountType1
	PenaltyClause                             []ClauseType
	RequiredFinancialGuarantee                []FinancialGuaranteeType
	ProcurementLegislationDocumentReference   *DocumentReferenceType
	FiscalLegislationDocumentReference        *DocumentReferenceType
	EnvironmentalLegislationDocumentReference *DocumentReferenceType
	EmploymentLegislationDocumentReference    *DocumentReferenceType
	ContractualDocumentReference              []DocumentReferenceType
	CallForTendersDocumentReference           *DocumentReferenceType
	WarrantyValidityPeriod                    *PeriodType
	PaymentTerms                              []PaymentTermsType
	TendererQualificationRequest              []TendererQualificationRequestType
	AllowedSubcontractTerms                   []SubcontractTermsType
	TenderPreparation                         []TenderPreparationType
	ContractExecutionRequirement              []ContractExecutionRequirementType
	AwardingTerms                             *AwardingTermsType
	AdditionalInformationParty                *PartyType
	DocumentProviderParty                     *PartyType
	TenderRecipientParty                      *PartyType
	ContractResponsibleParty                  *PartyType
	TenderEvaluationParty                     []PartyType
	TenderValidityPeriod                      *PeriodType
	ContractAcceptancePeriod                  *PeriodType
	AppealTerms                               *AppealTermsType
	Language                                  []LanguageType
	BudgetAccountLine                         []BudgetAccountLineType
	ReplacedNoticeDocumentReference           *DocumentReferenceType
}
type ClauseType struct {
	ID      *IDType
	Content []ContentType
}
type FinancialGuaranteeType struct {
	GuaranteeTypeCode  *CodeType1
	Description        []DescriptionType
	LiabilityAmount    *AmountType1
	AmountRate         *AmountRateType
	ConstitutionPeriod *PeriodType
}
type TendererQualificationRequestType struct {
	CompanyLegalFormCode                 *CodeType1
	CompanyLegalForm                     *CompanyLegalFormType
	PersonalSituation                    []PersonalSituationType
	OperatingYearsQuantity               *OperatingYearsQuantityType
	EmployeeQuantity                     *EmployeeQuantityType
	Description                          []DescriptionType
	RequiredBusinessClassificationScheme []ClassificationSchemeType
	TechnicalEvaluationCriterion         []EvaluationCriterionType
	FinancialEvaluationCriterion         []EvaluationCriterionType
	SpecificTendererRequirement          []TendererRequirementType
	EconomicOperatorRole                 []EconomicOperatorRoleType
}
type EvaluationCriterionType struct {
	EvaluationCriterionTypeCode *CodeType1
	Description                 []DescriptionType
	ThresholdAmount             *AmountType1
	ThresholdQuantity           *ThresholdQuantityType
	ExpressionCode              *CodeType1
	Expression                  []ExpressionType
	DurationPeriod              *PeriodType
	SuggestedEvidence           []EvidenceType
}
type EvidenceType struct {
	ID                   *IDType
	EvidenceTypeCode     *CodeType1
	Description          []DescriptionType
	CandidateStatement   []CandidateStatementType
	EvidenceIssuingParty *PartyType
	DocumentReference    *DocumentReferenceType
	Language             *LanguageType
}
type LanguageType struct {
	ID         *IDType
	Name       *NameType1
	LocaleCode *CodeType1
}
type TendererRequirementType struct {
	Name                        []NameType1
	TendererRequirementTypeCode *CodeType1
	Description                 []DescriptionType
	LegalReference              *LegalReferenceType
	SuggestedEvidence           []EvidenceType
}
type TenderPreparationType struct {
	TenderEnvelopeID          *IdentifierType1
	TenderEnvelopeTypeCode    *CodeType1
	Description               []DescriptionType
	OpenTenderID              *IdentifierType1
	ProcurementProjectLot     []ProcurementProjectLotType
	DocumentTenderRequirement []TenderRequirementType
}
type TenderRequirementType struct {
	Name                      *NameType1
	Description               []DescriptionType
	TemplateDocumentReference *DocumentReferenceType
}
type ContractExecutionRequirementType struct {
	Name                     []NameType1
	ExecutionRequirementCode *CodeType1
	Description              []DescriptionType
}
type AwardingTermsType struct {
	WeightingAlgorithmCode        *CodeType1
	Description                   []DescriptionType
	TechnicalCommitteeDescription []TechnicalCommitteeDescriptionType
	LowTendersDescription         []LowTendersDescriptionType
	PrizeIndicator                *PrizeIndicatorType
	PrizeDescription              []PrizeDescriptionType
	PaymentDescription            []PaymentDescriptionType
	FollowupContractIndicator     *FollowupContractIndicatorType
	BindingOnBuyerIndicator       *BindingOnBuyerIndicatorType
	AwardingCriterion             []AwardingCriterionType
	TechnicalCommitteePerson      []PersonType
}
type AwardingCriterionType struct {
	ID                           *IDType
	AwardingCriterionTypeCode    *CodeType1
	Description                  []DescriptionType
	WeightNumeric                *WeightNumericType
	Weight                       []WeightType
	CalculationExpression        []CalculationExpressionType
	CalculationExpressionCode    *CodeType1
	MinimumQuantity              *MinimumQuantityType
	MaximumQuantity              *MaximumQuantityType
	MinimumAmount                *AmountType1
	MaximumAmount                *AmountType1
	MinimumImprovementBid        []MinimumImprovementBidType
	SubordinateAwardingCriterion []AwardingCriterionType
}
type BudgetAccountLineType struct {
	ID            *IDType
	TotalAmount   *AmountType1
	BudgetAccount []BudgetAccountType
}
type BudgetAccountType struct {
	ID                           *IDType
	BudgetYearNumeric            *BudgetYearNumericType
	RequiredClassificationScheme *ClassificationSchemeType
}
type ProcurementProjectType struct {
	ID                                *IDType
	Name                              []NameType1
	Description                       []DescriptionType
	ProcurementTypeCode               *CodeType1
	ProcurementSubTypeCode            *CodeType1
	QualityControlCode                *CodeType1
	RequiredFeeAmount                 *AmountType1
	FeeDescription                    []FeeDescriptionType
	RequestedDeliveryDate             *RequestedDeliveryDateType
	EstimatedOverallContractQuantity  *EstimatedOverallContractQuantityType
	Note                              []NoteType
	RequestedTenderTotal              *RequestedTenderTotalType
	MainCommodityClassification       *CommodityClassificationType
	AdditionalCommodityClassification []CommodityClassificationType
	RealizedLocation                  []LocationType1
	PlannedPeriod                     *PeriodType
	ContractExtension                 *ContractExtensionType
	RequestForTenderLine              []RequestForTenderLineType
}
type RequestedTenderTotalType struct {
	EstimatedOverallContractAmount  *AmountType1
	TotalAmount                     *AmountType1
	TaxIncludedIndicator            *TaxIncludedIndicatorType
	MinimumAmount                   *AmountType1
	MaximumAmount                   *AmountType1
	MonetaryScope                   []MonetaryScopeType
	AverageSubsequentContractAmount *AmountType1
	ApplicableTaxCategory           []TaxCategoryType
}
type ContractExtensionType struct {
	OptionsDescription   []OptionsDescriptionType
	MinimumNumberNumeric *MinimumNumberNumericType
	MaximumNumberNumeric *MaximumNumberNumericType
	OptionValidityPeriod *PeriodType
	Renewal              []RenewalType
}
type RenewalType struct {
	Amount *AmountType2
	Period *PeriodType
}
type RequestForTenderLineType struct {
	ID                           *IDType
	UUID                         *UUIDType
	Note                         []NoteType
	Quantity                     *QuantityType2
	MinimumQuantity              *MinimumQuantityType
	MaximumQuantity              *MaximumQuantityType
	TaxIncludedIndicator         *TaxIncludedIndicatorType
	MinimumAmount                *AmountType1
	MaximumAmount                *AmountType1
	EstimatedAmount              *AmountType1
	DocumentReference            []DocumentReferenceType
	DeliveryPeriod               []PeriodType
	RequiredItemLocationQuantity []ItemLocationQuantityType
	WarrantyValidityPeriod       *PeriodType
	Item                         *ItemType
	SubRequestForTenderLine      []RequestForTenderLineType
}
type TenderLineType struct {
	ID                              *IDType
	Note                            []NoteType
	Quantity                        *QuantityType2
	LineExtensionAmount             *AmountType1
	TotalTaxAmount                  *AmountType1
	OrderableUnit                   *OrderableUnitType
	ContentUnitQuantity             *ContentUnitQuantityType
	OrderQuantityIncrementNumeric   *OrderQuantityIncrementNumericType
	MinimumOrderQuantity            *MinimumOrderQuantityType
	MaximumOrderQuantity            *MaximumOrderQuantityType
	WarrantyInformation             []WarrantyInformationType
	PackLevelCode                   *CodeType1
	DocumentReference               []DocumentReferenceType
	Item                            *ItemType
	OfferedItemLocationQuantity     []ItemLocationQuantityType
	ReplacementRelatedItem          []RelatedItemType
	WarrantyParty                   *PartyType
	WarrantyValidityPeriod          *PeriodType
	SubTenderLine                   []TenderLineType
	CallForTendersLineReference     *LineReferenceType
	CallForTendersDocumentReference *DocumentReferenceType
}
type AwardingCriterionResponseType struct {
	CallForTendersDocumentReference      *IDType
	AwardingCriterionID                  *IdentifierType1
	AwardingCriterionDescription         []AwardingCriterionDescriptionType
	Description                          []DescriptionType
	Quantity                             *QuantityType2
	Amount                               *AmountType2
	SubordinateAwardingCriterionResponse []AwardingCriterionResponseType
}
type DutyType1 struct {
	Amount      *AmountType2
	Duty        *DutyType
	DutyCode    *CodeType1
	TaxCategory *TaxCategoryType
}
type CardAccountType struct {
	PrimaryAccountNumberID *IdentifierType1
	NetworkID              *IdentifierType1
	CardTypeCode           *CodeType1
	ValidityStartDate      *ValidityStartDateType
	ExpiryDate             *ExpiryDateType
	IssuerID               *IdentifierType1
	IssueNumberID          *IdentifierType1
	CV2ID                  *IdentifierType1
	CardChipCode           *CodeType1
	ChipApplicationID      *IdentifierType1
	HolderName             *HolderNameType
}
type CatalogueItemSpecificationUpdateLineType struct {
	ID                      *IDType
	ContractorCustomerParty *CustomerPartyType
	SellerSupplierParty     *SupplierPartyType
	Item                    *ItemType
}
type CatalogueLineType struct {
	ID                            *IDType
	ActionCode                    *CodeType1
	LifeCycleStatusCode           *CodeType1
	ContractSubdivision           *ContractSubdivisionType
	Note                          []NoteType
	OrderableIndicator            *OrderableIndicatorType
	OrderableUnit                 *OrderableUnitType
	ContentUnitQuantity           *ContentUnitQuantityType
	OrderQuantityIncrementNumeric *OrderQuantityIncrementNumericType
	MinimumOrderQuantity          *MinimumOrderQuantityType
	MaximumOrderQuantity          *MaximumOrderQuantityType
	WarrantyInformation           []WarrantyInformationType
	PackLevelCode                 *CodeType1
	ContractorCustomerParty       *CustomerPartyType
	SellerSupplierParty           *SupplierPartyType
	WarrantyParty                 *PartyType
	WarrantyValidityPeriod        *PeriodType
	LineValidityPeriod            *PeriodType
	ItemComparison                []ItemComparisonType
	//Burayı kontrol et
	//ItemComparison                  []RelatedItemType
	AccessoryRelatedItem            []RelatedItemType
	RequiredRelatedItem             []RelatedItemType
	ReplacementRelatedItem          []RelatedItemType
	ComplementaryRelatedItem        []RelatedItemType
	ReplacedRelatedItem             []RelatedItemType
	RequiredItemLocationQuantity    []ItemLocationQuantityType
	DocumentReference               []DocumentReferenceType
	Item                            *ItemType
	KeywordItemProperty             []ItemPropertyType
	CallForTendersLineReference     LineReferenceType
	CallForTendersDocumentReference *DocumentReferenceType
}
type ItemComparisonType struct {
	PriceAmount *PriceAmountType
	Quantity    *QuantityType2
}
type CataloguePricingUpdateLineType struct {
	ID                           *IDType
	ContractorCustomerParty      *CustomerPartyType
	SellerSupplierParty          *SupplierPartyType
	RequiredItemLocationQuantity []ItemLocationQuantityType
}
type CatalogueReferenceType struct {
	ID                *IDType
	UUID              *UUIDType
	IssueDate         *IssueDateType
	IssueTime         *IssueTimeType
	RevisionDate      *RevisionDateType
	RevisionTime      *RevisionTimeType
	Note              []NoteType
	Description       []DescriptionType
	VersionID         *IdentifierType1
	PreviousVersionID *IdentifierType1
}
type CatalogueRequestLineType struct {
	ID                           *IDType
	ContractSubdivision          *ContractSubdivisionType
	Note                         []NoteType
	LineValidityPeriod           *PeriodType
	RequiredItemLocationQuantity []ItemLocationQuantityType
	Item                         *ItemType
}
type CertificateType struct {
	ID                  *IDType
	CertificateTypeCode *CodeType1
	CertificateType1    *CertificateTypeType
	Remarks             []RemarksType
	IssuerParty         *PartyType
	DocumentReference   []DocumentReferenceType
	Signature           []SignatureType
}
type CertificateOfOriginApplicationType struct {
	ReferenceID                 *IdentifierType1
	CertificateType             *CertificateTypeType
	ApplicationStatusCode       *CodeType1
	OriginalJobID               *IdentifierType1
	PreviousJobID               *IdentifierType1
	Remarks                     []RemarksType
	Shipment                    *ShipmentType
	EndorserParty               []EndorserPartyType
	PreparationParty            *PartyType
	IssuerParty                 *PartyType
	ExporterParty               *PartyType
	ImporterParty               *PartyType
	IssuingCountry              *CountryType
	DocumentDistribution        []DocumentDistributionType
	SupportingDocumentReference []DocumentReferenceType
	Signature                   []SignatureType
}
type EndorserPartyType struct {
	RoleCode         *CodeType1
	SequenceNumeric  *SequenceNumericType
	Party            *PartyType
	SignatoryContact *ContactType
}
type DocumentDistributionType struct {
	PrintQualifier       *PrintQualifierType
	MaximumCopiesNumeric *MaximumCopiesNumericType
	Party                *PartyType
}
type ConsignmentType struct {
	ID                          *IDType
	CarrierAssignedID           *IdentifierType1
	ConsigneeAssignedID         *IdentifierType1
	ConsignorAssignedID         *IdentifierType1
	FreightForwarderAssignedID  *IdentifierType1
	BrokerAssignedID            *IdentifierType1
	ContractedCarrierAssignedID *IdentifierType1
	PerformingCarrierAssignedID *IdentifierType1
	SummaryDescription          []SummaryDescriptionType
	TotalInvoiceAmount          *AmountType1
	DeclaredCustomsValueAmount  *AmountType1
	TariffDescription           []TariffDescriptionType
	TariffCode                  *CodeType1
	InsurancePremiumAmount      *AmountType1
	GrossWeightMeasure          *GrossWeightMeasureType
	NetWeightMeasure            *NetWeightMeasureType
	NetNetWeightMeasure         *NetNetWeightMeasureType
	ChargeableWeightMeasure     *ChargeableWeightMeasureType
	GrossVolumeMeasure          *GrossVolumeMeasureType
	NetVolumeMeasure            *NetVolumeMeasureType
	LoadingLengthMeasure        *LoadingLengthMeasureType
	//Burayı kontrol et
	//LoadingLengthMeasure                  []RemarksType
	HazardousRiskIndicator                *HazardousRiskIndicatorType
	AnimalFoodIndicator                   *AnimalFoodIndicatorType
	HumanFoodIndicator                    *HumanFoodIndicatorType
	LivestockIndicator                    *LivestockIndicatorType
	BulkCargoIndicator                    *BulkCargoIndicatorType
	ContainerizedIndicator                *ContainerizedIndicatorType
	GeneralCargoIndicator                 *GeneralCargoIndicatorType
	SpecialSecurityIndicator              *SpecialSecurityIndicatorType
	ThirdPartyPayerIndicator              *ThirdPartyPayerIndicatorType
	CarrierServiceInstructions            []CarrierServiceInstructionsType
	CustomsClearanceServiceInstructions   []CustomsClearanceServiceInstructionsType
	ForwarderServiceInstructions          []ForwarderServiceInstructionsType
	SpecialServiceInstructions            []SpecialServiceInstructionsType
	SequenceID                            *IdentifierType1
	ShippingPriorityLevelCode             *CodeType1
	HandlingCode                          *CodeType1
	HandlingInstructions                  []HandlingInstructionsType
	Information                           []InformationType
	TotalGoodsItemQuantity                *TotalGoodsItemQuantityType
	TotalTransportHandlingUnitQuantity    *TotalTransportHandlingUnitQuantityType
	InsuranceValueAmount                  *AmountType1
	DeclaredForCarriageValueAmount        *AmountType1
	DeclaredStatisticsValueAmount         *AmountType1
	FreeOnBoardValueAmount                *AmountType1
	SpecialInstructions                   []SpecialInstructionsType
	SplitConsignmentIndicator             *SplitConsignmentIndicatorType
	DeliveryInstructions                  []DeliveryInstructionsType
	ConsignmentQuantity                   *ConsignmentQuantityType
	ConsolidatableIndicator               *ConsolidatableIndicatorType
	HaulageInstructions                   []HaulageInstructionsType
	LoadingSequenceID                     *IdentifierType1
	ChildConsignmentQuantity              *ChildConsignmentQuantityType
	TotalPackagesQuantity                 *TotalPackagesQuantityType
	ConsolidatedShipment                  []ShipmentType
	CustomsDeclaration                    []CustomsDeclarationType
	RequestedPickupTransportEvent         *TransportEventType
	RequestedDeliveryTransportEvent       *TransportEventType
	PlannedPickupTransportEvent           *TransportEventType
	PlannedDeliveryTransportEvent         *TransportEventType
	Status                                []StatusType
	ChildConsignment                      []ConsignmentType
	ConsigneeParty                        *PartyType
	ExporterParty                         *PartyType
	ConsignorParty                        *PartyType
	ImporterParty                         *PartyType
	CarrierParty                          *PartyType
	freightForwarderParty                 *PartyType
	NotifyParty                           *PartyType
	OriginalDespatchParty                 *PartyType
	FinalDeliveryParty                    *PartyType
	PerformingCarrierParty                *PartyType
	SubstituteCarrierParty                *PartyType
	LogisticsOperatorParty                *PartyType
	TransportAdvisorParty                 *PartyType
	HazardousItemNotificationParty        *PartyType
	InsuranceParty                        *PartyType
	MortgageHolderParty                   *PartyType
	BillOfLadingHolderParty               *PartyType
	OriginalDepartureCountry              *CountryType
	FinalDestinationCountry               *CountryType
	TransitCountry                        []CountryType
	TransportContract                     *ContractType
	TransportEvent                        []TransportEventType
	OriginalDespatchTransportationService *TransportationServiceType
	FinalDeliveryTransportationService    *TransportationServiceType
	DeliveryTerms                         *DeliveryTermsType
	PaymentTerms                          *PaymentTermsType
	CollectPaymentTerms                   *PaymentTermsType
	DisbursementPaymentTerms              *PaymentTermsType
	PrepaidPaymentTerms                   *PaymentTermsType
	FreightAllowanceCharge                []AllowanceChargeType
	ExtraAllowanceCharge                  []AllowanceChargeType
	MainCarriageShipmentStage             []ShipmentStageType
	PreCarriageShipmentStage              []ShipmentStageType
	OnCarriageShipmentStage               []ShipmentStageType
	TransportHandlingUnit                 []TransportHandlingUnitType
	FirstArrivalPortLocation              *LocationType1
	LastExitPortLocation                  *LocationType1
}
type ContractType struct {
	ID                        *IDType
	IssueDate                 *IssueDateType
	IssueTime                 *IssueTimeType
	NominationDate            *NominationDateType
	NominationTime            *NominationTimeType
	ContractTypeCode          *CodeType1
	ContractType1             *ContractTypeType
	Note                      []NoteType
	VersionID                 *IdentifierType1
	Description               []DescriptionType
	ValidityPeriod            *PeriodType
	ContractDocumentReference []DocumentReferenceType
	NominationPeriod          *PeriodType
	ContractualDelivery       *DeliveryType
}
type PaymentType struct {
	ID            *IDType
	PaidAmount    *AmountType1
	ReceivedDate  *ReceivedDateType
	PaidDate      *PaidDateType
	PaidTime      *PaidTimeType
	InstructionID *IdentifierType1
}
type ConsumptionType struct {
	UtilityStatementTypeCode *CodeType1
	MainPeriod               *PeriodType
	AllowanceCharge          []AllowanceChargeType
	TaxTotal                 []TaxTotalType
	EnergyWaterSupply        *EnergyWaterSupplyType
	TelecommunicationsSupply *TelecommunicationsSupplyType
	LegalMonetaryTotal       *MonetaryTotalType
}
type EnergyWaterSupplyType struct {
	ConsumptionReport                []ConsumptionReportType
	EnergyTaxReport                  []EnergyTaxReportType
	ConsumptionAverage               []ConsumptionAverageType
	EnergyWaterConsumptionCorrection []ConsumptionCorrectionType
}
type ConsumptionReportType struct {
	ID                         *IDType
	ConsumptionType            *ConsumptionTypeType
	ConsumptionTypeCode        *CodeType1
	Description                []DescriptionType
	TotalConsumedQuantity      *TotalConsumedQuantityType
	BasicConsumedQuantity      *BasicConsumedQuantityType
	ResidentOccupantsNumeric   *ResidentOccupantsNumericType
	ConsumersEnergyLevelCode   *CodeType1
	ConsumersEnergyLevel       *ConsumersEnergyLevelType
	ResidenceType              *ResidenceTypeType
	ResidenceTypeCode          *CodeType1
	HeatingType                *HeatingTypeType
	HeatingTypeCode            *CodeType1
	Period                     *PeriodType
	GuidanceDocumentReference  *DocumentReferenceType
	DocumentReference          *DocumentReferenceType
	ConsumptionReportReference []ConsumptionReportReferenceType
	ConsumptionHistory         []ConsumptionHistoryType
}
type ConsumptionReportReferenceType struct {
	ConsumptionReportID   *IdentifierType1
	ConsumptionType       *ConsumptionTypeType
	ConsumptionTypeCode   *CodeType1
	TotalConsumedQuantity *TotalConsumedQuantityType
	Period                *PeriodType
}
type ConsumptionHistoryType struct {
	MeterNumber          *MeterNumberType
	Quantity             *QuantityType2
	Amount               *AmountType2
	ConsumptionLevelCode *CodeType1
	ConsumptionLevel     *ConsumptionLevelType
	Description          []DescriptionType
	Period               *PeriodType
}
type EnergyTaxReportType struct {
	TaxEnergyAmount          *AmountType1
	TaxEnergyOnAccountAmount *AmountType1
	TaxEnergyBalanceAmount   *AmountType1
	TaxScheme                *TaxSchemeType
}
type ConsumptionAverageType struct {
	AverageAmount *AmountType1
	Description   []DescriptionType
}
type ConsumptionCorrectionType struct {
	CorrectionType                         *CorrectionTypeType
	CorrectionTypeCode                     *CodeType1
	MeterNumber                            *MeterNumberType
	GasPressureQuantity                    *GasPressureQuantityType
	ActualTemperatureReductionQuantity     *ActualTemperatureReductionQuantityType
	NormalTemperatureReductionQuantity     *NormalTemperatureReductionQuantityType
	DifferenceTemperatureReductionQuantity *DifferenceTemperatureReductionQuantityType
	Description                            []DescriptionType
	CorrectionUnitAmount                   *AmountType1
	ConsumptionEnergyQuantity              *ConsumptionEnergyQuantityType
	ConsumptionWaterQuantity               *ConsumptionWaterQuantityType
	CorrectionAmount                       *AmountType1
}
type TelecommunicationsSupplyType struct {
	TelecommunicationsSupplyType1    *TelecommunicationsSupplyTypeType
	TelecommunicationsSupplyTypeCode *CodeType1
	PrivacyCode                      *CodeType1
	Description                      []DescriptionType
	TotalAmount                      *AmountType1
	TelecommunicationsSupplyLine     []TelecommunicationsSupplyLineType
}
type TelecommunicationsSupplyLineType struct {
	ID                        *IDType
	PhoneNumber               *PhoneNumberType
	Description               []DescriptionType
	LineExtensionAmount       *AmountType1
	ExchangeRate              []ExchangeRateType
	AllowanceCharge           []AllowanceChargeType
	TaxTotal                  []TaxTotalType
	TelecommunicationsService []TelecommunicationsServiceType
}
type TelecommunicationsServiceType struct {
	ID                                    *IDType
	CallDate                              *CallDateType
	CallTime                              *CallTimeType
	ServiceNumberCalled                   *ServiceNumberCalledType
	TelecommunicationsServiceCategory     *TelecommunicationsServiceCategoryType
	TelecommunicationsServiceCategoryCode *CodeType1
	MovieTitle                            *MovieTitleType
	RoamingPartnerName                    *RoamingPartnerNameType
	PayPerView                            *PayPerViewType
	Quantity                              *QuantityType2
	TelecommunicationsServiceCall         *TelecommunicationsServiceCallType
	TelecommunicationsServiceCallCode     *CodeType1
	CallBaseAmount                        *AmountType1
	CallExtensionAmount                   *AmountType1
	Price                                 *PriceType
	Country                               *CountryType
	ExchangeRate                          []ExchangeRateType
	AllowanceCharge                       []AllowanceChargeType
	TaxTotal                              []TaxTotalType
	CallDuty                              []DutyType1
	TimeDuty                              []DutyType1
}
type ConsumptionLineType struct {
	ID                            *IDType
	ParentDocumentLineReferenceID *IdentifierType1
	InvoicedQuantity              *InvoicedQuantityType
	LineExtensionAmount           *AmountType1
	Period                        *PeriodType
	Delivery                      []DeliveryType
	AllowanceCharge               []AllowanceChargeType
	TaxTotal                      []TaxTotalType
	UtilityItem                   *UtilityItemType
	Price                         *PriceType
	UnstructuredPrice             *UnstructuredPriceType
}
type UtilityItemType struct {
	ID                    *IDType
	SubscriberID          *IdentifierType1
	SubscriberType        *SubscriberTypeType
	SubscriberTypeCode    *CodeType1
	Description           []DescriptionType
	PackQuantity          *PackQuantityType
	PackSizeNumeric       *PackSizeNumericType
	ConsumptionType       *ConsumptionTypeType
	ConsumptionTypeCode   *CodeType1
	CurrentChargeType     *CurrentChargeTypeType
	CurrentChargeTypeCode *CodeType1
	OneTimeChargeType     *OneTimeChargeTypeType
	OneTimeChargeTypeCode *CodeType1
	TaxCategory           *TaxCategoryType
	Contract              *ContractType
}
type UnstructuredPriceType struct {
	PriceAmount *PriceAmountType
	TimeAmount  *TimeAmountType
}
type ConsumptionPointType struct {
	ID             *IDType
	Description    []DescriptionType
	SubscriberID   *IdentifierType1
	SubscriberType *SubscriberTypeType
	//Burayı kontrol et
	//SubscriberType         SubscriberTypeCodeType
	TotalDeliveredQuantity *TotalDeliveredQuantityType
	Address                *AddressType
	WebSiteAccess          *WebSiteAccessType
	UtilityMeter           []MeterType
}
type WebSiteAccessType struct {
	URI      *IdentifierType1
	Password *PasswordType
	Login    *LoginType
}
type MeterType struct {
	MeterNumber            *MeterNumberType
	MeterName              *MeterNameType
	MeterConstant          *MeterConstantType
	MeterConstantCode      *CodeType1
	TotalDeliveredQuantity *TotalDeliveredQuantityType
	MeterReading           []MeterReadingType
	MeterProperty          []MeterPropertyType
}
type MeterReadingType struct {
	ID                             *IDType
	MeterReadingType1              *MeterReadingTypeType
	MeterReadingTypeCode           *CodeType1
	PreviousMeterReadingDate       *PreviousMeterReadingDateType
	PreviousMeterQuantity          *PreviousMeterQuantityType
	LatestMeterReadingDate         *LatestMeterReadingDateType
	LatestMeterQuantity            *LatestMeterQuantityType
	PreviousMeterReadingMethod     *PreviousMeterReadingMethodType
	PreviousMeterReadingMethodCode *CodeType1
	LatestMeterReadingMethod       *LatestMeterReadingMethodType
	LatestMeterReadingMethodCode   *CodeType1
	MeterReadingComments           []MeterReadingCommentsType
	DeliveredQuantity              *DeliveredQuantityType
}
type MeterPropertyType struct {
	Name           *NameType1
	NameCode       *CodeType1
	Value          *ValueType
	ValueQuantity  *ValueQuantityType
	ValueQualifier []ValueQualifierType
}
type ContractingActivityType struct {
	ActivityTypeCode *CodeType1
	ActivityType     *ActivityTypeType
}
type ContractingPartyType struct {
	BuyerProfileURI       *IdentifierType1
	ContractingPartyType1 []ContractingPartyTypeType
	ContractingActivity   []ContractingActivityType
	Party                 *PartyType
}
type ContractingPartyTypeType struct {
	PartyTypeCode *CodeType1
	PartyType     *PartyTypeType
}
type CreditAccountType struct {
	AccountID *IdentifierType1
}
type CreditNoteLineType struct {
	ID                    *IDType
	UUID                  *UUIDType
	Note                  []NoteType
	CreditedQuantity      *CreditedQuantityType
	LineExtensionAmount   *AmountType1
	TaxPointDate          *TaxPointDateType
	AccountingCostCode    *CodeType1
	AccountingCost        *AccountingCostType
	PaymentPurposeCode    *CodeType1
	FreeOfChargeIndicator *FreeOfChargeIndicatorType
	InvoicePeriod         []PeriodType
	OrderLineReference    []OrderLineReferenceType
	DiscrepancyResponse   []ResponseType
	DespatchLineReference []LineReferenceType
	ReceiptLineReference  []LineReferenceType
	BillingReference      []BillingReferenceType
	DocumentReference     []DocumentReferenceType
	PricingReference      *PricingReferenceType
	OriginatorParty       *PartyType
	Delivery              []DeliveryType
	PaymentTerms          []PaymentTermsType
	TaxTotal              []TaxTotalType
	AllowanceCharge       []AllowanceChargeType
	Item                  *ItemType
	Price                 *PriceType
	DeliveryTerms         []DeliveryTermsType
	SubCreditNoteLine     []CreditNoteLineType
	ItemPriceExtension    *PriceExtensionType
}
type DebitNoteLineType struct {
	ID                    *IDType
	UUID                  *UUIDType
	Note                  []NoteType
	DebitedQuantity       *DebitedQuantityType
	LineExtensionAmount   *AmountType1
	TaxPointDate          *TaxPointDateType
	AccountingCostCode    *CodeType1
	AccountingCost        *AccountingCostType
	PaymentPurposeCode    *CodeType1
	DiscrepancyResponse   []ResponseType
	DespatchLineReference []LineReferenceType
	ReceiptLineReference  []LineReferenceType
	BillingReference      []BillingReferenceType
	DocumentReference     []DocumentReferenceType
	PricingReference      *PricingReferenceType
	Delivery              []DeliveryType
	TaxTotal              []TaxTotalType
	AllowanceCharge       []AllowanceChargeType
	Item                  *ItemType
	Price                 *PriceType
	SubDebitNoteLine      []DebitNoteLineType
}
type DespatchLineType struct {
	ID                  *IDType
	UUID                *UUIDType
	Note                []NoteType
	LineStatusCode      *LineStatusCodeType
	DeliveredQuantity   *DeliveredQuantityType
	BackorderQuantity   *BackorderQuantityType
	BackorderReason     []BackorderReasonType
	OutstandingQuantity *OutstandingQuantityType
	OutstandingReason   []OutstandingReasonType
	OversupplyQuantity  *OversupplyQuantityType
	OrderLineReference  []OrderLineReferenceType
	DocumentReference   []DocumentReferenceType
	Item                *ItemType
	Shipment            []ShipmentType
}
type EconomicOperatorShortListType struct {
	LimitationDescription []LimitationDescriptionType
	ExpectedQuantity      *ExpectedQuantityType
	MaximumQuantity       *MaximumQuantityType
	MinimumQuantity       *MinimumQuantityType
	PreSelectedParty      []PartyType
}
type EndorsementType struct {
	DocumentID     *IdentifierType1
	ApprovalStatus *ApprovalStatusType
	Remarks        []RemarksType
	EndorserParty  *EndorserPartyType
	Signature      []SignatureType
}
type EventType struct {
	IdentificationID    *IdentifierType1
	OccurrenceDate      *OccurrenceDateType
	OccurrenceTime      *OccurrenceTimeType
	TypeCode            *CodeType1
	Description         []DescriptionType
	CompletionIndicator *CompletionIndicatorType
	CurrentStatus       []StatusType
	Contact             []ContactType
	OccurenceLocation   *LocationType1
}
type EventCommentType struct {
	Comment   *CommentType
	IssueDate *IssueDateType
	IssueTime *IssueTimeType
}
type EventLineItemType struct {
	LineNumberNumeric              *LineNumberNumericType
	ParticipatingLocationsLocation *LocationType1
	RetailPlannedImpact            []RetailPlannedImpactType
	SupplyItem                     *ItemType
}
type RetailPlannedImpactType struct {
	Amount              *AmountType2
	ForecastPurposeCode *CodeType1
	ForecastTypeCode    *CodeType1
	Period              *PeriodType
}
type EventTacticType struct {
	Comment                *CommentType
	Quantity               *QuantityType2
	EventTacticEnumeration *EventTacticEnumerationType
	Period                 *PeriodType
}
type EventTacticEnumerationType struct {
	ConsumerIncentiveTacticTypeCode  *CodeType1
	DisplayTacticTypeCode            *CodeType1
	FeatureTacticTypeCode            *CodeType1
	TradeItemPackingLabelingTypeCode *CodeType1
}
type ExceptionCriteriaLineType struct {
	ID                             *IDType
	Note                           []NoteType
	ThresholdValueComparisonCode   *CodeType1
	ThresholdQuantity              *ThresholdQuantityType
	ExceptionStatusCode            *CodeType1
	CollaborationPriorityCode      *CodeType1
	ExceptionResolutionCode        *CodeType1
	SupplyChainActivityTypeCode    *CodeType1
	PerformanceMetricTypeCode      *CodeType1
	EffectivePeriod                *PeriodType
	SupplyItem                     []ItemType
	ForecastExceptionCriterionLine *ForecastExceptionCriterionLineType
}
type ForecastExceptionCriterionLineType struct {
	ForecastPurposeCode      *CodeType1
	ForecastTypeCode         *CodeType1
	ComparisonDataSourceCode *CodeType1
	DataSourceCode           *CodeType1
	TimeDeltaDaysQuantity    *TimeDeltaDaysQuantityType
}
type ExceptionNotificationLineType struct {
	ID                          *IDType
	Note                        []NoteType
	Description                 []DescriptionType
	ExceptionStatusCode         *CodeType1
	CollaborationPriorityCode   *CodeType1
	ResolutionCode              *CodeType1
	ComparedValueMeasure        *ComparedValueMeasureType
	SourceValueMeasure          *SourceValueMeasureType
	VarianceQuantity            *VarianceQuantityType
	SupplyChainActivityTypeCode *CodeType1
	PerformanceMetricTypeCode   *CodeType1
	ExceptionObservationPeriod  *PeriodType
	DocumentReference           []DocumentReferenceType
	ForecastException           *ForecastExceptionType
	SupplyItem                  *ItemType
}
type ForecastExceptionType struct {
	ForecastPurposeCode         *CodeType1
	ForecastTypeCode            *CodeType1
	IssueDate                   *IssueDateType
	IssueTime                   *IssueTimeType
	DataSourceCode              *CodeType1
	ComparisonDataCode          *CodeType1
	ComparisonForecastIssueTime *ComparisonForecastIssueTimeType
	ComparisonForecastIssueDate *ComparisonForecastIssueDateType
}
type ForecastLineType struct {
	ID                      *IDType
	Note                    []NoteType
	FrozenDocumentIndicator *FrozenDocumentIndicatorType
	//Burayı kontrol et
	//ForecastTypeCodeType
	ForecastPeriod *PeriodType
	SalesItem      *SalesItemType
}
type ForecastRevisionLineType struct {
	ID                      *IDType
	Note                    []NoteType
	Description             []DescriptionType
	RevisedForecastLineID   *IdentifierType1
	SourceForecastIssueDate *SourceForecastIssueDateType
	SourceForecastIssueTime *SourceForecastIssueTimeType
	AdjustmentReasonCode    *CodeType1
	ForecastPeriod          *PeriodType
	SalesItem               *SalesItemType
}
type FrameworkAgreementType struct {
	ExpectedOperatorQuantity           *ExpectedOperatorQuantityType
	MaximumOperatorQuantity            *MaximumOperatorQuantityType
	Justification                      []JustificationType
	Frequency                          []FrequencyType
	DurationPeriod                     *PeriodType
	SubsequentProcessTenderRequirement []TenderRequirementType
}
type GoodsItemContainerType struct {
	ID                 *IDType
	Quantity           *QuantityType2
	TransportEquipment []TransportEquipmentType
}
type TradingTermsType struct {
	Information       []InformationType
	Reference         *ReferenceType
	ApplicableAddress *AddressType
}
type HazardousItemType struct {
	ID                         *IDType
	PlacardNotation            *PlacardNotationType
	PlacardEndorsement         *PlacardEndorsementType
	AdditionalInformation      []AdditionalInformationType
	UNDGCode                   *CodeType1
	EmergencyProceduresCode    *CodeType1
	MedicalFirstAidGuideCode   *CodeType1
	TechnicalName              *TechnicalNameType
	CategoryName               *CategoryNameType
	HazardousCategoryCode      *CodeType1
	UpperOrangeHazardPlacardID *IdentifierType1
	LowerOrangeHazardPlacardID *IdentifierType1
	MarkingID                  *IdentifierType1
	HazardClassID              *IdentifierType1
	NetWeightMeasure           *NetWeightMeasureType
	NetVolumeMeasure           *NetVolumeMeasureType
	Quantity                   *QuantityType2
	ContactParty               *PartyType
	SecondaryHazard            []SecondaryHazardType
	HazardousGoodsTransit      []HazardousGoodsTransitType
	EmergencyTemperature       *TemperatureType
	FlashpointTemperature      *TemperatureType
	AdditionalTemperature      []TemperatureType
}
type SecondaryHazardType struct {
	ID                      *IDType
	PlacardNotation         *PlacardNotationType
	PlacardEndorsement      *PlacardEndorsementType
	EmergencyProceduresCode *CodeType1
	Extension               []ExtensionType
}
type ImmobilizedSecurityType struct {
	ImmobilizationCertificateID *IdentifierType1
	SecurityID                  *IdentifierType1
	IssueDate                   *IssueDateType
	FaceValueAmount             *AmountType1
	MarketValueAmount           *AmountType1
	SharesNumberQuantity        *SharesNumberQuantityType
	IssuerParty                 *PartyType
}
type InstructionForReturnsLineType struct {
	ID                *IDType
	Note              []NoteType
	Quantity          *QuantityType2
	ManufacturerParty *PartyType
	Item              *ItemType
}
type InventoryReportLineType struct {
	ID                     *IDType
	Note                   []NoteType
	Quantity               *QuantityType2
	InventoryValueAmount   *AmountType1
	AvailabilityDate       *AvailabilityDateType
	AvailabilityStatusCode *CodeType1
	Item                   *ItemType
	InventoryLocation      *LocationType1
}
type ItemInformationRequestLineType struct {
	TimeFrequencyCode           *CodeType1
	SupplyChainActivityTypeCode *CodeType1
	ForecastTypeCode            *CodeType1
	PerformanceMetricTypeCode   *CodeType1
	Period                      []PeriodType
	SalesItem                   []SalesItemType
}
type ItemManagementProfileType struct {
	FrozenPeriodDaysNumeric       *FrozenPeriodDaysNumericType
	MinimumInventoryQuantity      *MinimumInventoryQuantityType
	MultipleOrderQuantity         *MultipleOrderQuantityType
	OrderIntervalDaysNumeric      *OrderIntervalDaysNumericType
	ReplenishmentOwnerDescription []ReplenishmentOwnerDescriptionType
	TargetServicePercent          *TargetServicePercentType
	TargetInventoryQuantity       *TargetInventoryQuantityType
	EffectivePeriod               *PeriodType
	Item                          *ItemType
	ItemLocationQuantity          *ItemLocationQuantityType
}
type LocationCoordinateType struct {
	CoordinateSystemCode    *CodeType1
	LatitudeDegreesMeasure  *LatitudeDegreesMeasureType
	LatitudeMinutesMeasure  *LatitudeMinutesMeasureType
	LatitudeDirectionCode   *CodeType1
	LongitudeDegreesMeasure *LongitudeDegreesMeasureType
	LongitudeMinutesMeasure *LongitudeMinutesMeasureType
	LongitudeDirectionCode  *CodeType1
	AltitudeMeasure         *AltitudeMeasureType
}
type OnAccountPaymentType struct {
	EstimatedConsumedQuantity *EstimatedConsumedQuantityType
	Note                      []NoteType
	PaymentTerms              []PaymentTermsType
}
type MiscellaneousEventType struct {
	MiscellaneousEventTypeCode *CodeType1
	EventLineItem              []EventLineItemType
}
type NotificationRequirementType struct {
	NotificationTypeCode                 *CodeType1
	PostEventNotificationDurationMeasure *PostEventNotificationDurationMeasureType
	PreEventNotificationDurationMeasure  *PreEventNotificationDurationMeasureType
	NotifyParty                          []PartyType
	NotificationPeriod                   []PeriodType
	NotificationLocation                 []LocationType1
}
type OrderLineType struct {
	SubstitutionStatusCode           *CodeType1
	Note                             []NoteType
	LineItem                         *LineItemType
	SellerProposedSubstituteLineItem []LineItemType
	SellerSubstitutedLineItem        []LineItemType
	BuyerProposedSubstituteLineItem  []LineItemType
	CatalogueLineReference           *LineReferenceType
	QuotationLineReference           *LineReferenceType
	OrderLineReference               []OrderLineReferenceType
	DocumentReference                []DocumentReferenceType
}
type PaymentMandateType struct {
	ID                                *IDType
	MandateTypeCode                   *CodeType1
	MaximumPaymentInstructionsNumeric *MaximumPaymentInstructionsNumericType
	MaximumPaidAmount                 *AmountType1
	SignatureID                       *IdentifierType1
	PayerParty                        *PartyType
	PayerFinancialAccount             *FinancialAccountType
	ValidityPeriod                    *PeriodType
	PaymentReversalPeriod             *PeriodType
	Clause                            []ClauseType
}
type PerformanceDataLineType struct {
	ID                        *IDType
	Note                      []NoteType
	PerformanceValueQuantity  *PerformanceValueQuantityType
	PerformanceMetricTypeCode *CodeType1
	Period                    *PeriodType
	Item                      *ItemType
}
type PhysicalAttributeType struct {
	AttributeID     *IdentifierType1
	PositionCode    *CodeType1
	DescriptionCode *CodeType1
	Description     []DescriptionType
}
type PickupType struct {
	ID                 *IDType
	ActualPickupDate   *ActualPickupDateType
	ActualPickupTime   *ActualPickupTimeType
	EarliestPickupDate *EarliestPickupDateType
	EarliestPickupTime *EarliestPickupTimeType
	LatestPickupDate   *LatestPickupDateType
	LatestPickupTime   *LatestPickupTimeType
	PickupLocation     *LocationType1
	PickupParty        *PartyType
}
type PowerOfAttorneyType struct {
	ID                       *IDType
	IssueDate                *IssueDateType
	IssueTime                *IssueTimeType
	Description              []DescriptionType
	NotaryParty              *PartyType
	AgentParty               *PartyType
	WitnessParty             []PartyType
	MandateDocumentReference []DocumentReferenceType
}
type PriceListType struct {
	ID                *IDType
	StatusCode        *CodeType1
	ValidityPeriod    []PeriodType
	PreviousPriceList *PriceListType
}
type ProcessJustificationType struct {
	PreviousCancellationReasonCode *CodeType1
	ProcessReasonCode              *CodeType1
	ProcessReason                  []ProcessReasonType
	Description                    []DescriptionType
}
type ProjectReferenceType struct {
	ID                 *IDType
	UUID               *UUIDType
	IssueDate          *IssueDateType
	WorkPhaseReference []WorkPhaseReferenceType
}
type WorkPhaseReferenceType struct {
	ID                         *IDType
	WorkPhaseCode              *CodeType1
	WorkPhase                  []WorkPhaseType
	ProgressPercent            *ProgressPercentType
	StartDate                  *StartDateType
	EndDate                    *EndDateType
	WorkOrderDocumentReference []DocumentReferenceType
}
type PromotionalEventType struct {
	PromotionalEventTypeCode      *CodeType1
	SubmissionDate                *SubmissionDateType
	FirstShipmentAvailibilityDate *FirstShipmentAvailibilityDateType
	LatestProposalAcceptanceDate  *LatestProposalAcceptanceDateType
	PromotionalSpecification      []PromotionalSpecificationType
}
type PromotionalSpecificationType struct {
	SpecificationID          *IdentifierType1
	PromotionalEventLineItem []PromotionalEventLineItemType
	EventTactic              []EventTacticType
}
type PromotionalEventLineItemType struct {
	Amount        *AmountType2
	EventLineItem *EventLineItemType
}
type QualificationResolutionType struct {
	AdmissionCode         *CodeType1
	ExclusionReason       []ExclusionReasonType
	Resolution            []ResolutionType
	ResolutionDate        *ResolutionDateType
	ResolutionTime        *ResolutionTimeType
	ProcurementProjectLot *ProcurementProjectLotType
}
type QuotationLineType struct {
	ID                               *IDType
	Note                             []NoteType
	Quantity                         *QuantityType2
	LineExtensionAmount              *AmountType1
	TotalTaxAmount                   *AmountType1
	RequestForQuotationLineID        *IdentifierType1
	DocumentReference                []DocumentReferenceType
	LineItem                         *LineItemType
	SellerProposedSubstituteLineItem []LineItemType
	AlternativeLineItem              []LineItemType
	RequestLineReference             *LineReferenceType
}
type ReceiptLineType struct {
	ID                      *IDType
	UUID                    *UUIDType
	Note                    []NoteType
	ReceivedQuantity        *ReceivedQuantityType
	ShortQuantity           *ShortQuantityType
	ShortageActionCode      *CodeType1
	RejectedQuantity        *RejectedQuantityType
	RejectReasonCode        *CodeType1
	RejectReason            []RejectReasonType
	RejectActionCode        *CodeType1
	QuantityDiscrepancyCode *CodeType1
	OversupplyQuantity      *OversupplyQuantityType
	ReceivedDate            *ReceivedDateType
	TimingComplaintCode     *CodeType1
	TimingComplaint         *TimingComplaintType
	OrderLineReference      *OrderLineReferenceType
	DespatchLineReference   []LineReferenceType
	//Burayı kontrol et
	//DespatchLineReference   []DocumentReferenceType
	Item     []ItemType
	Shipment []ShipmentType
}
type ReminderLineType struct {
	Shipment                       *IDType
	Note                           []NoteType
	UUID                           *UUIDType
	BalanceBroughtForwardIndicator *BalanceBroughtForwardIndicatorType
	DebitLineAmount                *AmountType1
	CreditLineAmount               *AmountType1
	AccountingCostCode             *CodeType1
	AccountingCost                 *AccountingCostType
	PenaltySurchargePercent        *PenaltySurchargePercentType
	Amount                         *AmountType2
	PaymentPurposeCode             *CodeType1
	ReminderPeriod                 []PeriodType
	BillingReference               []BillingReferenceType
	ExchangeRate                   *ExchangeRateType
}
type RemittanceAdviceLineType struct {
	ID                      *IDType
	Note                    []NoteType
	UUID                    *UUIDType
	DebitLineAmount         *AmountType1
	CreditLineAmount        *AmountType1
	BalanceAmount           *AmountType1
	PaymentPurposeCode      *CodeType1
	InvoicingPartyReference *InvoicingPartyReferenceType
	AccountingSupplierParty *SupplierPartyType
	AccountingCustomerParty *CustomerPartyType
	BuyerCustomerParty      *CustomerPartyType
	SellerSupplierParty     *SupplierPartyType
	OriginatorCustomerParty *CustomerPartyType
	PayeeParty              *PartyType
	InvoicePeriod           []PeriodType
	BillingReference        []BillingReferenceType
	DocumentReference       []DocumentReferenceType
	ExchangeRate            *ExchangeRateType
}
type RequestForQuotationLineType struct {
	ID                        *IDType
	UUID                      *UUIDType
	Note                      []NoteType
	OptionalLineItemIndicator *OptionalLineItemIndicatorType
	PrivacyCode               *CodeType1
	//Burayı kontrol et
	//PrivacyCode               SecurityClassificationCodeType
	DocumentReference []DocumentReferenceType
	LineItem          *LineItemType
}
type ResultOfVerificationType struct {
	ValidatorID          *IdentifierType1
	ValidationResultCode *CodeType1
	ValidationDate       *ValidationDateType
	ValidationTime       *ValidationTimeType
	ValidateProcess      *ValidateProcessType
	ValidateTool         *ValidateToolType
	ValidateToolVersion  *ValidateToolVersionType
	SignatoryParty       *PartyType
}
type ServiceProviderPartyType struct {
	ID              *IDType
	ServiceTypeCode *CodeType1
	ServiceType     []ServiceTypeType
	Party           *PartyType
	SellerContact   *ContactType
}
type ShareholderPartyType struct {
	PartecipationPercent *PartecipationPercentType
	Party                *PartyType
}
type StatementLineType struct {
	ID                             *IDType
	Note                           []NoteType
	UUID                           *UUIDType
	BalanceBroughtForwardIndicator *BalanceBroughtForwardIndicatorType
	DebitLineAmount                *AmountType1
	CreditLineAmount               *AmountType1
	BalanceAmount                  *AmountType1
	PaymentPurposeCode             *CodeType1
	PaymentMeans                   *PaymentMeansType
	PaymentTerms                   []PaymentTermsType
	BuyerCustomerParty             *CustomerPartyType
	SellerSupplierParty            *SupplierPartyType
	OriginatorCustomerParty        *CustomerPartyType
	AccountingCustomerParty        *CustomerPartyType
	AccountingSupplierParty        *SupplierPartyType
	PayeeParty                     *PartyType
	InvoicePeriod                  []PeriodType
	BillingReference               []BillingReferenceType
	DocumentReference              []DocumentReferenceType
	ExchangeRate                   *ExchangeRateType
	AllowanceCharge                []AllowanceChargeType
	CollectedPayment               []PaymentType
}
type StockAvailabilityReportLineType struct {
	ID                     *IDType
	Note                   []NoteType
	Quantity               *QuantityType2
	ValueAmount            *AmountType1
	AvailabilityDate       *AvailabilityDateType
	AvailabilityStatusCode *CodeType1
	Item                   *ItemType
}
type SubscriberConsumptionType struct {
	ConsumptionID           *IdentifierType1
	SpecificationTypeCode   *CodeType1
	Note                    []NoteType
	TotalMeteredQuantity    *TotalMeteredQuantityType
	SubscriberParty         *PartyType
	UtilityConsumptionPoint *ConsumptionPointType
	OnAccountPayment        []OnAccountPaymentType
	Consumption             *ConsumptionType
	SupplierConsumption     []SupplierConsumptionType
}
type SupplierConsumptionType struct {
	Description          []DescriptionType
	UtilitySupplierParty *PartyType
	UtilityCustomerParty *PartyType
	Consumption          *ConsumptionType
	Contract             *ContractType
	ConsumptionLine      []ConsumptionLineType
}
type TenderResultType struct {
	TenderResultCode                 *CodeType1
	Description                      []DescriptionType
	AdvertisementAmount              *AmountType1
	AwardDate                        *AwardDateType
	AwardTime                        *AwardTimeType
	ReceivedTenderQuantity           *ReceivedTenderQuantityType
	LowerTenderAmount                *AmountType1
	HigherTenderAmount               *AmountType1
	StartDate                        *StartDateType
	ReceivedElectronicTenderQuantity *ReceivedElectronicTenderQuantityType
	ReceivedForeignTenderQuantity    *ReceivedForeignTenderQuantityType
	Contract                         *ContractType
	AwardedTenderedProject           *TenderedProjectType
	ContractFormalizationPeriod      *PeriodType
	SubcontractTerms                 []SubcontractTermsType
	WinningParty                     []WinningPartyType
}
type WinningPartyType struct {
	Rank  *RankType
	Party *PartyType
}
type TendererPartyQualificationType struct {
	InterestedProcurementProjectLot []ProcurementProjectLotType
	MainQualifyingParty             *QualifyingPartyType
	AdditionalQualifyingParty       []QualifyingPartyType
}
type TenderingProcessType struct {
	ID                                     *IDType
	OriginalContractingSystemID            *IdentifierType1
	Description                            []DescriptionType
	NegotiationDescription                 []NegotiationDescriptionType
	ProcedureCode                          *CodeType1
	UrgencyCode                            *CodeType1
	ExpenseCode                            *CodeType1
	PartPresentationCode                   *CodeType1
	ContractingSystemCode                  *CodeType1
	SubmissionMethodCode                   *CodeType1
	CandidateReductionConstraintIndicator  *CandidateReductionConstraintIndicatorType
	GovernmentAgreementConstraintIndicator *GovernmentAgreementConstraintIndicatorType
	DocumentAvailabilityPeriod             *PeriodType
	TenderSubmissionDeadlinePeriod         *PeriodType
	InvitationSubmissionPeriod             *PeriodType
	ParticipationRequestReceptionPeriod    *PeriodType
	NoticeDocumentReference                []DocumentReferenceType
	AdditionalDocumentReference            []DocumentReferenceType
	ProcessJustification                   []ProcessJustificationType
	EconomicOperatorShortList              *EconomicOperatorShortListType
	//Burayı kontrol et
	//EconomicOperatorShortList              []EventType
	AuctionTerms       *AuctionTermsType
	FrameworkAgreement *FrameworkAgreementType
}
type TradeFinancingType struct {
	FrameworkAgreement        *IDType
	FinancingInstrumentCode   *CodeType1
	ContractDocumentReference *DocumentReferenceType
	DocumentReference         []DocumentReferenceType
	FinancingParty            *PartyType
	FinancingFinancialAccount *FinancialAccountType
	Clause                    []ClauseType
}
type TransactionConditionsType struct {
	ID                *IDType
	ActionCode        *CodeType1
	Description       []DescriptionType
	DocumentReference []DocumentReferenceType
}
type TransportEquipmentSealType struct {
	ID                 *IDType
	SealIssuerTypeCode *CodeType1
	Condition          *ConditionType
	SealStatusCode     *CodeType1
	SealingPartyType   *SealingPartyTypeType
}
type TransportExecutionTermsType struct {
	TransportUserSpecialTerms            []TransportUserSpecialTermsType
	TransportServiceProviderSpecialTerms []TransportServiceProviderSpecialTermsType
	ChangeConditions                     []ChangeConditionsType
	PaymentTerms                         []PaymentTermsType
	DeliveryTerms                        []DeliveryTermsType
	BonusPaymentTerms                    *PaymentTermsType
	CommissionPaymentTerms               *PaymentTermsType
	PenaltyPaymentTerms                  *PaymentTermsType
	EnvironmentalEmission                []EnvironmentalEmissionType
	NotificationRequirement              []NotificationRequirementType
	ServiceChargePaymentTerms            *PaymentTermsType
}
type TransportScheduleType struct {
	SequenceNumeric                  *SequenceNumericType
	ReferenceDate                    *ReferenceDateType
	ReferenceTime                    *ReferenceTimeType
	ReliabilityPercent               *ReliabilityPercentType
	Remarks                          []RemarksType
	StatusLocation                   *LocationType1
	ActualArrivalTransportEvent      *TransportEventType
	ActualDepartureTransportEvent    *TransportEventType
	EstimatedDepartureTransportEvent *TransportEventType
	EstimatedArrivalTransportEvent   *TransportEventType
	PlannedDepartureTransportEvent   *TransportEventType
	PlannedArrivalTransportEvent     *TransportEventType
}
type TransportationSegmentType struct {
	SequenceNumeric                   *SequenceNumericType
	TransportExecutionPlanReferenceID *IdentifierType1
	TransportationService             *TransportationServiceType
	TransportServiceProviderParty     *PartyType
	ReferencedConsignment             *ConsignmentType
	ShipmentStage                     []ShipmentStageType
}
type UBLExtensionsType struct {
	UBLExtension []UBLExtensionType
}
type UBLDocumentSignaturesType struct {
	SignatureInformation []SignatureInformationType
}
type SignatureInformationType struct {
	ID                    *IDType
	ReferencedSignatureID *IdentifierType1
	Signature             *SignatureType1
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
	//ItemsElementName []ItemsChoiceType2
	Text []string
	Id   string
}
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
	//ItemsElementName []ItemsChoiceType1
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
	Any      string //System.Xml.XmlElement
}
type X509DataType struct {
	Items []string //object
	//Burayı kontrol et
	//ItemsElementName []ItemsChoiceType
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
type AnyType struct {
	Any     []string //System.Xml.XmlNode
	AnyAttr []string //System.Xml.XmlAttribute
}
type ObjectIdentifierType struct {
	Identifier              *IdentifierType2
	Description             string
	DocumentationReferences *DocumentationReferencesType
}
type IdentifierType2 struct {
	Qualifier          string //QualifierType
	QualifierSpecified bool
	Value              string
}

//QualifierType enum
//OIDAsURI,
//OIDAsURN,
//}
type DocumentationReferencesType struct {
	DocumentationReference []string
}
type EncapsulatedPKIDataType struct {
	Id       string
	Encoding string
	Value    []byte
}
type IncludeType struct {
	URI                     string
	referencedData          bool
	referencedDataSpecified bool
}
type ReferenceInfoType struct {
	DigestMethod *DigestMethodType
	DigestValue  []byte
	Id           string
	URI          string
}
type XAdESTimeStampType struct {
	GenericTimeStampType
}
type GenericTimeStampType struct {
	Items                  []string //object
	CanonicalizationMethod *CanonicalizationMethodType1
	Items1                 []string //object
	Id                     string
}

type OtherTimeStampType struct {
	GenericTimeStampType
}
type QualifyingPropertiesType struct {
	SignedProperties   *SignedPropertiesType
	UnsignedProperties *UnsignedPropertiesType
	Id                 string
}

type SignedPropertiesType struct {
	SignedSignatureProperties  *SignedSignaturePropertiesType
	SignedDataObjectProperties *SignedDataObjectPropertiesType
	Id                         string
}
type SignedSignaturePropertiesType struct {
	SigningTime               time.Time //System.DateTime
	SigningTimeSpecified      bool
	SigningCertificate        []CertIDType
	SignaturePolicyIdentifier *SignaturePolicyIdentifierType
	SignatureProductionPlace  *SignatureProductionPlaceType
	SignerRole                *SignerRoleType
	Id                        string
}
type CertIDType struct {
	CertDigest   *DigestAlgAndValueType
	IssuerSerial *X509IssuerSerialType
	URI          string
}
type DigestAlgAndValueType struct {
	DigestMethod *DigestMethodType
	DigestValue  []byte
}
type SignaturePolicyIdentifierType struct {
	Item string //object
}
type SignaturePolicyIdType struct {
	SigPolicyId         *ObjectIdentifierType
	Transforms          []TransformType
	SigPolicyHash       *DigestAlgAndValueType
	SigPolicyQualifiers []AnyType
}
type SignatureProductionPlaceType struct {
	City            string
	StateOrProvince string
	PostalCode      string
	CountryName     string
}
type SignerRoleType struct {
	ClaimedRoles   []AnyType
	CertifiedRoles []EncapsulatedPKIDataType
}
type SignedDataObjectPropertiesType struct {
	DataObjectFormat               []DataObjectFormatType
	CommitmentTypeIndication       []CommitmentTypeIndicationType
	AllDataObjectsTimeStamp        []XAdESTimeStampType
	IndividualDataObjectsTimeStamp []XAdESTimeStampType
	Id                             string
}
type DataObjectFormatType struct {
	Description      string
	ObjectIdentifier *ObjectIdentifierType
	MimeType         string
	Encoding         string
	ObjectReference  string
}
type CommitmentTypeIndicationType struct {
	CommitmentTypeId         *ObjectIdentifierType
	Items                    []string //object
	CommitmentTypeQualifiers []AnyType
}
type UnsignedPropertiesType struct {
	UnsignedSignatureProperties  *UnsignedSignaturePropertiesType
	UnsignedDataObjectProperties *UnsignedDataObjectPropertiesType
	Id                           string
}
type UnsignedSignaturePropertiesType struct {
	Items            []string //object
	ItemsElementName []string //ItemsChoiceType3
	Id               string
}
type CertificateValuesType struct {
	Items []string //object
	//Burayı kontrol et
	//Items string
}
type CompleteCertificateRefsType struct {
	CertRefs []CertIDType
	Id       string
}
type CompleteRevocationRefsType struct {
	CRLRefs   []CRLRefType
	OCSPRefs  []OCSPRefType
	OtherRefs []AnyType
	Id        string
}
type CRLRefType struct {
	DigestAlgAndValue *DigestAlgAndValueType
	CRLIdentifier     *CRLIdentifierType
}
type CRLIdentifierType struct {
	Issuer    string
	IssueTime time.Time //System.DateTime
	Number    string
	URI       string
}
type OCSPRefType struct {
	OCSPIdentifier    *OCSPIdentifierType
	DigestAlgAndValue *DigestAlgAndValueType
}
type OCSPIdentifierType struct {
	ResponderID *ResponderIDType
	ProducedAt  time.Time //System.DateTime
	URI         string
}
type ResponderIDType struct {
	Item string //object
}
type RevocationValuesType struct {
	CRLValues   []EncapsulatedPKIDataType
	OCSPValues  []EncapsulatedPKIDataType
	OtherValues []AnyType
	Id          string
}
type CounterSignatureType struct {
	Signature *SignatureType1
}

//ItemsChoiceType3 enum
//Item,
//ArchiveTimeStamp,
//AttrAuthoritiesCertValues,
//AttributeCertificateRefs,
//AttributeRevocationRefs,
//AttributeRevocationValues,
//CertificateValues,
//CompleteCertificateRefs,
//CompleteRevocationRefs,
//CounterSignature,
//RefsOnlyTimeStamp,
//RevocationValues,
//SigAndRefsTimeStamp,
//SignatureTimeStamp,
//}
type UnsignedDataObjectPropertiesType struct {
	UnsignedDataObjectProperty []AnyType
	Id                         string
}
type QualifyingPropertiesReferenceType struct {
	URI string
	Id  string
}
type CertIDListType struct {
	Cert []CertIDType
}
type SPUserNoticeType struct {
	NoticeRef    *NoticeReferenceType
	ExplicitText string
}
type NoticeReferenceType struct {
	Organization  string
	NoticeNumbers []string
}
type ValidationDataType struct {
	CertificateValues *CertificateValuesType
	RevocationValues  *RevocationValuesType
	Id                string
	UR                string
}
type TransformsType struct {
	Transform []TransformType
}
type ManifestType struct {
	Reference []ReferenceType1
	Id        string
}
type SignaturePropertiesType struct {
	SignatureProperty []SignaturePropertyType
	Id                string
}
type SignaturePropertyType struct {
	Items  []string //System.Xml.XmlElement
	Text   []string
	Target string
	Id     string
}
