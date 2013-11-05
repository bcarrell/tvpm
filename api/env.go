package api

import (
	"os"
)

var (
	TraktKey string = os.Getenv("TRAKT_API_KEY")
)
