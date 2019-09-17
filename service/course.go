package service

import "discuss/model"

var CourseService = &courseService{}

type courseService struct {}

// 课程列表
func (this *courseService) Select(start, limit int) (courses []*model.Course, err error) {
	courses, err = model.Course{}.Select(start, limit)
	return
}

// 课程详情
func (this *courseService) Detail(id int) (course model.Course, err error) {
	course, err = model.Course{}.Find(id)
	return
}

// 章节列表
func (this *courseService) Chapters(courseId int) (chapters []*model.CourseChapter, err error) {
	chapters, err = model.CourseChapter{}.Chapter(courseId)
	return
}

// 章节详情
func (this *courseService) Chapter(courseId, chapterId int) (chapter model.CourseChapter, err error) {
	if chapterId > 0 {
		chapter, err = model.CourseChapter{}.Find(chapterId)
	} else {
		chapter, err = model.CourseChapter{}.First(courseId)
	}
	return
}

// 创建课程
func (this *courseService) Create(userId int, title string, introduction string) (result bool, err error) {
	// 处理封面图片

	// 写入数据库
	result, err = model.Course{
		UserId: userId,
		Title: title,
		Introduction: introduction,
	}.Create()

	// 生成缓存数据
	return
}
