package core

const (
	RequestTypeJson RequestType = iota + 1
	RequestTypeForm
)

type RequestType uint8
