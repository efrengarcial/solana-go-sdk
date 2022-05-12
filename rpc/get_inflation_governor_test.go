package rpc

import (
	"context"
	"testing"
)

func TestGetInflationGovernor(t *testing.T) {
	tests := []testRpcCallParam{
		{
			RequestBody:  `{"jsonrpc":"2.0", "id":1, "method":"getInflationGovernor"}`,
			ResponseBody: `{"jsonrpc":"2.0","result":{"foundation":0.05,"foundationTerm":7.0,"initial":0.08,"taper":0.15,"terminal":0.015},"id":1}`,
			RpcCall: func(rc RpcClient) (interface{}, error) {
				return rc.GetInflationGovernor(
					context.TODO(),
				)
			},
			ExpectedResponse: GetInflationGovernorResponse{
				GeneralResponse: GeneralResponse{
					JsonRPC: "2.0",
					ID:      1,
					Error:   nil,
				},
				Result: GetInflationGovernorResponseResult{
					Foundation:     0.05,
					FoundationTerm: 7.0,
					Initial:        0.08,
					Taper:          0.15,
					Terminal:       0.015,
				},
			},
			ExpectedError: nil,
		},
		{
			RequestBody:  `{"jsonrpc":"2.0", "id":1, "method":"getInflationGovernor", "params":[{"commitment": "processed"}]}`,
			ResponseBody: `{"jsonrpc":"2.0","result":{"foundation":0.05,"foundationTerm":7.0,"initial":0.08,"taper":0.15,"terminal":0.015},"id":1}`,
			RpcCall: func(rc RpcClient) (interface{}, error) {
				return rc.GetInflationGovernorWithConfig(
					context.TODO(),
					GetInflationGovernorConfig{
						Commitment: CommitmentProcessed,
					},
				)
			},
			ExpectedResponse: GetInflationGovernorResponse{
				GeneralResponse: GeneralResponse{
					JsonRPC: "2.0",
					ID:      1,
					Error:   nil,
				},
				Result: GetInflationGovernorResponseResult{
					Foundation:     0.05,
					FoundationTerm: 7.0,
					Initial:        0.08,
					Taper:          0.15,
					Terminal:       0.015,
				},
			},
			ExpectedError: nil,
		},
	}
	for _, tt := range tests {
		t.Run("", func(t *testing.T) {
			testRpcCall(t, tt)
		})
	}
}
