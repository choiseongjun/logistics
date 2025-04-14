package supplychain

import (
	"math/rand"

	sdk "github.com/cosmos/cosmos-sdk/types"
	"github.com/cosmos/cosmos-sdk/types/module"
	simtypes "github.com/cosmos/cosmos-sdk/types/simulation"
	"github.com/cosmos/cosmos-sdk/x/simulation"

	"logistics/testutil/sample"
	supplychainsimulation "logistics/x/supplychain/simulation"
	"logistics/x/supplychain/types"
)

// avoid unused import issue
var (
	_ = supplychainsimulation.FindAccount
	_ = rand.Rand{}
	_ = sample.AccAddress
	_ = sdk.AccAddress{}
	_ = simulation.MsgEntryKind
)

const (
	opWeightMsgTransferOwnership = "op_weight_msg_transfer_ownership"
	// TODO: Determine the simulation weight value
	defaultWeightMsgTransferOwnership int = 100

	opWeightMsgUpdateStatus = "op_weight_msg_update_status"
	// TODO: Determine the simulation weight value
	defaultWeightMsgUpdateStatus int = 100

	opWeightMsgUpdateShipment = "op_weight_msg_update_shipment"
	// TODO: Determine the simulation weight value
	defaultWeightMsgUpdateShipment int = 100

	opWeightMsgCreateProduct = "op_weight_msg_create_product"
	// TODO: Determine the simulation weight value
	defaultWeightMsgCreateProduct int = 100

	opWeightMsgCreateOrder = "op_weight_msg_order"
	// TODO: Determine the simulation weight value
	defaultWeightMsgCreateOrder int = 100

	opWeightMsgUpdateOrder = "op_weight_msg_order"
	// TODO: Determine the simulation weight value
	defaultWeightMsgUpdateOrder int = 100

	opWeightMsgDeleteOrder = "op_weight_msg_order"
	// TODO: Determine the simulation weight value
	defaultWeightMsgDeleteOrder int = 100

	// this line is used by starport scaffolding # simapp/module/const
)

// GenerateGenesisState creates a randomized GenState of the module.
func (AppModule) GenerateGenesisState(simState *module.SimulationState) {
	accs := make([]string, len(simState.Accounts))
	for i, acc := range simState.Accounts {
		accs[i] = acc.Address.String()
	}
	supplychainGenesis := types.GenesisState{
		Params: types.DefaultParams(),
		OrderList: []types.Order{
			{
				Id:      0,
				Creator: sample.AccAddress(),
			},
			{
				Id:      1,
				Creator: sample.AccAddress(),
			},
		},
		OrderCount: 2,
		// this line is used by starport scaffolding # simapp/module/genesisState
	}
	simState.GenState[types.ModuleName] = simState.Cdc.MustMarshalJSON(&supplychainGenesis)
}

// RegisterStoreDecoder registers a decoder.
func (am AppModule) RegisterStoreDecoder(_ simtypes.StoreDecoderRegistry) {}

// WeightedOperations returns the all the gov module operations with their respective weights.
func (am AppModule) WeightedOperations(simState module.SimulationState) []simtypes.WeightedOperation {
	operations := make([]simtypes.WeightedOperation, 0)

	var weightMsgTransferOwnership int
	simState.AppParams.GetOrGenerate(opWeightMsgTransferOwnership, &weightMsgTransferOwnership, nil,
		func(_ *rand.Rand) {
			weightMsgTransferOwnership = defaultWeightMsgTransferOwnership
		},
	)
	operations = append(operations, simulation.NewWeightedOperation(
		weightMsgTransferOwnership,
		supplychainsimulation.SimulateMsgTransferOwnership(am.accountKeeper, am.bankKeeper, am.keeper),
	))

	var weightMsgUpdateStatus int
	simState.AppParams.GetOrGenerate(opWeightMsgUpdateStatus, &weightMsgUpdateStatus, nil,
		func(_ *rand.Rand) {
			weightMsgUpdateStatus = defaultWeightMsgUpdateStatus
		},
	)
	operations = append(operations, simulation.NewWeightedOperation(
		weightMsgUpdateStatus,
		supplychainsimulation.SimulateMsgUpdateStatus(am.accountKeeper, am.bankKeeper, am.keeper),
	))

	var weightMsgUpdateShipment int
	simState.AppParams.GetOrGenerate(opWeightMsgUpdateShipment, &weightMsgUpdateShipment, nil,
		func(_ *rand.Rand) {
			weightMsgUpdateShipment = defaultWeightMsgUpdateShipment
		},
	)
	operations = append(operations, simulation.NewWeightedOperation(
		weightMsgUpdateShipment,
		supplychainsimulation.SimulateMsgUpdateShipment(am.accountKeeper, am.bankKeeper, am.keeper),
	))

	var weightMsgCreateProduct int
	simState.AppParams.GetOrGenerate(opWeightMsgCreateProduct, &weightMsgCreateProduct, nil,
		func(_ *rand.Rand) {
			weightMsgCreateProduct = defaultWeightMsgCreateProduct
		},
	)
	operations = append(operations, simulation.NewWeightedOperation(
		weightMsgCreateProduct,
		supplychainsimulation.SimulateMsgCreateProduct(am.accountKeeper, am.bankKeeper, am.keeper),
	))

	var weightMsgCreateOrder int
	simState.AppParams.GetOrGenerate(opWeightMsgCreateOrder, &weightMsgCreateOrder, nil,
		func(_ *rand.Rand) {
			weightMsgCreateOrder = defaultWeightMsgCreateOrder
		},
	)
	operations = append(operations, simulation.NewWeightedOperation(
		weightMsgCreateOrder,
		supplychainsimulation.SimulateMsgCreateOrder(am.accountKeeper, am.bankKeeper, am.keeper),
	))

	var weightMsgUpdateOrder int
	simState.AppParams.GetOrGenerate(opWeightMsgUpdateOrder, &weightMsgUpdateOrder, nil,
		func(_ *rand.Rand) {
			weightMsgUpdateOrder = defaultWeightMsgUpdateOrder
		},
	)
	operations = append(operations, simulation.NewWeightedOperation(
		weightMsgUpdateOrder,
		supplychainsimulation.SimulateMsgUpdateOrder(am.accountKeeper, am.bankKeeper, am.keeper),
	))

	var weightMsgDeleteOrder int
	simState.AppParams.GetOrGenerate(opWeightMsgDeleteOrder, &weightMsgDeleteOrder, nil,
		func(_ *rand.Rand) {
			weightMsgDeleteOrder = defaultWeightMsgDeleteOrder
		},
	)
	operations = append(operations, simulation.NewWeightedOperation(
		weightMsgDeleteOrder,
		supplychainsimulation.SimulateMsgDeleteOrder(am.accountKeeper, am.bankKeeper, am.keeper),
	))

	// this line is used by starport scaffolding # simapp/module/operation

	return operations
}

// ProposalMsgs returns msgs used for governance proposals for simulations.
func (am AppModule) ProposalMsgs(simState module.SimulationState) []simtypes.WeightedProposalMsg {
	return []simtypes.WeightedProposalMsg{
		simulation.NewWeightedProposalMsg(
			opWeightMsgTransferOwnership,
			defaultWeightMsgTransferOwnership,
			func(r *rand.Rand, ctx sdk.Context, accs []simtypes.Account) sdk.Msg {
				supplychainsimulation.SimulateMsgTransferOwnership(am.accountKeeper, am.bankKeeper, am.keeper)
				return nil
			},
		),
		simulation.NewWeightedProposalMsg(
			opWeightMsgUpdateStatus,
			defaultWeightMsgUpdateStatus,
			func(r *rand.Rand, ctx sdk.Context, accs []simtypes.Account) sdk.Msg {
				supplychainsimulation.SimulateMsgUpdateStatus(am.accountKeeper, am.bankKeeper, am.keeper)
				return nil
			},
		),
		simulation.NewWeightedProposalMsg(
			opWeightMsgUpdateShipment,
			defaultWeightMsgUpdateShipment,
			func(r *rand.Rand, ctx sdk.Context, accs []simtypes.Account) sdk.Msg {
				supplychainsimulation.SimulateMsgUpdateShipment(am.accountKeeper, am.bankKeeper, am.keeper)
				return nil
			},
		),
		simulation.NewWeightedProposalMsg(
			opWeightMsgCreateProduct,
			defaultWeightMsgCreateProduct,
			func(r *rand.Rand, ctx sdk.Context, accs []simtypes.Account) sdk.Msg {
				supplychainsimulation.SimulateMsgCreateProduct(am.accountKeeper, am.bankKeeper, am.keeper)
				return nil
			},
		),
		simulation.NewWeightedProposalMsg(
			opWeightMsgCreateOrder,
			defaultWeightMsgCreateOrder,
			func(r *rand.Rand, ctx sdk.Context, accs []simtypes.Account) sdk.Msg {
				supplychainsimulation.SimulateMsgCreateOrder(am.accountKeeper, am.bankKeeper, am.keeper)
				return nil
			},
		),
		simulation.NewWeightedProposalMsg(
			opWeightMsgUpdateOrder,
			defaultWeightMsgUpdateOrder,
			func(r *rand.Rand, ctx sdk.Context, accs []simtypes.Account) sdk.Msg {
				supplychainsimulation.SimulateMsgUpdateOrder(am.accountKeeper, am.bankKeeper, am.keeper)
				return nil
			},
		),
		simulation.NewWeightedProposalMsg(
			opWeightMsgDeleteOrder,
			defaultWeightMsgDeleteOrder,
			func(r *rand.Rand, ctx sdk.Context, accs []simtypes.Account) sdk.Msg {
				supplychainsimulation.SimulateMsgDeleteOrder(am.accountKeeper, am.bankKeeper, am.keeper)
				return nil
			},
		),
		// this line is used by starport scaffolding # simapp/module/OpMsg
	}
}
