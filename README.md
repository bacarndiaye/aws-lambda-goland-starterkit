**Overview**
- Purpose: Gin-based AWS Lambda API, inspired by `nexus-iban-client-lambda`.
- Modes: Local Gin server (no SAM), and Lambda via the Gin adapter.

**Run Locally (no SAM)**
- Prereqs: Go 1.24+, `go mod download` to fetch deps.
- Recommended: `make run` (uses `run-local.sh` to free the port before start).
- Alternatives:
  - `bash ./run-local.sh`
  - `go run ./local-main.go`
- Env: `PORT` (default 3000), `LOG_LEVEL` (default info), `ENV` (default local)
- Custom port: `PORT=4000 make run`
- URLs: `http://localhost:3000/healthz`, `http://localhost:3000/readyz`, `http://localhost:3000/v1/hello/world`
- Note: `run-local.sh` will try `lsof`/`fuser`/`ss` to kill any process listening on `PORT` before launching.

**Environment (.env files)**
- Files: `.env` (base), `.env.local` (machine overrides)
- Local loader: automatically loads `.env`, then overrides with `.env.local`
- Variables:
  - `ENV`: runtime environment, default `local`
  - `PORT`: HTTP port for local server, default `3000`
  - `LOG_LEVEL`: `debug|info|warn|error`, default `info`
  - `AWS_REGION`: region for AWS clients, default `eu-west-3`
- Customize: copy `.env.example` to `.env`, then tweak. Optionally create `.env.local` to override on your machine.

**Build For Lambda**
- Build tag: `lambda` (see `build_constraints.go`).
- Binary: `bootstrap` for `provided.al2` runtime, ARM64.
- Example: `GOOS=linux GOARCH=arm64 CGO_ENABLED=0 go build -tags lambda -ldflags="-s -w" -o bootstrap ./main.go`
- Package: `zip function.zip bootstrap` or `make package`.

**Project Layout**
- `main.go`: Lambda entrypoint using `aws-lambda-go-api-proxy/gin` (build tag `lambda`).
- `local-main.go`: Local Gin server (no SAM needed; build tag `!lambda`).
- `internal/router`: Gin engine and routes.
- `internal/handlers`: Handlers (`/healthz`, `/readyz`, `/v1/*`).
- `internal/middleware`: `CORS`, `RequestID`, `Logger`.
- `internal/config`: Env loading (port, env, log level).
- `template.yaml`: SAM template (proxy all routes; runtime provided.al2).
- `buildspec.yaml`: CodeBuild pipeline (builds `./main.go`).

**Dependencies**
- `github.com/gin-gonic/gin`: Web framework.
- `github.com/aws/aws-lambda-go`: Lambda types and bootstrap.
- `github.com/awslabs/aws-lambda-go-api-proxy/gin`: Adapter for API Gateway → Gin.

**Next Steps**
- Add your domain routes under `router.New()` or split into packages (`handlers`, `services`, `models`).
- If mirroring the reference repo further, add middleware, utils, and database layers similarly.
