// faturago
package main

import (
	"encoding/xml"
	"fmt"
	"os"
)

func main() {

	var invoice Invoice

	err := xml.Unmarshal([]byte(xmlTicariFatura), &invoice)

	if err != nil {
		fmt.Printf("error: %v", err)

	}

	enc := xml.NewEncoder(os.Stdout)
	enc.Indent("  ", "  ")
	enc.Encode(invoice)
}
