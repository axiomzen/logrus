package logrus

func (entry *Entry) ID(id string) *Entry {
	return entry.WithField("id", id)
}

func ID(id string) *Entry {
	return std.WithField("id", id)
}

func (entry *Entry) WithError(err error) *Entry {
	return entry.WithField("error", err)
}

func WithError(err error) *Entry {
	return std.WithField("error", err)
}
