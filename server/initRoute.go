package server

import (
	_authController "github.com/jaayroots/habit-tracker-api/pkg/auth/controller"
	_authRepository "github.com/jaayroots/habit-tracker-api/pkg/auth/repository"
	_authService "github.com/jaayroots/habit-tracker-api/pkg/auth/service"
	_userRepository "github.com/jaayroots/habit-tracker-api/pkg/user/repository"
)

func (s *echoServer) initRoute() {

	authorizingMiddleware := s.getAuthorizingMiddleware()

	s.healthRouter()
	s.authRouter(authorizingMiddleware)
	s.usersRouter(authorizingMiddleware)
	s.habitRouter(authorizingMiddleware)
	s.checkinRouter(authorizingMiddleware)
}

func (s *echoServer) getAuthorizingMiddleware() *authorizingMiddleware {

	userRepository := _userRepository.NewUserRepositoryImpl(s.db, s.app.Logger)
	authRepository := _authRepository.NewSessionRepositoryImpl(s.db, s.app.Logger)

	authService := _authService.NewAuthServiceImpl(userRepository, authRepository)
	middlewareController := _authController.NewMiddlewareControllerImpl(authService)
	return &authorizingMiddleware{
		middlewareController,
	}
}
