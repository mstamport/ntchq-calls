
package main

type Filters struct {
	GroupOp string `json:"groupOp"`
	Rules   []struct {
		Field string `json:"field"`
		Op    string `json:"op"`
		Data  string `json:"data"`
	} `json:"rules"`
}