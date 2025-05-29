package server

import (
	_checkinController "github.com/jaayroots/habit-tracker-api/pkg/checkin/controller"
	_checkinRepository "github.com/jaayroots/habit-tracker-api/pkg/checkin/repository"
	_checkinService "github.com/jaayroots/habit-tracker-api/pkg/checkin/service"
	_habitRepository "github.com/jaayroots/habit-tracker-api/pkg/habit/repository"
	_userRepository "github.com/jaayroots/habit-tracker-api/pkg/user/repository"
)

func (s *echoServer) checkinRouter(m *authorizingMiddleware) {
	router := s.app.Group("/checkin")

	checkinRepository := _checkinRepository.NewCheckinRepositoryImpl(s.db, s.app.Logger)
	userRepository := _userRepository.NewUserRepositoryImpl(s.db, s.app.Logger)
	habitRepository := _habitRepository.NewHabitRepositoryImpl(s.db, s.app.Logger)

	checkinService := _checkinService.NewCheckinServiceImpl(habitRepository, checkinRepository, userRepository)
	checkinController := _checkinController.NewCheckinControllerImpl(checkinService)

	router.POST("", checkinController.Create, m.Authorizing)
	router.POST("/search", checkinController.FindAll, m.Authorizing)

}
