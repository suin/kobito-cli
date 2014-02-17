package main

import (
	"os/user"
)

func init() {
	usr, err := user.Current()
	failOnError(err)
	itemRepository = NewSqliteItemRepository(usr.HomeDir + "/Library/Kobito/Kobito.db")
}
