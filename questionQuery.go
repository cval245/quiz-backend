package main

import "context"

func getQuestionId(id int) QuestionFull {
	var question QuestionFull
	conn := connectDB()
	defer conn.Close(context.Background())
	allRows, err := conn.Query(context.Background(),
		"select answer.id, answer_text,  q.id, question_text, quiz_id from answer full join question q on q.id = answer.question_id where q.id = $1 order by q.id;", id)
	if err != nil {
		panic(err)
	}

	var answers []Answer
	for allRows.Next() {
		var quiz_id *int
		var answer_id *int
		var answer_text *string
		var question_id *int
		var question_text *string
		var answer Answer
		err = allRows.Scan(&answer_id, &answer_text, &question_id, &question_text, &quiz_id)
		if answer_id != nil {
			answer.Id = *answer_id
		}
		if answer_text != nil {
			answer.Text = *answer_text
		}

		if question_id != nil {
			question.Id = *question_id
			answer.Question_id = question.Id
		}
		if question_text != nil {
			question.Text = *question_text
		}
		if quiz_id != nil {
			question.Quiz_id = *quiz_id
		}
		if answer_id != nil {
			answers = append(answers, answer)
			question.Answers = answers
		}
	}
	return question
}
