package message

// Message is an abstract of an Email Message
type Message struct {
	From     string
	FromName string
	To       string
	ToName   string
	Subject  string
	HTML     string
	Type     string
}
