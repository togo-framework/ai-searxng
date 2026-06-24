// Package searxng is a togo AI data-source plugin: query a self-hosted SearXNG
// metasearch instance so agents and ai-rag can do web search. Registers an
// "ai-searxng" service + REST endpoint: POST /api/ai/searxng {"query":"…"}.
// Config: SEARXNG_URL (default http://localhost:8080).
package searxng

import (
	"context"
	"encoding/json"
	"net/http"
	"net/url"
	"os"
	"time"

	"github.com/go-chi/chi/v5"
	"github.com/togo-framework/togo"
)

// Result is a single search hit.
type Result struct {
	Title   string `json:"title"`
	URL     string `json:"url"`
	Content string `json:"content"`
}

// Source queries a SearXNG instance.
type Source struct {
	base   string
	client *http.Client
}

// New builds a Source from SEARXNG_URL.
func New() *Source {
	base := os.Getenv("SEARXNG_URL")
	if base == "" {
		base = "http://localhost:8080"
	}
	return &Source{base: base, client: &http.Client{Timeout: 20 * time.Second}}
}

// Search runs a query and returns results.
func (s *Source) Search(ctx context.Context, query string) ([]Result, error) {
	u := s.base + "/search?format=json&q=" + url.QueryEscape(query)
	req, err := http.NewRequestWithContext(ctx, http.MethodGet, u, nil)
	if err != nil {
		return nil, err
	}
	resp, err := s.client.Do(req)
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()
	var out struct {
		Results []Result `json:"results"`
	}
	if err := json.NewDecoder(resp.Body).Decode(&out); err != nil {
		return nil, err
	}
	return out.Results, nil
}

// FromKernel returns the registered Source, or nil.
func FromKernel(k *togo.Kernel) *Source {
	if v, ok := k.Get("ai-searxng"); ok {
		if s, ok := v.(*Source); ok {
			return s
		}
	}
	return nil
}

func init() {
	togo.RegisterProviderFunc("ai-searxng", togo.PriorityService, func(k *togo.Kernel) error {
		s := New()
		k.Set("ai-searxng", s)
		mount(k.Router, s)
		return nil
	})
}

func mount(r chi.Router, s *Source) {
	r.Post("/api/ai/searxng", func(w http.ResponseWriter, req *http.Request) {
		var body struct {
			Query string `json:"query"`
		}
		if err := json.NewDecoder(req.Body).Decode(&body); err != nil || body.Query == "" {
			http.Error(w, `{"error":"query required"}`, http.StatusBadRequest)
			return
		}
		res, err := s.Search(req.Context(), body.Query)
		if err != nil {
			http.Error(w, `{"error":"`+err.Error()+`"}`, http.StatusBadGateway)
			return
		}
		w.Header().Set("Content-Type", "application/json")
		_ = json.NewEncoder(w).Encode(map[string]any{"results": res})
	})
}
