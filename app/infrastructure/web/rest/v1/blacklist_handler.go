package v1

import (
	"bitbucket.org/fabribraguev/api-toolbox/app/application/usecase"
	"github.com/labstack/echo"
	"net/http"
)

type blackListHandler struct {
	blackListUsecase usecase.BlackListUsecase
}

func NewBlackListHandler(e *echo.Echo, blackListUsecase usecase.BlackListUsecase) *blackListHandler {
	h := &blackListHandler{blackListUsecase: blackListUsecase}
	g := e.Group("/tools")

	g.GET("/show_black_list", h.showBlackListUsecase)
	return h
}

func (h *blackListHandler) showBlackListUsecase(c echo.Context) error {
	option := c.QueryParam("option")

	blacklist, err := h.blackListUsecase.ShowBlackList(option)
	if err != nil {
		return c.JSON(http.StatusBadGateway, echo.Map{"message": err.Error()})
	}
	return c.JSON(http.StatusOK, blacklist)
}