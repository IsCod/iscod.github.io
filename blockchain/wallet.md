# 钱包

### BTC钱包

#### Bech32

#### Base58

1. 创建私钥（ecdsa椭圆算法）
2. 创建公钥
3. 生成公钥hash160: 将公钥进行一次 256 hash, 再进行一次 160 hash
4. 生成 checksum, sum = sha256(sha256(version + hash160))[:4] // 两次hash256(version + hash160),取出前四个字节
5. 生成地址：base58.Encode(version + hash160 + checksum)

```go
package main

import (
	"crypto/ecdsa"
	"crypto/elliptic"
	"crypto/rand"
	"crypto/sha256"
	"fmt"
	"github.com/btcsuite/btcutil/base58"
	"golang.org/x/crypto/ripemd160"
)

const VERSION = byte(0x00)
const CHECKBYTELEN = 4

type Wallet struct {
	privateKey *ecdsa.PrivateKey
	publicKey  []byte // privateKey.x, privateKey.y
}

func NewWallet() (*Wallet, error) {
	privateKey, err := ecdsa.GenerateKey(elliptic.P256(), rand.Reader)
	if err != nil {
		return nil, err
	}

	publicKey := append(privateKey.X.Bytes(), privateKey.Y.Bytes()...)

	return &Wallet{
		privateKey: privateKey,
		publicKey:  publicKey,
	}, nil
}

func (w *Wallet) GetPublicKeyHash160() []byte {
	hash256 := sha256.Sum256(w.publicKey)
	r := ripemd160.New()
	r.Write(hash256[:])
	return r.Sum(nil)
}

func (w *Wallet) GetCheckSum() []byte {
	hash160 := w.GetPublicKeyHash160()
	sum := append([]byte{VERSION}, hash160...)
	sum1 := sha256.Sum256(sum)
	sum2 := sha256.Sum256(sum1[:])
	return sum2[:CHECKBYTELEN]
}

func (w *Wallet) GetAddress() string {
	sum := append([]byte{VERSION}, w.GetPublicKeyHash160()...)
	sum = append(sum, w.GetCheckSum()...)
	return base58.Encode(sum)
}

func main() {
	w, _ := NewWallet()
	fmt.Println(w.GetAddress())
}
```

### 区块

### 创世区块

#### 配置选项

#### 从磁盘读取

### 数字签名

#### 哈希
#### 冲压
#### 签名
#### 寻址
#### 验证

* 参考
    * [JSON](https://developer.bitcoin.org/reference/wallets.html?highlight=wallet)