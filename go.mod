module github.com/chremoas/lookup-cmd

go 1.14

require (
	github.com/chremoas/chremoas v1.3.0
	github.com/chremoas/esi-srv v1.3.0
	github.com/chremoas/services-common v1.3.0
	github.com/micro/go-micro v1.9.1
	golang.org/x/net v0.0.0-20190724013045-ca1201d0de80
)

replace github.com/chremoas/lookup-cmd => ../lookup-cmd
replace github.com/hashicorp/consul => github.com/hashicorp/consul v1.5.1
