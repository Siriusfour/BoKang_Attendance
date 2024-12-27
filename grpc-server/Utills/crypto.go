package Utills

import (
	"bytes"
	"crypto/aes"
	"crypto/cipher"
	"crypto/hmac"
	"crypto/rand"
	"crypto/sha256"
	"errors"
	"fmt"
)

type AesKeyType int32

const (
	AES128 AesKeyType = iota
	AES192
	AES256
)

// NewAesKey create a aes key according to the give AesKeyType,
// Supported:
//   - AES128
//   - AES192
//   - AES256
func NewAesKey(t AesKeyType) ([]byte, error) {
	var bLen int32
	switch t {
	case AES128:
		bLen = 16
	case AES192:
		bLen = 24
	case AES256:
		bLen = 32
	default:
		return nil, errors.New("unsupported AES key length")
	}
	b := make([]byte, bLen)
	_, err := rand.Read(b)
	if err != nil {
		return nil, err
	}
	return b, nil
}

// Encrypt takes the key and used it to encrypt the given plain text.
// The result is actually:
// IV + HMAC + CipherText
func Encrypt(key []byte, plainText []byte) ([]byte, error) {
	block, err := aes.NewCipher(key)
	if err != nil {
		return nil, err
	}
	blockSize := block.BlockSize()

	plainText, err = PKCS7Pad(plainText, blockSize)
	if err != nil {
		return nil, err
	}

	// The IV needs to be unique, but not secure. Therefore it's common to
	// include it at the beginning of the ciphertext.
	cipherText := make([]byte, blockSize+sha256.Size+len(plainText))
	iv := cipherText[:blockSize]
	mac := cipherText[blockSize : blockSize+sha256.Size]
	payload := cipherText[blockSize+sha256.Size:]
	if _, err = rand.Read(iv); err != nil {
		return nil, err
	}

	mode := cipher.NewCBCEncrypter(block, iv)
	mode.CryptBlocks(payload, plainText)

	// we use Encrypt-then-MAC
	// https://crypto.stackexchange.com/questions/202/should-we-mac-then-encrypt-or-encrypt-then-mac
	hash := hmac.New(sha256.New, key)
	hash.Write(payload)
	copy(mac, hash.Sum(nil))

	return cipherText, nil
}

// Decrypt takes a key and use it to decrypt the given cipherText, finally returns the plain text.
func Decrypt(key []byte, cipherText []byte) ([]byte, error) {
	block, err := aes.NewCipher(key)
	if err != nil {
		return nil, err
	}
	blockSize := block.BlockSize()

	if len(cipherText) <= blockSize+sha256.Size {
		return nil, errors.New("ciphertext too short")
	}

	iv := cipherText[:blockSize]
	mac := cipherText[blockSize : blockSize+sha256.Size]
	cipherText = cipherText[blockSize+sha256.Size:]

	if len(cipherText)%blockSize != 0 {
		return nil, errors.New("ciphertext is not block-aligned, maybe corrupted")
	}

	hash := hmac.New(sha256.New, key)
	hash.Write(cipherText)
	if !hmac.Equal(hash.Sum(nil), mac) {
		return nil, errors.New("hmac failure, message corrupted")
	}

	plainText := make([]byte, len(cipherText))
	mode := cipher.NewCBCDecrypter(block, iv)
	mode.CryptBlocks(plainText, cipherText)

	plainText, err = PKCS7UnPad(plainText, blockSize)
	if err != nil {
		return nil, err
	}
	return plainText, nil
}

// PKCS7Pad applies PKCS#7Padding in the given byte array.
// Padding is necessary to ensure that plain text can be divisible by the block size.
//
// https://stackoverflow.com/a/13572751/11381693
func PKCS7Pad(data []byte, blockSize int) ([]byte, error) {
	if blockSize < 1 || blockSize >= 256 {
		return nil, fmt.Errorf("invalid block size: %d", blockSize)
	}

	// according to https://www.rfc-editor.org/rfc/rfc2315:
	//
	//			2.   Some content-encryption algorithms assume the
	//			input length is a multiple of k octets, where k > 1, and
	//			let the application define a method for handling inputs
	//			whose lengths are not a multiple of k octets. For such
	//			algorithms, the method shall be to pad the input at the
	//			trailing end with k - (l mod k) octets all having value k -
	//			(l mod k), where l is the length of the input. In other
	//			words, the input is padded at the trailing end with one of
	//			the following strings:
	//
	//			01 -- if l mod k = k-1
	//			02 02 -- if l mod k = k-2
	//			.
	//			.
	//			.
	//			k k ... k k -- if l mod k = 0
	//
	//			The padding can be removed unambiguously since all input is
	//			padded and no padding string is a suffix of another. This
	//			padding method is well-defined if and only if k < 256;
	//			methods for larger k are an open issue for further study.
	//

	// calculate the padding length, ranging from 1 to blockSize
	paddingLen := blockSize - len(data)%blockSize

	// build the padding text
	padding := bytes.Repeat([]byte{byte(paddingLen)}, paddingLen)
	return append(data, padding...), nil
}

// PKCS7UnPad removes PKCS#7 padding from the given byte array.
func PKCS7UnPad(data []byte, blockSize int) ([]byte, error) {
	length := len(data)
	if length == 0 { // empty
		return nil, errors.New("unpad called on zero length byte array")
	}
	if length%blockSize != 0 {
		return nil, errors.New("data is not block-aligned")
	}

	// just the number that the last byte represents
	paddingLen := int(data[length-1])
	padding := bytes.Repeat([]byte{byte(paddingLen)}, paddingLen)
	if paddingLen > blockSize || paddingLen == 0 || !bytes.HasSuffix(data, padding) {
		return nil, errors.New("invalid padding")
	}
	return data[:length-paddingLen], nil
}
