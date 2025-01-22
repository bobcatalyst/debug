//go:build debug

package debug

import "log/slog"

const Debug = true

const LogLevel = slog.LevelDebug

func Error(err error) error {
    return err
}
