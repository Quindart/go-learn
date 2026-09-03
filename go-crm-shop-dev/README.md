# go-crm-shop-dev

CRM shop service written in Go.

## Run

```sh
go run ./cmd/app
```

The API listens on `:8080` by default. Override it with `APP_ADDR`.

Health check:

```sh
curl http://localhost:8080/healthz
```