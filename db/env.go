package db

import (
	"os"
	"strings"
)

type DbInfo struct {
	Path     string
	Name     string
	FullPath string
}

var dbInfo DbInfo = GetDbInfo()

func (info *DbInfo) setUserPath(path string) {
	if !strings.HasSuffix(info.Path, "/") {
		path += "/"
		info.Path = path
		info.FullPath = path + "tvpmdb.db"
	} else {
		info.Path = path
		info.FullPath = path + "tvpmdb.db"
	}
}

func GetDbInfo() DbInfo {
	userDbPath := os.Getenv("TVPM_DB_PATH")

	dbConfig := DbInfo{
		Path:     "$HOME/tvpm/",
		Name:     "tvpmdb.db",
		FullPath: "$HOME/tvpm/tvpmdb.db",
	}

	if len(userDbPath) > 0 {
		dbConfig.setUserPath(userDbPath)
	}

	return dbConfig
}
