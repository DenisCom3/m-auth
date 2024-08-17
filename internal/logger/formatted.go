package logger

import "log/slog"

func formatArgs(args ...any) []any {
	castArgs := make([]any, len(args))

	for _, arg := range args {
		var castArg any
		switch _arg := arg.(type) {
		case error:
			castArg = slog.Attr{
				Key:   "error",
				Value: slog.StringValue(_arg.Error()),
			}
		default:
			castArg = _arg
		}

		castArgs = append(castArgs, castArg)
	}

	return castArgs
}
