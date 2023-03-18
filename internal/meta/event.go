package meta

// Meta event notify event structure

type Event struct {
	Entries []Entry `json:"entry"`
}

type Entry struct {
	Changes []Change `json:"changes"`
}

type Change struct {
	Value Value `json:"value"`
}

type Value struct {
	Messages []Message `json:"messages"`
}

type Message struct {
	Text Text `json:"text"`
}

type Text struct {
	Body string `json:"body"`
}

func (e Event) TextBody() string {
	return e.Entries[0].Changes[0].Value.Messages[0].Text.Body
}