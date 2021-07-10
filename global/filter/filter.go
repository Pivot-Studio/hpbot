package filter

import (
	"bufio"
	"net/http"

	"github.com/Mrs4s/go-cqhttp/global/config"
	filter "github.com/antlinker/go-dirtyfilter"
	"github.com/antlinker/go-dirtyfilter/store"
)

var FilterManager *filter.DirtyManager

func InitFilter(keywords []string) {
	memStore, err := store.NewMemoryStore(store.MemoryConfig{
		DataSource: keywords,
	})
	if err != nil {
		panic(err)
	}
	FilterManager = filter.NewDirtyManager(memStore)
}

func FetchBlockWord() {
	res, err := http.Get("https://raw.githubusercontent.com/Pivot-Studio/HUSTPORT_blockwords/master/blockwords.txt")
	if err != nil {
		return
	}
	defer res.Body.Close()
	c := config.Get()
	s := bufio.NewScanner(res.Body)
	c.Keywords = []string{}
	for s.Scan() {
		c.Keywords = append(c.Keywords, s.Text())
	}
}
