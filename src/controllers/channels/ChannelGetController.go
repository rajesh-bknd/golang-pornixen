package controller

import (
	"github.com/gin-gonic/gin"
	"github.com/my/repo/src"
	"github.com/my/repo/src/database/dao"
	"github.com/my/repo/src/models"
	"strconv"
)

func GetAllChannels(ctx *gin.Context) {
	var (
		responseModel models.ResponseModel
		err           error
		sortType      string
		pageNo        int64
		channels      models.ChannelsModel
	)

	sortType = ctx.Query("sort-type")
	pageNo_ := ctx.Query("page-no")
	pageNo, pageNoerr := strconv.ParseInt(pageNo_, 10, 64)
	if pageNoerr != nil {
		pageNo = 1
	}
	if len(sortType) == 0 {
		sortType = src.CHANNELTABLE["rank"]
	}
	sortBy, orderBy := GetSort(sortType)
	channels, err = dao.GetAllChannelsDao(sortBy, orderBy, pageNo)
	if err != nil {
		responseModel.Status = `fail`
		responseModel.SetError(`fail`, err.Error())
		ctx.JSON(200, responseModel)
	} else {
		responseModel.Status = `success`
		responseModel.Result = channels
		ctx.JSON(200, responseModel)
	}
}

func GetSort(sortType string) (string, string) {
	return src.CHANNELTABLE[sortType], "ASC"
}
