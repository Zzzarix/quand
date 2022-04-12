package parser

import (
	"context"
	"log"
	"quand/app/domain"

	"github.com/vartanbeno/go-reddit/v2/reddit"
)

type RedditParser struct {
	ID       string `json:"id"`
	Secret   string `json:"secret"`
	Username string `json:"username"`
	Password string `json:"password"`
}

var ctx = context.Background()

func (rp RedditParser) run() []domain.IQuestion {
	var quests []domain.IQuestion

	credentials := reddit.Credentials{
		ID:       rp.ID,
		Secret:   rp.Secret,
		Username: rp.Username,
		Password: rp.Password,
	}

	rd, err := reddit.NewClient(credentials)
	if err != nil {
		log.Fatal(err)
	}

	posts, _, err := rd.Subreddit.NewPosts(ctx, "AskReddit+AskWomen", &reddit.ListOptions{Limit: 100})
	if err != nil {
		log.Fatal(err)
	}

	for _, p := range posts {
		quests = append(quests, domain.IQuestion{
			Text:   p.Title,
			Kind:   domain.JustAsk,
			Author: p.Author,
		})
	}

	return quests
}
