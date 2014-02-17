package main

type Item struct {
	Id       int
	Title    string
	Html     string
	Markdown string
}

func NewItem(id int, title, html, markdown string) (this *Item) {
	this = new(Item)
	this.Id = id
	this.Title = title
	this.Html = html
	this.Markdown = markdown
	return
}

type ItemRepository interface {
	Items() (items []*Item, err error)
	ItemOfId(id int) (item *Item, err error)
}
