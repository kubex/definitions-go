package app

type Rpc struct {
	App          GlobalAppID
	Key          string
	Inputs       []Attribute
	ResponseCode int32
	Outputs      []Attribute
}
