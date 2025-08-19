package server

import (
	_controller "github.com/habit-tracker-api/controller"
	_repository "github.com/habit-tracker-api/repository"
	_service "github.com/habit-tracker-api/service"
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

	userRepository := _repository.NewUserRepository(s.db, s.app.Logger)
	userContactRepository := _repository.NewUserContactRepository(s.db, s.app.Logger)
	authRepository := _repository.NewSessionRepositoryImpl(s.db, s.app.Logger)

	authService := _service.NewAuthServiceImpl(s.db, userRepository, userContactRepository, authRepository)
	middlewareController := _controller.NewMiddlewareControllerImpl(authService)
	return &authorizingMiddleware{
		middlewareController,
	}
}
