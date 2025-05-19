package moex

import (
	"encoding/json"
	"fmt"
	"io"
	"log/slog"
	"net/http"
)

const (
	bondsURL           = `https://iss.moex.com/iss/engines/stock/markets/bonds/boardgroups/%d/securities.json?iss.dp=comma&iss.meta=off&iss.only=securities,marketdata&securities.columns=SECID,SECNAME,PREVLEGALCLOSEPRICE&marketdata.columns=SECID,YIELD,DURATION`
	stockBonds         = "stock_bonds"           // Т0: Основной режим - безадрес.
	stockBondsTPlus    = "stock_bonds_tplus"     // Т+: Основной режим - безадрес.
	stockBondsTPlusUSD = "stock_bonds_tplus_usd" // Т+: Основной режим (USD) - безадрес.
)

var (
	boardGroups = map[string]int{
		stockBonds:         7,
		stockBondsTPlus:    58,
		stockBondsTPlusUSD: 193,
	}
)

type Moex struct {
	Bonds SecuritiesList
}

func NewMoex() *Moex {
	return &Moex{
		Bonds: SecuritiesList{
			Securities: make(map[string]Security),
		},
	}
}

func (p *Moex) Fetch() error {
	for name, group := range boardGroups {
		var bonds rawSecuritiesList
		slog.Debug("fetching bonds for the group", "name", name, "index", group)
		data, err := fetch(fmt.Sprintf(bondsURL, group))
		if err != nil {
			return err
		}
		slog.Debug("fetched bonds", "size_bytes", len(*data))
		if err = json.Unmarshal([]byte(*data), &bonds); err != nil {
			return err
		}
		slog.Debug("unmarshalled bonds", "length", len(bonds.Securities.Data))

		// Transform securities array to securities map
		for _, b := range bonds.Securities.Data {
			p.Bonds.Securities[b.ID] = Security{
				ID:    b.ID,
				Name:  b.Name,
				Price: b.Price,
			}
		}
		for _, b := range bonds.MarketData.Data {
			if entry, ok := p.Bonds.Securities[b.ID]; ok {
				entry.Yield = b.Yield
				entry.Duration = b.Duration

				p.Bonds.Securities[b.ID] = entry
			}
		}
	}

	return nil
}

func fetch(url string) (*string, error) {
	req, err := http.NewRequest("GET", url, nil)
	if err != nil {
		return nil, err
	}

	resp, err := http.DefaultClient.Do(req)
	if err != nil {
		return nil, err
	}
	defer func(Body io.ReadCloser) {
		err = Body.Close()
		if err != nil {
			slog.Error("failed to close response body", "url", url, "err", err)
		}
	}(resp.Body)

	data, err := io.ReadAll(resp.Body)
	if err != nil {
		return nil, err
	}
	result := string(data)

	return &result, nil
}
