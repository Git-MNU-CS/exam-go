package models

import (
	"github.com/goexam"
)

type (

	// ProblemResponse is
	ProblemResponse struct {
		ID         uint                 `json:"id"`
		Level      goexam.ProblemLevel  `json:"level"`
		Type       goexam.ProblemType   `json:"type"`
		Name       string               `json:"name"`
		CourseID   uint                 `json:"course_id"`
		CourseName string               `json:"course_name"`
		Describe   string               `json:"describe"`
		Status     goexam.ProblemStatus `json:"status"`
		CreatedAt  string               `json:"created_at"`
		UpdatedAt  string               `json:"updated_at"`
	}
)

// BuildResponse is
func (u *ProblemResponse) BuildResponse(problem *goexam.Problem) {
	*u = ProblemResponse{
		ID:         problem.ID,
		Level:      problem.Level,
		Type:       problem.Type,
		Name:       problem.Name,
		CourseID:   problem.CourseID,
		CourseName: problem.Course.Name,
		Describe:   problem.Describe,
		Status:     problem.Status,
		CreatedAt:  problem.CreatedAt.Format(TimeFormat),
		UpdatedAt:  problem.UpdatedAt.Format(TimeFormat),
	}
}
