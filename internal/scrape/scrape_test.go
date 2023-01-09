package scrape_test

import (
	"encoding/json"
	"fmt"
	"getdl/internal/scrape"
	"os"
	"testing"

	"github.com/stretchr/testify/require"
)

func writeToJson(data scrape.ScrapeResponse, website string) error {
	result, err := json.Marshal(data)
	if err != nil {
		return err
	}

	return os.WriteFile(fmt.Sprintf("result-%s.json", website), result, 0666)
}

func TestScrape(t *testing.T) {
	tests := []struct {
		name   string
		link   string
		scrape func(link string) scrape.ScrapeResponse
	}{
		{
			name:   "doronime",
			link:   "https://doronime.id/anime/mieruko-chan/episode-12",
			scrape: scrape.Doronime,
		},
		{
			name:   "lendrive",
			link:   "https://lendrive.web.id/akebi-chan-no-sailor-fuku-ep-01-dual-subs-x265-hevc-subtitle-indonesia-english/",
			scrape: scrape.Lendrive,
		},
		{
			name:   "oploverz",
			link:   "https://oploverz.co.in/danmachi-s4-episode-12-subtitle-indonesia/?utm_source=rss&utm_medium=rss&utm_campaign=danmachi-s4-episode-12-subtitle-indonesia",
			scrape: scrape.Oploverz,
		},
		{
			name:   "samehadaku",
			link:   "https://194.163.183.129/saihate-no-paladin-episode-10/",
			scrape: scrape.Samehadaku,
		},
		{
			name:   "animekompi",
			link:   "https://animekompi.net/hyouken-no-majutsushi-ga-sekai-wo-suberu-episode-01-subtitle-indonesia/",
			scrape: scrape.Animekompi,
		},
	}

	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			result := test.scrape(test.link)
			require.NotEmpty(t, result.Downloads)
			writeToJson(result, test.name)
		})

	}

}
