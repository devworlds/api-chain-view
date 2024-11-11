package ethereum

import (
	"os"

	"github.com/ethereum/go-ethereum/ethclient"
)

func newClient() (*ethclient.Client, error) {
	client, err := ethclient.Dial(os.Getenv("NODE_URL"))
	if err != nil {
		return nil, err
	}
	return client, nil
}
