package command

import (
	"encoding/json"
	"encoding/xml"
	"errors"
)

// Get one or many scanners.
// The client uses the get_scanners command to get scanner information.

type GetScanner struct {
	XMLName   xml.Name `xml:"get_scanners"`
	ScannerId string   `xml:"scanner_id,attr,omitempty" json:"scanner_id"` // ID of single scanner to get.
	Filter    string   `xml:"filter,attr,omitempty" json:"filter"`         // Filter term to use to filter query.
	FiltId    string   `xml:"filt_id,attr,omitempty" json:"filt_id"`       // ID of filter to use to filter query.
	Trash     bool     `xml:"trash,attr,omitempty" json:"trash"`           // Whether to get the trashcan scanners instead.
	Details   bool     `xml:"details,attr,omitempty" json:"details"`       // Whether to include extra details like tasks using this scanner.
}

type GetScannerResp struct {
	*CommandResp
	XMLName      xml.Name                    `xml:"get_scanners_response"`
	Scanner      []*Scanner                  `xml:"scanner" json:"scanner"`
	Filters      *GetScannerRespFilter       `xml:"filters" json:"filters"`
	Sort         *GetScannerRespSort         `xml:"sort" json:"sort"`
	Scanners     *GetScannerRespScanners     `xml:"scanners" json:"scanners"`
	ScannerCount *GetScannerRespScannerCount `xml:"scanner_count" json:"scanner_count"`
}

type Scanner struct {
	Id               string                     `xml:"id,attr" json:"id"`
	Owner            *Owner                     `xml:"owner" json:"owner"`     // Owner of the scanner.
	Name             string                     `xml:"name" json:"name"`       // The name of the scanner.
	Comment          string                     `xml:"comment" json:"comment"` // The comment on the scanner.
	Copy             string                     `xml:"copy" json:"copy"`       // The UUID of an existing scanner.
	CreationTime     string                     `xml:"creation_time" json:"creation_time"`
	ModificationTime string                     `xml:"modification_time" json:"modification_time"`
	Writable         bool                       `xml:"writable" json:"writable"`                 // Whether the scanner is writable.
	InUse            bool                       `xml:"in_use" json:"in_use"`                     // Whether any tasks are using the scanner.
	Permissions      *GetScannerRespPermissions `xml:"permissions" json:"permissions"`           // Permissions that the current user has on the scanner.
	UserTags         *GetScannerRespUserTags    `xml:"user_tags" json:"user_tags"`               // Info on tags attached to the chedule.
	CaPubInfo        *CertificateInfo           `xml:"ca_pub_info" json:"ca_pub_info"`           // Info about the CA certificate.
	CertificateInfo  *CertificateInfo           `xml:"certificate_info" json:"certificate_info"` // Info about the certificate.
	Host             string                     `xml:"host" json:"host"`                         // Host of the scanner.
	Port             uint16                     `xml:"port" json:"port"`                         // Port of the scanner.
	Type             string                     `xml:"type" json:"type"`                         // Type of the scanner.
	CaPub            string                     `xml:"ca_pub" json:"ca_pub"`                     // CA Certificate to verify the scanner's certificate.
	Credential       *GetScannerRespCredential  `xml:"credential" json:"credential"`             // Client certificate credential for the Scanner.
	Configs          *GetScannerRespConfigs     `xml:"configs" json:"configs"`
}

type GetScannerRespPermissions struct {
	Permission []*GetScannerRespPermission `xml:"permission" json:"permission"`
}

type GetScannerRespPermission struct {
	Name string `xml:"name" json:"name"` // The name of the permission.
}

type GetScannerRespUserTags struct {
	Count int                  `xml:"count" json:"count"` // Number of attached tags.
	Tags  []*GetScannerRespTag `xml:"tag" json:"tag"`     // Short info on an individual tag (only if details were requested).
}

type GetScannerRespTag struct {
	Id      string `xml:"id,attr" json:"id"`      // UUID of the tag.
	Name    string `xml:"name" json:"name"`       // Name of the tag (usually namespace:predicate).
	Value   string `xml:"value" json:"value"`     // Value of the tag.
	Comment string `xml:"comment" json:"comment"` // Comment for the tag.
}

type GetScannerRespCredential struct {
	Id    string `xml:"id,attr" json:"id"`
	Name  string `xml:"name" json:"name"`   // Name of the credential.
	Trash bool   `xml:"trash" json:"trash"` // Whether the credential is in the trashcan.
}

type GetScannerRespConfigs struct {
	Config []*GetScannerRespConfig `xml:"config" json:"config"`
}

type GetScannerRespConfig struct {
	Id          string                     `xml:"id,attr" json:"id"`
	Name        string                     `xml:"name" json:"name"`               // The name of the config.
	Permissions *GetScannerRespPermissions `xml:"permissions" json:"permissions"` // Permissions the user has on the config.
}

type GetScannerRespTask struct {
	XMLName     xml.Name                   `xml:"task"`
	Id          string                     `xml:"id,attr" json:"id"`
	Name        string                     `xml:"name" json:"name"`               // The name of the task.
	Permissions *GetScannerRespPermissions `xml:"permissions" json:"permissions"` // Permissions the user has on the task.
}

type GetScannerRespFilter struct {
	Id       string                  `xml:"id,attr" json:"id"`        // UUID of filter if any, else 0.
	Term     string                  `xml:"term" json:"term"`         // Filter term.
	Name     string                  `xml:"name" json:"name"`         // Filter name, if applicable.
	Keywords *GetScannerRespKeywords `xml:"keywords" json:"keywords"` // Filter broken down into keywords.
}

type GetScannerRespKeywords struct {
	Keyword []*GetScannerRespKeyword `xml:"keyword" json:"keyword"`
}

type GetScannerRespKeyword struct {
	Column   string `xml:"column" json:"column"`     // Column prefix.
	Relation string `xml:"relation" json:"relation"` // Relation operator.
	Value    string `xml:"value" json:"value"`       // The filter text.
}

type GetScannerRespSort struct {
	Field GetScannerRespField `xml:"field" json:"field"`
}

type GetScannerRespField struct {
	Order Order `xml:"order" json:"order"`
}

type GetScannerRespScanners struct {
	Start int `xml:"start,attr" json:"start"` // First scanner.
	Max   int `xml:"max,attr" json:"max"`     // Maximum number of scanners.
}

type GetScannerRespScannerCount struct {
	Filtered int `xml:"filtered" json:"filtered"` // Number of scanners after filtering.
	Page     int `xml:"page" json:"page"`         // Number of scanners on current page.
}

func NewGetAllScanner() *GetScanner {
	return &GetScanner{}
}

func NewGetScannerWithId(id string) *GetScanner {
	return &GetScanner{
		ScannerId: id,
	}
}

func NewGetScanner(filter *GetScanner) *GetScanner {
	return filter
}

func (a *GetScanner) Command() string {
	return "get_scanner"
}

func (a *GetScanner) String() string {
	data, _ := xml.Marshal(a)
	return string(data)
}

func (a *GetScanner) StringWithJson() string {
	data, _ := json.Marshal(a)
	return string(data)
}

func (a *GetScanner) GetRespStruct() interface{} {
	return &GetScannerResp{}
}

func (a *GetScanner) Handler(data interface{}) (interface{}, error) {
	resp := data.(*GetScannerResp)

	if !HasSuccess(resp.Status) {
		return nil, errors.New(resp.StatusText)
	}

	return resp, nil
}
