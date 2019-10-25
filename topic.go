package main

import(
  "fmt"
  "github.com/go-pg/pg"
)

type Topic struct {
  ID        int
  Name      string
  Questions []*Question
}

type TopicRepo interface {
  Find(ID int) (error, *Topic)
}

type PGRepo struct {
  db *pg.DB
}

func (repo *PGRepo) Find(topicID int) (error, *Topic) {
  topic := &Topic{ID: topicID}
  err := repo.db.Model(topic).
                 Relation("Questions").
                 Where("topic.id = ?", topic.ID).
                 Select()
  fmt.Println(topic)
  return err, topic
}

func NewPGRepo(db *pg.DB) TopicRepo {
  return &PGRepo{
    db: db,
  }
}
