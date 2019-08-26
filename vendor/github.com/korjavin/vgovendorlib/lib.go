package vgovendorlib

import (
	jsoniter "github.com/json-iterator/go"
)

func Something() ([]byte, error) {
	return jsoniter.Marshal(1)
}
