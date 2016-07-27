package logrus

// IDKey defines the key when adding the id for the request.
var IDKey = "id"

// ID allows you to ad an id to the log entry
func (entry *Entry) ID(id string) *Entry {
	return entry.WithField(IDKey, id)
}

// ID allows you to ad an id to the log entry
func ID(id string) *Entry {
	return std.WithField(IDKey, id)
}

// WithLocation allows you to add a location in the source code
func WithLocation() *Entry {
	return std.WithField("loc", getLocation())
}

// WithLocation allows you to add a location in the source code
func (entry *Entry) WithLocation() *Entry {
	return entry.WithField("loc", getLocation())
}
