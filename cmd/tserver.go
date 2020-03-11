package cmd

import (
	"context"
	"fmt"
	"os"
	"os/signal"
	"time"

	"github.com/goexam/internal/controllers"

	"github.com/labstack/echo"

	"github.com/labstack/echo/middleware"
	"github.com/spf13/cobra"
)

var teacherServerCmd = &cobra.Command{
	Use:   "t-server",
	Short: "t-server",
	Long:  `t-server`,
	Run: func(cmd *cobra.Command, args []string) {
		opts := loadApplocationOps()
		boot := newBootStrap(opts)
		problemCtrl := controllers.NewProblemController(boot.ProblemSvc)
		classCtrl := controllers.NewClassController(boot.ClassSvc)
		userCtrl := controllers.NewUserController(boot.UserSvc)
		courseCtrl := controllers.NewCourseController(boot.CourseSvc)
		contentCtrl := controllers.NewContentController(boot.ContentSvc)
		e := echo.New()
		e.Use(middleware.Logger())
		v1 := e.Group("/v1")
		problem := v1.Group("/problem")
		{
			problem.GET("/:id", problemCtrl.Get)
			problem.GET("/s", problemCtrl.GetList)
			problem.POST("", problemCtrl.Create)
			problem.PUT("/:id", problemCtrl.Update)
			problem.DELETE("/:id", problemCtrl.Delele)
		}

		class := v1.Group("/class")
		{
			class.GET("/:id", classCtrl.Get)
			class.DELETE("/:id", classCtrl.Delete)
			class.POST("", classCtrl.Create)
			class.PUT("/:id", classCtrl.Update)
			class.GET("/s", classCtrl.GetList)
		}

		user := v1.Group("/user")
		{
			user.GET("/:id", userCtrl.Get)
			user.GET("/s", userCtrl.GetList)
			user.POST("", userCtrl.Create)
			user.PUT("/:id", userCtrl.Update)
			user.DELETE("/:id", userCtrl.Delete)
			user.POST("/login", userCtrl.Login)
		}

		course := v1.Group("/course")
		{
			course.GET("/:id", courseCtrl.Get)
			course.GET("/s", courseCtrl.GetList)
			course.POST("", courseCtrl.Create)
			course.PUT("/:id", courseCtrl.Update)
			course.DELETE("/:id", courseCtrl.Delele)
		}

		content := v1.Group("/content")
		{
			content.GET("/:id", contentCtrl.Get)
			content.GET("/s", contentCtrl.GetList)
			content.POST("", contentCtrl.Create)
			content.PUT("/:id", contentCtrl.Update)
			content.DELETE("/:id", contentCtrl.Delete)
		}
		go func() {
			address := fmt.Sprintf("%s:%d", opts.Server.TServer.Host, opts.Server.TServer.Port)
			e.Start(address)
		}()
		quit := make(chan os.Signal, 1)
		signal.Notify(quit, os.Interrupt, os.Kill)
		<-quit

		ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
		defer cancel()
		if err := e.Shutdown(ctx); err != nil {
			e.Logger.Fatal(err)
		}
	},
}

func init() {
	rootCmd.AddCommand(teacherServerCmd)
}

// TeacherOps is ...
type TeacherOps = ServerOptions
