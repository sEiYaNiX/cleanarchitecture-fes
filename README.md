# cleanarchitecture-fes
CleanArchitectureのサンプル

## 起動

```
go run src/entrypoint/main.go
```

## GraphQL のスキーマ更新 & codegen

- `graphql_schema`以下の`*.schema`を編集する
- 以下を実行

```shell script
go run github.com/99designs/gqlgen --config graphql_schema/gqlgen.yml
```


#### lint 手動実行
```shell script
golangci-lint run
```
