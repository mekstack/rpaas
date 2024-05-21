package xdsmodels

type DomainEndpoint struct {
	Domain string `validate:"hostname"`
	Host   string `validate:"ip"`
	Port   string `validate:"hostname_port"`
}
