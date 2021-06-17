package command

import (
	"encoding/xml"
	"strings"
)

type CommandResp struct {
	Status     string `xml:"status,attr" json:"status"`
	StatusText string `xml:"status_text,attr" json:"status_text"`
}

type GmpResponse struct {
	XMLName xml.Name `xml:"gmp_response"`
	*CommandResp
}

func HasSuccess(status string) bool {
	return strings.HasPrefix(status, "2")
}

type UsageType string

const (
	Scan   UsageType = "scan"
	Audit  UsageType = "audit"
	Policy UsageType = "policy"
	All    UsageType = ""
)

type Owner struct {
	Name string `xml:"name" json:"name"` // name
}

type CertificateInfo struct {
	TimeStatus     CertificateStatus `xml:"time_status" json:"time_status"`         // Whether the certificate is valid, expired or not active yet.
	ActivationTime string            `xml:"activation_time" json:"activation_time"` // Time before which the certificate is not valid.
	ExpirationTime string            `xml:"expiration_time" json:"expiration_time"` // Time after which the certificate is no longer valid.
	Issuer         string            `xml:"issuer" json:"issuer"`                   // DN of the issuer of the certificate.
	Md5Fingerprint string            `xml:"md5_fingerprint" json:"md5_fingerprint"` // MD5 fingerprint of the certificate.
}

type CertificateStatus string

const (
	Expired  CertificateStatus = "expired"
	Inactive CertificateStatus = "inactive"
	Unknown  CertificateStatus = "unknown"
	Valid    CertificateStatus = "valid"
)

type Order string

const (
	OrderAscending  Order = "ascending"
	OrderDescending Order = "descending"
)
