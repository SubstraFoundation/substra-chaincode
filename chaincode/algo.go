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
)

// Set is a method of the receiver Algo. It uses inputAlgo fields to set the Algo
// Returns the algoKey
func (algo *Algo) Set(db *LedgerDB, inp inputAlgo) (err error) {
	// find associated owner
	owner, err := GetTxCreator(db.cc)
	if err != nil {
		return
	}

	permissions, err := NewPermissions(db, inp.Permissions)
	if err != nil {
		return
	}

	algo.Key = inp.Key
	algo.AssetType = AlgoType
	algo.Name = inp.Name
	algo.Checksum = inp.Checksum
	algo.StorageAddress = inp.StorageAddress
	algo.Description = &ChecksumAddress{
		Checksum:       inp.DescriptionChecksum,
		StorageAddress: inp.DescriptionStorageAddress,
	}
	algo.Owner = owner
	algo.Permissions = permissions
	algo.Metadata = inp.Metadata
	return
}

// -------------------------------------------------------------------------------------------
// Smart contracts related to an algo
// -------------------------------------------------------------------------------------------
// registerAlgo stores a new algo in the ledger.
// If the key exists, it will override the value with the new one
func registerAlgo(db *LedgerDB, args []string) (resp outputKey, err error) {
	inp := inputAlgo{}
	err = AssetFromJSON(args, &inp)
	if err != nil {
		return
	}
	// check validity of input args and convert it to Algo
	algo := Algo{}
	err = algo.Set(db, inp)
	if err != nil {
		return
	}
	// submit to ledger
	err = db.Add(algo.Key, algo)
	if err != nil {
		return
	}
	// create composite key
	err = db.CreateIndex("algo~owner~key", []string{"algo", algo.Owner, algo.Key})
	if err != nil {
		return
	}
	return outputKey{Key: algo.Key}, nil
}

// queryAlgo returns an algo of the ledger given its key
func queryAlgo(db *LedgerDB, args []string) (out outputAlgo, err error) {
	inp := inputKey{}
	err = AssetFromJSON(args, &inp)
	if err != nil {
		return
	}
	algo, err := db.GetAlgo(inp.Key)
	if err != nil {
		return
	}
	out.Fill(algo)
	return
}

// queryAlgos returns all algos of the ledger
func queryAlgos(db *LedgerDB, args []string) (outAlgos []outputAlgo, bookmark string, err error) {
	inp := inputBookmark{}
	outAlgos = []outputAlgo{}

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

	elementsKeys, bookmark, err := db.GetIndexKeysWithPagination("algo~owner~key", []string{"algo"}, OutputPageSize, inp.Bookmark)

	if err != nil {
		return
	}

	for _, key := range elementsKeys {
		algo, err := db.GetAlgo(key)
		if err != nil {
			return outAlgos, bookmark, err
		}
		var out outputAlgo
		out.Fill(algo)
		outAlgos = append(outAlgos, out)
	}
	return
}
