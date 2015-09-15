package main

type stamper interface {
	lick() []byte
}

type envelope struct {
	content []byte
}

type stamp struct {
	content []byte
}
