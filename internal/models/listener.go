package xdsmodels

type Listener struct {
	Name            string
	Address         string
	Port            uint32
	RouteConfigName string

	Cert *Cert
}
