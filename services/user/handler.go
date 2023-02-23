package user

import (
	"net/http"

	"github.com/labstack/echo"
)

type UserGetResponse struct {
	User *User `json:"user"`
}

func (s *UserService) HandleUserGet(ctx echo.Context) error {
	userID := ctx.Param("id")
	if userID == "" {
		return ctx.JSON(http.StatusBadRequest, "user ID is required")
	}

	user, err := s.processUserGet(userID)
	if err != nil {
		return ctx.JSON(http.StatusInternalServerError, err)
	}

	return ctx.JSON(http.StatusOK, UserGetResponse{
		User: user,
	})
}

type UserSetRequest struct {
	FirstName string `json:"firstName"`
	LastName  string `json:"lastName"`
	Email     string `json:"email"`
	Age       int    `json:"age"`
}

type UserSetResponse struct {
	Success bool `json:"success"`
}

func (s *UserService) HandleUserSet(ctx echo.Context) error {
	req := new(UserSetRequest)
	err := ctx.Bind(req)
	if err != nil {
		return ctx.JSON(http.StatusBadRequest, err)
	}

	err = ctx.Validate(req)
	if err != nil {
		return ctx.JSON(http.StatusBadRequest, err)
	}

	err = s.processUserSet(&User{
		Email:     req.Email,
		FirstName: req.FirstName,
		LastName:  req.LastName,
		Age:       req.Age,
	})

	return ctx.JSON(http.StatusOK, UserSetResponse{
		Success: true,
	})
}

func (s *UserService) RegisterRoutes(e *echo.Echo) {
	e.GET(RouteUser, s.HandleUserGet)
	e.POST(RouteUser, s.HandleUserSet)
}
