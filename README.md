# Yonks API

```
$ cp .env.example .env  # 実行後、.envに適切な値を設定する
$ make deps
$ make up
$ make ngrok
```

### DBマイグレーション

##### マイグレーションの実行

```
$ make migrate/up
$ make migrate/down
```

##### マイグレーションファイルを追加する

`migrations/targets.go` の `registerMigrationTargets()` に該当するモデルを追記する

```
func registerMigrationTargets() []interface{} {
	return []interface{}{
		&models.Foo{},
        &models.Bar{},
        &models.Baz{},
	}
}
```
