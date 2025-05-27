package server

import (
	_userController "github.com/jaayroots/habit-tracker-api/pkg/user/controller"
	_userRepository "github.com/jaayroots/habit-tracker-api/pkg/user/repository"
	_userService "github.com/jaayroots/habit-tracker-api/pkg/user/service"
)

func (s *echoServer) usersRouter(m *authorizingMiddleware) {
	router := s.app.Group("/user")

	userRepository := _userRepository.NewUserRepositoryImpl(s.db, s.app.Logger)
	userService := _userService.NewUserServiceImpl(userRepository)
	userController := _userController.NewUserControllerImpl(userService)

	router.GET("/:userID", userController.FindByID, m.Authorizing)
	router.PATCH("/:userID", userController.Update, m.Authorizing)
	router.DELETE("/:userID", userController.Delete, m.Authorizing)
}
