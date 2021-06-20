package command

import (
	"encoding/json"
	"encoding/xml"
	"errors"
)

// Modify an existing scanner.
// The client uses the modify_scanner command to change an existing scanner.
type ModifyScanner struct {
	XMLName    xml.Name                 `xml:"modify_scanner"`
	ScannerId  string                   `xml:"scanner_id,attr" json:"scanner_id"`      // ID of scanner to modify.
	Comment    string                   `xml:"comment,omitempty" json:"comment"`       // Comment on scanner.
	Name       string                   `xml:"name,omitempty" json:"name"`             // Name of scanner.
	Host       string                   `xml:"host,omitempty" json:"host"`             // Host of the scanner.
	Port       uint16                   `xml:"port,omitempty" json:"port"`             // Port of the scanner.
	Type       int                      `xml:"type,omitempty" json:"type"`             // Type of the scanner. '1' for OSP, '2' for OpenVAS (classic) Scanner.
	CaPub      *CertificateInfo         `xml:"ca_pub,omitempty" json:"ca_pub"`         // Certificate of CA to verify scanner's certificate.
	Credential *CreateScannerCredential `xml:"credential,omitempty" json:"credential"` // Client certificate credential for the Scanner.
}

type ModifyScannerResp struct {
	*CommandResp
	XMLName xml.Name `xml:"modify_scanner_response"`
}

func NewModifyScanner(scanner *ModifyScanner) *ModifyScanner {
	return scanner
}

func (a *ModifyScanner) Command() string {
	return "modify_scanner"
}

func (a *ModifyScanner) String() string {
	data, _ := xml.Marshal(a)
	return string(data)
}

func (a *ModifyScanner) StringWithJson() string {
	data, _ := json.Marshal(a)
	return string(data)
}

func (a *ModifyScanner) GetRespStruct() interface{} {
	return &ModifyScannerResp{}
}

func (a *ModifyScanner) Handler(data interface{}) (interface{}, error) {
	resp := data.(*ModifyScannerResp)

	if !HasSuccess(resp.Status) {
		return nil, errors.New(resp.StatusText)
	}

	return nil, nil
}
