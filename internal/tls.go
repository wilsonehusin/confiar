package internal

import (
	"github.com/wilsonehusin/confiar/internal/cryptographer"
)

func NewTLSSelfAuthority(cryptographerBackend string, fqdn string) error {
	gostd := cryptographer.GoStd{}
	return gostd.NewTLSSelfAuthority(fqdn)
}
