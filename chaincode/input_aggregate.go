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

// inputAggregatetuple is the representation of input args to register an aggregate Tuple
type inputAggregatetuple struct {
	Key            string            `validate:"required,len=36" json:"key"`
	AlgoKey        string            `validate:"required,len=36" json:"algo_key"`
	InModels       []string          `validate:"omitempty,dive,len=36" json:"in_models"`
	ComputePlanKey string            `validate:"required_with=Rank" json:"compute_plan_key"`
	Metadata       map[string]string `validate:"lte=100,dive,keys,lte=50,endkeys,lte=100" json:"metadata"`
	Rank           string            `json:"rank"`
	Tag            string            `validate:"omitempty,lte=64" json:"tag"`
	Worker         string            `validate:"required" json:"worker"`
}

type inputAggregateAlgo struct {
	inputAlgo
}
