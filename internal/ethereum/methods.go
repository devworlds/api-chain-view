package ethereum

import (
	"crypto/ecdsa"
	"encoding/hex"
	"fmt"

	"github.com/ethereum/go-ethereum/crypto"
)

// GenerateEvmWallet creates a new Ethereum wallet and returns the address and store private key in memory.
func GenerateWallet() (Wallet, error) {
	// Generate a new private key
	privateKey, err := crypto.GenerateKey()
	if err != nil {
		return Wallet{}, fmt.Errorf("failed to generate private key: %v", err)
	}

	// Extract the public key from the private key
	privateKeyBytes := crypto.FromECDSA(privateKey)
	privateKeyHex := hex.EncodeToString(privateKeyBytes)

	// Generate the public key from the private key
	publicKeyECDSA := privateKey.Public().(*ecdsa.PublicKey)
	publicKeyBytes := crypto.FromECDSAPub(publicKeyECDSA)
	publicKeyHex := hex.EncodeToString(publicKeyBytes)

	// Generate the address from the public key
	address := crypto.PubkeyToAddress(*publicKeyECDSA).Hex()

	wallet := Wallet{
		Address:    address,
		PrivateKey: privateKeyHex,
		PublicKey:  publicKeyHex,
	}

	return wallet, nil
}
