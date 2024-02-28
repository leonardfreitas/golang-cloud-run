# Run Tests

```bash
make test
```

# Build

```bash
make build
```

### Dev

### Docker

```bash
go run cmd/main.go
```

#### OK

curl -X GET http://localhost:8080?cep=89216310

#### Not Found

curl -X GET http://localhost:8080?cep=89216369

#### Unprocessable Entity

curl -X GET http://localhost:8080?cep=892169

### Cloud run

`https://.run.app`
