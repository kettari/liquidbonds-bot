package console

import (
	prov "github.com/kettari/liquidbonds-bot/internal/provider"
	"log/slog"
)

type FetchBondsCommand struct {
}

func NewFetchCommand() *FetchBondsCommand {
	cmd := FetchBondsCommand{}
	return &cmd
}

func (cmd *FetchBondsCommand) Name() string {
	return "fetch:bonds"
}

func (cmd *FetchBondsCommand) Description() string {
	return "fetches bonds from the MOEX API"
}

func (cmd *FetchBondsCommand) Run() error {
	slog.SetLogLoggerLevel(slog.LevelDebug)
	
	provider, err := prov.NewProvider("moex")
	if err != nil {
		return err
	}
	if err = provider.Fetch(); err != nil {
		return err
	}

	return nil
}
