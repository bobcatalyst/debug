//go:build !debug

package debug

import "log/slog"

const Debug = false

const LogLevel = slog.LevelInfo

func Error(error) error {
    return nil
}
