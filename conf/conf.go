package conf

type EncryptStruct struct {
	Type       string
	Key        string
	PublicKey  string
	PrivateKey string
	Mode       string
}

var myConf map[string]EncryptStruct

func init() {
	myConf = make(map[string]EncryptStruct)
	doLoadConfig()
}

func GetMyConf() map[string]EncryptStruct {
	return myConf
}

func doLoadConfig() {
	myConf["email1"] = EncryptStruct{Type: "des", Mode: "CBC", Key: "1234567%"}
	myConf["email2"] = EncryptStruct{Type: "des3", Mode: "CBC", Key: "1ef8sdfsd04a44cc8a7fab24"}
	myConf["email3"] = EncryptStruct{Type: "aes", Mode: "CBC", Key: "!23qazwsx%rts7_u"}
	myConf["email4"] = EncryptStruct{
		Type: "rsa",
		PublicKey: `
-----BEGIN PUBLIC KEY-----
MIGfMA0GCSqGSIb3DQEBAQUAA4GNADCBiQKBgQC38FLQSBPKnCcLb6puvo7RV6w3
XDVkV5bMEmSgaLyq1WJug3NyF88FdtEs0OhGJThSlQd7YvZQ/yswID8vL6rM7yvd
oCX+or02BV+e42l4OHAyG6E0qttwG4kfWUMchq4xGVx26fCUHG/cvsEZNqbC7qI/
PtHSlXPzksQ698qTpwIDAQAB
-----END PUBLIC KEY-----
`,
		PrivateKey: `
-----BEGIN RSA PRIVATE KEY-----
MIICXQIBAAKBgQC38FLQSBPKnCcLb6puvo7RV6w3XDVkV5bMEmSgaLyq1WJug3Ny
F88FdtEs0OhGJThSlQd7YvZQ/yswID8vL6rM7yvdoCX+or02BV+e42l4OHAyG6E0
qttwG4kfWUMchq4xGVx26fCUHG/cvsEZNqbC7qI/PtHSlXPzksQ698qTpwIDAQAB
AoGBALPkPynihDCwfWMq57V9erH0m8Jc1P74xst45Z8YMASnwqewudSIwnhmlvbM
rY4E0su5YuLii2H13OgpcYIVjWVNhBVbuy/XaAixF1cxeXNVekDjdhIyBghMyT74
4r0j17vnlX17dmUVPESsqT4JwTT8PY95h73UZ2ZETgQY/SwpAkEA4JIveSDNHVzp
J/Vi2HmWp2pGNGnZIwVPhLBmTlof9AmGB5Cw+juz6X+84T/JquHK0i/uvAhcrM99
H+E9q1LtIwJBANGuYtQ3jm/L9MBOWc2kNdBQmOErXFhTd8vo3SP4VfIiTYHS+Wd2
FjhjdCHobUchXG+4gu182lTD2G2zhK+LEa0CQBO4YpjVa42nigXrQ3nU+4jKCU4y
+VFc9wWk4+b/fQzWBUIGz7O5qGvZvsFc852g5Eme0e0LIvA97DPveJXh21cCQQCP
VaY5a8Dyq3mj++tnp9khjfbz9bAqsOQzf/urmxclbk2NcasMDq0h3tqEU15gQW3u
8TKPJcoaJ8bLNux9E5QBAkAnXWAuXaEDd4Q/8QpydHn/lQ7ZmKYLrdPOMOHJhGhD
3qwMDqc9JmwFT/AIGUR9671v6OCipg8VQ0qVJf4oypqV
-----END RSA PRIVATE KEY-----
`}
}
