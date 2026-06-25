# ai-searxng — documentation

  <img src=".github/assets/togo-mark.svg" alt="togo" height="64" />

## Overview

Package searxng is a togo AI data-source plugin: query a self-hosted SearXNG
metasearch instance so agents and ai-rag can do web search. Registers an
"ai-searxng" service + REST endpoint: POST /api/ai/searxng {"query":"…"}.
Config: SEARXNG_URL (default http://localhost:8080).

## Install

```bash
togo install togo-framework/ai-searxng
```

A capability plugin — it self-registers on boot; no driver selector needed.

## Configuration

Environment variables read by this plugin (extracted from the source):

| Env var | Notes |
|---|---|
| `G` | _see provider docs_ |
| `SEARXNG_URL` | _see provider docs_ |

## Usage

```go
// A data source for ai-rag / agents: fetch/scrape/search web content.
src := searxng.FromKernel(k)
docs, err := src.Fetch(ctx, "https://example.com")
```

## Links

- Marketplace: https://to-go.dev/marketplace
- Source: https://github.com/togo-framework/ai-searxng
- README: ../README.md
