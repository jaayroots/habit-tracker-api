package server

import (
	_controller "github.com/jaayroots/habit-tracker-api/controller"
	_repository "github.com/jaayroots/habit-tracker-api/repository"
	_service "github.com/jaayroots/habit-tracker-api/service"
)

func (s *echoServer) habitRouter(m *authorizingMiddleware) {
	router := s.app.Group("/habit")

	habitRepository := _repository.NewHabitRepositoryImpl(s.db, s.app.Logger)
	userRepository := _repository.NewUserRepository(s.db, s.app.Logger)
	checkinRepository := _repository.NewCheckinRepositoryImpl(s.db, s.app.Logger)

	habitService := _service.NewHabitServiceImpl(habitRepository, userRepository, checkinRepository)
	habitController := _controller.NewHabitControllerImpl(habitService)

	router.POST("", habitController.Create, m.Authorizing)
	router.GET("/:habitID", habitController.FindByID, m.Authorizing)
	router.PATCH("/:habitID", habitController.Update, m.Authorizing)
	router.DELETE("/:habitID", habitController.Delete, m.Authorizing)
	router.POST("/search", habitController.FindAll, m.Authorizing)

}
