// Copyright 2018 Owkin, inc.
//
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
//
//      http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License.

package main

import (
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

var (

	// Default compute plan:
	//
	//     ø     ø             ø      ø
	//     |     |             |      |
	//     hd    tr            tr     hd
	//   -----------         -----------
	//  | Composite |       | Composite |
	//   -----------         -----------
	//     hd    tr            tr     hd
	//     |      \            /      |
	//     |       \          /       |
	//     |        -----------       |
	//     |       | Aggregate |      |
	//     |        -----------       |
	//     |            |             |
	//     |     ,_____/ \_____       |
	//     |     |             |      |
	//     hd    tr            tr     hd
	//   -----------         -----------
	//  | Composite |       | Composite |
	//   -----------         -----------
	//     hd    tr            tr     hd
	//
	//
	defaultComputePlan = inputComputePlan{
		ObjectiveKey: objectiveDescriptionHash,
		Traintuples: []inputComputePlanTraintuple{
			inputComputePlanTraintuple{
				DataManagerKey: dataManagerOpenerHash,
				DataSampleKeys: []string{trainDataSampleHash1},
				AlgoKey:        algoHash,
				ID:             traintupleID1,
			},
			inputComputePlanTraintuple{
				DataManagerKey: dataManagerOpenerHash,
				DataSampleKeys: []string{trainDataSampleHash2},
				ID:             traintupleID2,
				AlgoKey:        algoHash,
				InModelsIDs:    []string{traintupleID1},
			},
		},
		Testtuples: []inputComputePlanTesttuple{
			inputComputePlanTesttuple{
				DataManagerKey: dataManagerOpenerHash,
				DataSampleKeys: []string{testDataSampleHash1, testDataSampleHash2},
				TraintupleID:   traintupleID2,
			},
		},
	}
)

func TestCreateComputePlanCompositeAggregate(t *testing.T) {
	scc := new(SubstraChaincode)
	mockStub := NewMockStubWithRegisterNode("substra", scc)
	registerItem(t, *mockStub, "aggregateAlgo")

	mockStub.MockTransactionStart("42")
	db := NewLedgerDB(mockStub)

	tag := []string{"compositeTraintuple1", "compositeTraintuple2", "aggregatetuple1", "aggregatetuple2"}

	inCP := inputComputePlan{
		ObjectiveKey: objectiveDescriptionHash,
		CompositeTraintuples: []inputComputePlanCompositeTraintuple{
			{
				DataManagerKey: dataManagerOpenerHash,
				DataSampleKeys: []string{trainDataSampleHash1},
				AlgoKey:        compositeAlgoHash,
				ID:             tag[0],
			},
			{
				DataManagerKey: dataManagerOpenerHash,
				DataSampleKeys: []string{trainDataSampleHash1},
				AlgoKey:        compositeAlgoHash,
				ID:             tag[1],
				InTrunkModelID: tag[0],
				InHeadModelID:  tag[0],
			},
		},
		Aggregatetuples: []inputComputePlanAggregatetuple{
			{
				AlgoKey: aggregateAlgoHash,
				ID:      tag[2],
			},
			{
				AlgoKey:     aggregateAlgoHash,
				ID:          tag[3],
				InModelsIDs: []string{tag[2]},
			},
		},
	}

	outCP, err := createComputePlanInternal(db, inCP)
	assert.NoError(t, err)

	// Check the composite traintuples
	traintuples, err := queryCompositeTraintuples(db, []string{})
	assert.NoError(t, err)
	require.Len(t, traintuples, 2)
	require.Contains(t, outCP.CompositeTraintupleKeys, traintuples[0].Key)
	require.Contains(t, outCP.CompositeTraintupleKeys, traintuples[1].Key)

	// Check the aggregate traintuples
	aggtuples, err := queryAggregatetuples(db, []string{})
	assert.NoError(t, err)
	require.Len(t, aggtuples, 2)
	require.Contains(t, outCP.AggregatetupleKeys, aggtuples[0].Key)
	require.Contains(t, outCP.AggregatetupleKeys, aggtuples[1].Key)

	// Query the compute plan
	cp, err := queryComputePlan(db, assetToArgs(inputHash{Key: outCP.ComputePlanID}))
	assert.NoError(t, err, "calling queryComputePlan should succeed")
	assert.NotNil(t, cp)
	assert.Equal(t, 2, len(cp.CompositeTraintupleKeys))
	assert.Equal(t, 2, len(cp.AggregatetupleKeys))

	// Query compute plans
	cps, err := queryComputePlans(db, []string{})
	assert.NoError(t, err, "calling queryComputePlans should succeed")
	assert.Len(t, cps, 1, "queryComputePlans should return one compute plan")
	assert.Equal(t, 2, len(cps[0].CompositeTraintupleKeys))
	assert.Equal(t, 2, len(cps[0].AggregatetupleKeys))
}

func TestCreateComputePlan(t *testing.T) {
	scc := new(SubstraChaincode)
	mockStub := NewMockStubWithRegisterNode("substra", scc)
	registerItem(t, *mockStub, "algo")

	mockStub.MockTransactionStart("42")
	db := NewLedgerDB(mockStub)

	// Simply test method and return values
	inCP := defaultComputePlan
	outCP, err := createComputePlanInternal(db, inCP)
	assert.NoError(t, err)
	validateDefaultComputePlan(t, outCP, defaultComputePlan)

	// Check the traintuples
	traintuples, err := queryTraintuples(db, []string{})
	assert.NoError(t, err)
	assert.Len(t, traintuples, 2)
	require.Contains(t, outCP.TraintupleKeys, traintuples[0].Key)
	require.Contains(t, outCP.TraintupleKeys, traintuples[1].Key)
	var first, second outputTraintuple
	for _, el := range traintuples {
		switch el.Key {
		case outCP.TraintupleKeys[0]:
			first = el
		case outCP.TraintupleKeys[1]:
			second = el
		}
	}

	// check first traintuple
	assert.NotZero(t, first)
	assert.EqualValues(t, first.Key, first.ComputePlanID)
	assert.Equal(t, inCP.Traintuples[0].AlgoKey, first.Algo.Hash)
	assert.Equal(t, StatusTodo, first.Status)

	// check second traintuple
	assert.NotZero(t, second)
	assert.EqualValues(t, first.Key, second.InModels[0].TraintupleKey)
	assert.EqualValues(t, first.ComputePlanID, second.ComputePlanID)
	assert.Len(t, second.InModels, 1)
	assert.Equal(t, inCP.Traintuples[1].AlgoKey, second.Algo.Hash)
	assert.Equal(t, StatusWaiting, second.Status)

	// Check the testtuples
	testtuples, err := queryTesttuples(db, []string{})
	assert.NoError(t, err)
	require.Len(t, testtuples, 1)
	testtuple := testtuples[0]
	require.Contains(t, outCP.TesttupleKeys, testtuple.Key)
	assert.EqualValues(t, second.Key, testtuple.TraintupleKey)
	assert.True(t, testtuple.Certified)
}

func TestQueryComputePlan(t *testing.T) {
	scc := new(SubstraChaincode)
	mockStub := NewMockStubWithRegisterNode("substra", scc)
	registerItem(t, *mockStub, "algo")

	mockStub.MockTransactionStart("42")
	db := NewLedgerDB(mockStub)

	// Simply test method and return values
	inCP := defaultComputePlan
	outCP, err := createComputePlanInternal(db, inCP)
	assert.NoError(t, err)
	assert.NotNil(t, outCP)
	assert.Equal(t, outCP.ComputePlanID, outCP.TraintupleKeys[0])

	cp, err := queryComputePlan(db, assetToArgs(inputHash{Key: outCP.ComputePlanID}))
	assert.NoError(t, err, "calling queryComputePlan should succeed")
	assert.NotNil(t, cp)
	validateDefaultComputePlan(t, cp, defaultComputePlan)
}

func TestQueryComputePlans(t *testing.T) {
	scc := new(SubstraChaincode)
	mockStub := NewMockStubWithRegisterNode("substra", scc)
	registerItem(t, *mockStub, "algo")

	mockStub.MockTransactionStart("42")
	db := NewLedgerDB(mockStub)

	// Simply test method and return values
	inCP := defaultComputePlan
	outCP, err := createComputePlanInternal(db, inCP)
	assert.NoError(t, err)
	assert.NotNil(t, outCP)
	assert.Equal(t, outCP.ComputePlanID, outCP.TraintupleKeys[0])

	cps, err := queryComputePlans(db, []string{})
	assert.NoError(t, err, "calling queryComputePlans should succeed")
	assert.Len(t, cps, 1, "queryComputePlans should return one compute plan")
	validateDefaultComputePlan(t, cps[0], defaultComputePlan)
}

func validateDefaultComputePlan(t *testing.T, cp outputComputePlan, in inputComputePlan) {
	assert.Len(t, cp.TraintupleKeys, 2)
	cpID := cp.TraintupleKeys[0]

	assert.Equal(t, in.ObjectiveKey, cp.ObjectiveKey)

	assert.Equal(t, cpID, cp.ComputePlanID)
	assert.Equal(t, in.ObjectiveKey, cp.ObjectiveKey)

	assert.NotEmpty(t, cp.TraintupleKeys[0])
	assert.NotEmpty(t, cp.TraintupleKeys[1])

	require.Len(t, cp.TesttupleKeys, 1)
	assert.NotEmpty(t, cp.TesttupleKeys[0])
}

func TestComputePlanEmptyTesttuples(t *testing.T) {
	scc := new(SubstraChaincode)
	mockStub := NewMockStubWithRegisterNode("substra", scc)
	registerItem(t, *mockStub, "algo")

	mockStub.MockTransactionStart("42")
	db := NewLedgerDB(mockStub)

	inCP := inputComputePlan{
		ObjectiveKey: objectiveDescriptionHash,
		Traintuples: []inputComputePlanTraintuple{
			inputComputePlanTraintuple{
				DataManagerKey: dataManagerOpenerHash,
				DataSampleKeys: []string{trainDataSampleHash1},
				AlgoKey:        algoHash,
				ID:             traintupleID1,
			},
			inputComputePlanTraintuple{
				DataManagerKey: dataManagerOpenerHash,
				DataSampleKeys: []string{trainDataSampleHash2},
				ID:             traintupleID2,
				AlgoKey:        algoHash,
				InModelsIDs:    []string{traintupleID1},
			},
		},
		Testtuples: []inputComputePlanTesttuple{},
	}

	outCP, err := createComputePlanInternal(db, inCP)
	assert.NoError(t, err)
	assert.NotNil(t, outCP)
	assert.Equal(t, outCP.ComputePlanID, outCP.TraintupleKeys[0])
	assert.Equal(t, []string{}, outCP.TesttupleKeys)

	cp, err := queryComputePlan(db, assetToArgs(inputHash{Key: outCP.ComputePlanID}))
	assert.NoError(t, err, "calling queryComputePlan should succeed")
	assert.NotNil(t, cp)
	assert.Equal(t, []string{}, outCP.TesttupleKeys)

	cps, err := queryComputePlans(db, []string{})
	assert.NoError(t, err, "calling queryComputePlans should succeed")
	assert.Len(t, cps, 1, "queryComputePlans should return one compute plan")
	assert.Equal(t, []string{}, cps[0].TesttupleKeys)
}

func TestQueryComputePlanEmpty(t *testing.T) {
	scc := new(SubstraChaincode)
	mockStub := NewMockStubWithRegisterNode("substra", scc)
	registerItem(t, *mockStub, "algo")

	mockStub.MockTransactionStart("42")
	db := NewLedgerDB(mockStub)

	cps, err := queryComputePlans(db, []string{})
	assert.NoError(t, err, "calling queryComputePlans should succeed")
	assert.Equal(t, []outputComputePlan{}, cps)
}
