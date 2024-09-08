package cryptor

import (
	"fmt"
	"github.com/tiyee/gokit/assert"
	"os"
	"testing"
)

func TestE(t *testing.T) {
	fmt.Println("test error")
	as := assert.NewAssert(t, "Testerror")
	assert.Equal(as, "123", "123")

}
func TestAesEcbEncrypt1(t *testing.T) {
	t.Parallel()
	data := "hello world"
	key := "abcdefghijklmnop"
	as := assert.NewAssert(t, "TestAesEcbEncrypt")
	aesEcbEncrypt, err := AesEcbEncrypt([]byte(data), []byte(key))
	assert.IsNil(as, err)
	aesEcbDecrypt, err := AesEcbDecrypt(aesEcbEncrypt, []byte(key))
	assert.IsNil(as, err)
	assert.Equal(as, data, string(aesEcbDecrypt))
}

func TestAesCbcEncrypt(t *testing.T) {
	t.Parallel()

	data := "hello world"
	key := "abcdefghijklmnop"
	as := assert.NewAssert(t, "TestAesCbcEncrypt")
	aesCbcEncrypt, err := AesCbcEncrypt([]byte(data), []byte(key))
	assert.IsNil(as, err)
	aesCbcDecrypt, err := AesCbcDecrypt(aesCbcEncrypt, []byte(key))
	assert.IsNil(as, err)

	assert.Equal(as, data, string(aesCbcDecrypt))
}

func TestAesCtrCrypt(t *testing.T) {
	t.Parallel()

	data := "hello world"
	key := "abcdefghijklmnop"
	as := assert.NewAssert(t, "TestAesCtrCrypt")
	aesCtrCrypt, err := AesCtrCrypt([]byte(data), []byte(key))
	assert.IsNil(as, err)
	aesCtrDeCrypt, err := AesCtrCrypt(aesCtrCrypt, []byte(key))
	assert.IsNil(as, err)

	assert.Equal(as, data, string(aesCtrDeCrypt))
}

func TestAesCfbEncrypt(t *testing.T) {
	t.Parallel()

	data := "hello world"
	key := "abcdefghijklmnop"
	as := assert.NewAssert(t, "TestAesCfbEncrypt")
	aesCfbEncrypt, err := AesCfbEncrypt([]byte(data), []byte(key))
	assert.IsNil(as, err)
	aesCfbDecrypt, err := AesCfbDecrypt(aesCfbEncrypt, []byte(key))
	assert.IsNil(as, err)
	assert.Equal(as, data, string(aesCfbDecrypt))
}

func TestAesOfbEncrypt(t *testing.T) {
	t.Parallel()

	data := "hello world"
	key := "abcdefghijklmnop"
	as := assert.NewAssert(t, "TestAesOfbEncrypt")
	aesOfbEncrypt, err := AesOfbEncrypt([]byte(data), []byte(key))
	assert.IsNil(as, err)
	aesOfbDecrypt, err := AesOfbDecrypt(aesOfbEncrypt, []byte(key))
	assert.IsNil(as, err)
	assert.Equal(as, data, string(aesOfbDecrypt))
}

func TestDesEcbEncrypt(t *testing.T) {
	t.Parallel()

	data := "hello world"
	key := "abcdefgh"
	as := assert.NewAssert(t, "TestDesEcbEncrypt")
	desEcbEncrypt, err := DesEcbEncrypt([]byte(data), []byte(key))
	assert.IsNil(as, err)
	desEcbDecrypt, err := DesEcbDecrypt(desEcbEncrypt, []byte(key))
	assert.IsNil(as, err)
	assert.Equal(as, data, string(desEcbDecrypt))
}

func TestDesCbcEncrypt(t *testing.T) {
	t.Parallel()

	data := "hello world"
	key := "abcdefgh"
	as := assert.NewAssert(t, "TestDesCbcEncrypt")
	desCbcEncrypt, err := DesCbcEncrypt([]byte(data), []byte(key))
	assert.IsNil(as, err)
	desCbcDecrypt, err := DesCbcDecrypt(desCbcEncrypt, []byte(key))
	assert.IsNil(as, err)
	assert.Equal(as, data, string(desCbcDecrypt))
}

func TestDesCtrCrypt(t *testing.T) {
	t.Parallel()

	data := "hello world"
	key := "abcdefgh"
	as := assert.NewAssert(t, "TestDesCtrCrypt")
	desCtrCrypt, err := DesCtrCrypt([]byte(data), []byte(key))
	assert.IsNil(as, err)
	desCtrDeCrypt, err := DesCtrCrypt(desCtrCrypt, []byte(key))
	assert.IsNil(as, err)
	assert.Equal(as, data, string(desCtrDeCrypt))
}

func TestDesCfbEncrypt(t *testing.T) {
	t.Parallel()

	data := "hello world"
	key := "abcdefgh"
	as := assert.NewAssert(t, "TestDesCfbEncrypt")
	desCfbEncrypt, err := DesCfbEncrypt([]byte(data), []byte(key))
	assert.IsNil(as, err)
	desCfbDecrypt, err := DesCfbDecrypt(desCfbEncrypt, []byte(key))
	assert.IsNil(as, err)
	assert.Equal(as, data, string(desCfbDecrypt))
}

func TestDesOfbEncrypt(t *testing.T) {
	t.Parallel()

	data := "hello world"
	key := "abcdefgh"
	as := assert.NewAssert(t, "TestDesOfbEncrypt")
	desOfbEncrypt, err := DesOfbEncrypt([]byte(data), []byte(key))
	assert.IsNil(as, err)
	desOfbDecrypt, err := DesOfbDecrypt(desOfbEncrypt, []byte(key))
	assert.IsNil(as, err)

	assert.Equal(as, data, string(desOfbDecrypt))
}

func TestRsaEncrypt(t *testing.T) {
	t.Parallel()
	as := assert.NewAssert(t, "TestRsaEncrypt")
	err := GenerateRsaKey(4096, "rsa_private.pem", "rsa_public.pem")
	if err != nil {
		t.FailNow()
	}
	defer func() {
		os.Remove("rsa_private.pem")
		os.Remove("rsa_public.pem")
	}()
	data := []byte("hello world")
	encrypted, err := RsaEncrypt(data, "rsa_public.pem")

	decrypted, err := RsaDecrypt(encrypted, "rsa_private.pem")
	assert.IsNil(as, err)
	assert.Equal(as, string(data), string(decrypted))
}

func TestRsaEncryptOAEP(t *testing.T) {
	as := assert.NewAssert(t, "TestRsaEncrypt")
	t.Parallel()

	pri, pub := GenerateRsaKeyPair(1024)

	data := []byte("hello world")
	label := []byte("123456")

	encrypted, err := RsaEncryptOAEP(data, label, *pub)
	assert.IsNil(as, err)

	decrypted, err := RsaDecryptOAEP([]byte(encrypted), label, *pri)

	assert.IsNil(as, err)
	assert.Equal(as, "hello world", string(decrypted))
}
