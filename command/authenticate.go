package command

import (
	"encoding/json"
	"encoding/xml"
	"errors"
)

// Authenticate with the manager.
// The client uses the authenticate command to authenticate with the Manager.
// The client normally executes this command at the beginning of each connection. The only command permitted before authentication is get_version.
type Authenticate struct {
	XMLName     xml.Name     `xml:"authenticate"`
	Credentials *Credentials `xml:"credentials" json:"credentials"`
}

type AuthenticateResp struct {
	*CommandResp
	XMLName  xml.Name `xml:"authenticate_response"`
	Role     string   `xml:"role" json:"role"`
	Timezone string   `xml:"timezone" json:"timezone"`
}

func NewAuthenticate(username string, password string) *Authenticate {

	return &Authenticate{
		Credentials: &Credentials{
			Username: username,
			Password: password,
		},
	}
}

func (a *Authenticate) Command() string {
	return "authenticate"
}

func (a *Authenticate) String() string {
	data, _ := xml.Marshal(a)
	return string(data)
}

func (a *Authenticate) StringWithJson() string {
	data, _ := json.Marshal(a)
	return string(data)
}

func (a *Authenticate) GetRespStruct() interface{} {
	return &AuthenticateResp{}
}

func (a *Authenticate) Handler(data interface{}) (interface{}, error) {
	resp := data.(*AuthenticateResp)

	if resp.Status != "200" {
		return nil, errors.New(resp.StatusText)
	}

	return nil, nil
}

type Credentials struct {
	Username string `xml:"username" json:"username"`
	Password string `xml:"password" json:"password"`
}
