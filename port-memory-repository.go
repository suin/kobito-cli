package main

// SQLiteと結合すると重いので、デバッグ用に作ったfakeクラス
type OnMemoryItemRepository struct {
}

func NewOnMemoryItemRepository() (this *OnMemoryItemRepository) {
	this = new(OnMemoryItemRepository)
	return
}

func (this *OnMemoryItemRepository) Items() (items []*Item, err error) {
	for i := 10; i > 0; i-- {
		items = append(items, NewItem(345*i, "Go言語のクロスコンパイル設定値 $GOOS, $GOARCH 一覧リスト", "html...", "markdown..."))
		items = append(items, NewItem(344*i, "Go言語でコマンドが実行可能かチェックする", "html...", "markdown..."))
		items = append(items, NewItem(343*i, "Go言語でオブジェクトを作ってみる: オブジェクト・メソッド・コンストラクタなど", "html...", "markdown..."))
		items = append(items, NewItem(340*i, "Go言語でファイル操作: ディレクトリの作り方", "html...", "markdown..."))
		items = append(items, NewItem(341*i, "リンク内のシングルクォーテーションがエスケープされる", "html...", "markdown..."))
	}

	return
}

func (this *OnMemoryItemRepository) ItemOfId(id int) (item *Item, err error) {
	item = NewItem(345, "Go言語のクロスコンパイル設定値/$GOOS,/$GOARCH/一覧リスト", `<!DOCTYPE HTML>
<html>
  <head>
    <meta http-equiv="content-type" content="text/html; charset=UTF-8" />
    <title>Kobito Memo</title>
    <script type="text/javascript" src="highlight.pack.js"></script>
    <script>hljs.initHighlightingOnLoad();</script>
    <link type="text/css" rel="stylesheet" href="github.min.css">
    <link type="text/css" rel="stylesheet" href="markdown.css">
  </head>
  <body>
<h1>Markdown記法チートシート</h1>

<p>Kobitoで利用可能なMarkdown記法のチートシートです。</p>

<p>Markdown記法の原文については、<a href="http://daringfireball.net/projects/markdown/syntax.php">Daring Fireball: Markdown Syntax Documentation</a>をご覧下さい。<br>
また、コードに関する記法は<a href="http://github.github.com/github-flavored-markdown/">GitHub Flavored Markdown</a>に準拠しています。</p>

<h2>Code - コードの挿入</h2>

<p>たとえば、Rubyで記述したコードをファイル名「kobito.rb」として記録したいときは、<strong>バッククオート</strong>を使用して以下のように投稿するとシンタックスハイライトが適用されます。<br>
<strong>コードブロック上下に空行を挿入しないと正しく表示されないことがあります。</strong></p>

<div class="code-frame"><div class="code-lang"><span class="bold">kobito.rb</span></div><div class="highlight"><pre><code class="ruby">puts &#39;The best app to log and share programming knowledge.&#39;
</code></pre></div></div>

  </body>
</html>`, "タイトル\n# 見出し1\n\n* list\n* list\n* list\n")
	return
}
