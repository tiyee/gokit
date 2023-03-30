package helps

import (
	"bytes"
	"crypto/cipher"
	"crypto/des"
	"github.com/tiyee/gokit/pkg/consts"
)

const key = consts.RsaKey
const iv = consts.RsaIV

func DecryptByDESAndCBC(cipherText []byte) ([]byte, error) {

	textBytes := cipherText
	keyBytes := []byte(key)
	ivBytes := []byte(iv)

	//1. 创建并返回一个使用DES算法的cipher.Block接口。
	block, err := des.NewCipher(keyBytes)
	if err != nil {
		return nil, err
	}

	//2. 创建CBC分组模式
	blockMode := cipher.NewCBCDecrypter(block, ivBytes)

	//3. 解密
	out := make([]byte, len(cipherText))
	blockMode.CryptBlocks(out, textBytes)

	//4. 去掉填充数据 (注意去掉填充的顺序是在解密之后)
	out = PKCS5UnPadding(out)

	return out, nil
}

func EncryptByDESAndCBC(textBytes []byte) ([]byte, error) {

	keyBytes := []byte(key)
	ivBytes := []byte(iv)

	//加密
	//1. 创建并返回一个使用DES算法的cipher.Block接口
	//使用des调用NewCipher获取block接口
	block, err := des.NewCipher(keyBytes)
	if err != nil {
		return nil, err
	}

	//2. 填充数据，将输入的明文构造成8的倍数
	textBytes = PKCS5Padding(textBytes, block.BlockSize())

	//3. 创建CBC分组模式,返回一个密码分组链接模式的、底层用b加密的BlockMode接口，初始向量iv的长度必须等于b的块尺寸
	//使用cipher调用NewCBCDecrypter获取blockMode接口
	blockMode := cipher.NewCBCEncrypter(block, ivBytes)

	//3. 加密
	//这里的两个参数为什么都是textBytes
	//第一个是目标,第二个是源
	//也就是说将第二个进行加密,然后放到第一个里面
	//如果我们重新定义一个密文cipherTextBytes
	//那么就是blockMode.CryptBlocks(cipherTextBytes, textBytes)
	out := make([]byte, len(textBytes))
	blockMode.CryptBlocks(out, textBytes)
	return out, nil

}

func PKCS5Padding(ciphertext []byte, blockSize int) []byte {
	padding := blockSize - len(ciphertext)%blockSize
	padText := bytes.Repeat([]byte{byte(padding)}, padding)
	return append(ciphertext, padText...)
}
func PKCS5UnPadding(origData []byte) []byte {
	length := len(origData)
	unPadding := int(origData[length-1])
	return origData[:(length - unPadding)]
}
