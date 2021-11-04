# Yonks API

```bash
$ cp .env.example .env  # 実行後、.envに適切な値を設定する
$ make deps
$ make up
$ make ngrok
```

### DBマイグレーション

##### マイグレーションの実行

```bash
$ make migrate/up
$ make migrate/down
```

##### マイグレーションファイルを追加する

`migrations/targets.go` の `registerMigrationTargets()` に該当するモデルを追記する

```go
func registerMigrationTargets() []interface{} {
	return []interface{}{
		&models.Foo{},
		&models.Bar{},
		&models.Baz{},
	}
}
```

### モックデータの作成

```bash
$ make seed
```
