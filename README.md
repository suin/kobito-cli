# Kobito cli tools

Kobitoをちょっぴり便利にするCLIツール

![](https://raw2.github.com/suin/kobito-cli/master/image.png)

## インストール

### Homebrewで入れる方法

```
brew install --HEAD https://raw.github.com/suin/kobito-cli/master/homebrew/kobito-cli.rb
```

### wgetで入れる方法

```
wget -O /usr/local/bin/kobito https://github.com/suin/kobito-cli/raw/master/bin/kobito
chmod +x /usr/local/bin/kobito
```


## 使い方

<usage>

```

Usage: 
  kobito [command]

Available Commands: 
  ls                                            :: List all items
  show [item id]                                :: Show item
  html [item id]                                :: Show item as HTML
  print [item id]                               :: Print out item
  link [markdown file]                          :: Link markdown file to Kobito
  version                                       :: Print kobito cli tools version
  pdf [item id] | pdf [item id] [pdf file name] :: Save item as PDF
  password                                      :: Show Kobito password
  sticker                                       :: Do you want Kobito sticker?
  help [command]                                :: Help about any command


Use "kobito help [command]" for more information about that command.
```

</usage>


### Kobito へ文字列を送信する

パイプやリダイレクトなどで入力を渡すとKobitoのエントリーになります

historyの5件からKobitoのエントリーを作る

```
history | tail -n 5 | kobito
```

main.goのソースコードでKobitoのエントリーを作る

```
kobito < main.go
```

「[Ruby - ターミナルから Kobito へ文字列を送信する - Qiita](http://qiita.com/watson1978/items/c6ad9417298367aa9b9b)」を参考にしました。


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
