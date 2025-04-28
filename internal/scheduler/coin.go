package scheduler

import (
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"sort"
	"strings"

	"github.com/bestkkii/saedori-api-server/internal/model"
)

const (
	UPBIT_URL        = "https://api.upbit.com/v1/market/all"
	UPBIT_TICKER_URL = "https://api.upbit.com/v1/ticker?markets="
)

type CoinScheduler struct {
}

func (c *CoinScheduler) GetCoinChangeRate() ([]string, error) {
	return GetCoinChangeRate()
}

func abs(x float64) float64 {
	if x < 0 {
		return -x
	}
	return x
}

func GetCoinChangeRate() ([]string, error) {
	markets, err := fetchMarkets()
	if err != nil {
		return nil, err
	}

	marketMap, krwMarkets := filterKRWMarkets(markets)
	tickers, err := fetchTickers(krwMarkets)
	if err != nil {
		return nil, err
	}

	changes := calculateChanges(tickers, marketMap)
	return formatChanges(changes), nil
}

func fetchMarkets() ([]model.Market, error) {
	resp, err := http.Get(UPBIT_URL)
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()

	body, err := io.ReadAll(resp.Body)
	if err != nil {
		return nil, err
	}

	var markets []model.Market
	if err := json.Unmarshal(body, &markets); err != nil {
		return nil, err
	}

	return markets, nil
}

func filterKRWMarkets(markets []model.Market) (map[string]string, []string) {
	marketMap := make(map[string]string)
	var krwMarkets []string
	for _, m := range markets {
		if strings.HasPrefix(m.Market, "KRW-") {
			krwMarkets = append(krwMarkets, m.Market)
			marketMap[m.Market] = m.KoreanName
		}
	}
	return marketMap, krwMarkets
}

func fetchTickers(krwMarkets []string) ([]model.Ticker, error) {
	marketsParam := strings.Join(krwMarkets, ",")
	tickerURL := UPBIT_TICKER_URL + marketsParam

	resp, err := http.Get(tickerURL)
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()

	body, err := io.ReadAll(resp.Body)
	if err != nil {
		return nil, err
	}

	var tickers []model.Ticker
	if err := json.Unmarshal(body, &tickers); err != nil {
		return nil, err
	}

	return tickers, nil
}

func calculateChanges(tickers []model.Ticker, marketMap map[string]string) []model.ChangeInfo {
	var changes []model.ChangeInfo
	for _, ticker := range tickers {
		changeRate := ticker.SignedChangeRate * 100
		if changeRate > 3 || changeRate < -3 {
			koreanName := marketMap[ticker.Market]
			symbol := strings.TrimPrefix(ticker.Market, "KRW-")
			changes = append(changes, model.ChangeInfo{
				KoreanName: koreanName,
				Symbol:     symbol,
				ChangeRate: changeRate,
			})
		}
	}

	sort.Slice(changes, func(i, j int) bool {
		return abs(changes[i].ChangeRate) > abs(changes[j].ChangeRate)
	})

	if len(changes) > 3 {
		changes = changes[:3]
	}

	return changes
}

func formatChanges(changes []model.ChangeInfo) []string {
	var formattedChanges []string
	for _, c := range changes {
		changeType := "급등"
		if c.ChangeRate < 0 {
			changeType = "급락"
		}
		formattedChange := fmt.Sprintf("%s %s %s (%.1f%%)", c.KoreanName, c.Symbol, changeType, float64(int(c.ChangeRate*10))/10.0)
		formattedChanges = append(formattedChanges, formattedChange)
	}
	return formattedChanges
}
