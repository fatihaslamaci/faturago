package main

import (
	"testing"

	"encoding/xml"

	//"os"
)

func stringKarsilastir(t *testing.T, alanAdi, gerceklesen, beklenen string) {
	if gerceklesen != beklenen {
		t.Errorf("Hata : '%s' alanı '%s' beklenirken '%s' geldi.", alanAdi, beklenen, gerceklesen)
	}
}

func TestInvoiceDeserialize(t *testing.T) {

	var invoice Invoice

	err := xml.Unmarshal([]byte(xmlTicariFatura), &invoice)

	if err != nil {
		t.Errorf("XML dosya hatalı: %v", err)
	}

	stringKarsilastir(t, "invoice.UBLVersionID", invoice.UBLVersionID, "2.1")
	stringKarsilastir(t, "invoice.CustomizationID", invoice.CustomizationID, "TR1.2")
	stringKarsilastir(t, "invoice.IssueDate", invoice.IssueDate.String(), "2009-01-05 00:00:00 +0000 UTC")
	stringKarsilastir(t, "invoice.IssueTime", invoice.IssueTime.String(), "0000-01-01 14:42:00 +0000 UTC")
	stringKarsilastir(t, "invoice.ID.Value", invoice.ID.Value, "GIB2009000000011")

}
