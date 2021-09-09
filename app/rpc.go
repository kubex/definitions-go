package app

type Rpc struct {
	Key          string
	Inputs       []Attribute
	ResponseCode int32
	Outputs      []Attribute
}
