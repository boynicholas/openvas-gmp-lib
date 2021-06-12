package command

import "errors"

var ErrInvalidUsernameOrPassword = errors.New("invalid username or password")
var ErrGvmCACertNotFound = errors.New("the GVM CA cert has not found")
var ErrGvmClientCertNotFound = errors.New("the GVM Client cert has not found")
var ErrGvmClientKeyNotFound = errors.New("the GVM Client key has not found")
var ErrInvalidGvmCACert = errors.New("invalid GVM CA certificate")
var ErrUnsupportGVMVersion = errors.New("unsupport GVM Version")
