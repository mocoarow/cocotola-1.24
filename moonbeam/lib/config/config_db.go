package config

import (
	"context"
	"database/sql"
	"io/fs"
	"log/slog"

	"gorm.io/gorm"

	libdomain "github.com/mocoarow/cocotola-1.24/moonbeam/lib/domain"
	liberrors "github.com/mocoarow/cocotola-1.24/moonbeam/lib/errors"
	libgateway "github.com/mocoarow/cocotola-1.24/moonbeam/lib/gateway"
)

type DBConfig struct {
	DriverName string                     `yaml:"driverName"`
	MySQL      *libgateway.MySQLConfig    `yaml:"mysql"`
	Postgres   *libgateway.PostgresConfig `yaml:"postgres"`
	SQLite3    *libgateway.SQLite3Config  `yaml:"sqlite3"`
	Migration  bool                       `yaml:"migration"`
}

type MergedFS struct {
	fss     []fs.FS
	entries []fs.DirEntry
}

func MergeFS(driverName string, fss ...fs.FS) (*MergedFS, error) {
	entries := make([]fs.DirEntry, 0)
	for i := range fss {
		e, err := fs.ReadDir(fss[i], driverName)
		if err != nil {
			return nil, liberrors.Errorf("read %q directory: %w", driverName, err)
		}
		entries = append(entries, e...)
	}

	return &MergedFS{
		fss:     fss,
		entries: entries,
	}, nil
}

func (f *MergedFS) Open(name string) (fs.File, error) {
	var file fs.File
	var err error
	for i := range f.fss {
		file, err = f.fss[i].Open(name)
		if err == nil {
			return file, nil
		}
	}

	return nil, err
}

func (f *MergedFS) ReadDir(_ string) ([]fs.DirEntry, error) {
	return f.entries, nil
}

func InitDB(ctx context.Context, cfg *DBConfig, logConfig *LogConfig, appName string, sqlFSs ...fs.FS) (libgateway.DialectRDBMS, *gorm.DB, *sql.DB, error) {
	mergedFS, err := MergeFS(cfg.DriverName, sqlFSs...)
	if err != nil {
		return nil, nil, nil, liberrors.Errorf("merge sql files in %q directory: %w", cfg.DriverName, err)
	}

	initDBFunc, ok := initDBs[cfg.DriverName]
	if !ok {
		return nil, nil, nil, libdomain.ErrInvalidArgument
	}
	dbLogLevel := slog.LevelWarn
	if level, ok := logConfig.Levels["db"]; ok {
		dbLogLevel = stringToLogLevel(level)
	}

	return initDBFunc(ctx, cfg, dbLogLevel, mergedFS, appName)
}
