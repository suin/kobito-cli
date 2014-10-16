package main

import (
	"os"
	"os/user"
)

func init() {
	usr, err := user.Current()
	failOnError(err)

	// Kobito version 2.x
	kobito2 := usr.HomeDir + "/Library/Containers/com.qiita.Kobito/Data/Library/Kobito/Kobito.db"
	// Kobito version 1.x
	kobito1 := usr.HomeDir + "/Library/Kobito/Kobito.db"

	var kobitoDbFile string

	if _, err := os.Stat(kobito2); err == nil {
		kobitoDbFile = kobito2
	} else {
		kobitoDbFile = kobito1
	}

	itemRepository = NewSqliteItemRepository(kobitoDbFile)
}
