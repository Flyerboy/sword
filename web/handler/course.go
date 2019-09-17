package handler

import (
	"discuss/service"
	"github.com/gin-gonic/gin"
	"net/http"
	"strconv"
)

var CourseHandler = &courseHandler{}

type courseHandler struct {
}

func (this *courseHandler) Index(c *gin.Context) {
	data := getCommonData(c)
	data["Title"] = "课程"
	data["Courses"], _ = service.CourseService.Select(0, 10)
	c.HTML(http.StatusOK, "course/index.html", data)
}

func (this *courseHandler) Detail(c *gin.Context) {
	id := c.Param("id")
	courseId, _ := strconv.Atoi(id)

	course, err := service.CourseService.Detail(courseId)
	if err != nil {

	} else {
		data := getCommonData(c)
		data["Title"] = "课程 - " + course.Title
		data["Course"] = course

		// 目录列表
		chapters, _ := service.CourseService.Chapters(courseId)
		data["Chapters"] = chapters

		// 章节详情
		chapterId := c.DefaultQuery("chapter_id", "0")
		newChapterId, _ := strconv.Atoi(chapterId)
		data["Chapter"], _ = service.CourseService.Chapter(courseId, newChapterId)
		c.HTML(http.StatusOK, "course/detail.html", data)
	}
}
