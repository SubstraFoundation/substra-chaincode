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
	"chaincode/errors"
	"strconv"
)

// -------------------------------------------------------------------------------------------
// Methods on receivers traintuple
// -------------------------------------------------------------------------------------------

// SetFromInput is a method of the receiver Traintuple.
// It uses the inputTraintuple to check and set the traintuple's parameters
// which don't depend on previous traintuples values :
//  - AssetType
//  - Creator & permissions
//  - Tag
//  - AlgoKey & ObjectiveKey
//  - Dataset
func (traintuple *Traintuple) SetFromInput(db *LedgerDB, inp inputTraintuple) error {

	// TODO later: check permissions
	// find associated creator and check permissions (TODO later)
	creator, err := GetTxCreator(db.cc)
	if err != nil {
		return err
	}
	traintuple.Key = inp.Key
	traintuple.AssetType = TraintupleType
	traintuple.Creator = creator
	traintuple.ComputePlanKey = inp.ComputePlanKey
	traintuple.Metadata = inp.Metadata
	traintuple.Tag = inp.Tag
	algo, err := db.GetAlgo(inp.AlgoKey)
	if err != nil {
		return errors.BadRequest(err, "could not retrieve algo with key %s", inp.AlgoKey)
	}
	if !algo.Permissions.CanProcess(algo.Owner, creator) {
		return errors.Forbidden("not authorized to process algo %s", inp.AlgoKey)
	}
	traintuple.AlgoKey = inp.AlgoKey

	// check if DataSampleKeys are from the same dataManager and if they are not test only dataSample
	_, trainOnly, err := checkSameDataManager(db, inp.DataManagerKey, inp.DataSampleKeys)
	if err != nil {
		return err
	}
	if !trainOnly {
		return errors.BadRequest("not possible to create a traintuple with test only data")
	}

	dataManager, err := db.GetDataManager(inp.DataManagerKey)
	if err != nil {
		return errors.BadRequest(err, "could not retrieve dataManager with key \"%s\"", inp.DataManagerKey)
	}
	if !dataManager.Permissions.CanProcess(dataManager.Owner, creator) {
		return errors.Forbidden("not authorized to process dataManager %s", inp.DataManagerKey)
	}

	traintuple.Permissions = MergePermissions(dataManager.Permissions, algo.Permissions)

	// fill traintuple.Dataset from dataManager and dataSample
	traintuple.Dataset = &Dataset{
		DataManagerKey: inp.DataManagerKey,
		DataSampleKeys: inp.DataSampleKeys,
	}
	traintuple.Dataset.Worker, err = getDataManagerOwner(db, traintuple.Dataset.DataManagerKey)
	return err
}

// SetFromParents set the status of the traintuple depending on its "parents",
// i.e. the traintuples from which it received the outModels as inModels.
// Also it's InModelKeys are set.
func (traintuple *Traintuple) SetFromParents(db *LedgerDB, inModels []string) error {
	var parentStatuses []string
	inModelKeys := traintuple.InModelKeys

	for _, parentTraintupleKey := range inModels {
		tuple, err := db.GetGenericTuple(parentTraintupleKey)
		if err != nil {
			return errors.BadRequest(err, "could not retrieve parent traintuple with key %s", parentTraintupleKey)
		}
		if !typeInSlice(tuple.AssetType, []AssetType{TraintupleType, CompositeTraintupleType, AggregatetupleType}) {
			return errors.Internal("aggregate.SetFromParents: Unsupported parent type %s", tuple.AssetType)
		}
		parentStatuses = append(parentStatuses, tuple.Status)
		inModelKeys = append(inModelKeys, parentTraintupleKey)
	}
	traintuple.Status = determineStatusFromInModels(parentStatuses)
	traintuple.InModelKeys = inModelKeys
	return nil
}

// AddToComputePlan set the traintuple's parameters that determines if it's part of on ComputePlan and how.
// It uses the inputTraintuple values as follow:
//  - If neither ComputePlanKey nor rank is set it returns immediately
//  - If rank is 0 and ComputePlanKey empty, it's start a new one using this traintuple key
//  - If rank and ComputePlanKey are set, it checks if there are coherent with previous ones and set it.
// Use checkComputePlanAvailability to ensure the compute plan exists and no other tuple is registered with the same worker/rank
func (traintuple *Traintuple) AddToComputePlan(db *LedgerDB, inp inputTraintuple, traintupleKey string, checkComputePlanAvailability bool) error {
	// check ComputePlanKey and Rank and set it when required
	var err error
	if inp.Rank == "" {
		if inp.ComputePlanKey != "" {
			return errors.BadRequest("invalid inputs, a ComputePlan should have a rank")
		}
		return nil
	}
	traintuple.Rank, err = strconv.Atoi(inp.Rank)
	if err != nil {
		return err
	}
	traintuple.ComputePlanKey = inp.ComputePlanKey
	computePlan, err := db.GetComputePlan(inp.ComputePlanKey)
	if err != nil {
		return err
	}
	err = computePlan.AddTuple(db, TraintupleType, traintupleKey, traintuple.Status, traintuple.Dataset.Worker)
	if err != nil {
		return err
	}
	err = computePlan.Save(db, traintuple.ComputePlanKey)
	if err != nil {
		return err
	}

	if !checkComputePlanAvailability {
		return nil
	}
	var ttKeys []string
	ttKeys, err = db.GetIndexKeys("computePlan~computeplankey~worker~rank~key", []string{"computePlan", inp.ComputePlanKey, traintuple.Dataset.Worker, inp.Rank})
	if err != nil {
		return err
	} else if len(ttKeys) > 0 {
		err = errors.BadRequest("ComputePlanKey %s with worker %s rank %d already exists", inp.ComputePlanKey, traintuple.Dataset.Worker, traintuple.Rank)
		return err
	}
	return nil
}

// Save will put in the legder interface both the traintuple with its key
// and all the associated composite keys
func (traintuple *Traintuple) Save(db *LedgerDB, traintupleKey string) error {

	// store in ledger
	if err := db.Add(traintupleKey, traintuple); err != nil {
		return err
	}

	// create composite keys
	if err := db.CreateIndex("traintuple~algo~key", []string{"traintuple", traintuple.AlgoKey, traintupleKey}); err != nil {
		return err
	}
	if err := db.CreateIndex("traintuple~worker~status~key", []string{"traintuple", traintuple.Dataset.Worker, traintuple.Status, traintupleKey}); err != nil {
		return err
	}
	for _, inModelKey := range traintuple.InModelKeys {
		if err := db.CreateIndex("tuple~inModel~key", []string{"tuple", inModelKey, traintupleKey}); err != nil {
			return err
		}
	}
	if traintuple.ComputePlanKey != "" {
		if err := db.CreateIndex("computePlan~computeplankey~worker~rank~key", []string{"computePlan", traintuple.ComputePlanKey, traintuple.Dataset.Worker, strconv.Itoa(traintuple.Rank), traintupleKey}); err != nil {
			return err
		}
		if err := db.CreateIndex("algo~computeplankey~key", []string{"algo", traintuple.ComputePlanKey, traintuple.AlgoKey}); err != nil {
			return err
		}
	}
	if traintuple.Tag != "" {
		err := db.CreateIndex("traintuple~tag~key", []string{"traintuple", traintuple.Tag, traintupleKey})
		if err != nil {
			return err
		}
	}
	return nil
}

// -------------------------------------------------------------------------------------------
// Smart contracts related to traintuples
// -------------------------------------------------------------------------------------------

// createTraintuple adds a Traintuple in the ledger
func createTraintuple(db *LedgerDB, args []string) (outputKey, error) {
	inp := inputTraintuple{}
	err := AssetFromJSON(args, &inp)
	if err != nil {
		return outputKey{}, err
	}

	key, err := createTraintupleInternal(db, inp, true)

	if err != nil {
		return outputKey{}, err
	}

	return outputKey{Key: key}, nil
}

func createTraintupleInternal(db *LedgerDB, inp inputTraintuple, checkComputePlanAvailability bool) (string, error) {
	traintuple := Traintuple{}
	err := traintuple.SetFromInput(db, inp)
	if err != nil {
		return "", err
	}
	err = traintuple.SetFromParents(db, inp.InModels)
	if err != nil {
		return "", err
	}
	// Test if the key (ergo the traintuple) already exists
	tupleExists, err := db.KeyExists(traintuple.Key)
	if err != nil {
		return "", err
	}
	if tupleExists {
		return "", errors.Conflict("traintuple already exists").WithKey(traintuple.Key)
	}
	err = traintuple.AddToComputePlan(db, inp, traintuple.Key, checkComputePlanAvailability)
	if err != nil {
		return "", err
	}

	err = traintuple.Save(db, traintuple.Key)
	if err != nil {
		return "", err
	}

	err = db.AddTupleEvent(traintuple.Key)
	if err != nil {
		return "", err
	}

	return traintuple.Key, nil
}

// logStartTrain modifies a traintuple by changing its status from todo to doing
func logStartTrain(db *LedgerDB, args []string) (o outputTraintuple, err error) {
	status := StatusDoing
	inp := inputKey{}
	err = AssetFromJSON(args, &inp)
	if err != nil {
		return
	}

	// get traintuple, check validity of the update
	traintuple, err := db.GetTraintuple(inp.Key)
	if err != nil {
		return
	}

	if err = validateTupleOwner(db, traintuple.Dataset.Worker); err != nil {
		return
	}
	if err = traintuple.commitStatusUpdate(db, inp.Key, status); err != nil {
		return
	}
	err = o.Fill(db, traintuple)
	return
}

// logSuccessTrain modifies a traintuple by changing its status from doing to done
// reports logs and associated performances
func logSuccessTrain(db *LedgerDB, args []string) (o outputTraintuple, err error) {
	status := StatusDone
	inp := inputLogSuccessTrain{}
	err = AssetFromJSON(args, &inp)
	if err != nil {
		return
	}
	traintupleKey := inp.Key

	// get, update and commit traintuple
	traintuple, err := db.GetTraintuple(traintupleKey)
	if err != nil {
		return
	}

	traintuple.OutModel = &KeyChecksumAddress{
		Key:            inp.OutModel.Key,
		Checksum:       inp.OutModel.Checksum,
		StorageAddress: inp.OutModel.StorageAddress}
	traintuple.Log += inp.Log

	err = createModelIndex(db, inp.OutModel.Key, traintupleKey)
	if err != nil {
		return
	}

	if err = validateTupleOwner(db, traintuple.Dataset.Worker); err != nil {
		return
	}

	if err = traintuple.commitStatusUpdate(db, traintupleKey, status); err != nil {
		return
	}

	err = TryAddIntermediaryModel(db, traintuple.ComputePlanKey, traintuple.Dataset.Worker, traintupleKey, traintuple.OutModel.Key)
	if err != nil {
		return
	}

	// update depending tuples
	err = UpdateTraintupleChildren(db, traintupleKey, traintuple.Status, []string{})
	if err != nil {
		return
	}

	err = UpdateTesttupleChildren(db, traintupleKey, traintuple.Status)
	if err != nil {
		return
	}

	err = o.Fill(db, traintuple)
	return
}

// logFailTrain modifies a traintuple by changing its status to fail and reports associated logs
func logFailTrain(db *LedgerDB, args []string) (o outputTraintuple, err error) {
	status := StatusFailed
	inp := inputLogFailTrain{}
	err = AssetFromJSON(args, &inp)
	if err != nil {
		return
	}

	// get, update and commit traintuple
	traintuple, err := db.GetTraintuple(inp.Key)
	if err != nil {
		return
	}

	traintuple.Log += inp.Log

	if err = validateTupleOwner(db, traintuple.Dataset.Worker); err != nil {
		return
	}
	if err = traintuple.commitStatusUpdate(db, inp.Key, status); err != nil {
		return
	}

	if err = o.Fill(db, traintuple); err != nil {
		return
	}

	// Do not propagate failure if we are in a compute plan
	if traintuple.ComputePlanKey != "" {
		return
	}
	// update depending tuples
	err = UpdateTesttupleChildren(db, inp.Key, traintuple.Status)
	if err != nil {
		return
	}

	err = UpdateTraintupleChildren(db, inp.Key, traintuple.Status, []string{})
	return
}

// queryTraintuple returns info about a traintuple given its key
func queryTraintuple(db *LedgerDB, args []string) (outputTraintuple outputTraintuple, err error) {
	inp := inputKey{}
	err = AssetFromJSON(args, &inp)
	if err != nil {
		return
	}
	traintuple, err := db.GetTraintuple(inp.Key)
	if err != nil {
		return
	}
	if traintuple.AssetType != TraintupleType {
		err = errors.NotFound("no element with key %s", inp.Key)
		return
	}
	err = outputTraintuple.Fill(db, traintuple)
	return
}

// queryTraintuples returns all traintuples
func queryTraintuples(db *LedgerDB, args []string) (outTraintuples []outputTraintuple, bookmark string, err error) {
	inp := inputBookmark{}
	outTraintuples = []outputTraintuple{}

	if len(args) > 1 {
		err = errors.BadRequest("incorrect number of arguments, expecting at most one argument")
		return
	}

	if len(args) == 1 && args[0] != "" {
		err = AssetFromJSON(args, &inp)
		if err != nil {
			return
		}
	}

	elementsKeys, bookmark, err := db.GetIndexKeysWithPagination("traintuple~algo~key", []string{"traintuple"}, OutputPageSize, inp.Bookmark)

	if err != nil {
		return
	}

	for _, key := range elementsKeys {
		outputTraintuple, err := getOutputTraintuple(db, key)
		if err != nil {
			return outTraintuples, bookmark, err
		}
		outTraintuples = append(outTraintuples, outputTraintuple)
	}
	return
}

// -----------------------------------------------
// Utils for smartcontracts related to traintuples
// -----------------------------------------------

// getOutputTraintuple takes as input a traintuple key and returns the outputTraintuple
func getOutputTraintuple(db *LedgerDB, traintupleKey string) (outTraintuple outputTraintuple, err error) {
	traintuple, err := db.GetTraintuple(traintupleKey)
	if err != nil {
		return
	}
	err = outTraintuple.Fill(db, traintuple)
	return
}

// getOutputTraintuples takes as input a list of keys and returns a paylaod containing a list of associated retrieved elements
func getOutputTraintuples(db *LedgerDB, traintupleKeys []string) (outTraintuples []outputTraintuple, err error) {
	nb := getLimitedNbSliceElements(traintupleKeys)
	for _, key := range traintupleKeys[:nb] {
		var outputTraintuple outputTraintuple
		outputTraintuple, err = getOutputTraintuple(db, key)
		if err != nil {
			return
		}
		outTraintuples = append(outTraintuples, outputTraintuple)
	}
	return
}

// validateNewStatus verifies that the new status is consistent with the tuple current status
func (traintuple *Traintuple) validateNewStatus(db *LedgerDB, status string) error {
	// check validity of worker and change of status
	return checkUpdateTuple(db, traintuple.Dataset.Worker, traintuple.Status, status)
}

// UpdateTraintupleChildren updates the status of waiting trainuples  InModels of traintuples once they have been trained (succesfully or failed)
func UpdateTraintupleChildren(db *LedgerDB, traintupleKey string, traintupleStatus string, alreadyUpdatedKeys []string) error {
	// get keys from tuple having as inModels the input traintuple
	allChildKeys, err := db.GetIndexKeys("tuple~inModel~key", []string{"tuple", traintupleKey})
	if err != nil {
		return errors.Internal("error while getting associated tuples to update their inModel, tupleKey=%s tupleStatus=%s %s", traintupleKey, traintupleStatus, err)
	}
	for _, childTraintupleKey := range allChildKeys {
		if stringInSlice(childTraintupleKey, alreadyUpdatedKeys) {
			continue
		}
		child, err := db.GetGenericTuple(childTraintupleKey)
		if err != nil {
			return err
		}

		if stringInSlice(child.Status, []string{StatusFailed, StatusAborted}) {
			// traintuple is already failed, don't update it
			continue
		}

		if child.Status != StatusWaiting {
			return errors.Internal("traintuple %s has invalid status : '%s' instead of waiting", childTraintupleKey, child.Status)
		}

		childTraintupleStatus := child.Status

		// Update the child traintuple and get its new status
		switch child.AssetType {
		case TraintupleType:
			childTraintupleStatus, err = UpdateTraintupleChild(db, traintupleKey, childTraintupleKey, traintupleStatus)
			if err != nil {
				return err
			}
		case CompositeTraintupleType:
			childTraintupleStatus, err = UpdateCompositeTraintupleChild(db, traintupleKey, childTraintupleKey, traintupleStatus)
			if err != nil {
				return err
			}
		case AggregatetupleType:
			childTraintupleStatus, err = UpdateAggregatetupleChild(db, traintupleKey, childTraintupleKey, traintupleStatus)
			if err != nil {
				return err
			}
		default:
			return errors.Internal("Unknown child traintuple type: %s", child.AssetType)
		}

		alreadyUpdatedKeys = append(alreadyUpdatedKeys, childTraintupleKey)
		if stringInSlice(traintupleStatus, []string{StatusFailed, StatusAborted}) {
			// Recursively call for an update on this child's children
			err = UpdateTesttupleChildren(db, childTraintupleKey, childTraintupleStatus)
			if err != nil {
				return err
			}

			err = UpdateTraintupleChildren(db, childTraintupleKey, childTraintupleStatus, alreadyUpdatedKeys)
			if err != nil {
				return err
			}
		}
	}
	return nil
}

// UpdateTraintupleChild updates the status of a waiting trainuple, given the new parent traintuple status
func UpdateTraintupleChild(db *LedgerDB, parentTraintupleKey string, childTraintupleKey string, traintupleStatus string) (childStatus string, err error) {
	// get and update traintuple
	childTraintuple, err := db.GetTraintuple(childTraintupleKey)
	if err != nil {
		return
	}

	childStatus = childTraintuple.Status

	// get traintuple new status
	var newStatus string
	if traintupleStatus == StatusFailed {
		newStatus = StatusFailed
	} else if traintupleStatus == StatusDone {
		ready, _err := childTraintuple.isReady(db, parentTraintupleKey)
		if _err != nil {
			err = _err
			return
		}
		if ready {
			newStatus = StatusTodo
		}
	}

	// commit new status
	if newStatus == "" {
		return
	}
	if err = childTraintuple.commitStatusUpdate(db, childTraintupleKey, newStatus); err != nil {
		return
	}

	// update return value after status update
	childStatus = childTraintuple.Status

	err = db.AddTupleEvent(childTraintupleKey)

	return
}

func (traintuple *Traintuple) isReady(db *LedgerDB, newDoneTraintupleKey string) (ready bool, err error) {
	return IsReady(db, traintuple.InModelKeys, newDoneTraintupleKey)
}

// IsReady checks if inModels of a traintuple have been trained, except the newDoneTraintupleKey (since the transaction is not commited)
func IsReady(db *LedgerDB, inModelKeys []string, newDoneTraintupleKey string) (ready bool, err error) {
	for _, key := range inModelKeys {
		// don't check newly done traintuple
		if key == newDoneTraintupleKey {
			continue
		}
		tuple, err := db.GetGenericTuple(key)
		if err != nil {
			return false, err
		}
		if tuple.Status != StatusDone {
			return false, nil
		}
	}
	return true, nil
}

// commitStatusUpdate update the traintuple status in the ledger
func (traintuple *Traintuple) commitStatusUpdate(db *LedgerDB, traintupleKey string, newStatus string) error {
	if traintuple.Status == newStatus {
		return nil
	}

	// do not update if previous status is already Done, Failed, Todo, Doing
	if StatusAborted == newStatus && traintuple.Status != StatusWaiting {
		return nil
	}

	if err := traintuple.validateNewStatus(db, newStatus); err != nil {
		return errors.Internal("update traintuple %s failed: %s", traintupleKey, err.Error())
	}

	oldStatus := traintuple.Status
	traintuple.Status = newStatus
	if err := db.Put(traintupleKey, traintuple); err != nil {
		return errors.Internal("failed to update traintuple %s - %s", traintupleKey, err.Error())
	}

	// update associated composite keys
	indexName := "traintuple~worker~status~key"
	oldAttributes := []string{"traintuple", traintuple.Dataset.Worker, oldStatus, traintupleKey}
	newAttributes := []string{"traintuple", traintuple.Dataset.Worker, traintuple.Status, traintupleKey}
	if err := db.UpdateIndex(indexName, oldAttributes, newAttributes); err != nil {
		return err
	}
	if err := UpdateComputePlanState(db, traintuple.ComputePlanKey, newStatus, traintupleKey, traintuple.Dataset.Worker); err != nil {
		return err
	}
	logger.Infof("traintuple %s status updated: %s (from=%s)", traintupleKey, newStatus, oldStatus)
	return nil
}

// UpdateTesttupleChildren update testtuples status associated with a done or failed traintuple
func UpdateTesttupleChildren(db *LedgerDB, traintupleKey string, traintupleStatus string) error {
	var newStatus string
	switch {
	case traintupleStatus == StatusFailed:
		newStatus = StatusFailed
	case traintupleStatus == StatusDone:
		newStatus = StatusTodo
	default:
		return nil
	}

	indexName := "testtuple~traintuple~certified~key"
	// get testtuple associated with this traintuple and updates its status
	testtupleKeys, err := db.GetIndexKeys(indexName, []string{"testtuple", traintupleKey})
	if err != nil {
		return err
	}
	for _, testtupleKey := range testtupleKeys {
		// get and update testtuple
		testtuple, err := db.GetTesttuple(testtupleKey)
		if err != nil {
			return err
		}

		if testtuple.Status == StatusAborted {
			continue
		}

		testtuple.TraintupleKey = traintupleKey

		if err := testtuple.commitStatusUpdate(db, testtupleKey, newStatus); err != nil {
			return err
		}

		err = db.AddTupleEvent(testtupleKey)
		if err != nil {
			return err
		}
	}
	return nil
}
