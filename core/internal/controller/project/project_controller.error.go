package project_controller

import "errors"

var (
	ErrEpNotValid    = errors.New("Endpoint is not valid")
	ErrIpV4NotValid  = errors.New("IPv4 address is not valid")
	ErrPortNotValid  = errors.New("Port is not valid")
	ErrSubAlrTaken   = errors.New("Subdomain is already taken")
	ErrDomNotInPool  = errors.New("The domain of this subdomain is not part of the pool of valid domains")
	ErrRouteNotValid = errors.New("Route is not correct")
	ErrSubNotValid   = errors.New("Subdomain is not valid")
)
