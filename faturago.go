// faturago
package main

import (
	"encoding/xml"
	//"fmt"
	"os"
)

type Person struct {
	Person2

	//XMLName xml.Name `xml:"ns1 person"`
	//NS2     string   `xml:"xmlns:ns2,attr"`

	//Name  string `xml:"name"`
	Phone string `xml:"ns2:phone",omitempty`
}

type Person2 struct {
	XMLName xml.Name `xml:"ns1 person"`
	NS2     string   `xml:"xmlns:ns2,attr"`

	Name  string `xml:"name"`
	Phone string `xml:"ns2:phone",omitempty`
}

func main() {

	//var invoice Invoice

	//invoice.N1 = "urn:oasis:names:specification:ubl:schema:xsd:Invoice-2"
	//invoice.N4 = "http://www.altova.com/samplexml/other-namespace"
	//invoice.UBLVersionID = "1.2"

	//enc := xml.NewEncoder(os.Stdout)
	//enc.Indent("  ", "  ")
	//enc.Encode(invoice)

	//fmt.Println("")

	//err := xml.Unmarshal([]byte(xmlTest), &invoice)

	//if err != nil {
	//	fmt.Printf("error: %v", err)

	//}

	//enc := xml.NewEncoder(os.Stdout)
	//enc.Indent("  ", "  ")
	//enc.Encode(invoice)

	//fmt.Println(invoice.N4)

	//var person Person
	var person2 Person2

	//invoice.N1 = "urn:oasis:names:specification:ubl:schema:xsd:Invoice-2"
	//invoice.N4 = "http://www.altova.com/samplexml/other-namespace"
	//invoice.UBLVersionID = "1.2"

	//enc := xml.NewEncoder(os.Stdout)
	//enc.Indent("  ", "  ")
	//enc.Encode(invoice)

	//fmt.Println("")

	xml.Unmarshal([]byte(xmlTest2), &person2)

	enc2 := xml.NewEncoder(os.Stdout)
	enc2.Indent("  ", "  ")
	enc2.Encode(person2)

}

type Invoice struct {
	//XMLName      xml.Name `xml:"aaa Invoice"`

	XMLName xml.Name `xml:xmlns="urn:oasis:names:specification:ubl:schema:xsd:Invoice-2"  Invoice`
	//Xsi     string   `xml:"schemaLocation,attr"`
	//N1      string   `xml:"n1,attr"`
	//N4      string   `xml:"n4,attr"`
	CBC string `xml:"xmlns:cbc,attr"`

	//UBLVersionID string `xml:cbc: "UBLVersionID"`
	UBLVersionID string `xml:"cbc:UBLVersionID,omitempty"`
}
