package cmd

import (
	"log"

	"github.com/MNU/exam-go"
	"github.com/MNU/exam-go/internal/services"
)

type bootStrap struct {
	UserSvc    goexam.UserService
	ClassSvc   goexam.ClassService
	ContentSvc goexam.ContentService
	CourseSvc  goexam.CourseService
	ProblemSvc goexam.ProblemService
	RecordSvc  goexam.RecordService
	CollageSvc goexam.CollageService
}

func newBootStrap(opts *ApplicationOps) *bootStrap {
	//logger := log.New(opts.Logging, os.Stdout)
	db, err := services.NewDatabase(opts.Database)
	db.LogMode(true)
	if err != nil {
		log.Panic(err)
	}
	classSvc := services.NewClassService(db)
	userSvc := services.NewUserService(db)
	courseSvc := services.NewCourseService(db)
	problemSvc := services.NewProblemService(db)
	contentSvc := services.NewContentService(db)
	recordSvc := services.NewRecordService(db)
	collageSvc := services.NewCollageService(db)
	return &bootStrap{
		UserSvc:    userSvc,
		ClassSvc:   classSvc,
		ContentSvc: contentSvc,
		CourseSvc:  courseSvc,
		ProblemSvc: problemSvc,
		RecordSvc:  recordSvc,
		CollageSvc: collageSvc,
	}
}
