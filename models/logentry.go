package models

import (
	"time"
)

type LogEntry struct {
	BookID             uint
	Book               Book
	StartDate          *time.Time
	EndDate            *time.Time
	ProgressPercentage float32
	ProgressPage       uint
	ProgressDate       time.Time
	ExcludeFromStats   bool
}

func (LogEntry) TableName() string {
	return "library_logentry"
}

func (entry LogEntry) String() (result string) {
	result = entry.Book.FirstAuthor.DisplayName() + ", " + entry.Book.FullTitle() + ", "

	if entry.StartDate != nil {
		result += "from " + entry.StartDate.Format("2006-01-02") + ", "
	}

	if entry.EndDate != nil {
		result += "to " + entry.EndDate.Format("2006-01-02")
	} else {
		result += "unfinished"
	}

	return
}

func (entry *LogEntry) StartNow() {
	now := time.Now()
	entry.StartAt(&now)
}

func (entry *LogEntry) StartAt(at *time.Time) {
	entry.StartDate = at
}

func (entry *LogEntry) EndNow() {
	now := time.Now()
	entry.EndAt(&now)
}

func (entry *LogEntry) EndAt(at *time.Time) {
	entry.EndDate = at
}
