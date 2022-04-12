package domain

type Answer struct {
	Author User
	Text   string
	Topic  Topic
}
