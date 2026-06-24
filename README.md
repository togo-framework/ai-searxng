<!-- togo-header -->
<div align="center">
  <img src=".github/assets/togo-mark.svg" alt="togo" height="64" />
  <h1>togo-framework/ai-searxng</h1>
  <p>
    <a href="https://to-go.dev/marketplace"><img src="https://img.shields.io/badge/marketplace-to--go.dev-1FC7DC" alt="marketplace" /></a>
    <a href="https://pkg.go.dev/github.com/togo-framework/ai-searxng"><img src="https://pkg.go.dev/badge/github.com/togo-framework/ai-searxng.svg" alt="pkg.go.dev" /></a>
    <img src="https://img.shields.io/badge/license-MIT-blue" alt="MIT" />
  </p>
  <p><strong>Part of the <a href="https://to-go.dev">togo</a> framework.</strong></p>
</div>

## Install

```bash
togo install togo-framework/ai-searxng
```

<!-- /togo-header -->

# ai-searxng

A togo **AI data-source** plugin — web search via a self-hosted **SearXNG** metasearch instance, for agents and `ai-rag`.

```
togo install togo-framework/ai-searxng
```
Set `SEARXNG_URL` (default `http://localhost:8080`).

## Use
- Go: `searxng.FromKernel(k).Search(ctx, "togo framework")` → `[]Result{Title,URL,Content}`
- REST: `POST /api/ai/searxng` `{"query":"…"}`

Part of the [togo AI kit](https://to-go.dev/ai). MIT.

<!-- togo-sponsors -->
---

<div align="center">
  <h3>Premium sponsors</h3>
  <p>
    <a href="https://id8media.com"><strong>ID8 Media</strong></a> &nbsp;·&nbsp;
    <a href="https://one-studio.co"><strong>One Studio</strong></a>
  </p>
  <p><sub>Support togo — <a href="https://github.com/sponsors/fadymondy">become a sponsor</a>.</sub></p>
</div>
<!-- /togo-sponsors -->
