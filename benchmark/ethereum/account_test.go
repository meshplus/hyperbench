package ethereum

import (
	"crypto/ecdsa"
	"encoding/hex"
	"io"
	"os"
	"testing"

	"github.com/ethereum/go-ethereum/crypto"
	"github.com/stretchr/testify/assert"
)

var keys = []string{
	"0xac0974bec39a17e36ba4a6b4d238ff944bacb478cbed5efcae784d7bf4f2ff80",
	"0x59c6995e998f97a5a0044966f0945389dc9e86dae88c7a8412f4603b6b78690d",
	"0x5de4111afa1a4b94908f83103eb1f1706367c2e68ca870fc3fb9a804cdab365a",
	"0x7c852118294e51e653712a81e05800f419141751be58f605c371e15141b007a6",
	"0x47e179ec197488593b187f80a00eb0da91f1b9d0b13f8733639f19c30a34926a",
	"0x8b3a350cf5c34c9194ca85829a2df0ec3153be0318b5e2d3348e872092edffba",
	"0x92db14e403b83dfe3df233f83dfa3a0d7096f21ca9b0d6d6b8d88b2b4ec1564e",
	"0x4bbbf85ce3377467afe5d46f804f221813b2bb87f24d81f60f1fcdbf7cbf4356",
	"0xdbda1821b80551c9d65939329250298aa3472ba22feea921c0cf5d620ea67b97",
	"0x2a871d0798f97d79848a013d4936a73bf4cc922c825d33c1cf7073dff6d409c6",
}

var keyFilePaths = []string{
	"transfer/eth/keystore/keys",
	"invoke/eth/keystore/keys",
}

const TotalAccount = 2000000

// func TestAccount(t *testing.T) {
// 	for _, key := range keys {
// 		priKey := &keystore.Key{}
// 		id, err := uuid.NewRandom()
// 		assert.Nil(t, err)
// 		priKey.Id = id
// 		sk, err := crypto.HexToECDSA(strings.TrimPrefix(key, "0x"))
// 		assert.Nil(t, err)
// 		priKey.PrivateKey = sk
// 		priKey.Address = crypto.PubkeyToAddress(sk.PublicKey)

// 		t.Logf("key is: %+v", priKey)

// 		encryptKey, err := keystore.EncryptKey(priKey, "", 262144, 1)
// 		assert.Nil(t, err)

// 		for _, keyFilePath := range keyFilePaths {
// 			filename := fmt.Sprintf("%s--%s", time.Now().Format("UTC--2006-01-02T15-04-05.999999999Z"), strings.TrimPrefix(priKey.Address.Hex(), "0x"))
// 			err = os.WriteFile(filepath.Join(keyFilePath, filename), encryptKey, os.ModePerm)
// 			assert.Nil(t, err)
// 		}
// 	}
// }

func TestGenerateAccount(t *testing.T) {
	var dstFile *os.File
	var err error
	for i, keyFilePath := range keyFilePaths {
		if i == 0 {
			dstFile, err = os.OpenFile(keyFilePath, os.O_RDWR|os.O_CREATE|os.O_APPEND, 0666)
			assert.Nil(t, err)

			var sk *ecdsa.PrivateKey
			for i := 0; i < TotalAccount; i++ {
				sk, err = crypto.GenerateKey()
				assert.Nil(t, err)

				//t.Logf("key is: %+v", *sk)
				privKey := hex.EncodeToString(crypto.FromECDSA(sk))
				//t.Logf("priv key: %s", privKey)

				_, err = dstFile.Write([]byte(privKey))
				assert.Nil(t, err)
				_, err = dstFile.Write([]byte("\n"))
				assert.Nil(t, err)
			}
			err = dstFile.Close()
			assert.Nil(t, err)
		} else {
			f, err := os.OpenFile(keyFilePath, os.O_RDWR|os.O_CREATE|os.O_APPEND, 0666)
			assert.Nil(t, err)
			_, err = io.Copy(f, dstFile)
			assert.Nil(t, err)
			err = f.Close()
			assert.Nil(t, err)
		}
	}
}
