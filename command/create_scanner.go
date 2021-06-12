package command

import (
	"encoding/json"
	"encoding/xml"
	"errors"
)

// Create a scanner
// The client uses the create_scanner command to create a new scanner.
type CreateScanner struct {
	XMLName xml.Name `xml:"create_scanner"`

	// A name for the scanner.
	Name string `xml:"name" json:"name"`
	// A comment on the scanner.
	Comment string `xml:"comment" json:"comment"`
	// The UUID of an existing scanner.
	Copy string `xml:"copy" json:"copy"`
	// The host of the scanner.
	Host string `xml:"host" json:"host"`
	// The port of the scanner
	Port uint16 `xml:"port" json:"port"`
	// The type of the scanner.
	Type string `xml:"type" json:"type"`
	// Certificate of CA to verify scanner certificate..
	CaPub      string                  `xml:"ca_pub" json:"ca_pub"`
	Credential CreateScannerCredential `xml:"credential" json:"credential"`
}

type CreateScannerCredential struct {
	Id string `xml:"id,attr" json:"id"`
}

type CreateScannerResp struct {
	*CommandResp
	XMLName xml.Name `xml:"create_scanner_response"`
	Id      string   `xml:"id,attr" json:"id"`
}

func NewCreateScanner(scanner *CreateScanner) *CreateScanner {
	return scanner
}

func (a *CreateScanner) Command() string {
	return "create_scanner"
}

func (a *CreateScanner) String() string {
	data, _ := xml.Marshal(a)
	return string(data)
}

func (a *CreateScanner) StringWithJson() string {
	data, _ := json.Marshal(a)
	return string(data)
}

func (a *CreateScanner) GetRespStruct() interface{} {
	return &AuthenticateResp{}
}

func (a *CreateScanner) Handler(data interface{}) (interface{}, error) {
	resp := data.(*CreateScannerResp)

	if resp.Status != "200" {
		return nil, errors.New(resp.StatusText)
	}

	return resp.Id, nil
}
