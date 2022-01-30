package main

//package quiz_backend

import (
	"context"
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"regexp"
	"strconv"
	"strings"
)

type route struct {
	method  string
	regex   *regexp.Regexp
	handler http.HandlerFunc
}

var routes = []route{
	newRoute("GET", "/quiz/", returnAllQuizzes),
	newRoute("GET", "/quiz/([0-9]+)", returnFullQuiz),
	newRoute("GET", "/question/", returnAllQuestions),
	newRoute("GET", "/question/([0-9]+)", returnFullQuestion),
	newRoute("GET", "/answer/", returnAllAnswers),
	newRoute("GET", "/course/", returnAllCourses),
	newRoute("GET", "/course/([0-9]+)", returnCourseById),
	newRoute("GET", "/chapter/([0-9]+)", returnChapterById),
}

func enableCors(w *http.ResponseWriter) {
	(*w).Header().Set("Access-Control-Allow-Origin", "*")
}

func newRoute(method, pattern string, handler http.HandlerFunc) route {
	return route{method, regexp.MustCompile("^" + pattern + "$"), handler}
}

type ctxKey struct{}

func getField(r *http.Request, index int) string {
	fields := r.Context().Value(ctxKey{}).([]string)
	return fields[index]
}

func Serve(w http.ResponseWriter, r *http.Request) {
	var allow []string
	enableCors(&w)
	for _, route := range routes {
		matches := route.regex.FindStringSubmatch(r.URL.Path)
		if len(matches) > 0 {
			if r.Method != route.method {
				allow = append(allow, route.method)
				continue
			}
			ctx := context.WithValue(r.Context(), ctxKey{}, matches[1:])
			route.handler(w, r.WithContext(ctx))
			return
		}
	}
	if len(allow) > 0 {
		w.Header().Set("Allow", strings.Join(allow, ", "))
		http.Error(w, "405 method not allowed", http.StatusMethodNotAllowed)
		return
	}
	http.NotFound(w, r)
}
func main() {
	log.Fatal(http.ListenAndServe(":8080", http.HandlerFunc(Serve)))
}

func returnFullQuiz(writer http.ResponseWriter, request *http.Request) {
	id, _ := strconv.Atoi(getField(request, 0))
	fullQuiz := getQuiz(id)
	json.NewEncoder(writer).Encode(fullQuiz)
}

func returnFullQuestion(writer http.ResponseWriter, request *http.Request) {
	id, _ := strconv.Atoi(getField(request, 0))
	fullQuestion := getQuestionId(id)
	json.NewEncoder(writer).Encode(fullQuestion)
}

func returnAllQuizzes(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	quizzes := getAllQuizzes()
	json.NewEncoder(w).Encode(quizzes)
}
func returnAllQuestions(w http.ResponseWriter, r *http.Request) {
	quizzes := getAllQuizzes()
	json.NewEncoder(w).Encode(quizzes)
}
func returnAllAnswers(w http.ResponseWriter, r *http.Request) {
	quizzes := getAllQuizzes()
	json.NewEncoder(w).Encode(quizzes)
}
func returnAllCourses(w http.ResponseWriter, r *http.Request) {
	courses := getAllCourses()
	json.NewEncoder(w).Encode(courses)
}
func returnCourseById(w http.ResponseWriter, r *http.Request) {
	id, _ := strconv.Atoi(getField(r, 0))
	course := getCourseById(id)
	json.NewEncoder(w).Encode(course)
}
func returnChapterById(w http.ResponseWriter, r *http.Request) {
	id, _ := strconv.Atoi(getField(r, 0))
	chapter := getChapterById(id)
	json.NewEncoder(w).Encode(chapter)
}

func getAll(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	var quizzes []Quiz
	quiz := Quiz{Id: 1, Name: "name"}

	quizzes = append(quizzes, quiz)
	fmt.Println(quizzes)
	err := json.NewEncoder(w).Encode(quizzes)
	if err != nil {
		//e := Error{Message: "Internal Server Error"}
		e := 0
		w.WriteHeader(http.StatusInternalServerError)
		_ = json.NewEncoder(w).Encode(e)
	}
}
