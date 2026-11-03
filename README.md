# m11s/redis-http-proxy

Small, bearer-token-authenticated HTTP proxy for a Redis instance. It exposes
Redis commands over HTTP for clients that cannot use the Redis protocol.

## Tags

| Tag | Use |
| --- | --- |
| `latest` | Most recently published version release |
| `0.1.0` | Pinned published release |
| `sha-<commit>` | Image built from a commit pushed to `main` |

Version tags are published from matching `vX.Y.Z` Git tags. Commits pushed to
`main` publish only their `sha-<commit>` tag; they do not move `latest` or
rewrite a version release.

## Usage

```bash
docker run --rm -p 8080:80 \
  -e TOKEN=replace-with-a-long-secret \
  -e REDIS_URL=redis://redis:6379 \
  m11s/redis-http-proxy:latest
```

Run a command with a bearer token:

```bash
curl -X POST http://localhost:8080/SET/example/value \
  -H 'Authorization: Bearer replace-with-a-long-secret'
curl -X POST http://localhost:8080/GET/example \
  -H 'Authorization: Bearer replace-with-a-long-secret'
```

`TOKEN` is required. `REDIS_URL` defaults to `redis://localhost:6379`, and
`PORT` defaults to `80`. The unauthenticated `GET /healthz` endpoint reports
whether Redis is reachable. Source and license: [m11s-io/redis-http-proxy](https://github.com/m11s-io/redis-http-proxy).
