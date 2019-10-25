package main

type Question struct {
  ID      int
  Content string
  TopicID int
  Topic   *Topic
  Answers []*Answer
}

type QuestionRepo interface {
  Find(ID int) *Question
}
