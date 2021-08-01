// Code generated by protoc-gen-micro. DO NOT EDIT.
// source: exorder.proto

package fotune_srv_exchange

import (
	fmt "fmt"
	proto "github.com/golang/protobuf/proto"
	empty "github.com/golang/protobuf/ptypes/empty"
	math "math"
)

import (
	context "context"
	api "github.com/micro/go-micro/v2/api"
	client "github.com/micro/go-micro/v2/client"
	server "github.com/micro/go-micro/v2/server"
)

// Reference imports to suppress errors if they are not otherwise used.
var _ = proto.Marshal
var _ = fmt.Errorf
var _ = math.Inf

// This is a compile-time assertion to ensure that this generated file
// is compatible with the proto package it is being compiled against.
// A compilation error at this line likely means your copy of the
// proto package needs to be updated.
const _ = proto.ProtoPackageIsVersion3 // please upgrade the proto package

// Reference imports to suppress errors if they are not otherwise used.
var _ api.Endpoint
var _ context.Context
var _ client.Option
var _ server.Option

// Api Endpoints for ExOrder service

func NewExOrderEndpoints() []*api.Endpoint {
	return []*api.Endpoint{}
}

// Client API for ExOrder service

type ExOrderService interface {
	ExChangeInfo(ctx context.Context, in *empty.Empty, opts ...client.CallOption) (*ExChangeList, error)
	AddExchangeApi(ctx context.Context, in *ExchangeApi, opts ...client.CallOption) (*empty.Empty, error)
	GetExchangeApiList(ctx context.Context, in *GetExApiReq, opts ...client.CallOption) (*ExApiResp, error)
	UpdateExchangeApi(ctx context.Context, in *UpdateExchangeApiReq, opts ...client.CallOption) (*empty.Empty, error)
	DeleteExchangeApi(ctx context.Context, in *UserApiReq, opts ...client.CallOption) (*empty.Empty, error)
	GetTradeList(ctx context.Context, in *GetTradeListReq, opts ...client.CallOption) (*TradeListResp, error)
	GetProfitRealTime(ctx context.Context, in *ProfitReq, opts ...client.CallOption) (*ProfitRealTimeResp, error)
	//  rpc GetUserStrategyList (GetStrategyReq) returns (UserStrategyListResp) {
	//  }
	//  rpc SetUserStrategyApi (SetUserStrategyApiReq) returns (google.protobuf.Empty) {
	//  }
	//  rpc GetUserStrategyDetail (UserStrategyDetailReq) returns (UserStrategy) {
	//  }
	//  rpc GetStrategyList (StrategyReq) returns (StrategyList) {
	//  }
	//  rpc SetUserStrategyBalance (SetUserStrategyBalanceReq) returns (google.protobuf.Empty) {
	//  }
	//  rpc GetStrategy (GetStrategyDetail) returns (Strategy) {
	//  }
	//  rpc CreateStrategy (CreateStrategyReq) returns (google.protobuf.Empty) {
	//  }
	Evaluation(ctx context.Context, in *TradeReq, opts ...client.CallOption) (*empty.Empty, error)
	EvaluationSpot(ctx context.Context, in *OrderReq, opts ...client.CallOption) (*empty.Empty, error)
	GetExchangePos(ctx context.Context, in *GetExchangePosReq, opts ...client.CallOption) (*ExchangePosResp, error)
	//  rpc RunUserStrategy (UserStrategyRunReq) returns (google.protobuf.Empty) {
	//  }
	GetTradeSymbols(ctx context.Context, in *TradeSymbolReq, opts ...client.CallOption) (*GetSymbolsResp, error)
	GetUserStrategyEva(ctx context.Context, in *UserStrategyDetailReq, opts ...client.CallOption) (*UserStrategyEvaResp, error)
	GetApiKeyInfo(ctx context.Context, in *UserApiKeyReq, opts ...client.CallOption) (*ExchangeApiResp, error)
	GetAssetsByAllApiKey(ctx context.Context, in *GetExApiReq, opts ...client.CallOption) (*AssertsResp, error)
	GetSymbolRankWithRateYear(ctx context.Context, in *empty.Empty, opts ...client.CallOption) (*SymbolRankWithRateYearResp, error)
	StrategyProfitCompensate(ctx context.Context, in *StrategyProfitCompensateReq, opts ...client.CallOption) (*empty.Empty, error)
}

type exOrderService struct {
	c    client.Client
	name string
}

func NewExOrderService(name string, c client.Client) ExOrderService {
	return &exOrderService{
		c:    c,
		name: name,
	}
}

func (c *exOrderService) ExChangeInfo(ctx context.Context, in *empty.Empty, opts ...client.CallOption) (*ExChangeList, error) {
	req := c.c.NewRequest(c.name, "ExOrder.ExChangeInfo", in)
	out := new(ExChangeList)
	err := c.c.Call(ctx, req, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *exOrderService) AddExchangeApi(ctx context.Context, in *ExchangeApi, opts ...client.CallOption) (*empty.Empty, error) {
	req := c.c.NewRequest(c.name, "ExOrder.AddExchangeApi", in)
	out := new(empty.Empty)
	err := c.c.Call(ctx, req, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *exOrderService) GetExchangeApiList(ctx context.Context, in *GetExApiReq, opts ...client.CallOption) (*ExApiResp, error) {
	req := c.c.NewRequest(c.name, "ExOrder.GetExchangeApiList", in)
	out := new(ExApiResp)
	err := c.c.Call(ctx, req, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *exOrderService) UpdateExchangeApi(ctx context.Context, in *UpdateExchangeApiReq, opts ...client.CallOption) (*empty.Empty, error) {
	req := c.c.NewRequest(c.name, "ExOrder.UpdateExchangeApi", in)
	out := new(empty.Empty)
	err := c.c.Call(ctx, req, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *exOrderService) DeleteExchangeApi(ctx context.Context, in *UserApiReq, opts ...client.CallOption) (*empty.Empty, error) {
	req := c.c.NewRequest(c.name, "ExOrder.DeleteExchangeApi", in)
	out := new(empty.Empty)
	err := c.c.Call(ctx, req, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *exOrderService) GetTradeList(ctx context.Context, in *GetTradeListReq, opts ...client.CallOption) (*TradeListResp, error) {
	req := c.c.NewRequest(c.name, "ExOrder.GetTradeList", in)
	out := new(TradeListResp)
	err := c.c.Call(ctx, req, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *exOrderService) GetProfitRealTime(ctx context.Context, in *ProfitReq, opts ...client.CallOption) (*ProfitRealTimeResp, error) {
	req := c.c.NewRequest(c.name, "ExOrder.GetProfitRealTime", in)
	out := new(ProfitRealTimeResp)
	err := c.c.Call(ctx, req, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *exOrderService) Evaluation(ctx context.Context, in *TradeReq, opts ...client.CallOption) (*empty.Empty, error) {
	req := c.c.NewRequest(c.name, "ExOrder.Evaluation", in)
	out := new(empty.Empty)
	err := c.c.Call(ctx, req, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *exOrderService) EvaluationSpot(ctx context.Context, in *OrderReq, opts ...client.CallOption) (*empty.Empty, error) {
	req := c.c.NewRequest(c.name, "ExOrder.EvaluationSpot", in)
	out := new(empty.Empty)
	err := c.c.Call(ctx, req, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *exOrderService) GetExchangePos(ctx context.Context, in *GetExchangePosReq, opts ...client.CallOption) (*ExchangePosResp, error) {
	req := c.c.NewRequest(c.name, "ExOrder.GetExchangePos", in)
	out := new(ExchangePosResp)
	err := c.c.Call(ctx, req, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *exOrderService) GetTradeSymbols(ctx context.Context, in *TradeSymbolReq, opts ...client.CallOption) (*GetSymbolsResp, error) {
	req := c.c.NewRequest(c.name, "ExOrder.GetTradeSymbols", in)
	out := new(GetSymbolsResp)
	err := c.c.Call(ctx, req, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *exOrderService) GetUserStrategyEva(ctx context.Context, in *UserStrategyDetailReq, opts ...client.CallOption) (*UserStrategyEvaResp, error) {
	req := c.c.NewRequest(c.name, "ExOrder.GetUserStrategyEva", in)
	out := new(UserStrategyEvaResp)
	err := c.c.Call(ctx, req, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *exOrderService) GetApiKeyInfo(ctx context.Context, in *UserApiKeyReq, opts ...client.CallOption) (*ExchangeApiResp, error) {
	req := c.c.NewRequest(c.name, "ExOrder.GetApiKeyInfo", in)
	out := new(ExchangeApiResp)
	err := c.c.Call(ctx, req, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *exOrderService) GetAssetsByAllApiKey(ctx context.Context, in *GetExApiReq, opts ...client.CallOption) (*AssertsResp, error) {
	req := c.c.NewRequest(c.name, "ExOrder.GetAssetsByAllApiKey", in)
	out := new(AssertsResp)
	err := c.c.Call(ctx, req, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *exOrderService) GetSymbolRankWithRateYear(ctx context.Context, in *empty.Empty, opts ...client.CallOption) (*SymbolRankWithRateYearResp, error) {
	req := c.c.NewRequest(c.name, "ExOrder.GetSymbolRankWithRateYear", in)
	out := new(SymbolRankWithRateYearResp)
	err := c.c.Call(ctx, req, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *exOrderService) StrategyProfitCompensate(ctx context.Context, in *StrategyProfitCompensateReq, opts ...client.CallOption) (*empty.Empty, error) {
	req := c.c.NewRequest(c.name, "ExOrder.StrategyProfitCompensate", in)
	out := new(empty.Empty)
	err := c.c.Call(ctx, req, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// Server API for ExOrder service

type ExOrderHandler interface {
	ExChangeInfo(context.Context, *empty.Empty, *ExChangeList) error
	AddExchangeApi(context.Context, *ExchangeApi, *empty.Empty) error
	GetExchangeApiList(context.Context, *GetExApiReq, *ExApiResp) error
	UpdateExchangeApi(context.Context, *UpdateExchangeApiReq, *empty.Empty) error
	DeleteExchangeApi(context.Context, *UserApiReq, *empty.Empty) error
	GetTradeList(context.Context, *GetTradeListReq, *TradeListResp) error
	GetProfitRealTime(context.Context, *ProfitReq, *ProfitRealTimeResp) error
	//  rpc GetUserStrategyList (GetStrategyReq) returns (UserStrategyListResp) {
	//  }
	//  rpc SetUserStrategyApi (SetUserStrategyApiReq) returns (google.protobuf.Empty) {
	//  }
	//  rpc GetUserStrategyDetail (UserStrategyDetailReq) returns (UserStrategy) {
	//  }
	//  rpc GetStrategyList (StrategyReq) returns (StrategyList) {
	//  }
	//  rpc SetUserStrategyBalance (SetUserStrategyBalanceReq) returns (google.protobuf.Empty) {
	//  }
	//  rpc GetStrategy (GetStrategyDetail) returns (Strategy) {
	//  }
	//  rpc CreateStrategy (CreateStrategyReq) returns (google.protobuf.Empty) {
	//  }
	Evaluation(context.Context, *TradeReq, *empty.Empty) error
	EvaluationSpot(context.Context, *OrderReq, *empty.Empty) error
	GetExchangePos(context.Context, *GetExchangePosReq, *ExchangePosResp) error
	//  rpc RunUserStrategy (UserStrategyRunReq) returns (google.protobuf.Empty) {
	//  }
	GetTradeSymbols(context.Context, *TradeSymbolReq, *GetSymbolsResp) error
	GetUserStrategyEva(context.Context, *UserStrategyDetailReq, *UserStrategyEvaResp) error
	GetApiKeyInfo(context.Context, *UserApiKeyReq, *ExchangeApiResp) error
	GetAssetsByAllApiKey(context.Context, *GetExApiReq, *AssertsResp) error
	GetSymbolRankWithRateYear(context.Context, *empty.Empty, *SymbolRankWithRateYearResp) error
	StrategyProfitCompensate(context.Context, *StrategyProfitCompensateReq, *empty.Empty) error
}

func RegisterExOrderHandler(s server.Server, hdlr ExOrderHandler, opts ...server.HandlerOption) error {
	type exOrder interface {
		ExChangeInfo(ctx context.Context, in *empty.Empty, out *ExChangeList) error
		AddExchangeApi(ctx context.Context, in *ExchangeApi, out *empty.Empty) error
		GetExchangeApiList(ctx context.Context, in *GetExApiReq, out *ExApiResp) error
		UpdateExchangeApi(ctx context.Context, in *UpdateExchangeApiReq, out *empty.Empty) error
		DeleteExchangeApi(ctx context.Context, in *UserApiReq, out *empty.Empty) error
		GetTradeList(ctx context.Context, in *GetTradeListReq, out *TradeListResp) error
		GetProfitRealTime(ctx context.Context, in *ProfitReq, out *ProfitRealTimeResp) error
		Evaluation(ctx context.Context, in *TradeReq, out *empty.Empty) error
		EvaluationSpot(ctx context.Context, in *OrderReq, out *empty.Empty) error
		GetExchangePos(ctx context.Context, in *GetExchangePosReq, out *ExchangePosResp) error
		GetTradeSymbols(ctx context.Context, in *TradeSymbolReq, out *GetSymbolsResp) error
		GetUserStrategyEva(ctx context.Context, in *UserStrategyDetailReq, out *UserStrategyEvaResp) error
		GetApiKeyInfo(ctx context.Context, in *UserApiKeyReq, out *ExchangeApiResp) error
		GetAssetsByAllApiKey(ctx context.Context, in *GetExApiReq, out *AssertsResp) error
		GetSymbolRankWithRateYear(ctx context.Context, in *empty.Empty, out *SymbolRankWithRateYearResp) error
		StrategyProfitCompensate(ctx context.Context, in *StrategyProfitCompensateReq, out *empty.Empty) error
	}
	type ExOrder struct {
		exOrder
	}
	h := &exOrderHandler{hdlr}
	return s.Handle(s.NewHandler(&ExOrder{h}, opts...))
}

type exOrderHandler struct {
	ExOrderHandler
}

func (h *exOrderHandler) ExChangeInfo(ctx context.Context, in *empty.Empty, out *ExChangeList) error {
	return h.ExOrderHandler.ExChangeInfo(ctx, in, out)
}

func (h *exOrderHandler) AddExchangeApi(ctx context.Context, in *ExchangeApi, out *empty.Empty) error {
	return h.ExOrderHandler.AddExchangeApi(ctx, in, out)
}

func (h *exOrderHandler) GetExchangeApiList(ctx context.Context, in *GetExApiReq, out *ExApiResp) error {
	return h.ExOrderHandler.GetExchangeApiList(ctx, in, out)
}

func (h *exOrderHandler) UpdateExchangeApi(ctx context.Context, in *UpdateExchangeApiReq, out *empty.Empty) error {
	return h.ExOrderHandler.UpdateExchangeApi(ctx, in, out)
}

func (h *exOrderHandler) DeleteExchangeApi(ctx context.Context, in *UserApiReq, out *empty.Empty) error {
	return h.ExOrderHandler.DeleteExchangeApi(ctx, in, out)
}

func (h *exOrderHandler) GetTradeList(ctx context.Context, in *GetTradeListReq, out *TradeListResp) error {
	return h.ExOrderHandler.GetTradeList(ctx, in, out)
}

func (h *exOrderHandler) GetProfitRealTime(ctx context.Context, in *ProfitReq, out *ProfitRealTimeResp) error {
	return h.ExOrderHandler.GetProfitRealTime(ctx, in, out)
}

func (h *exOrderHandler) Evaluation(ctx context.Context, in *TradeReq, out *empty.Empty) error {
	return h.ExOrderHandler.Evaluation(ctx, in, out)
}

func (h *exOrderHandler) EvaluationSpot(ctx context.Context, in *OrderReq, out *empty.Empty) error {
	return h.ExOrderHandler.EvaluationSpot(ctx, in, out)
}

func (h *exOrderHandler) GetExchangePos(ctx context.Context, in *GetExchangePosReq, out *ExchangePosResp) error {
	return h.ExOrderHandler.GetExchangePos(ctx, in, out)
}

func (h *exOrderHandler) GetTradeSymbols(ctx context.Context, in *TradeSymbolReq, out *GetSymbolsResp) error {
	return h.ExOrderHandler.GetTradeSymbols(ctx, in, out)
}

func (h *exOrderHandler) GetUserStrategyEva(ctx context.Context, in *UserStrategyDetailReq, out *UserStrategyEvaResp) error {
	return h.ExOrderHandler.GetUserStrategyEva(ctx, in, out)
}

func (h *exOrderHandler) GetApiKeyInfo(ctx context.Context, in *UserApiKeyReq, out *ExchangeApiResp) error {
	return h.ExOrderHandler.GetApiKeyInfo(ctx, in, out)
}

func (h *exOrderHandler) GetAssetsByAllApiKey(ctx context.Context, in *GetExApiReq, out *AssertsResp) error {
	return h.ExOrderHandler.GetAssetsByAllApiKey(ctx, in, out)
}

func (h *exOrderHandler) GetSymbolRankWithRateYear(ctx context.Context, in *empty.Empty, out *SymbolRankWithRateYearResp) error {
	return h.ExOrderHandler.GetSymbolRankWithRateYear(ctx, in, out)
}

func (h *exOrderHandler) StrategyProfitCompensate(ctx context.Context, in *StrategyProfitCompensateReq, out *empty.Empty) error {
	return h.ExOrderHandler.StrategyProfitCompensate(ctx, in, out)
}