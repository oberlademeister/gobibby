package main

import (
	"fmt"
	"log/slog"
	"strconv"
)

type BTType int

const (
	TWikipedia BTType = iota
	TSAPInternal
	TInternet
	TBook
)

type Item struct {
	T  BTType
	Id string
	M  map[string]any
}

type ItemList []Item

func (i Item) RenderBibTex() string {
	switch i.T {
	case TWikipedia:
		ad := DecodeDate(i.M, "accessedDate")
		return RenderBibTex("misc", i.Id, [][2]string{
			{"author", "Wikipedia"},
			{"title", fmt.Sprintf("%s --- {W}ikipedia{,} The Free Encyclopedia", DecodeString(i.M, "title"))},
			{"year", strconv.Itoa(ad.Year())},
			{"howpublished", fmt.Sprintf("\\url{%s}", DecodeString(i.M, "url"))},
			{"note", fmt.Sprintf("[Online; accessed %s]", ad.Format("2006-01-02"))},
		})
	case TInternet:
		ad := DecodeDate(i.M, "accessedDate")
		as := DecodeStringSlice(i.M, "authors")
		return RenderBibTex("misc", i.Id, [][2]string{
			{"author", MakeAuthorString(as)},
			{"title", fmt.Sprintf("%s", DecodeString(i.M, "title"))},
			{"year", strconv.Itoa(ad.Year())},
			{"howpublished", fmt.Sprintf("\\url{%s}", DecodeString(i.M, "url"))},
			{"note", fmt.Sprintf("[Online; accessed %s]", ad.Format("2006-01-02"))},
		})
	case TSAPInternal:
		ad := DecodeDate(i.M, "accessedDate")
		as := DecodeStringSlice(i.M, "authors")
		return RenderBibTex("misc", i.Id, [][2]string{
			{"author", MakeAuthorString(as)},
			{"title", fmt.Sprintf("%s", DecodeString(i.M, "title"))},
			{"year", strconv.Itoa(ad.Year())},
			{"howpublished", fmt.Sprintf("\\url{%s}", DecodeString(i.M, "url"))},
			{"note", fmt.Sprintf("[Online; accessed %s]", ad.Format("2006-01-02"))},
		})

	case TBook:
		as := DecodeStringSlice(i.M, "authors")
		paras := [][2]string{
			{"author", MakeAuthorString(as)},
			{"title", fmt.Sprintf("%s", DecodeString(i.M, "title"))},
			{"year", DecodeInt(i.M, "year")},
			{"publisher", DecodeString(i.M, "publisher")},
		}
		if es := DecodeStringSlice(i.M, "editors"); len(es) > 0 {
			paras = append(paras, [2]string{"editor", MakeAuthorString(es)})
		}
		paras = optionalStringArgument(i.M, "month", paras)
		paras = optionalStringArgument(i.M, "isbn", paras)
		paras = optionalStringArgument(i.M, "edition", paras)
		paras = optionalStringArgument(i.M, "address", paras)
		paras = optionalStringArgument(i.M, "abstract", paras)
		return RenderBibTex("book", i.Id, paras)

	}
	slog.Warn("no rendering for type", "t", i.T)
	return ""
}

func optionalStringArgument(m map[string]any, key string, paras [][2]string) [][2]string {
	if s := DecodeString(m, key); s != "" {
		return append(paras, [2]string{key, s})
	}
	return paras
}
