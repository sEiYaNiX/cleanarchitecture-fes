# cleanarchitecture-fes
CleanArchitectureのサンプル

## DBの初期化・マイグレーション
```
export MYSQL_ARGS=cafes:cafes99@(localhost:3306)/cleanarchitecture_fes
docker-compose up -d
go run maintain/initdb.go
```

## 起動
(先にDBの初期化マイグレーションを実施してください)
```
go run src/entrypoint/*.go
```

## 終了
Ctrl+Cでgoプロセスを停止
MySQLサーバーを停止
```
docker-compose down
```

### GraphQL codegen

```
go run github.com/99designs/gqlgen --config graphql_schema/gqlgen.yml
```

#### lint
```
golangci-lint run
```
