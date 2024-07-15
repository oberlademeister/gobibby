package main

import (
	"log/slog"
	"strconv"
	"time"
)

func DecodeString(m map[string]any, key string) string {
	a, ok := m[key]
	if !ok {
		//slog.Warn("key does not exist", "key", key)
		return ""
	}
	if t, ok := a.(string); ok {
		return t
	}
	slog.Warn("interface can't be cast to string", "key", key)
	return ""
}

func DecodeInt(m map[string]any, key string) string {
	a, ok := m[key]
	if !ok {
		//slog.Warn("key does not exist", "key", key)
		return ""
	}
	if t, ok := a.(float64); ok {
		return strconv.Itoa(int(t))
	}
	slog.Warn("interface can't be cast to int", "key", key)
	return ""
}

func DecodeStringSlice(m map[string]any, key string) []string {
	a, ok := m[key]
	if !ok {
		//slog.Warn("key does not exist", "key", key)
		return nil
	}
	t, ok := a.([]any)
	if ok {
		var ret []string
		for _, s := range t {
			if ts, tok := s.(string); tok {
				ret = append(ret, ts)
			}
		}
		return ret
	}
	slog.Warn("interface can't be cast to string slice", "key", key, "a", a)
	return nil
}

func DecodeDate(m map[string]any, key string) time.Time {
	s := DecodeString(m, key)
	if s == "" {
		slog.Warn("zero date from zero string", "key", key)
		return time.Time{}
	}
	t, err := time.Parse("2006-01-02", s)
	if err != nil {
		slog.Warn("time parse error", "err", err, "t", t, "s", s, "key", key)
		return time.Time{}
	}
	return t
}
