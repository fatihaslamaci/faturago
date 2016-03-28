// faturago
package main

import (
	"encoding/xml"
	"fmt"
	"os"

	"faturago/invoice"
	"faturago/invoice/serialized"
)

func main() {

	var invoice invoice.InvoiceType

	var invoice2 serialized.InvoiceType

	err := xml.Unmarshal([]byte(xmlTicariFatura), &invoice)

	if err != nil {
		fmt.Printf("error: %v", err)

	}

	invoice2.UBLVersionID = new(serialized.IdentifierType1)
	invoice2.UBLVersionID.Value = invoice.UBLVersionID.Value

	//invoice.UBLExtensions[0].XMLName = xml.Name{Local: "ext:UBLExtensions"}

	//invoice.UBLVersionID.XMLName = xml.Name{Local: "cbc:UBLVersionID"}
	//invoice.CustomizationID.XMLName = xml.Name{Local: "cbc:CustomizationID"}

	//if invoice.ProfileID != nil {
	//	invoice.ProfileID.XMLName = xml.Name{Local: "cbc:ProfileID"}
	//}

	//if invoice.IssueDate != nil {
	//	invoice.IssueDate.XMLName = xml.Name{Local: "cbc:IssueDate"}
	//}

	enc := xml.NewEncoder(os.Stdout)
	enc.Indent("  ", "  ")
	enc.Encode(invoice2)
}
