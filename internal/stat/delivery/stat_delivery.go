package delivery

import (
	"encoding/json"
	"github.com/buaazp/fasthttprouter"
	"github.com/sirupsen/logrus"
	"github.com/valyala/fasthttp"
	"github.com/yletamitlu/trainee/internal/models"
	"github.com/yletamitlu/trainee/internal/stat"
)

type StatDelivery struct {
	statUcase stat.StatUsecase
}

func NewStatDelivery(statUsecase stat.StatUsecase) *StatDelivery {
	return &StatDelivery{
		statUcase: statUsecase,
	}
}

func (sd *StatDelivery) Configure(router *fasthttprouter.Router) {
	router.POST("/save", sd.SaveStatisticHandler())
	router.GET("/get", sd.GetStatisticHandler())
	router.POST("/clear", sd.ClearStatisticHandler())
}

func (sd *StatDelivery) SaveStatisticHandler() fasthttp.RequestHandler {
	return func(ctx *fasthttp.RequestCtx) {
		var data *models.Data

		err := json.Unmarshal(ctx.Request.Body(), &data)

		if err != nil {
			logrus.Info(err)
			sd.sendResponse(ctx, 500, "internal server error")
			return
		}
		err = sd.statUcase.AddNewStatistic(data)

		if err != nil {
			logrus.Info(err)
			sd.sendResponse(ctx, 500, "internal server error")
			return
		}

		sd.sendResponse(ctx, 200, "success")
	}
}

func (sd *StatDelivery) GetStatisticHandler() fasthttp.RequestHandler {
	return func(ctx *fasthttp.RequestCtx) {
		since := string(ctx.QueryArgs().Peek("since"))
		until := string(ctx.QueryArgs().Peek("until"))

		statistic, err := sd.statUcase.GetStatisticByPeriod(since, until)
		if err != nil {
			logrus.Info(err)
			sd.sendResponse(ctx, 500, "internal server error")
			return
		}

		sd.sendResponse(ctx, 200, statistic)
	}
}

func (sd *StatDelivery) ClearStatisticHandler() fasthttp.RequestHandler {
	return func(ctx *fasthttp.RequestCtx) {
		err := sd.statUcase.RemoveStatistic()
		if err != nil {
			logrus.Info(err)
			sd.sendResponse(ctx, 500, "internal server error")
			return
		}
		sd.sendResponse(ctx, 200, "success")
	}
}

func (sd *StatDelivery) sendResponse(ctx *fasthttp.RequestCtx, code int, content interface{}) {
	ctx.SetStatusCode(code)
	body, _ := json.Marshal(&content)
	if body != nil {
		ctx.SetBody(body)
	}
}
