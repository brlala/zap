package benchmarks

import (
	"io"

	"github.com/ssgreg/logf"
)

type DiscardEntryWriter struct{}

func (d DiscardEntryWriter) WriteEntry(e logf.Entry) {
	io.Discard.Write([]byte{})
}

func newDisabledSsgreg() *logf.Logger {
	// Create an instance of your custom EntryWriter
	discardWriter := DiscardEntryWriter{}

	// Initialize the logger
	logger := logf.NewLogger(logf.LevelError, discardWriter) // Adjust the level as needed
	return logger
}

func fakeSsGregFields() []logf.Field {
	return []logf.Field{
		logf.Int("int", _tenInts[0]),
		logf.Ints("ints", _tenInts),
		logf.String("string", _tenStrings[0]),
		logf.Strings("strings", _tenStrings),
		logf.Time("time", _tenTimes[0]),
		logf.Any("times", _tenTimes),
		logf.Any("user1", _oneUser),
		logf.Any("user2", _oneUser),
		logf.Any("users", _tenUsers),
		logf.Error(errExample),
	}
}
