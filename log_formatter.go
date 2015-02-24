package logrus

import (
	"bytes"
	"fmt"
	"sort"
	"strings"
	"time"
)

func init() {
	baseTimestamp = time.Now()
	isTerminal = IsTerminal()
}

type LogFormatter struct {
	// Set to true to bypass checking for a TTY before outputting colors.
	ForceColors   bool
	DisableColors bool
	// Set to true to disable timestamp logging (useful when the output
	// is redirected to a logging system already adding a timestamp)
}

func (f *LogFormatter) Format(entry *Entry) ([]byte, error) {
	var keys []string = make([]string, 0, len(entry.Data))
	for k := range entry.Data {
		keys = append(keys, k)
	}
	sort.Strings(keys)

	b := &bytes.Buffer{}

	prefixFieldClashes(entry.Data)

	isColored := (f.ForceColors || isTerminal) && !f.DisableColors

	if isColored {
		printColoredWithTime(b, entry, keys)
	} else {
		f.appendKeyValue(b, "time", entry.Time.Format(time.RFC3339))
		f.appendKeyValue(b, "level", entry.Level.String())
		f.appendKeyValue(b, "msg", entry.Message)
		for _, key := range keys {
			f.appendKeyValue(b, key, entry.Data[key])
		}
	}

	b.WriteByte('\n')
	return b.Bytes(), nil
}

func (f *LogFormatter) appendKeyValue(b *bytes.Buffer, key, value interface{}) {
	switch value.(type) {
	case string:
		if needsQuoting(value.(string)) {
			fmt.Fprintf(b, "%v=%s ", key, value)
		} else {
			fmt.Fprintf(b, "%v=%q ", key, value)
		}
	case error:
		if needsQuoting(value.(error).Error()) {
			fmt.Fprintf(b, "%v=%s ", key, value)
		} else {
			fmt.Fprintf(b, "%v=%q ", key, value)
		}
	default:
		fmt.Fprintf(b, "%v=%v ", key, value)
	}
}

func fullTime(t time.Time) string {
	return t.Format(time.Stamp)
}

func printColoredWithTime(b *bytes.Buffer, entry *Entry, keys []string) {
	var levelColor int
	switch entry.Level {
	case WarnLevel:
		levelColor = yellow
	case ErrorLevel, FatalLevel, PanicLevel:
		levelColor = red
	default:
		levelColor = blue
	}

	levelText := strings.ToUpper(entry.Level.String())[0:4]

	fmt.Fprintf(b, "\x1b[%dm%s\x1b[0m[%15s] %-44s ", levelColor, levelText, fullTime(entry.Time), entry.Message)
	for _, k := range keys {
		v := entry.Data[k]
		fmt.Fprintf(b, " \x1b[%dm%s\x1b[0m=%v", levelColor, k, v)
	}
}
