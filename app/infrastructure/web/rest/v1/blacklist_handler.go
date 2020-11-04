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
	g := e.Group("/bl")

	g.GET("/show_black_list", h.showBlackListUsecase)
	g.GET("/show_exclude_items", h.showBlackListExcludeUsecase)
	g.GET("/show_black_list_exclude", h.showBlackListWithoutExcludeItemsUsecase)
	g.DELETE("/delete_excludes", h.deleteBlackListUsecase)

	return h
}

func (h *blackListHandler) showBlackListUsecase(c echo.Context) error {
	option := "ORIGINAL"

	blacklist, err := h.blackListUsecase.ShowBlackList(option)
	if err != nil {
		return c.JSON(http.StatusBadGateway, echo.Map{"message": err.Error()})
	}
	return c.JSON(http.StatusOK, blacklist)
}

func (h *blackListHandler) showBlackListExcludeUsecase(c echo.Context) error {
	option := "EXCLUDEITEMS"

	blacklist, err := h.blackListUsecase.ShowBlackList(option)
	if err != nil {
		return c.JSON(http.StatusBadGateway, echo.Map{"message": err.Error()})
	}
	return c.JSON(http.StatusOK, blacklist)
}

func (h *blackListHandler) showBlackListWithoutExcludeItemsUsecase(c echo.Context) error {
	option := "WITHOUTEXCLUDE"

	blacklist, err := h.blackListUsecase.ShowBlackList(option)
	if err != nil {
		return c.JSON(http.StatusBadGateway, echo.Map{"message": err.Error()})
	}
	return c.JSON(http.StatusOK, blacklist)
}

func (h *blackListHandler) deleteBlackListUsecase(c echo.Context) error {

	err := h.blackListUsecase.DeleteExcludeItems()
	if err != nil {
		return c.JSON(http.StatusBadGateway, echo.Map{"message": err.Error()})
	}
	return c.JSON(http.StatusOK, "All exclude items remove")
}