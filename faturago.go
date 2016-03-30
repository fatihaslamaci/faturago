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

	var invoice1 invoice.InvoiceType

	var invoice2 serialized.InvoiceType

	err := xml.Unmarshal([]byte(xmlTicariFatura), &invoice1)

	if err != nil {
		fmt.Printf("error: %v", err)

	}

	invoice2 = invoice.InvoiceMap(&invoice1)

	enc := xml.NewEncoder(os.Stdout)
	enc.Indent("  ", "  ")
	enc.Encode(invoice2)
}
