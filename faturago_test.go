package main

import (
	"testing"

	"encoding/xml"

	//"os"
)

func stringKarsilastir(t *testing.T, alanAdi, beklenen, gerceklesen string) {
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

	stringKarsilastir(t, "invoice.UBLVersionID", "2.1", invoice.UBLVersionID)
	stringKarsilastir(t, "invoice.CustomizationID", "TR1.2", invoice.CustomizationID)

}
