package service

import (
	"context"
	"github.com/golang/protobuf/ptypes/empty"
	pb "wq-fotune-backend/api/exchange"
	"wq-fotune-backend/libs/helper"
)

func (e *ExOrderService) Evaluation(ctx context.Context, req *pb.TradeReq, resp *empty.Empty) error {
	//合约的用不上了
	return e.ExOrderSrv.EvaluationSwap(req)
}

func (e *ExOrderService) EvaluationSpot(ctx context.Context, req *pb.OrderReq, resp *empty.Empty) error {
	return e.ExOrderSrv.EvaluationSpot(req)
}

func (e *ExOrderService) StrategyProfitCompensate(ctx context.Context, req *pb.StrategyProfitCompensateReq, empty *empty.Empty) error {
	return e.ExOrderSrv.StrategyProfitCompensate(req.StrategyId, req.Price)
}

func (e *ExOrderService) GetUserStrategyEva(ctx context.Context, req *pb.UserStrategyDetailReq, resp *pb.UserStrategyEvaResp) error {
	profit, err := e.ExOrderSrv.GetProfitByStrID(req.UserId, req.StrategyId)
	if err != nil {
		resp.RealizeProfit = "0"
		resp.RateReturnYear = "0"
		resp.RateReturn = "0"
		return nil
	}
	resp.RateReturn = helper.Float64ToString(profit.RateReturn)
	resp.RealizeProfit = profit.RealizeProfit
	resp.RateReturnYear = helper.Float64ToString(profit.RateReturnYear)
	profitDailyList := e.ExOrderSrv.GetProfitDailyByStrID(req.UserId, req.StrategyId, 365)

	dateMap := make(map[string]bool, 0)
	for _, daily := range profitDailyList {
		dateStr := daily.Date.Format("2006-01-02")
		if _, ok := dateMap[dateStr]; ok { //日期去重数据
			continue
		}
		resp.EvaDailyList = append(resp.EvaDailyList, &pb.EvaDaily{
			Date:        dateStr,
			ProfitDaily: daily.RealizeProfit,
		})
		dateMap[dateStr] = true
	}
	return nil
}
