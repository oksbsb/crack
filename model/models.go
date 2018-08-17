package model

type ServerResult struct {
	Hostport string
	Open     bool
	Server   string
	Banner   string
}

type ScanResult struct {
	Hostport string
	Server   func(result ScanResult) ScanResult
	Username string
	Password string
	Success  bool
}
