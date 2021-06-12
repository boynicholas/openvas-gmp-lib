package command

import "encoding/xml"

type CommandResp struct {
	Status     string `xml:"status,attr" json:"status"`
	StatusText string `xml:"status_text,attr" json:"status_text"`
}

type GmpResponse struct {
	XMLName xml.Name `xml:"gmp_response"`
	*CommandResp
}
