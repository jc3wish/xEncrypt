# xEncrypt

- 对邮箱，电话，地址等进行加解解密的服务
- 通过加密密钥配在源码文件中，进行编译，可以让普通开发人员并不知道具体的加密key，以及加密算法
- 现在支持 AES,DES,DES3,RSA 加密算法

## 配置密钥或者KEY
源码文件:conf/conf.go
`````go
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
`````

myConf 中的key 用于客户端进行传过来指定用哪个加密配置

## 编译
当前用hprose 做为中间件传输，也可用于其他 RPC 中间件
先安装 hprose-go ,再编译
`````go
go get github.com/hprose/hprose-golang

go build server/hprose/xEncryptServer_hprose.go
`````

## 启动
./xEncryptServer_hprose -host 0.0.0.0 -port 1036

## 客户端使用
`````php
$client = \Hprose\Client::create('tcp://127.0.0.1:1036', false);

$encryptData = $client->DoEncrypt("email1","I am xEncrypt");

$decryptData = $client->DoDecrypt("email1","N1/u1DamOOOHzxOExyJHQA==");

`````


