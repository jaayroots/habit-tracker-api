package server

import (
	_controller "github.com/jaayroots/habit-tracker-api/controller"
	_service "github.com/jaayroots/habit-tracker-api/service"
	_repository "github.com/jaayroots/habit-tracker-api/repository"
)

func (s *echoServer) checkinRouter(m *authorizingMiddleware) {
	router := s.app.Group("/checkin")

	checkinRepository := _repository.NewCheckinRepositoryImpl(s.db, s.app.Logger)
	userRepository := _repository.NewUserRepository(s.db, s.app.Logger)
	habitRepository := _repository.NewHabitRepositoryImpl(s.db, s.app.Logger)

	checkinService := _service.NewCheckinServiceImpl(habitRepository, checkinRepository, userRepository)
	checkinController := _controller.NewCheckinControllerImpl(checkinService)

	router.POST("", checkinController.Create, m.Authorizing)
	router.POST("/search", checkinController.FindAll, m.Authorizing)

}
