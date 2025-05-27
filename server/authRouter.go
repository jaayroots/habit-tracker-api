package server

import (
	_authController "github.com/jaayroots/go_base/pkg/auth/controller"
	_authRepository "github.com/jaayroots/go_base/pkg/auth/repository"
	_authService "github.com/jaayroots/go_base/pkg/auth/service"
	_userRepository "github.com/jaayroots/go_base/pkg/user/repository"
)

func (s *echoServer) authRouter(m *authorizingMiddleware) {
	router := s.app.Group("/auth")

	userRepository := _userRepository.NewUserRepositoryImpl(s.db, s.app.Logger)
	sessionRepository := _authRepository.NewSessionRepositoryImpl(s.db, s.app.Logger)

	authService := _authService.NewAuthServiceImpl(userRepository, sessionRepository)
	authController := _authController.NewAuthControllerImpl(authService)

	router.POST("/register", authController.Register)
	router.POST("/login", authController.Login)
	router.POST("/logout", authController.Logout, m.Authorizing)
	router.POST("/refreash", authController.Refresh, m.Authorizing)
}
