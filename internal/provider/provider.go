package provider

import (
	"fmt"
	"github.com/kettari/liquidbonds-bot/internal/provider/moex"
)

type Provider interface {
	Fetch() error
}

func NewProvider(provider string) (Provider, error) {
	switch provider {
	case "moex":
		return moex.NewMoex(), nil
	}
	return nil, fmt.Errorf("unknown provider: %s", provider)
}
