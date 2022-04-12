package parser

import (
	"quand/app/domain"
	rep "quand/app/repository"
	"sync"
)

type ParserInterface interface {
	run() []domain.IQuestion
}

func ParseQuestions(parsers []ParserInterface) {
	wg := sync.WaitGroup{}
	for _, parser := range parsers {
		p := parser
		go func() {
			wg.Add(1)
			rep.NewQuestions(p.run())
			wg.Done()
		}()
	}
	wg.Wait()
}
