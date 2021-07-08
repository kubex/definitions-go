package app

type Event struct {
	App        GlobalAppID
	Key        string
	Attributes []Attribute
}
