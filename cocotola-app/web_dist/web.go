package web

import (
	"embed"
	_ "embed"
)

//go:embed react/**
//go:embed flutter/**
var Web embed.FS
