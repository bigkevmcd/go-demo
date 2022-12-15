# go-demo

This is a simple demo service for a Kubernetes pipeline.

# Requirements

This requires a configured Redis server.

# Building

```shell
$ go build ./cmd/go-demo
```

# Testing

```shell
TEST_REDIS_URL=redis://localhost:6379  go test ./...
```

If `TEST_REDIS_URL` is _not set_ this will default to `redis://localhost:6379/9`.

## Acorn

Runninng this with [Acorn](https://acorn.io/) is as simple as:

```shell
$ acorn install
$ acorn run -n go-demo .
```
