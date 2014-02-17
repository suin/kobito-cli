# SQLite結合テスト用
run: fmt
	go run port-sqlite3-*.go cmd-*go model.go main.go ${ARGS}

# デバッグ用
mem: fmt
	go run port-memory-*.go cmd-*go model.go main.go ${ARGS}

fmt:
	go fmt ./...

build: fmt
	# バージョンをつける
	sed -i '' -e "s/__KOBITO_CLI_VERSION__/$$(git rev-parse HEAD)/g" cmd-version.go

	go build -o bin/kobito port-sqlite3-*.go cmd-*go model.go main.go

	# 元に戻す
	sed -i '' -e "s/$$(git rev-parse HEAD)/__KOBITO_CLI_VERSION__/g" cmd-version.go

	# README の使い方を更新する
	perl -pi -e "BEGIN{undef $$/;} s#<usage>.*?</usage>#<usage>\n\n\`\`\`\n$$(./bin/kobito help 2>&1)\n\`\`\`\n\n</usage>#s"  README.md

install: build
	cp bin/* /usr/local/bin/
