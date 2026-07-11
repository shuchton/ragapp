# ragapp

A Go implementation of a Retrieval-Augmented Generation (RAG) application, built
while following Trevor Sawler's Udemy course
[*Building a RAG Application in Go (Golang)*](https://www.udemy.com/).

This repository tracks the course end-to-end: it starts as a minimal terminal
chat client wired to an LLM and grows into a full RAG pipeline with a vector
store, document ingestion, and a web frontend with streaming responses.

## What this covers

- **RAG pipeline fundamentals** — chunking, embedding, vector search, query
  rewriting, context injection, and streaming generation, end to end.
- **Streaming output** — streaming LLM tokens to a terminal REPL and to a
  browser via Server-Sent Events (SSE).
- **Vector search with Postgres** — using Postgres + `pgvector` for
  production-grade similarity search, including HNSW indexes and handling
  embedding-dimension migrations.
- **Multimodal content** — image upload, captioning with a vision model, and
  rendering images in the chat UI.
- **Swappable provider interfaces** — Go interfaces for the LLM, embedder, and
  vector store so any implementation can be swapped without touching the rest
  of the codebase.
- **Multiple LLM providers** — running against OpenAI, Ollama, LM Studio, or
  Groq, including mixing providers (e.g. a hosted chat model with local
  embeddings).
- **Reactive document ingestion** — watching a documents directory with
  `fsnotify`, debouncing half-written files, and re-ingesting idempotently.

## Project layout

```
cmd/rag/       entry point (main.go)
app/           application wiring / orchestration
config/        environment-based configuration loading
llm/           LLM client (OpenAI-compatible)
```

This layout will expand as the course introduces embeddings, vector storage,
document ingestion, and the web server/frontend.

## Status

Work in progress, following the course chapter by chapter. Setup and usage
instructions will be added as the application becomes runnable end-to-end.
