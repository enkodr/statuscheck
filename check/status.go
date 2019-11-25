package check

import "enkodr.dev/status/config"

// Status is the interface to be used for creating testing endpoints
type Status interface {
	Check(config config.Check) (bool, error)
}

// Make will create a configured endpoint
func Make(typeName string) Status {
	switch typeName {
	case "http":
		return HTTP{}
	case "port":
		return Port{}
	default:
		return nil
	}
}
