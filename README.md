# cleanarchitecture-fes
CleanArchitectureのサンプル

# DBの初期化・マイグレーション
```
export MYSQL_ARGS=cafes:cafes99@(localhost:3306)/cleanarchitecture_fes
go run maintain/initdb.go
```

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
