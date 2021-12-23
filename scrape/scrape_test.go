package scrape_test

import (
	"encoding/json"
	"fmt"
	"getdl/scrape"
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

}

func TestOploverz(t *testing.T) {
	result := scrape.Oploverz("https://oploverz.fan/kimetsu-no-yaiba-s2-episode-03-subtitle-indonesia/")
	tojson, _ := json.Marshal(result)
	t.Log(string(tojson))

}

func TestDoronime(t *testing.T) {
	result := scrape.Doronime("https://doronime.id/anime/mieruko-chan/episode-12")
	tojson, _ := json.Marshal(result)
	t.Log(string(tojson))
	writeToJson(tojson, "doronime")

}

// func TestEgao(t *testing.T) {
// 	result := scrape.BypassEgao("https://egao.in?id=eyJpdiI6IkFZUVZ2RjZGVWFTdFI5RlhoNDBlTGc9PSIsInZhbHVlIjoiRWtwZGNKRjArNTBYTnUrMHpVcjVlQT09IiwibWFjIjoiNGM0MGRmZTQ2ZWZhNDViN2ZlNzg3NDg4YWU3NTllZDNjODYxNWE5MGJhYzNjZjZjYzkyZWJiOGRkMjk1MmFlMyJ9")
// 	t.Log(result)
//
// }
//
