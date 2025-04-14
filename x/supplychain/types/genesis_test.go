package types_test

import (
	"testing"

	"logistics/x/supplychain/types"

	"github.com/stretchr/testify/require"
)

func TestGenesisState_Validate(t *testing.T) {
	tests := []struct {
		desc     string
		genState *types.GenesisState
		valid    bool
	}{
		{
			desc:     "default is valid",
			genState: types.DefaultGenesis(),
			valid:    true,
		},
		{
			desc: "valid genesis state",
			genState: &types.GenesisState{

				ProductList: []types.Product{
					{
						Id: 0,
					},
					{
						Id: 1,
					},
				},
				ProductCount: 2,
				ShipmentList: []types.Shipment{
					{
						Id: 0,
					},
					{
						Id: 1,
					},
				},
				ShipmentCount: 2,
				OrderList: []types.Order{
					{
						Id: 0,
					},
					{
						Id: 1,
					},
				},
				OrderCount: 2,
				// this line is used by starport scaffolding # types/genesis/validField
			},
			valid: true,
		},
		{
			desc: "duplicated product",
			genState: &types.GenesisState{
				ProductList: []types.Product{
					{
						Id: 0,
					},
					{
						Id: 0,
					},
				},
			},
			valid: false,
		},
		{
			desc: "invalid product count",
			genState: &types.GenesisState{
				ProductList: []types.Product{
					{
						Id: 1,
					},
				},
				ProductCount: 0,
			},
			valid: false,
		},
		{
			desc: "duplicated shipment",
			genState: &types.GenesisState{
				ShipmentList: []types.Shipment{
					{
						Id: 0,
					},
					{
						Id: 0,
					},
				},
			},
			valid: false,
		},
		{
			desc: "invalid shipment count",
			genState: &types.GenesisState{
				ShipmentList: []types.Shipment{
					{
						Id: 1,
					},
				},
				ShipmentCount: 0,
			},
			valid: false,
		},
		{
			desc: "duplicated order",
			genState: &types.GenesisState{
				OrderList: []types.Order{
					{
						Id: 0,
					},
					{
						Id: 0,
					},
				},
			},
			valid: false,
		},
		{
			desc: "invalid order count",
			genState: &types.GenesisState{
				OrderList: []types.Order{
					{
						Id: 1,
					},
				},
				OrderCount: 0,
			},
			valid: false,
		},
		// this line is used by starport scaffolding # types/genesis/testcase
	}
	for _, tc := range tests {
		t.Run(tc.desc, func(t *testing.T) {
			err := tc.genState.Validate()
			if tc.valid {
				require.NoError(t, err)
			} else {
				require.Error(t, err)
			}
		})
	}
}
