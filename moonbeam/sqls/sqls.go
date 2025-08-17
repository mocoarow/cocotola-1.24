package sqls

import (
	"embed"
	"os"
	// _ "embed"
)

//go:embed mysql/*.sql
//go:embed postgres/*.sql
//go:embed sqlite3/*.sql
var SQL embed.FS

func A() string {
	return os.Getenv("A")
}
