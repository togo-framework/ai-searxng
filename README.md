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
