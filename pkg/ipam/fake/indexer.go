package fake

import (
	"k8s.io/client-go/tools/cache"
)

type Indexer struct {
	cache.Store
	IndexFunc               func(indexName string, obj interface{}) ([]interface{}, error)
	IndexKeysFunc           func(indexName, indexedValue string) ([]string, error)
	ListIndexFuncValuesFunc func(indexName string) []string
	ByIndexFunc             func(indexName, indexedValue string) ([]interface{}, error)
	GetIndexersFunc         func() cache.Indexers
	AddIndexersFunc         func(newIndexers cache.Indexers) error
}

func (f Indexer) Index(indexName string, obj interface{}) ([]interface{}, error) {
	if f.IndexFunc != nil {
		return f.IndexFunc(indexName, obj)
	}
	return nil, nil
}

func (f Indexer) IndexKeys(indexName, indexedValue string) ([]string, error) {
	if f.IndexKeysFunc != nil {
		return f.IndexKeysFunc(indexName, indexedValue)
	}
	return nil, nil
}

func (f Indexer) ListIndexFuncValues(indexName string) []string {
	if f.ListIndexFuncValuesFunc != nil {
		return f.ListIndexFuncValuesFunc(indexName)
	}
	return nil
}

func (f Indexer) ByIndex(indexName, indexedValue string) ([]interface{}, error) {
	if f.ByIndexFunc != nil {
		return f.ByIndexFunc(indexName, indexedValue)
	}
	return nil, nil
}

func (f Indexer) GetIndexers() cache.Indexers {
	if f.GetIndexersFunc != nil {
		return f.GetIndexersFunc()
	}
	return nil
}

func (f Indexer) AddIndexers(newIndexers cache.Indexers) error {
	if f.AddIndexersFunc != nil {
		return f.AddIndexersFunc(newIndexers)
	}
	return nil
}
