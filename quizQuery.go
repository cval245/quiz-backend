package main

import (
	"context"
	"fmt"
	"github.com/jackc/pgx/v4"
	"os"
)

func connectDB() *pgx.Conn {
	conn, err := pgx.Connect(context.Background(), "user=djangoconnect password=Belgrade2010 host=localhost port=5432 dbname=quizbackend")
	if err != nil {
		fmt.Fprintf(os.Stderr, "Unable to connect to database: %v\n", err)
		os.Exit(1)
	}
	return conn
}

func getQuiz(id int) QuizFull {
	conn := connectDB()
	defer conn.Close(context.Background())
	var quizFull QuizFull
	allRows, err := conn.Query(context.Background(),
		"select answer.id, answer_text, correct, q.id, question_text, quiz_id, name "+
			"from answer full join question q on q.id = answer.question_id join quiz q2 on q2.id = q.quiz_id "+
			"where quiz_id = $1 order by question_id;", id)
	if err != nil {
		// handle this error better than this
		panic(err)
	}
	//var quizFuller QuizFull
	var questions []QuestionFullSansQId
	var answers []AnswerFullSansQId
	var oldQuestion QuestionFullSansQId
	for allRows.Next() {
		var answer AnswerFullSansQId
		var question QuestionFullSansQId
		var answer_id *int
		var answer_text *string
		var answer_correct *bool
		var question_id *int
		var question_text *string
		var quizFuller_id *int
		var quizFuller_name *string
		//var answer_id int
		//err = allRows.Scan(&answer.Id, &answer.Text, &question.Id, &question.Text, &quizFuller.Id)
		err = allRows.Scan(&answer_id, &answer_text, &answer_correct, &question_id, &question_text,
			&quizFuller_id, &quizFuller_name)
		if *question_id == oldQuestion.Id {
			if answer_id != nil {
				answer.Id = *answer_id
			}
			if answer_text != nil {
				answer.Text = *answer_text
			}
			if answer_correct != nil {
				answer.Correct = *answer_correct
			}
			answers = append(answers, answer)
			oldQuestion.Answers = answers
			questions[len(questions)-1] = oldQuestion
		} else {
			answers = nil
			if answer_id != nil {
				answer.Id = *answer_id
			}
			if answer_text != nil {
				answer.Text = *answer_text
			}
			if answer_correct != nil {
				answer.Correct = *answer_correct
			}
			if answer.Id != 0 {
				answers = append(answers, answer)
			}
			if question_id != nil {
				question.Id = *question_id
			}
			if question_text != nil {
				question.Text = *question_text
			}
			question.Answers = answers
			questions = append(questions, question)
			oldQuestion = question
		}
		quizFull.Id = *quizFuller_id
		quizFull.Name = *quizFuller_name
		if err != nil {
			// handle this error
			panic(err)
		}
	}
	quizFull.Questions = questions
	return quizFull
}
func getAllQuizzes() []Quiz {
	conn, err := pgx.Connect(context.Background(), "user=djangoconnect password=Belgrade2010 host=localhost port=5432 dbname=quizbackend")
	var quizzes []Quiz
	if err != nil {
		fmt.Fprintf(os.Stderr, "Unable to connect to database: %v\n", err)
		os.Exit(1)
	}
	defer conn.Close(context.Background())
	rows, err := conn.Query(context.Background(), "select * from quiz;")
	if err != nil {
		// handle this error better than this
		panic(err)
	}
	defer rows.Close()
	for rows.Next() {
		var quiz Quiz
		err = rows.Scan(&quiz.Id, &quiz.Name, &quiz.Chapter_id)
		if err != nil {
			// handle this error
			panic(err)
		}
		quizzes = append(quizzes, quiz)
	}
	// get any error encountered during iteration
	err = rows.Err()
	if err != nil {
		panic(err)
	}
	if err != nil {
		fmt.Fprintf(os.Stderr, "QueryRow failed: %v\n", err)
		os.Exit(1)
	}
	return quizzes
}
