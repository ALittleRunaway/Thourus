package gateway

import (
	"context"
	"crypto/aes"
	"crypto/cipher"
	"encoding/base64"
	"google.golang.org/grpc"
	"google.golang.org/grpc/grpclog"
	pb "thourus-api/infrastructure/grpc/proto"
)

type CryptoGw interface {
	EncryptPassword(text, MySecret string) (string, error)
	DecryptPassword(text, MySecret string) (string, error)
	GetDocumentPoW(docBytes []byte) (int64, error)
	ValidateDocument(docBytes []byte, hash string, PoW int64) (bool, error)
}

type CryptoGateway struct {
	secretString string
	cryptoBytes  []byte
	miningRule   int64
	grpcConn     *grpc.ClientConn
}

func NewCryptoGateway(secretString string, miningRule int64, grpcConn *grpc.ClientConn) *CryptoGateway {
	return &CryptoGateway{
		secretString: secretString,
		cryptoBytes:  []byte{35, 46, 57, 24, 85, 35, 24, 74, 87, 35, 88, 98, 66, 32, 14, 05},
		miningRule:   miningRule,
		grpcConn:     grpcConn,
	}
}

// EncryptPassword method is to encrypt or hide any classified text
func (gw *CryptoGateway) EncryptPassword(text, MySecret string) (string, error) {
	block, err := aes.NewCipher([]byte(MySecret))
	if err != nil {
		return "", err
	}
	plainText := []byte(text)
	cfb := cipher.NewCFBEncrypter(block, gw.cryptoBytes)
	cipherText := make([]byte, len(plainText))
	cfb.XORKeyStream(cipherText, plainText)

	return base64.StdEncoding.EncodeToString(cipherText), nil
}

// DecryptPassword method is to extract back the encrypted text
func (gw *CryptoGateway) DecryptPassword(text, MySecret string) (string, error) {
	block, err := aes.NewCipher([]byte(MySecret))
	if err != nil {
		return "", err
	}
	cipherText, err := base64.StdEncoding.DecodeString(text)
	if err != nil {
		return "", err
	}
	cfb := cipher.NewCFBDecrypter(block, gw.cryptoBytes)
	plainText := make([]byte, len(cipherText))
	cfb.XORKeyStream(plainText, cipherText)

	return string(plainText), nil
}

func (gw *CryptoGateway) GetDocumentPoW(docBytes []byte) (int64, error) {
	client := pb.NewCryptoCoreClient(gw.grpcConn)

	request := &pb.MineRequest{
		Bytes: docBytes,
		Rule:  gw.miningRule,
	}
	response, err := client.Mine(context.Background(), request)
	if err != nil {
		grpclog.Fatalf("fail to dial: %v", err)
	}
	return response.GetPow(), nil
}

func (gw *CryptoGateway) ValidateDocument(docBytes []byte, hash string, PoW int64) (bool, error) {
	client := pb.NewCryptoCoreClient(gw.grpcConn)

	request := &pb.ValidateRequest{
		Bytes: docBytes,
		Hash:  hash,
		Pow:   PoW,
	}
	response, err := client.Validate(context.Background(), request)
	if err != nil {
		grpclog.Fatalf("fail to dial: %v", err)
	}
	return response.GetValid(), nil
}
