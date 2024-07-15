package main

import (
	"fmt"
	"log/slog"
	"os"
	"sort"

	"cuelang.org/go/cue"
	"cuelang.org/go/cue/cuecontext"
	"github.com/urfave/cli/v2"
)

func run(clic *cli.Context) error {
	dbpath := clic.String("dbpath")
	outfile := clic.String("outfile")
	slog.Info("starting run", "dbpath", dbpath, "outfile", outfile)
	ctx := cuecontext.New()
	var items []Item

	compiledBytes := ctx.CompileBytes(cuefile)
	{
		cueValue := compiledBytes.LookupPath(cue.ParsePath("#Book"))
		tItems := ReadItemsMap(dbpath, "book-*.json", "book", TBook, ctx, cueValue)
		items = append(items, tItems...)
	}
	{
		cueValue := compiledBytes.LookupPath(cue.ParsePath("#Wikipedia"))
		tItems := ReadItemsMap(dbpath, "wikipedia-*.json", "wikipedia", TWikipedia, ctx, cueValue)
		items = append(items, tItems...)
	}
	{
		cueValue := compiledBytes.LookupPath(cue.ParsePath("#SAPInternal"))
		tItems := ReadItemsMap(dbpath, "sapinternal-*.json", "sapinternal", TSAPInternal, ctx, cueValue)
		items = append(items, tItems...)
	}
	{
		cueValue := compiledBytes.LookupPath(cue.ParsePath("#Internet"))
		tItems := ReadItemsMap(dbpath, "internet-*.json", "internet", TInternet, ctx, cueValue)
		items = append(items, tItems...)
	}
	// sort
	sort.SliceStable(items, func(i, j int) bool {
		return items[i].Id < items[j].Id
	})
	out, err := os.Create(outfile)
	if err != nil {
		return err
	}
	defer out.Close()
	slog.Info("opened bibtex out file", "file", outfile)
	for _, it := range items {
		fmt.Fprintln(out, it.RenderBibTex())
	}
	return nil
}
