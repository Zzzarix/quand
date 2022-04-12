package repository

import (
	"go.mongodb.org/mongo-driver/bson"
	"log"
	. "quand/app/domain"
	"quand/pkg/errors"

	"go.mongodb.org/mongo-driver/mongo"
)

type Db struct {
	Name     string `json:"name"`
	Password string `json:"password"`
	Username string `json:"username"`
}

type Storage struct {
	Users     *mongo.Collection
	Questions *mongo.Collection
	Topics    *mongo.Collection
}

var storage *Storage

func Init(dbUsers Db, dbQuests Db) {
	storage = &Storage{
		Users:     mongoClient(dbUsers).Database("RootDb").Collection("Users"),
		Questions: mongoClient(dbQuests).Database("RootDb").Collection("Questions"),
		Topics:    mongoClient(dbQuests).Database("RootDb").Collection("Topics"),
	}
}

func CheckUniqueEmail(email string) bool {
	r := storage.Users.FindOne(ctx, bson.D{{"email", email}})
	return r.Err() == mongo.ErrNoDocuments
}

func CheckUniqueUserName(name string) bool {
	r := storage.Users.FindOne(ctx, bson.D{{"userName", name}})
	return r.Err() == mongo.ErrNoDocuments
}

func NewUser(user User) error {
	if !CheckUniqueEmail(user.Email) {
		return errors.UserWithEmailAlreadyExists
	}
	if !CheckUniqueUserName(user.UserName) {
		return errors.UserWithNameAlreadyExists
	}

	_, err := storage.Users.InsertOne(ctx, user)
	return err
}

func UpdateUser(before, after User) error {
	_, err := storage.Users.UpdateOne(ctx, bson.D{{"userName", before.UserName}}, after)
	return err
}

func GetUser(name string) (User, error) {
	var user User
	r := storage.Users.FindOne(ctx, bson.D{{"userName", name}})
	if r.Err() != nil {
		return User{}, r.Err()
	}
	r.Decode(&user)
	return user, nil
}

func GetUserHistory(name string) (History, error) {
	user, err := GetUser(name)
	return user.History, err
}

func NewQuestion(quest IQuestion) {
	doc := Question{
		Text:   quest.Text,
		Kind:   quest.Kind,
		Author: quest.Author,
	}

	_, err := storage.Questions.InsertOne(ctx, doc)
	if err != nil {
		log.Fatal(err)
	}
}

func NewQuestions(quests []IQuestion) {
	var docs []interface{}
	for _, q := range quests {
		docs = append(docs, Question{
			Text:   q.Text,
			Kind:   q.Kind,
			Author: q.Author,
		})
	}

	_, err := storage.Questions.InsertMany(ctx, docs)
	if err != nil {
		log.Fatal(err)
	}
}

func GetQuestion() Question {
	var quest Question

	err := storage.Questions.FindOneAndDelete(ctx, bson.D{}).Decode(&quest)
	if err != nil {
		log.Fatal(err)
	}

	return quest
}
