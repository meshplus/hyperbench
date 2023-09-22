package ethereum

import (
	"bufio"
	"crypto/ecdsa"
	"encoding/hex"
	"io"
	"os"
	"strings"
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

func TestAccount(t *testing.T) {
	for _, key := range keys {
		sk, err := crypto.HexToECDSA(strings.TrimPrefix(key, "0x"))
		assert.Nil(t, err)
		addr := crypto.PubkeyToAddress(sk.PublicKey)

		t.Logf("address is: %s", addr.String())
	}
}

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

// 将百万账户三等分
func TestSplitAndSaveKeys(t *testing.T) {
	// 打开现有的 keys 文件
	srcFile, err := os.Open("./keys")
	if err != nil {
		panic(err)
	}
	defer srcFile.Close()

	// 使用 scanner 逐行读取文件
	var keys []string
	scanner := bufio.NewScanner(srcFile)
	for scanner.Scan() {
		line := scanner.Text()
		keys = append(keys, line)
	}

	if err := scanner.Err(); err != nil {
		panic(err)
	}

	// 分割 keys 切片
	totalKeys := len(keys)
	partSize := totalKeys / 3

	part1 := keys[:partSize]
	part2 := keys[partSize : 2*partSize]
	part3 := keys[2*partSize:]

	// 定义要保存到的三个文件路径
	paths := []string{
		"stability-erc20/eth/keystore/keys",
		"stability-transfer/eth/keystore/keys",
		"stability-uniswap/eth/keystore/keys"}

	// 分别保存到三个文件中
	for i, part := range [][]string{part1, part2, part3} {
		file, err := os.Create(paths[i])
		if err != nil {
			panic(err)
		}
		defer file.Close()

		for _, key := range part {
			_, err := file.WriteString(key + "\n")
			if err != nil {
				panic(err)
			}
		}
	}
}
