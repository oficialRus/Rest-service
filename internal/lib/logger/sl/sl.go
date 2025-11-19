package sl

import "log/slog"

func Err(err error) slog.Attr {
	return slog.Attr{
		Key:   "error",
		Value: slog.StringValue(err.Error()),
	}
<<<<<<< HEAD

=======
>>>>>>> 5ab71d5a7dd844656f6f44610dd94b3860f24d70
}
