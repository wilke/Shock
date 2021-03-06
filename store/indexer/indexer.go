package indexer

import (
	"github.com/MG-RAST/Shock/store/indexer/record"
	"github.com/MG-RAST/Shock/store/type/index"
	"io"
)

type indexerFunc func(io.ReadCloser) index.Indexer

var (
	indexers = map[string]indexerFunc{
		"record": record.NewIndexer,
		"size":   record.NewIndexer,
	}
)

func Indexer(i string) indexerFunc {
	return indexers[i]
}
