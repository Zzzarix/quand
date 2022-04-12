package app

import (
	"log"
	"quand/pkg/parser"
	"time"
)

func timelyParsing(parsers ...parser.ParserInterface) {
	for {
		func() {
			defer func() {
				if err := recover(); err != nil {
					log.Fatal(err)
				}
			}()
			time.Sleep(time.Hour * 10)
			parser.ParseQuestions(parsers)
		}()
	}
}

//func timelyClearing(db *mongo.Client) {
//	for {
//		func() {
//			defer func() {
//				if err := recover(); err != nil {
//					log.Fatal(err)
//				}
//			}()
//			parser.ParseQuestions(parsers)
//			time.Sleep(time.Minute * 10)
//		}()
//	}
//}

func timelyTodayQuestion() {
	for {
		func() {
			defer func() {
				if err := recover(); err != nil {
					log.Fatal(err)
				}
			}()
			time.Sleep(time.Hour * 24)
		}()
	}
}
