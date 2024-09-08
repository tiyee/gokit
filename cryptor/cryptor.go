package cryptor

import (
	"bytes"
	"crypto/aes"
	"crypto/cipher"
	"crypto/des"
	"crypto/rand"
	"crypto/rsa"
	"crypto/sha256"
	"crypto/x509"
	"encoding/pem"
	"errors"
	"io"
	"os"
)

// AesEcbEncrypt encrypt data with key use AES ECB algorithm
// len(key) should be 16, 24 or 32.
func AesEcbEncrypt(data, key []byte) ([]byte, error) {
	size := len(key)
	if !isOneOf(size, []int{16, 24, 32}) {
		return nil, errors.New("key length should be 16 or 24 or 32")
	}

	length := (len(data) + aes.BlockSize) / aes.BlockSize
	plain := make([]byte, length*aes.BlockSize)

	copy(plain, data)

	pad := byte(len(plain) - len(data))
	for i := len(data); i < len(plain); i++ {
		plain[i] = pad
	}

	encrypted := make([]byte, len(plain))
	cp, err := aes.NewCipher(generateAesKey(key, size))
	if err != nil {
		return nil, err
	}

	for bs, be := 0, cp.BlockSize(); bs <= len(data); bs, be = bs+cp.BlockSize(), be+cp.BlockSize() {
		cp.Encrypt(encrypted[bs:be], plain[bs:be])
	}

	return encrypted, nil
}

// AesEcbDecrypt decrypt data with key use AES ECB algorithm
// len(key) should be 16, 24 or 32.
func AesEcbDecrypt(encrypted, key []byte) ([]byte, error) {
	size := len(key)
	if !isOneOf(size, []int{16, 24, 32}) {
		return nil, errors.New("key length should be 16 or 24 or 32")
	}
	newCipher, _ := aes.NewCipher(generateAesKey(key, size))
	decrypted := make([]byte, len(encrypted))

	for bs, be := 0, newCipher.BlockSize(); bs < len(encrypted); bs, be = bs+newCipher.BlockSize(), be+newCipher.BlockSize() {
		newCipher.Decrypt(decrypted[bs:be], encrypted[bs:be])
	}

	trim := 0
	if len(decrypted) > 0 {
		trim = len(decrypted) - int(decrypted[len(decrypted)-1])
	}

	return decrypted[:trim], nil
}

// AesCbcEncrypt encrypt data with key use AES CBC algorithm
// len(key) should be 16, 24 or 32.
func AesCbcEncrypt(data, key []byte) ([]byte, error) {
	size := len(key)
	if !isOneOf(size, []int{16, 24, 32}) {
		return nil, errors.New("key length should be 16 or 24 or 32")
	}

	block, err := aes.NewCipher(key)
	if err != nil {
		return nil, err
	}
	data = pkcs7Padding(data, block.BlockSize())

	encrypted := make([]byte, aes.BlockSize+len(data))
	iv := encrypted[:aes.BlockSize]
	if _, err := io.ReadFull(rand.Reader, iv); err != nil {
		panic(err)
	}

	mode := cipher.NewCBCEncrypter(block, iv)
	mode.CryptBlocks(encrypted[aes.BlockSize:], data)

	return encrypted, nil
}

// AesCbcDecrypt decrypt data with key use AES CBC algorithm
// len(key) should be 16, 24 or 32.
func AesCbcDecrypt(encrypted, key []byte) ([]byte, error) {
	size := len(key)
	if !isOneOf(size, []int{16, 24, 32}) {
		return nil, errors.New("key length should be 16 or 24 or 32")
	}
	block, _ := aes.NewCipher(key)

	iv := encrypted[:aes.BlockSize]
	encrypted = encrypted[aes.BlockSize:]

	mode := cipher.NewCBCDecrypter(block, iv)
	mode.CryptBlocks(encrypted, encrypted)

	decrypted := pkcs7UnPadding(encrypted)
	return decrypted, nil
}

// AesCtrCrypt encrypt data with key use AES CTR algorithm
// len(key) should be 16, 24 or 32.
func AesCtrCrypt(data, key []byte) ([]byte, error) {
	size := len(key)
	if !isOneOf(size, []int{16, 24, 32}) {
		return nil, errors.New("key length should be 16 or 24 or 32")
	}
	block, _ := aes.NewCipher(key)

	iv := bytes.Repeat([]byte("1"), block.BlockSize())
	stream := cipher.NewCTR(block, iv)

	dst := make([]byte, len(data))
	stream.XORKeyStream(dst, data)

	return dst, nil
}

// AesCfbEncrypt encrypt data with key use AES CFB algorithm
// len(key) should be 16, 24 or 32.
func AesCfbEncrypt(data, key []byte) ([]byte, error) {
	size := len(key)
	if !isOneOf(size, []int{16, 24, 32}) {
		return nil, errors.New("key length should be 16 or 24 or 32")
	}
	block, err := aes.NewCipher(key)
	if err != nil {
		return nil, err
	}

	encrypted := make([]byte, aes.BlockSize+len(data))
	iv := encrypted[:aes.BlockSize]

	if _, err := io.ReadFull(rand.Reader, iv); err != nil {
		return nil, err
	}

	stream := cipher.NewCFBEncrypter(block, iv)
	stream.XORKeyStream(encrypted[aes.BlockSize:], data)

	return encrypted, nil
}

// AesCfbDecrypt decrypt data with key use AES CFB algorithm
// len(encrypted) should be great than 16, len(key) should be 16, 24 or 32.
func AesCfbDecrypt(encrypted, key []byte) ([]byte, error) {
	size := len(key)
	if !isOneOf(size, []int{16, 24, 32}) {
		return nil, errors.New("key length should be 16 or 24 or 32")
	}

	if len(encrypted) < aes.BlockSize {
		return nil, errors.New("encrypted data is too short")
	}

	block, _ := aes.NewCipher(key)
	iv := encrypted[:aes.BlockSize]
	encrypted = encrypted[aes.BlockSize:]

	stream := cipher.NewCFBDecrypter(block, iv)

	stream.XORKeyStream(encrypted, encrypted)

	return encrypted, nil
}

// AesOfbEncrypt encrypt data with key use AES OFB algorithm
// len(key) should be 16, 24 or 32.
func AesOfbEncrypt(data, key []byte) ([]byte, error) {
	size := len(key)
	if !isOneOf(size, []int{16, 24, 32}) {
		return nil, errors.New("key length should be 16 or 24 or 32")
	}

	block, err := aes.NewCipher(key)
	if err != nil {
		return nil, err
	}

	data = pkcs7Padding(data, aes.BlockSize)
	encrypted := make([]byte, aes.BlockSize+len(data))
	iv := encrypted[:aes.BlockSize]
	if _, err := io.ReadFull(rand.Reader, iv); err != nil {
		return nil, err
	}

	stream := cipher.NewOFB(block, iv)
	stream.XORKeyStream(encrypted[aes.BlockSize:], data)

	return encrypted, nil
}

// AesOfbDecrypt decrypt data with key use AES OFB algorithm
// len(key) should be 16, 24 or 32.
func AesOfbDecrypt(data, key []byte) ([]byte, error) {
	size := len(key)
	if !isOneOf(size, []int{16, 24, 32}) {
		return nil, errors.New("key length should be 16 or 24 or 32")
	}

	block, err := aes.NewCipher(key)
	if err != nil {
		return nil, err
	}

	iv := data[:aes.BlockSize]
	data = data[aes.BlockSize:]
	if len(data)%aes.BlockSize != 0 {
		return nil, errors.New("data is not a multiple of the block size")
	}

	decrypted := make([]byte, len(data))
	mode := cipher.NewOFB(block, iv)
	mode.XORKeyStream(decrypted, data)

	decrypted = pkcs7UnPadding(decrypted)

	return decrypted, nil
}

// DesEcbEncrypt encrypt data with key use DES ECB algorithm
// len(key) should be 8.
func DesEcbEncrypt(data, key []byte) ([]byte, error) {
	length := (len(data) + des.BlockSize) / des.BlockSize
	plain := make([]byte, length*des.BlockSize)
	copy(plain, data)

	pad := byte(len(plain) - len(data))
	for i := len(data); i < len(plain); i++ {
		plain[i] = pad
	}

	encrypted := make([]byte, len(plain))
	cip, err := des.NewCipher(generateDesKey(key))
	if err != nil {
		return nil, err
	}
	for bs, be := 0, cip.BlockSize(); bs <= len(data); bs, be = bs+cip.BlockSize(), be+cip.BlockSize() {
		cip.Encrypt(encrypted[bs:be], plain[bs:be])
	}

	return encrypted, nil
}

// DesEcbDecrypt decrypt data with key use DES ECB algorithm
// len(key) should be 8.
func DesEcbDecrypt(encrypted, key []byte) ([]byte, error) {
	cip, err := des.NewCipher(generateDesKey(key))
	if err != nil {
		return nil, err
	}
	decrypted := make([]byte, len(encrypted))

	for bs, be := 0, cip.BlockSize(); bs < len(encrypted); bs, be = bs+cip.BlockSize(), be+cip.BlockSize() {
		cip.Decrypt(decrypted[bs:be], encrypted[bs:be])
	}

	trim := 0
	if len(decrypted) > 0 {
		trim = len(decrypted) - int(decrypted[len(decrypted)-1])
	}

	return decrypted[:trim], nil
}

// DesCbcEncrypt encrypt data with key use DES CBC algorithm
// len(key) should be 8.
func DesCbcEncrypt(data, key []byte) ([]byte, error) {
	size := len(key)
	if size != 8 {
		return nil, errors.New("key length shoud be 8")
	}

	block, err := des.NewCipher(key)
	if err != nil {
		return nil, err
	}
	data = pkcs7Padding(data, block.BlockSize())

	encrypted := make([]byte, des.BlockSize+len(data))
	iv := encrypted[:des.BlockSize]

	if _, err = io.ReadFull(rand.Reader, iv); err != nil {
		return nil, err
	}

	mode := cipher.NewCBCEncrypter(block, iv)
	mode.CryptBlocks(encrypted[des.BlockSize:], data)

	return encrypted, nil
}

// DesCbcDecrypt decrypt data with key use DES CBC algorithm
// len(key) should be 8.
func DesCbcDecrypt(encrypted, key []byte) ([]byte, error) {
	size := len(key)
	if size != 8 {
		return nil, errors.New("key length shoud be 8")
	}

	block, err := des.NewCipher(key)
	if err != nil {
		return nil, err
	}

	iv := encrypted[:des.BlockSize]
	encrypted = encrypted[des.BlockSize:]

	mode := cipher.NewCBCDecrypter(block, iv)
	mode.CryptBlocks(encrypted, encrypted)

	decrypted := pkcs7UnPadding(encrypted)
	return decrypted, nil
}

// DesCtrCrypt encrypt data with key use DES CTR algorithm
// len(key) should be 8.
func DesCtrCrypt(data, key []byte) ([]byte, error) {
	size := len(key)
	if size != 8 {
		return nil, errors.New("key length shoud be 8")
	}

	block, err := des.NewCipher(key)
	if err != nil {
		return nil, err
	}
	iv := bytes.Repeat([]byte("1"), block.BlockSize())
	stream := cipher.NewCTR(block, iv)

	dst := make([]byte, len(data))
	stream.XORKeyStream(dst, data)

	return dst, nil
}

// DesCfbEncrypt encrypt data with key use DES CFB algorithm
// len(key) should be 8.

func DesCfbEncrypt(data, key []byte) ([]byte, error) {
	size := len(key)
	if size != 8 {
		return nil, errors.New("key length shoud be 8")
	}

	block, err := des.NewCipher(key)
	if err != nil {
		return nil, err
	}

	encrypted := make([]byte, des.BlockSize+len(data))
	iv := encrypted[:des.BlockSize]
	if _, err = io.ReadFull(rand.Reader, iv); err != nil {
		return nil, err
	}

	stream := cipher.NewCFBEncrypter(block, iv)
	stream.XORKeyStream(encrypted[des.BlockSize:], data)

	return encrypted, nil
}

// DesCfbDecrypt decrypt data with key use DES CFB algorithm
// len(encrypted) should be great than 16, len(key) should be 8.
func DesCfbDecrypt(encrypted, key []byte) ([]byte, error) {
	size := len(key)
	if size != 8 {
		return nil, errors.New("key length shoud be 8")

	}

	block, err := des.NewCipher(key)
	if err != nil {
		return nil, err
	}
	if len(encrypted) < des.BlockSize {
		return nil, errors.New("encrypted data is too short")
	}
	iv := encrypted[:des.BlockSize]
	encrypted = encrypted[des.BlockSize:]

	stream := cipher.NewCFBDecrypter(block, iv)
	stream.XORKeyStream(encrypted, encrypted)

	return encrypted, nil
}

// DesOfbEncrypt encrypt data with key use DES OFB algorithm
// len(key) should be 8.
func DesOfbEncrypt(data, key []byte) ([]byte, error) {
	size := len(key)
	if size != 8 {
		return nil, errors.New("key length shoud be 8")
	}

	block, err := des.NewCipher(key)
	if err != nil {
		return nil, err
	}
	data = pkcs7Padding(data, des.BlockSize)
	encrypted := make([]byte, des.BlockSize+len(data))
	iv := encrypted[:des.BlockSize]
	if _, err = io.ReadFull(rand.Reader, iv); err != nil {
		return nil, err
	}

	stream := cipher.NewOFB(block, iv)
	stream.XORKeyStream(encrypted[des.BlockSize:], data)

	return encrypted, nil
}

// DesOfbDecrypt decrypt data with key use DES OFB algorithm
// len(key) should be 8.
func DesOfbDecrypt(data, key []byte) ([]byte, error) {
	size := len(key)
	if size != 8 {
		return nil, errors.New("key length shoud be 8")
	}

	block, err := des.NewCipher(key)
	if err != nil {
		return nil, err
	}

	iv := data[:des.BlockSize]
	data = data[des.BlockSize:]
	if len(data)%des.BlockSize != 0 {
		return nil, errors.New("data is not a multiple of the block size")
	}

	decrypted := make([]byte, len(data))
	mode := cipher.NewOFB(block, iv)
	mode.XORKeyStream(decrypted, data)

	decrypted = pkcs7UnPadding(decrypted)

	return decrypted, nil
}

// GenerateRsaKey create rsa private and public pemo file.
func GenerateRsaKey(keySize int, priKeyFile, pubKeyFile string) error {
	// private key
	privateKey, err := rsa.GenerateKey(rand.Reader, keySize)
	if err != nil {
		return err
	}

	derText := x509.MarshalPKCS1PrivateKey(privateKey)

	block := pem.Block{
		Type:  "rsa private key",
		Bytes: derText,
	}

	file, err := os.Create(priKeyFile)
	if err != nil {
		return nil
	}
	err = pem.Encode(file, &block)
	if err != nil {
		return err
	}

	file.Close()

	// public key
	publicKey := privateKey.PublicKey

	derpText, err := x509.MarshalPKIXPublicKey(&publicKey)
	if err != nil {
		return err
	}

	block = pem.Block{
		Type:  "rsa public key",
		Bytes: derpText,
	}

	file, err = os.Create(pubKeyFile)
	if err != nil {
		return err
	}

	err = pem.Encode(file, &block)
	if err != nil {
		return err
	}

	file.Close()

	return nil
}

// RsaEncrypt encrypt data with ras algorithm.
func RsaEncrypt(data []byte, pubKeyFileName string) ([]byte, error) {
	file, err := os.Open(pubKeyFileName)
	if err != nil {
		return nil, err
	}
	fileInfo, err := file.Stat()
	if err != nil {
		return nil, err
	}
	defer file.Close()
	buf := make([]byte, fileInfo.Size())

	_, err = file.Read(buf)
	if err != nil {
		return nil, err
	}

	block, _ := pem.Decode(buf)

	pubInterface, err := x509.ParsePKIXPublicKey(block.Bytes)
	if err != nil {
		return nil, err
	}
	pubKey := pubInterface.(*rsa.PublicKey)

	cipherText, err := rsa.EncryptPKCS1v15(rand.Reader, pubKey, data)
	if err != nil {
		return nil, err
	}
	return cipherText, nil
}

// RsaDecrypt decrypt data with ras algorithm.
func RsaDecrypt(data []byte, privateKeyFileName string) ([]byte, error) {
	file, err := os.Open(privateKeyFileName)
	if err != nil {
		return nil, err
	}
	fileInfo, err := file.Stat()
	if err != nil {
		return nil, err
	}
	buf := make([]byte, fileInfo.Size())
	defer file.Close()

	_, err = file.Read(buf)
	if err != nil {
		return nil, err
	}

	block, _ := pem.Decode(buf)

	priKey, err := x509.ParsePKCS1PrivateKey(block.Bytes)
	if err != nil {
		return nil, err
	}

	plainText, err := rsa.DecryptPKCS1v15(rand.Reader, priKey, data)
	if err != nil {
		return nil, err
	}
	return plainText, nil
}

// GenerateRsaKeyPair create rsa private and public key.
func GenerateRsaKeyPair(keySize int) (*rsa.PrivateKey, *rsa.PublicKey) {
	privateKey, _ := rsa.GenerateKey(rand.Reader, keySize)
	return privateKey, &privateKey.PublicKey
}

// RsaEncryptOAEP encrypts the given data with RSA-OAEP.
func RsaEncryptOAEP(data []byte, label []byte, key rsa.PublicKey) ([]byte, error) {
	encryptedBytes, err := rsa.EncryptOAEP(sha256.New(), rand.Reader, &key, data, label)
	if err != nil {
		return nil, err
	}

	return encryptedBytes, nil
}

// RsaDecryptOAEP decrypts the data with RSA-OAEP.
func RsaDecryptOAEP(ciphertext []byte, label []byte, key rsa.PrivateKey) ([]byte, error) {
	decryptedBytes, err := rsa.DecryptOAEP(sha256.New(), rand.Reader, &key, ciphertext, label)
	if err != nil {
		return nil, err
	}

	return decryptedBytes, nil
}
