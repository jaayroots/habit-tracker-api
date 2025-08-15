package server

import (
	_controller "github.com/jaayroots/habit-tracker-api/controller"
	_repository "github.com/jaayroots/habit-tracker-api/repository"
	_service "github.com/jaayroots/habit-tracker-api/service"
)

func (s *echoServer) usersRouter(m *authorizingMiddleware) {
	router := s.app.Group("/user")

	userRepository := _repository.NewUserRepository(s.db, s.app.Logger)
	userContactRepository := _repository.NewUserContactRepository(s.db, s.app.Logger)
	userService := _service.NewUserService(s.db, userRepository, userContactRepository)
	userController := _controller.NewUserController(userService)

	router.GET("/:userID", userController.FindByID, m.Authorizing)
	router.PATCH("/:userID", userController.Update, m.Authorizing)
	router.DELETE("/:userID", userController.Delete, m.Authorizing)
}
