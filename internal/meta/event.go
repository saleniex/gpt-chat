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
	Contacts []Contact `json:"contacts"`
	Messages []Message `json:"messages"`
}

type Message struct {
	Text Text `json:"text"`
}

type Contact struct {
	Profile Profile `json:"profile"`
	WaId    string  `json:"wa_id"`
}

type Profile struct {
	Name string `json:"name"`
}

type Text struct {
	Body string `json:"body"`
}

func (e Event) TextBody() string {
	if e.firstMessage() == nil {
		return ""
	}
	return e.firstMessage().Text.Body
}

func (e Event) ProfileName() string {
	if e.firstContact() == nil {
		return ""
	}
	return e.firstContact().Profile.Name
}

func (e Event) WhatsappId() string {
	if e.firstContact() == nil {
		return ""
	}
	return e.firstContact().WaId
}

func (e Event) firstMessage() *Message {
	if len(e.Entries) == 0 {
		return nil
	}
	if len(e.Entries[0].Changes) == 0 {
		return nil
	}
	if len(e.Entries[0].Changes[0].Value.Messages) == 0 {
		return nil
	}
	return &e.Entries[0].Changes[0].Value.Messages[0]
}

func (e Event) firstContact() *Contact {
	if len(e.Entries) == 0 {
		return nil
	}
	if len(e.Entries[0].Changes) == 0 {
		return nil
	}
	if len(e.Entries[0].Changes[0].Value.Contacts) == 0 {
		return nil
	}
	return &e.Entries[0].Changes[0].Value.Contacts[0]
}
