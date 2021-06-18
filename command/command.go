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

type Permissions struct {
	Permission []*Permission `xml:"permission" json:"permission"`
}

type Permission struct {
	Name string `xml:"name" json:"name"` // The name of the permission.
}

type UserTags struct {
	Count int    `xml:"count" json:"count"` // Number of attached tags.
	Tags  []*Tag `xml:"tag" json:"tag"`     // Short info on an individual tag (only if details were requested).
}

type Tag struct {
	Id      string `xml:"id,attr" json:"id"`      // UUID of the tag.
	Name    string `xml:"name" json:"name"`       // Name of the tag (usually namespace:predicate).
	Value   string `xml:"value" json:"value"`     // Value of the tag.
	Comment string `xml:"comment" json:"comment"` // Comment for the tag.
}

type Filters struct {
	Id       string    `xml:"id,attr" json:"id"`        // UUID of filter if any, else 0.
	Term     string    `xml:"term" json:"term"`         // Filter term.
	Name     string    `xml:"name" json:"name"`         // Filter name, if applicable.
	Keywords *Keywords `xml:"keywords" json:"keywords"` // Filter broken down into keywords.
}

type Keywords struct {
	Keyword []*Keyword `xml:"keyword" json:"keyword"`
}

type Keyword struct {
	Column   string `xml:"column" json:"column"`     // Column prefix.
	Relation string `xml:"relation" json:"relation"` // Relation operator.
	Value    string `xml:"value" json:"value"`       // The filter text.
}

type Sort struct {
	Field Field `xml:"field" json:"field"`
}

type Field struct {
	Order Order `xml:"order" json:"order"`
}
