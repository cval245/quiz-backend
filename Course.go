package main

type Course struct {
	Name string `json:"name"`
	Id   int    `json:"id"`
}

type CourseFull struct {
	Name     string                `json:"name"`
	Id       int                   `json:"id"`
	Chapters []ChapterSansCourseID `json:"chapters"`
}

type ChapterSansCourseID struct {
	Name string `json:"name"`
	Id   int    `json:"id"`
}

type Chapter struct {
	Name      string `json:"name"`
	Course_id int    `json:"course_id"`
	Id        int    `json:"id"`
}

type ChapterFull struct {
	Name      string                 `json:"name"`
	Course_id int                    `json:"course_id"`
	Id        int                    `json:"id"`
	Sections  []SectionSansChapterId `json:"sections"`
}

type Section struct {
	Name       string `json:"name"`
	Chapter_id int    `json:"chapter_id"`
	Id         int    `json:"id"`
}

type SectionSansChapterId struct {
	Name string `json:"name"`
	Id   int    `json:"id"`
}
