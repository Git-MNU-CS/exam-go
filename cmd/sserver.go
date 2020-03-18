package cmd

import (
	"context"
	"fmt"
	studentController "github.com/MNU/exam-go/internal/controllers/student"
	"github.com/labstack/echo"
	"github.com/labstack/echo/middleware"
	"github.com/spf13/cobra"
	"os"
	"os/signal"
	"time"
)

// StudentOps is ...
type StudentOps = ServerOptions

var studentServerCmd = &cobra.Command{
	Use:   "s-server",
	Short: "s-server",
	Long:  `s-server`,
	Run: func(cmd *cobra.Command, args []string) {
		opts := loadApplicationOps()
		boot := newBootStrap(opts)
		userCtrl := studentController.NewUserController(boot.UserSvc, boot.ClassSvc)
		contentCtrl := studentController.NewContentController(boot.ContentSvc, boot.ProblemSvc, boot.UserSvc)
		recordCtrl := studentController.NewRecordController(boot.UserSvc, boot.ProblemSvc, boot.ContentSvc, boot.RecordSvc)
		e := echo.New()
		e.Use(middleware.Logger())
		v1 := e.Group("/v1")
		// down 除了登陆
		user := v1.Group("/user")
		{
			user.GET("/:id", userCtrl.Get)
			user.POST("/login", userCtrl.Login)
			user.POST("/logout", userCtrl.Logout)
		}

		content := v1.Group("/content")
		{
			content.GET("/:id", contentCtrl.Get)
			content.GET("/s", contentCtrl.GetList)
		}

		record := v1.Group("/record")
		{
			record.GET("/:id", recordCtrl.Get)
			record.POST("/submit", recordCtrl.Submit)
			record.PUT("/:id/changeResult", recordCtrl.ChangeResult)
		}

		go func() {
			address := fmt.Sprintf("%s:%d", opts.Server.SServer.Host, opts.Server.SServer.Port)
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
	rootCmd.AddCommand(studentServerCmd)
}
