package main

type Quiz struct {
	Name       string `json:"name"`
	Id         int    `json:"id"`
	Chapter_id int    `json:"chapter_id"`
}
type QuizFull struct {
	Name       string                `json:"name"`
	Id         int                   `json:"id"`
	Chapter_id int                   `json:"chapter_id"`
	Questions  []QuestionFullSansQId `json:"questions"`
}

type Question struct {
	Id      int    `json:"id"`
	Text    string `json:"text"`
	Quiz_id int    `json:"quiz_id"`
}

type QuestionFull struct {
	Id      int      `json:"id"`
	Text    string   `json:"text"`
	Quiz_id int      `json:"quiz_id"`
	Answers []Answer `json:"answers"`
}
type QuestionFullSansQId struct {
	Id      int                 `json:"id"`
	Text    string              `json:"text"`
	Answers []AnswerFullSansQId `json:"answers"`
}

type Answer struct {
	Id          int    `json:"id"`
	Text        string `json:"text"`
	Question_id int    `json:"question_id"`
	Correct     bool   `json:"correct"`
}
type AnswerFullSansQId struct {
	Id      int    `json:"id"`
	Text    string `json:"text"`
	Correct bool   `json:"correct"`
}
