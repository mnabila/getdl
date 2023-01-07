package scrape_test

import (
	"encoding/json"
	"fmt"
	"getdl/internal/scrape"
	"os"
	"testing"
)

func writeToJson(data scrape.ScrapeResponse, website string) error {
	result, err := json.Marshal(data)
	if err != nil {
		return err
	}

	return os.WriteFile(fmt.Sprintf("result-%s.json", website), result, 0666)
}

func TestDoronime(t *testing.T) {
	result := scrape.Doronime("https://doronime.id/anime/mieruko-chan/episode-12")
	writeToJson(result, "doronime")
}

func TestLendrive(t *testing.T) {
	result := scrape.Lendrive("https://lendrive.web.id/akebi-chan-no-sailor-fuku-ep-01-dual-subs-x265-hevc-subtitle-indonesia-english/")
	writeToJson(result, "lendrive")
}

func TestOploverz(t *testing.T) {
	result := scrape.Oploverz("https://oploverz.co.in/danmachi-s4-episode-12-subtitle-indonesia/?utm_source=rss&utm_medium=rss&utm_campaign=danmachi-s4-episode-12-subtitle-indonesia")
	writeToJson(result, "oploverz")
}

func TestSamehadaku(t *testing.T) {
	result := scrape.Samehadaku("https://194.163.183.129/saihate-no-paladin-episode-10/")
	writeToJson(result, "samehadaku")
}
