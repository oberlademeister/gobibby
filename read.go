package main

import (
	"fmt"
	"log/slog"
	"os"
	"path/filepath"

	"encoding/json"

	"cuelang.org/go/cue"
	cjson "cuelang.org/go/encoding/json"
)

func ReadItemsMap(path, globstring, idprefix string,
	t BTType, cueContext *cue.Context, cueSchema cue.Value) []Item {

	globpath := filepath.Join(path, globstring)
	fl, err := filepath.Glob(globpath)
	if err != nil {
		slog.Warn("failed to glob %s: %w", globpath, err)
		return nil
	}
	var ret []Item
	for _, fl := range fl {
		// validate first
		slog.Info("validate json", "file", fl)
		b, err := os.ReadFile(fl)
		if err != nil {
			slog.Warn("failed to read; skipping read", "fl", fl, "err", err)
			continue
		}
		jsonFile, err := cjson.Extract(fl, b)
		if err != nil {
			slog.Warn("can't extract json; skipping read", "file", fl)
			continue
		}
		jsonAsCUE := cueContext.BuildExpr(jsonFile)
		unified := cueSchema.Unify(jsonAsCUE)
		if err := unified.Validate(); err != nil {
			slog.Warn("json validation failed; skipping read", "file", fl, "err", err)
			continue
		}
		var m map[string]any
		err = json.Unmarshal(b, &m)
		if err != nil {
			slog.Warn("failed to unmarshal", "fl", fl, "err", err)
			continue
		}
		id, ok := m["id"]
		if !ok {
			slog.Warn("id not found", "fl", fl)
			continue
		}
		ret = append(
			ret,
			Item{
				T:  t,
				Id: fmt.Sprintf("%s:%s", idprefix, id),
				M:  m,
			},
		)
	}
	return ret
}
