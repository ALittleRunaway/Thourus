package usecase

//
//import (
//	gobytes "bytes"
//	"strconv"
//	"strings"
//	"thourus-api/gateway"
//)
//
//type cryptoUseCase struct {
//	cryptoGw gateway.CryptoGw
//	// TODO: add logger
//}
//
//func NewCryptoUseCase(cryptoGw gateway.CryptoGw) *cryptoUseCase {
//	return &cryptoUseCase{
//		cryptoGw: cryptoGw,
//	}
//}
//
//func (uc *cryptoUseCase) Mine(bytes []byte, rule int64) (int64, error) {
//
//	var PoW int64
//	var hash string
//
//	// TODO: add a bunch of goroutines to speed it up
//	for {
//		PoW += 1
//		finalString := gobytes.Join([][]byte{bytes, []byte(strconv.FormatInt(PoW, 10))}, []byte("-"))
//		hash = uc.cryptoGw.GenerateHash(finalString)
//		if strings.HasPrefix(hash, strconv.FormatInt(rule, 10)) {
//			return PoW, nil
//		}
//	}
//}
//
//func (uc *cryptoUseCase) Validate(bytes []byte, hash string, PoW int64) (bool, error) {
//
//	// TODO: add errors (but only declarations, not the real ones)
//
//	finalString := gobytes.Join([][]byte{bytes, []byte(strconv.FormatInt(PoW, 10))}, []byte("-"))
//	if uc.cryptoGw.GenerateHash(finalString) != hash {
//		return false, nil
//	}
//	return true, nil
//
//}
