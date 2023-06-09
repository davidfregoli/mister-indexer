package main

import (
	"fmt"
	"sort"
	"strings"
	"unicode"

	mr "fregoli.dev/mister"
)

type Indexer struct{}

func (app *Indexer) Map(document string, value string) []mr.KeyValue {
	res := []mr.KeyValue{}
	m := make(map[string]bool)
	words := strings.FieldsFunc(value, func(x rune) bool { return !unicode.IsLetter(x) })
	for _, w := range words {
		m[w] = true
	}
	for w := range m {
		kv := mr.KeyValue{Key: w, Value: document}
		res = append(res, kv)
	}
	return res
}

func (app *Indexer) Reduce(key string, values []string) string {
	sort.Strings(values)
	return fmt.Sprintf("%d %s", len(values), strings.Join(values, ","))
}

func main() {
	app := new(Indexer)
	mr.RegisterApp(app)
}
