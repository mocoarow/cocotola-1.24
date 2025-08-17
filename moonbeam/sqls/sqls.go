package sqls

import (
	"embed"
	_ "embed"
)

//go:embed mysql/*.sql
//go:embed postgres/*.sql
//go:embed sqlite3/*.sql
var SQL embed.FS
