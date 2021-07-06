package filter

import (
	"github.com/antlinker/go-dirtyfilter"
	"github.com/antlinker/go-dirtyfilter/store"
)

var FilterManager *filter.DirtyManager

func InitFilter (keywords []string) {
	memStore, err := store.NewMemoryStore(store.MemoryConfig{
		DataSource: keywords,
	})
	if err != nil {
		panic(err)
	}
	FilterManager = filter.NewDirtyManager(memStore)
}