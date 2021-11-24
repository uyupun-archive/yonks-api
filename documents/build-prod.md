# 環境構築(本番環境)

```bash
$ cp .env.example .env  # 実行後、.envに適切な値を設定する
# サーバとDBの起動
$ make up -f Makefile.prod
# サーバとDBの終了
$ make down -f Makefile.prod
# プロセスの確認
$ make ps -f Makefile.prod
# DBマイグレーションの実行
$ make migrate/up -f Makefile.prod
$ make migrate/down -f Makefile.prod
# 初期化データの追加
$ make seed/init -f Makefile.prod
# モックデータの追加
$ make seed/mock -f Makefile.prod
```

- その他のコマンドの詳細に関しては[Makefile.prod](../Makefile.prod)を参照
