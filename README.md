# cleanarchitecture-fes
CleanArchitectureのサンプル

## 起動

```
go run src/entrypoint/*.go
```

## GraphQL codegen

```
go run github.com/99designs/gqlgen --config graphql_schema/gqlgen.yml
```

#### lint
```
golangci-lint run
```
