package domain

type Topic struct {
	Title   string
	Author  User
	Answers []Answer
	Votes   int64
}
