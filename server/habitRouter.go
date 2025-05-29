package server

import (
	_checkinRepository "github.com/jaayroots/habit-tracker-api/pkg/checkin/repository"
	_habitController "github.com/jaayroots/habit-tracker-api/pkg/habit/controller"
	_habitRepository "github.com/jaayroots/habit-tracker-api/pkg/habit/repository"
	_habitService "github.com/jaayroots/habit-tracker-api/pkg/habit/service"
	_userRepository "github.com/jaayroots/habit-tracker-api/pkg/user/repository"
)

func (s *echoServer) habitRouter(m *authorizingMiddleware) {
	router := s.app.Group("/habit")

	habitRepository := _habitRepository.NewHabitRepositoryImpl(s.db, s.app.Logger)
	userRepository := _userRepository.NewUserRepositoryImpl(s.db, s.app.Logger)
	checkinRepository := _checkinRepository.NewCheckinRepositoryImpl(s.db, s.app.Logger)

	habitService := _habitService.NewHabitServiceImpl(habitRepository, userRepository, checkinRepository)
	habitController := _habitController.NewHabitControllerImpl(habitService)

	router.POST("", habitController.Create, m.Authorizing)
	router.GET("/:habitID", habitController.FindByID, m.Authorizing)
	router.PATCH("/:habitID", habitController.Update, m.Authorizing)
	router.DELETE("/:habitID", habitController.Delete, m.Authorizing)
	router.POST("/search", habitController.FindAll, m.Authorizing)

}
