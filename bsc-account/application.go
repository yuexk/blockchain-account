package main

import (
	"fmt"
	"github.com/ethereum/go-ethereum/common/hexutil"
	"github.com/ethereum/go-ethereum/crypto"
)

// 生成地址和私钥
func genAccount() (string, string, error) {
	privateKey, err := crypto.GenerateKey()
	if err != nil {
		return "", "", err
	}

	address := crypto.PubkeyToAddress(privateKey.PublicKey)
	fmt.Println(address.Hex())

	sk := crypto.FromECDSA(privateKey)
	hexSk := hexutil.Encode(sk)

	fmt.Println(hexSk)


	return address.Hex(), "", nil
}

func main() {
	genAccount()
}
