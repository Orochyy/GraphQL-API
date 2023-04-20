# GraphQL-API

### 1) Add gqlgen to your project's tools.go
```bash
printf '// +build tools\npackage tools\nimport (_ "github.com/99designs/gqlgen"\n _ "github.com/99designs/gqlgen/graphql/introspection")' | gofmt > tools.go

```


### 2)
```bash
go run github.com/99designs/gqlgen init
```

### 3)
```bash
go run github.com/99designs/gqlgen generate
```