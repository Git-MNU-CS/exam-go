package cmd

import (
	"log"

	"github.com/MNU/exam-go"
	"github.com/MNU/exam-go/internal/services"
)

type bootStrap struct {
	UserSvc           goexam.UserService
	ClassSvc          goexam.ClassService
	ContentAuthSvc    goexam.ContentAuthService
	ContentProblemSvc goexam.ContentProblemService
	ContentSvc        goexam.ContentService
	CourseSvc         goexam.CourseService
	ProblemSvc        goexam.ProblemService
	RecordSvc         goexam.RecordService
}

func newBootStrap(opts *ApplicationOps) *bootStrap {
	//logger := log.New(opts.Logging, os.Stdout)
	db, err := services.NewDatabase(opts.Database)
	if err != nil {
		log.Panic(err)
	}
	classSvc := services.NewClassService(db)
	userSvc := services.NewUserService(db)
	courseSvc := services.NewCourseService(db)
	problemSvc := services.NewProblemService(db)
	contentAuthSvc := services.NewContentAuthService(db)
	contentProblemSvc := services.NewContentProblemService(db, problemSvc)
	contentSvc := services.NewContentService(db, problemSvc, contentProblemSvc)
	recordSvc := services.NewRecordService(db)
	return &bootStrap{
		UserSvc:           userSvc,
		ClassSvc:          classSvc,
		ContentSvc:        contentSvc,
		ContentAuthSvc:    contentAuthSvc,
		ContentProblemSvc: contentProblemSvc,
		CourseSvc:         courseSvc,
		ProblemSvc:        problemSvc,
		RecordSvc:         recordSvc,
	}
}
