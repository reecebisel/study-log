package main

type Answer struct {
  ID         int
  Content    string
  QuestionID int
  Question   *Question
  Correct    bool
}

type AnswerRepo interface {
  Find(ID int) *Answer
}
