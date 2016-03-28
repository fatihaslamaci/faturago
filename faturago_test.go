package main

import (
	"testing"

	"encoding/xml"

	"faturago/invoice"
	//"os"
)

func stringKarsilastir(t *testing.T, alanAdi, gerceklesen, beklenen string) {
	if gerceklesen != beklenen {
		t.Errorf("Hata : '%s' alanı '%s' beklenirken '%s' geldi.", alanAdi, beklenen, gerceklesen)
	}
}

func TestInvoiceDeserialize(t *testing.T) {

	var invoice invoice.InvoiceType

	err := xml.Unmarshal([]byte(xmlTicariFatura), &invoice)

	if err != nil {
		t.Errorf("XML dosya hatalı: %v", err)
	}

	stringKarsilastir(t, "invoice.UBLVersionID", invoice.UBLVersionID.Value, "2.1")
	stringKarsilastir(t, "invoice.CustomizationID", invoice.CustomizationID.Value, "TR1.2")
	stringKarsilastir(t, "invoice.IssueDate", invoice.IssueDate.Value, "2009-01-05")
	stringKarsilastir(t, "invoice.IssueTime", invoice.IssueTime.Value, "14:42:00")
	stringKarsilastir(t, "invoice.ID.Value", invoice.ID.Value, "GIB2009000000011")

}
