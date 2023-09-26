package app

type Rpc struct {
	Key          string
	Inputs       []Property
	ResponseCode int32
	Outputs      []Property
}
