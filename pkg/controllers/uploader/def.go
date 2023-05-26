package uploader

import (
	"fmt"
	"net/http"
	"sort"
	"strings"
)

type OptionFn func(IUploader)
type IUploader interface {
	Reset(meta *Meta)
	Url() string
	Init() (string, error)
	Upload(index int, chunk []byte) (string, error)
	Merge(chunks []ChunkETag) error
}

type HeaderPair struct {
	Key, Value string
}

func (hp *HeaderPair) Less(o *HeaderPair) bool {
	return hp.Key < o.Key
}

type Headers []*HeaderPair

func (items Headers) Len() int {
	return len(items)
}
func (items Headers) Swap(i, j int)      { items[i], items[j] = items[j], items[i] }
func (items Headers) Less(i, j int) bool { return items[i].Less(items[j]) }
func (items Headers) Sort() {
	sort.Sort(items)
}
func (items Headers) Pair(key, value string) *HeaderPair {
	return &HeaderPair{Key: key, Value: value}
}
func (items Headers) Request(r *http.Request) {
	for _, p := range items {
		r.Header.Add(p.Key, p.Value)
	}
}
func (items Headers) String() string {
	items.Sort()
	arr := make([]string, 0, len(items))
	for _, p := range items {
		p.Key = strings.ToLower(p.Key)
		arr = append(arr, fmt.Sprintf("%s:%s", p.Key, p.Value))
	}
	return strings.Join(arr, "\n")
}
