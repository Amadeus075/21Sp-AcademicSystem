package model

import (
	"gorm.io/gorm"
)

type Course struct {
	ID     uint   `gorm:"primarykey"`
	Number string `json:"number" form:"number" gorm:"index" example:"0121"` // 课号
	Name   string `json:"name" form:"name" example:"数据库原理"`                 // 课名
	Credit uint8  `json:"credit" form:"credit" example:"4"`                 // 学分

	OfferedCourses []OfferedCourse `json:"offered_courses"`

	DepartmentID uint `json:"department_id" form:"department_id"` // 院系号
}

type CourseRes struct {
	ID             uint   `json:"id"`
	Number         string `json:"number"`
	Name           string `json:"name"`
	Credit         uint8  `json:"credit"`
	DepartmentName string `json:"department_name"` // 院系名
}

//// CourseRes ...
//type CourseRes struct {
//	ID          uint   `json:"id"`
//	Score       int    `json:"score"`        // 成绩
//	Number      string `json:"number"`       // 课号
//	Name        string `json:"name"`         // 课名
//	Credit      uint8  `json:"credit"`       // 学分
//	Department  string `json:"department"`   // 所属院系
//	Term        string `json:"term"`         // 学期
//	TeacherName string `json:"teacher_name"` // 老师名
//	StudentName string `json:"student_name"` // 学生名
//}

// BeforeDelete 删除课程时级联删除开课记录
func (c *Course) BeforeDelete(tx *gorm.DB) (err error) {
	_ = tx.Where("course_id = ?", c.ID).Delete(&OfferedCourse{}).RowsAffected
	return
}
