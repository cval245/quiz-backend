package main

import (
	"context"
	"fmt"
	"os"
)

func getCourseById(id int) CourseFull {
	conn := connectDB()
	defer conn.Close(context.Background())
	allRows, err := conn.Query(context.Background(),
		"select chapter.id, chapter.name, c.id, c.name "+
			"from chapter "+
			"join course c on c.id = chapter.course_id "+
			"where course_id = $1;", id)
	if err != nil {
		panic(err)
	}
	var course CourseFull
	var chapters []ChapterSansCourseID
	for allRows.Next() {
		var chapter_id *int
		var chapter_name *string
		var course_id *int
		var course_name *string
		var chapter ChapterSansCourseID

		err = allRows.Scan(&chapter_id, &chapter_name, &course_id, &course_name)
		if chapter_id != nil {
			chapter.Id = *chapter_id
		}
		if chapter_name != nil {
			chapter.Name = *chapter_name
		}
		if course_id != nil {
			course.Id = *course_id
		}
		if course_name != nil {
			course.Name = *course_name
		}
		chapters = append(chapters, chapter)
	}
	course.Chapters = chapters
	return course
}

func getAllCourses() []Course {
	conn := connectDB()
	var courses []Course

	defer conn.Close(context.Background())
	rows, err := conn.Query(context.Background(), "select * from course;")
	if err != nil {
		// handle this error better than this
		panic(err)
	}
	defer rows.Close()
	for rows.Next() {
		var course Course
		err = rows.Scan(&course.Id, &course.Name)
		if err != nil {
			// handle this error
			panic(err)
		}
		courses = append(courses, course)
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
	return courses
}
