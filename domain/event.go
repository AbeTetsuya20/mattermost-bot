package domain

// 日付とイベント内容と行数の情報を持つ
type Event struct {
	Date    string
	Content string
	Line    int
}

func NewEvent(date string, content string, line int) *Event {
	return &Event{
		Date:    date,
		Content: content,
		Line:    line,
	}
}

type Events []*Event
