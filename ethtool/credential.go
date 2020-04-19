package ethtool

import (
	"crypto/ecdsa"
	"fmt"
	"math/big"

	"github.com/ethereum/go-ethereum/accounts/abi/bind"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/common/hexutil"
	"github.com/ethereum/go-ethereum/core/types"
	"github.com/ethereum/go-ethereum/crypto"
)

type Credential struct {
	PrivateKey *ecdsa.PrivateKey
	PublicKey  *ecdsa.PublicKey
	Address    common.Address
}

func KeyToCredential(privateKey *ecdsa.PrivateKey) (*Credential, error) {
	publicKey, ok := privateKey.Public().(*ecdsa.PublicKey)
	if !ok {
		return nil, fmt.Errorf("Public key cast error.")
	}
	address := crypto.PubkeyToAddress(*publicKey)

	credential := &Credential{
		PrivateKey: privateKey,
		PublicKey:  publicKey,
		Address:    address,
	}
	return credential, nil
}

func HexToCredential(hexkey string) (*Credential, error) {
	privateKey, err := crypto.HexToECDSA(hexkey[2:])
	if err != nil {
		return nil, err
	}
	return KeyToCredential(privateKey)
}

func NewCredential() (*Credential, error) {
	privateKey, err := crypto.GenerateKey()
	if err != nil {
		return nil, err
	}
	return KeyToCredential(privateKey)
}

func (c *Credential) PrivateKeyHex() string {
	return hexutil.Encode(crypto.FromECDSA(c.PrivateKey))
}

func (c *Credential) PublicKeyHex() string {
	return hexutil.Encode(crypto.FromECDSAPub(c.PublicKey))
}

func (c *Credential) AddressHex() string {
	return c.Address.Hex()
}

func (c *Credential) Sign(data []byte) ([]byte, error) {
	hash := crypto.Keccak256(data)
	return crypto.Sign(hash, c.PrivateKey)
}

func (c *Credential) SignTx(tx *types.Transaction, chainid *big.Int) (*types.Transaction, error) {
	signer := types.NewEIP155Signer(chainid)
	return types.SignTx(tx, signer, c.PrivateKey)
}

func (c *Credential) Verify(data, signature []byte) bool {
	hash := crypto.Keccak256(data)
	pubkey := crypto.CompressPubkey(c.PublicKey)
	return crypto.VerifySignature(pubkey, hash, signature)
}

func (c *Credential) GetTransactOpts() *bind.TransactOpts {
	return bind.NewKeyedTransactor(c.PrivateKey)
}
