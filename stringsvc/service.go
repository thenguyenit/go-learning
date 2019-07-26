package main

import "errors"

//Interface
type StringService interface {
	Uppercase(string) (string, error)
}

type stringService struct{}

var ErrEmtpy = errors.New("Empty string")
