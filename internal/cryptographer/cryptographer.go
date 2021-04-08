package cryptographer

const certFileName = "cert.pem"
const keyFileName = "key.pem"

type Cryptographer interface {
	NewTLSSelfAuthority([]string, []string) error
}
