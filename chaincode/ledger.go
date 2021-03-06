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

// ---------------------------------------------------------------------------------
// Representation of elements stored in the ledger
// ---------------------------------------------------------------------------------

// AssetType is use to check the type of an asset
type AssetType uint8

// Const representing the types of asset findable in the ledger
const (
	ObjectiveType AssetType = iota
	DataManagerType
	DataSampleType
	AlgoType
	CompositeAlgoType
	AggregateAlgoType
	TraintupleType
	CompositeTraintupleType
	AggregatetupleType
	TesttupleType
	ComputePlanType
	// when adding a new type here, don't forget to update
	// the String() function in utils.go
)

// StatusUpdater is exported
type StatusUpdater interface {
	commitStatusUpdate(db *LedgerDB, key string, status string) error
}

// Objective is the representation of one of the element type stored in the ledger
type Objective struct {
	Key         string               `json:"key"`
	Name        string               `json:"name"`
	AssetType   AssetType            `json:"asset_type"`
	Description *ChecksumAddress     `json:"description"`
	Metrics     *ChecksumAddressName `json:"metrics"`
	Owner       string               `json:"owner"`
	TestDataset *Dataset             `json:"test_dataset"`
	Permissions Permissions          `json:"permissions"`
	Metadata    map[string]string    `json:"metadata"`
}

// DataManager is the representation of one of the elements type stored in the ledger
type DataManager struct {
	Key          string            `json:"key"`
	Name         string            `json:"name"`
	AssetType    AssetType         `json:"asset_type"`
	Opener       *ChecksumAddress  `json:"opener"`
	Type         string            `json:"type"`
	Description  *ChecksumAddress  `json:"description"`
	Owner        string            `json:"owner"`
	ObjectiveKey string            `json:"objective_key"`
	Permissions  Permissions       `json:"permissions"`
	Metadata     map[string]string `json:"metadata"`
}

// DataSample is the representation of one of the element type stored in the ledger
type DataSample struct {
	AssetType       AssetType `json:"asset_type"`
	DataManagerKeys []string  `json:"data_manager_keys"`
	Owner           string    `json:"owner"`
	TestOnly        bool      `json:"testOnly"`
}

// Algo is the representation of one of the element type stored in the ledger
type Algo struct {
	Key            string            `json:"key"`
	Name           string            `json:"name"`
	AssetType      AssetType         `json:"asset_type"`
	Checksum       string            `json:"checksum"`
	StorageAddress string            `json:"storage_address"`
	Description    *ChecksumAddress  `json:"description"`
	Owner          string            `json:"owner"`
	Permissions    Permissions       `json:"permissions"`
	Metadata       map[string]string `json:"metadata"`
}

// CompositeAlgo is the representation of one of the element type stored in the ledger
type CompositeAlgo struct {
	Algo
}

// AggregateAlgo is the representation of one of the element type stored in the ledger
type AggregateAlgo struct {
	Algo
}

// GenericTuple is a structure that contains the fields
// that are common to Traintuple, CompositeTraintuple and
// AggregateTuple
type GenericTuple struct {
	AssetType      AssetType         `json:"asset_type"`
	AlgoKey        string            `json:"algo_key"`
	ComputePlanKey string            `json:"compute_plan_key"`
	Creator        string            `json:"creator"`
	Log            string            `json:"log"`
	Metadata       map[string]string `json:"metadata"`
	Rank           int               `json:"rank"`
	Status         string            `json:"status"`
	Tag            string            `json:"tag"`
}

// Traintuple is the representation of one the element type stored in the ledger. It describes a training task occuring on the platform
type Traintuple struct {
	Key            string              `json:"key"`
	AssetType      AssetType           `json:"asset_type"`
	AlgoKey        string              `json:"algo_key"`
	ComputePlanKey string              `json:"compute_plan_key"`
	Creator        string              `json:"creator"`
	Log            string              `json:"log"`
	Metadata       map[string]string   `json:"metadata"`
	Rank           int                 `json:"rank"`
	Status         string              `json:"status"`
	Tag            string              `json:"tag"`
	Dataset        *Dataset            `json:"dataset"`
	InModelKeys    []string            `json:"in_models"`
	OutModel       *KeyChecksumAddress `json:"out_model"`
	Permissions    Permissions         `json:"permissions"`
}

// CompositeTraintuple is like a traintuple, but for composite model composition
type CompositeTraintuple struct {
	Key            string                          `json:"key"`
	AssetType      AssetType                       `json:"asset_type"`
	AlgoKey        string                          `json:"algo_key"`
	ComputePlanKey string                          `json:"compute_plan_key"`
	Creator        string                          `json:"creator"`
	Log            string                          `json:"log"`
	Metadata       map[string]string               `json:"metadata"`
	Rank           int                             `json:"rank"`
	Status         string                          `json:"status"`
	Tag            string                          `json:"tag"`
	Dataset        *Dataset                        `json:"dataset"`
	InHeadModel    string                          `json:"in_head_model"`
	InTrunkModel   string                          `json:"in_trunk_model"`
	OutHeadModel   CompositeTraintupleOutHeadModel `json:"out_head_model"`
	OutTrunkModel  CompositeTraintupleOutModel     `json:"out_trunk_model"`
}

// Aggregatetuple is like a traintuple, but for aggregate model composition
type Aggregatetuple struct {
	Key            string              `json:"key"`
	AssetType      AssetType           `json:"asset_type"`
	AlgoKey        string              `json:"algo_key"`
	ComputePlanKey string              `json:"compute_plan_key"`
	Creator        string              `json:"creator"`
	Log            string              `json:"log"`
	Metadata       map[string]string   `json:"metadata"`
	Rank           int                 `json:"rank"`
	Status         string              `json:"status"`
	Tag            string              `json:"tag"`
	InModelKeys    []string            `json:"in_models"`
	OutModel       *KeyChecksumAddress `json:"out_model"`
	Permissions    Permissions         `json:"permissions"` // TODO (aggregate): what do permissions mean here?
	Worker         string              `json:"worker"`
}

// CompositeTraintupleOutModel is the out-model of a CompositeTraintuple
type CompositeTraintupleOutModel struct {
	OutModel    *KeyChecksumAddress `json:"out_model"`
	Permissions Permissions         `json:"permissions"`
}

// CompositeTraintupleOutHeadModel is the out-model of a CompositeTraintuple
type CompositeTraintupleOutHeadModel struct {
	OutModel    *KeyChecksum `json:"out_model"`
	Permissions Permissions  `json:"permissions"`
}

// Testtuple is the representation of one the element type stored in the ledger. It describes a training task occuring on the platform
type Testtuple struct {
	Key            string            `json:"key"`
	AlgoKey        string            `json:"algo"`
	AssetType      AssetType         `json:"asset_type"`
	Certified      bool              `json:"certified"`
	ComputePlanKey string            `json:"compute_plan_key"`
	Creator        string            `json:"creator"`
	Dataset        *TtDataset        `json:"dataset"`
	Log            string            `json:"log"`
	Metadata       map[string]string `json:"metadata"`
	TraintupleKey  string            `json:"traintuple_key"`
	ObjectiveKey   string            `json:"objective"`
	Permissions    Permissions       `json:"permissions"`
	Rank           int               `json:"rank"`
	Status         string            `json:"status"`
	Tag            string            `json:"tag"`
}

// ComputePlan is the ledger's representation of a compute plan.
type ComputePlan struct {
	Key                     string               `json:"key"`
	AggregatetupleKeys      []string             `json:"aggregatetuple_keys"`
	AssetType               AssetType            `json:"asset_type"`
	CleanModels             bool                 `json:"clean_models"` // whether or not to delete intermediary models
	CompositeTraintupleKeys []string             `json:"composite_traintuple_keys"`
	IDToTrainTask           map[string]TrainTask `json:"id_to_train_task"`
	Metadata                map[string]string    `json:"metadata"`
	State                   ComputePlanState     `json:"-"` // "-" means this field is excluded from JSON (de)serialization
	StateKey                string               `json:"state_key"`
	Tag                     string               `json:"tag"`
	TesttupleKeys           []string             `json:"testtuple_keys"`
	TraintupleKeys          []string             `json:"traintuple_keys"`
	Workers                 []string             `json:"workers"`
}

// ComputePlanState is the ledger's representation of the compute plan state.
// To minimize the size of every compute plan, update its state record under another
// key in the ledger. It will reduce the growing rate of the blockchain size.
type ComputePlanState struct {
	Status string `json:"status"`
}

// ComputePlanWorkerState contains state information for a given
// compute plan and worker
type ComputePlanWorkerState struct {
	IntermediaryModelsInUse []string `json:"intermediary_models_in_use"`
	DoneCount               int      `json:"done_count"`  // the number of tuples in the "done" state for this compute plan and worker
	TupleCount              int      `json:"tuple_count"` // the total number of tuples registered for this compute plan and worker
}

// TrainTask is represent the information for one tuple in a Compute Plan
type TrainTask struct {
	Depth int    `json:"depth"`
	Key   string `json:"key"`
}

// ---------------------------------------------------------------------------------
// Struct used in the representation of elements stored in the ledger
// ---------------------------------------------------------------------------------

// KeyChecksum ...
type KeyChecksum struct {
	Key      string `json:"key"`
	Checksum string `json:"checksum"`
}

// ChecksumAddress stores a checksum and a Storage Address
type ChecksumAddress struct {
	Checksum       string `json:"checksum"`
	StorageAddress string `json:"storage_address"`
}

// KeyChecksumAddress ...
type KeyChecksumAddress struct {
	Key            string `json:"key"`
	Checksum       string `json:"checksum"`
	StorageAddress string `json:"storage_address"`
}

// ChecksumAddressName stores a checksum, a storage address, and a name
type ChecksumAddressName struct {
	Checksum       string `json:"checksum"`
	StorageAddress string `json:"storage_address"`
	Name           string `json:"name"`
}

// KeyChecksumAddressName ...
type KeyChecksumAddressName struct {
	Key            string `json:"key"`
	Checksum       string `json:"checksum"`
	StorageAddress string `json:"storage_address"`
	Name           string `json:"name"`
}

// Model stores the traintupleKey leading to the model, its checksum and storage address
type Model struct {
	Key            string `json:"key"`
	TraintupleKey  string `json:"traintuple_key"`
	Checksum       string `json:"checksum"`
	StorageAddress string `json:"storage_address"`
}

// Dataset stores info about a dataManagerKey and a list of associated dataSample
type Dataset struct {
	DataManagerKey string            `json:"data_manager_key"`
	DataSampleKeys []string          `json:"data_sample_keys"`
	Metadata       map[string]string `json:"metadata"`
	Worker         string            `json:"worker"`
}

// ---------------------------------------------------------------------------------
// Struct used in the representation of outputs when querying some elements
// ---------------------------------------------------------------------------------

// TtDataset stores info about dataset in a Traintyple (train or test data) and in a PredTuple (later)
type TtDataset struct {
	Key            string   `json:"key"`
	Worker         string   `json:"worker"`
	DataSampleKeys []string `json:"data_sample_keys"`
	OpenerChecksum string   `json:"opener_checksum"`
	Perf           float32  `json:"perf"`
}

// TtObjective stores info about a objective in a Traintuple
type TtObjective struct {
	Key     string           `json:"key"`
	Metrics *ChecksumAddress `json:"metrics"`
}

// Node stores informations about node registered into the network,
// would be used to list authorized nodes for permissions
type Node struct {
	ID string `json:"id"`
}
