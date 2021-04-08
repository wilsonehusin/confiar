package internal

import (
	"fmt"

	"github.com/wilsonehusin/confiar/internal/cryptographer"
)

var cryptoBackend cryptographer.Cryptographer

func NewTLSSelfAuthority(backendType string, names []string, ips []string) error {
	switch backendType {
	case "gostd":
		cryptoBackend = &cryptographer.GoStd{}
	default:
		return fmt.Errorf("unknown cryptographer backend type: %s", backendType)
	}
	return cryptoBackend.NewTLSSelfAuthority(names, ips)
}
