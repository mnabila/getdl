package scrape_test

import (
	"encoding/json"
	"fmt"
	"getdl/internal/scrape"
	"log"
	"os"
	"testing"
)

func writeToJson(data []byte, website string) {
	if err := os.WriteFile(fmt.Sprintf("result-%s.json", website), data, 0666); err != nil {
		log.Fatal(err)
	}
}

func TestSamehadaku(t *testing.T) {
	result := scrape.Samehadaku("https://194.163.183.129/saihate-no-paladin-episode-10/")
	tojson, _ := json.Marshal(result)
	t.Log(string(tojson))
	writeToJson(tojson, "samehadaku")

}

func TestOploverz(t *testing.T) {
	result := scrape.Oploverz("https://oploverz.fan/kimetsu-no-yaiba-s2-episode-03-subtitle-indonesia/")
	tojson, _ := json.Marshal(result)
	t.Log(string(tojson))
	writeToJson(tojson, "oploverz")

}

func TestDoronime(t *testing.T) {
	result := scrape.Doronime("https://doronime.id/anime/mieruko-chan/episode-12")
	tojson, _ := json.Marshal(result)
	t.Log(string(tojson))
	writeToJson(tojson, "doronime")

}

func TestLendrive(t *testing.T) {
	result := scrape.Lendrive("https://lendrive.web.id/akebi-chan-no-sailor-fuku-ep-01-dual-subs-x265-hevc-subtitle-indonesia-english/")
	tojson, _ := json.Marshal(result)
	t.Log(string(tojson))
	writeToJson(tojson, "lendrive")

}

func TestMirrored(t *testing.T) {
	result := scrape.Mirrored("https://www.mirrored.to/files/Y15EKT9F/[_LENDRIVE_]_Tensai_Ouji_no_Akaji_Kokka_Saisei_Jutsu_-_08_[720p_HEVC][Dualsubs].mkv_links")
	// tojson, _ := json.Marshal(result)
	// t.Log(string(tojson))
	// writeToJson(tojson, "lendrive")
	t.Log(result)

}
