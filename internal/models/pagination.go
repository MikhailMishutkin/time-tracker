package models

type FiltAndPagin struct {
	FilterMap   map[string]string
	Key         string
	ValueString string
	ValueInt    int
	Limit       int
	Offset      int
}
