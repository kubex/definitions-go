package app

type Rpc struct {
	Key          ScopedKey
	Inputs       []Attribute
	ResponseCode int32
	Outputs      []Attribute
}
