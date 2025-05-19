package moex

import "encoding/json"

type rawSecuritiesList struct {
	Securities struct {
		Columns []string      `json:"columns"`
		Data    []rawSecurity `json:"data"`
	} `json:"securities"`
	MarketData struct {
		Columns []string    `json:"columns"`
		Data    []rawMarket `json:"data"`
	} `json:"marketdata"`
}

type SecuritiesList struct {
	Securities map[string]Security
}

type rawSecurity struct {
	ID    string  // SECID
	Name  string  // SECNAME
	Price float64 // PREVLEGALCLOSEPRICE
}

type rawMarket struct {
	ID       string  // SECID
	Yield    float64 // YIELD
	Duration int     // DURATION
}

type Security struct {
	ID       string
	Name     string
	Price    float64
	Yield    float64
	Duration int
}

func (s *rawSecurity) UnmarshalJSON(p []byte) error {
	var tmp []json.RawMessage
	if err := json.Unmarshal(p, &tmp); err != nil {
		return err
	}
	if err := json.Unmarshal(tmp[0], &s.ID); err != nil {
		return err
	}
	if err := json.Unmarshal(tmp[1], &s.Name); err != nil {
		return err
	}
	if err := json.Unmarshal(tmp[2], &s.Price); err != nil {
		return err
	}

	return nil
}

func (m *rawMarket) UnmarshalJSON(p []byte) error {
	var tmp []json.RawMessage
	if err := json.Unmarshal(p, &tmp); err != nil {
		return err
	}
	if err := json.Unmarshal(tmp[0], &m.ID); err != nil {
		return err
	}
	if err := json.Unmarshal(tmp[1], &m.Yield); err != nil {
		return err
	}
	if err := json.Unmarshal(tmp[2], &m.Duration); err != nil {
		return err
	}

	return nil
}
