package getdl_test

import (
	"encoding/json"
	"fmt"
	"os"
	"testing"

	"github.com/mnabila/getdl/entities"
	"github.com/mnabila/getdl/pkg/getdl"
	"github.com/stretchr/testify/require"
)

func writeToJson(data entities.ScrapeResponse, website string) error {
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
		scrape func(link string) entities.ScrapeResponse
	}{
		{
			name:   "doronime",
			link:   "https://doronime.id/anime/mieruko-chan/episode-12",
			scrape: getdl.Doronime,
		},
		{
			name:   "lendrive",
			link:   "https://lendrive.web.id/akebi-chan-no-sailor-fuku-ep-01-dual-subs-x265-hevc-subtitle-indonesia-english/",
			scrape: getdl.Lendrive,
		},
		{
			name:   "oploverz",
			link:   "https://oploverz.red/jujutsu-kaisen-s2-episode-10-subtitle-indonesia/",
			scrape: getdl.Oploverz,
		},
		{
			name:   "samehadaku",
			link:   "https://194.163.183.129/saihate-no-paladin-episode-10/",
			scrape: getdl.Samehadaku,
		},
		{
			name:   "animekompi",
			link:   "https://animekompi.net/hyouken-no-majutsushi-ga-sekai-wo-suberu-episode-01-subtitle-indonesia/",
			scrape: getdl.Animekompi,
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
