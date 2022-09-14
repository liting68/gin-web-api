package utils

import (
	"bytes"
	"crypto/aes"
	"crypto/cipher"
	"crypto/md5"
	"fmt"
)

//AesEncryptCBC CBC加密
func AesEncryptCBC(origData []byte) (encrypted []byte) {
	str := []byte("RCS-Aes-key")
	has := md5.Sum(str)
	key := []byte(fmt.Sprintf("%x", has))
	block, _ := aes.NewCipher(key)                              // NewCipher该函数限制了输入k的长度必须为16, 24或者32
	blockSize := block.BlockSize()                              // 获取秘钥块的长度
	origData = pkcs5Padding(origData, blockSize)                // 补全码
	blockMode := cipher.NewCBCEncrypter(block, key[:blockSize]) // 加密模式
	encrypted = make([]byte, len(origData))                     // 创建数组
	blockMode.CryptBlocks(encrypted, origData)                  // 加密
	return encrypted
}

//AesDecryptCBC CBC解密
func AesDecryptCBC(encrypted []byte) (decrypted []byte, err error) {
	str := []byte("RCS-Aes-key")
	has := md5.Sum(str)
	key := []byte(fmt.Sprintf("%x", has))
	block, _ := aes.NewCipher(key)                              // 分组秘钥
	blockSize := block.BlockSize()                              // 获取秘钥块的长度
	blockMode := cipher.NewCBCDecrypter(block, key[:blockSize]) // 加密模式
	decrypted = make([]byte, len(encrypted))                    // 创建数组
	if len(decrypted) > 16 {
		return []byte(""), fmt.Errorf("decrypted too long")
	}
	blockMode.CryptBlocks(decrypted, encrypted) // 解密
	decrypted, err = pkcs5UnPadding(decrypted)  // 去除补全码
	if err != nil {
		return []byte(""), err
	}
	return decrypted, nil
}
func pkcs5Padding(ciphertext []byte, blockSize int) []byte {
	padding := blockSize - len(ciphertext)%blockSize
	padtext := bytes.Repeat([]byte{byte(padding)}, padding)
	return append(ciphertext, padtext...)
}
func pkcs5UnPadding(origData []byte) ([]byte, error) {
	length := len(origData)
	unpadding := int(origData[length-1])
	if length < unpadding {
		return []byte(""), fmt.Errorf("origData length too short")
	}
	return origData[:(length - unpadding)], nil
}
