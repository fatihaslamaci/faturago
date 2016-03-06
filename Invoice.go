package main

import (
	"encoding/xml"
)

type Invoice struct {
	XMLName         xml.Name `xml:"Invoice"`
	UBLVersionID    string
	CustomizationID string
}
