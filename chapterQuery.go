package main

import "context"

func getChapterById(id int) ChapterFull {
	conn := connectDB()
	defer conn.Close(context.Background())
	allRows, err := conn.Query(context.Background(),
		"select section.id, section.name, c.id, c.name, course_id "+
			"from section "+
			"join chapter c on section.chapter_id = c.id "+
			"where chapter_id = $1; ", id)
	if err != nil {
		// handle this error better than this
		panic(err)
	}
	//var quizFuller QuizFull
	var chapter ChapterFull
	var sections []SectionSansChapterId

	for allRows.Next() {
		var section SectionSansChapterId
		var section_id *int
		var section_name *string
		var chapter_id *int
		var chapter_name *string
		var course_id *int

		err = allRows.Scan(&section_id, &section_name, &chapter_id, &chapter_name, &course_id)
		if section_id != nil {
			section.Id = *section_id
		}
		if section_name != nil {
			section.Name = *section_name
		}
		if chapter_id != nil {
			chapter.Id = *chapter_id
		}
		if chapter_name != nil {
			chapter.Name = *chapter_name
		}
		if course_id != nil {
			chapter.Course_id = *course_id
		}
		sections = append(sections, section)
	}
	chapter.Sections = sections
	return chapter
}
