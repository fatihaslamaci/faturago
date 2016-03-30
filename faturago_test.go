package main

import (
	"testing"

	"encoding/xml"

	"faturago/invoice"
	"faturago/invoice/serialized"
	//"os"
)

func stringKarsilastir(t *testing.T, alanAdi, gerceklesen, beklenen string) {
	if gerceklesen != beklenen {
		t.Errorf("Hata : '%s' alanı '%s' beklenirken '%s' geldi.", alanAdi, beklenen, gerceklesen)
	}
}

func TestInvoiceDeserialize(t *testing.T) {

	var invoice1 invoice.InvoiceType
	var invoice2 serialized.InvoiceType

	err := xml.Unmarshal([]byte(xmlTicariFatura), &invoice1)

	invoice2 = invoice.InvoiceMap(&invoice1)

	if err != nil {
		t.Errorf("XML dosya hatalı: %v", err)
	}

	stringKarsilastir(t, "invoice.UBLVersionID", invoice1.UBLVersionID.Value, "2.1")
	stringKarsilastir(t, "map invoice.UBLVersionID", invoice1.UBLVersionID.Value, invoice2.UBLVersionID.Value)

	stringKarsilastir(t, "invoice.CustomizationID", invoice1.CustomizationID.Value, "TR1.2")
	stringKarsilastir(t, "map invoice.CustomizationID", invoice1.CustomizationID.Value, invoice2.CustomizationID.Value)

	stringKarsilastir(t, "invoice.IssueDate", invoice1.IssueDate.Value, "2009-01-05")
	stringKarsilastir(t, "map invoice.IssueDate", invoice1.IssueDate.Value, invoice1.IssueDate.Value)

	stringKarsilastir(t, "invoice.IssueTime", invoice1.IssueTime.Value, "14:42:00")
	stringKarsilastir(t, "invoice.ID.Value", invoice1.ID.Value, "GIB2009000000011")

}
