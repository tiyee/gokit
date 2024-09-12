package jwt

type Header struct {
	Typ string `json:"typ"`
	Alg Signer `json:"alg"`
}
