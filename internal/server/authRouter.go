package server

import (
	_controller "github.com/jaayroots/habit-tracker-api/controller"
	_repository "github.com/jaayroots/habit-tracker-api/repository"
	_service "github.com/jaayroots/habit-tracker-api/service"
)

func (s *echoServer) authRouter(m *authorizingMiddleware) {
	router := s.app.Group("/auth")

	userRepository := _repository.NewUserRepository(s.db, s.app.Logger)
	userContactRepository := _repository.NewUserContactRepository(s.db, s.app.Logger)
	sessionRepository := _repository.NewSessionRepositoryImpl(s.db, s.app.Logger)

	authService := _service.NewAuthServiceImpl(s.db, userRepository, userContactRepository, sessionRepository)
	authController := _controller.NewAuthControllerImpl(authService)

	router.POST("/register", authController.Register)
	router.POST("/login", authController.Login)
	router.POST("/logout", authController.Logout, m.Authorizing)
	router.POST("/refreash", authController.Refresh, m.Authorizing)
}
