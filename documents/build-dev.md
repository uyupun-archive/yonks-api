# 環境構築(開発環境)

```bash
# .envの作成
# 実行後、.envに適切な値を設定する
$ cp .env.example .env
# 依存関係のライブラリをインストール
$ make deps
# MySQLのDockerコンテナとAPIの開発用サーバを起動
$ make up
# MySQLのDockerコンテナを停止
$ make down
# Dockerのプロセスを確認
$ make ps
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

### DBシーダー

##### シーダーの実行

```bash
$ make seed/init
$ make seed/mock
```

##### テーブルの初期化データを追加する

- 本番環境においても使用されるような初期化用のデータを配置する
- `seeds/init` 以下にJSONを定義する
- `seeds/targets.go` の `registerInitTargets()` に該当するモデルとファイルを追記する

```go
func registerInitTargets() []Target {
	return []Target{
		{
			Model: models.Foo{},
			Seed:  "foos.json",
		},
		{
			Model: models.Bar{},
			Seed:  "bars.json",
		},
	}
}
```

##### テーブルのモックデータを追加する

- `seeds/mock` 以下にJSONを定義する
- 開発環境でのみ使用するようなテスト用のデータを配置する
- `seeds/targets.go` の `registerMockTargets()` に該当するモデルとファイルを追記する

```go
func registerMockTargets() []Target {
	return []Target{
		{
			Model: models.Foo{},
			Seed:  "foos.json",
		},
		{
			Model: models.Bar{},
			Seed:  "bars.json",
		},
	}
}
```

### テスト用ユーザ

|ユーザID|パスワード|
|:--|:--|
|test001|testtest|
|test002|testtest|
