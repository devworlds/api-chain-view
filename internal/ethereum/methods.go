package ethereum

import (
	"crypto/ecdsa"
	"encoding/hex"
	"fmt"

	"github.com/ethereum/go-ethereum/crypto"
)

// GenerateWallet creates a new Ethereum wallet and returns the address, private key, and public key.
func GenerateWallet() (Wallet, error) {

	// Generate a new private key from ECDSA
	privateKey, err := crypto.GenerateKey()
	if err != nil {
		return Wallet{}, fmt.Errorf("failed to generate private key: %v", err)
	}

	// Convert the private key to bytes and hex
	privateKeyBytes := crypto.FromECDSA(privateKey)
	privateKeyHex := hex.EncodeToString(privateKeyBytes)

	// Get the public key from the private key
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
