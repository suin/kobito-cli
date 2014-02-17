# Kobito cli tools

Kobitoをちょっぴり便利にするCLIツール

## インストール



## 使い方

<usage>

```

Usage: 
  kobito [command]

Available Commands: 
  ls                        :: List all items
  show [item id]            :: Show item
  html [item id]            :: Show item as HTML
  print [item id]           :: Print out item
  link [markdown file]      :: Link markdown file to Kobito
  version                   :: Print kobito cli tools version
  help [command]            :: Help about any command


Use "kobito help [command]" for more information about that command.
```

</usage>


## コントリビュート

go1.2をbrewでインストールしておいてください

チェックアウトしたら依存するライブラリを入れます:

```
go get ./...
```

### デバッグ方法

とりあえず make すれば実行できます

```
make
```

サブコマンドの実行は下記のように `ARGS` で指定できます

```
make ARGS="show 123"
```

SQLiteのライブラリをコンパイルするのに時間がかかるため、SQLite非依存の開発は下記のように `mem` でやることをおすすめします

```
make mem

make mem ARGS="show 123"
```



