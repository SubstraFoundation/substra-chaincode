### Examples

#### ------------ Add Node ------------
##### Command peer example:
```bash
peer chaincode invoke -n mycc -c '{"Args":["registerNode"]}' -C myc
```
##### Command output:
```json
{
 "id": "SampleOrg"
}
```
#### ------------ Add DataManager ------------
Smart contract: `registerDataManager`

##### JSON Inputs:
```go
{
 "key": string (required,len=36),
 "name": string (required,gte=1,lte=100),
 "opener_checksum": string (required,len=64,hexadecimal),
 "opener_storage_address": string (required,url),
 "type": string (required,gte=1,lte=30),
 "description_checksum": string (required,len=64,hexadecimal),
 "description_storage_address": string (required,url),
 "objective_key": string (omitempty,len=36),
 "permissions": (required){
   "process": (required){
     "public": bool (required),
     "authorized_ids": [string] (required),
   },
 },
 "metadata": map (lte=100,dive,keys,lte=50,endkeys,lte=100),
}
```
##### Command peer example:
```bash
peer chaincode invoke -n mycc -c '{"Args":["registerDataManager","{\"key\":\"da1bb7c3-1f62-244c-0f3a-761cc1688042\",\"name\":\"liver slide\",\"opener_checksum\":\"da1bb7c31f62244c0f3a761cc168804227115793d01c270021fe3f7935482dcc\",\"opener_storage_address\":\"https://toto/dataManager/42234/opener\",\"type\":\"images\",\"description_checksum\":\"8d4bb7c31f62244c0f3a761cc168804227115793d01c270021fe3f7935482eee\",\"description_storage_address\":\"https://toto/dataManager/42234/description\",\"objective_key\":\"\",\"permissions\":{\"process\":{\"public\":true,\"authorized_ids\":[]}},\"metadata\":null}"]}' -C myc
```
##### Command output:
```json
{
 "key": "da1bb7c3-1f62-244c-0f3a-761cc1688042"
}
```
#### ------------ Query DataManager From key ------------
Smart contract: `queryDataManager`

##### JSON Inputs:
```go
{
 "key": string (required,len=36),
}
```
##### Command peer example:
```bash
peer chaincode invoke -n mycc -c '{"Args":["queryDataManager","{\"key\":\"da1bb7c3-1f62-244c-0f3a-761cc1688042\"}"]}' -C myc
```
##### Command output:
```json
{
 "description": {
  "checksum": "8d4bb7c31f62244c0f3a761cc168804227115793d01c270021fe3f7935482eee",
  "storage_address": "https://toto/dataManager/42234/description"
 },
 "key": "da1bb7c3-1f62-244c-0f3a-761cc1688042",
 "metadata": {},
 "name": "liver slide",
 "objective_key": "",
 "opener": {
  "checksum": "da1bb7c31f62244c0f3a761cc168804227115793d01c270021fe3f7935482dcc",
  "storage_address": "https://toto/dataManager/42234/opener"
 },
 "owner": "SampleOrg",
 "permissions": {
  "process": {
   "authorized_ids": [],
   "public": true
  }
 },
 "type": "images"
}
```
#### ------------ Add test DataSample ------------
Smart contract: `registerDataSample`

##### JSON Inputs:
```go
{
 "keys": [string] (required,dive,len=36),
 "data_manager_keys": [string] (omitempty,dive,len=36),
 "testOnly": string (required,oneof=true false),
}
```
##### Command peer example:
```bash
peer chaincode invoke -n mycc -c '{"Args":["registerDataSample","{\"keys\":[\"bb1bb7c3-1f62-244c-0f3a-761cc1688042\",\"bb2bb7c3-1f62-244c-0f3a-761cc1688042\"],\"data_manager_keys\":[\"da1bb7c3-1f62-244c-0f3a-761cc1688042\"],\"testOnly\":\"true\"}"]}' -C myc
```
##### Command output:
```json
{
 "keys": [
  "bb1bb7c3-1f62-244c-0f3a-761cc1688042",
  "bb2bb7c3-1f62-244c-0f3a-761cc1688042"
 ]
}
```
#### ------------ Add Objective ------------
Smart contract: `registerObjective`

##### JSON Inputs:
```go
{
 "key": string (required,len=36),
 "name": string (required,gte=1,lte=100),
 "description_checksum": string (required,len=64,hexadecimal),
 "description_storage_address": string (required,url),
 "metrics_name": string (required,gte=1,lte=100),
 "metrics_checksum": string (required,len=64,hexadecimal),
 "metrics_storage_address": string (required,url),
 "test_dataset": (omitempty){
   "data_manager_key": string (omitempty,len=36),
   "data_sample_keys": [string] (omitempty,dive,len=36),
 },
 "permissions": (required){
   "process": (required){
     "public": bool (required),
     "authorized_ids": [string] (required),
   },
 },
 "metadata": map (lte=100,dive,keys,lte=50,endkeys,lte=100),
}
```
##### Command peer example:
```bash
peer chaincode invoke -n mycc -c '{"Args":["registerObjective","{\"key\":\"5c1d9cd1-c2c1-082d-de09-21b56d11030c\",\"name\":\"MSI classification\",\"description_checksum\":\"5c1d9cd1c2c1082dde0921b56d11030c81f62fbb51932758b58ac2569dd0b379\",\"description_storage_address\":\"https://toto/objective/222/description\",\"metrics_name\":\"accuracy\",\"metrics_checksum\":\"4a1d9cd1c2c1082dde0921b56d11030c81f62fbb51932758b58ac2569dd0b379\",\"metrics_storage_address\":\"https://toto/objective/222/metrics\",\"test_dataset\":{\"data_manager_key\":\"da1bb7c3-1f62-244c-0f3a-761cc1688042\",\"data_sample_keys\":[\"bb1bb7c3-1f62-244c-0f3a-761cc1688042\",\"bb2bb7c3-1f62-244c-0f3a-761cc1688042\"]},\"permissions\":{\"process\":{\"public\":true,\"authorized_ids\":[]}},\"metadata\":null}"]}' -C myc
```
##### Command output:
```json
{
 "key": "5c1d9cd1-c2c1-082d-de09-21b56d11030c"
}
```
#### ------------ Add Algo ------------
Smart contract: `registerAlgo`

##### JSON Inputs:
```go
{
 "key": string (required,len=36),
 "name": string (required,gte=1,lte=100),
 "checksum": string (required,len=64,hexadecimal),
 "storage_address": string (required,url),
 "description_checksum": string (required,len=64,hexadecimal),
 "description_storage_address": string (required,url),
 "permissions": (required){
   "process": (required){
     "public": bool (required),
     "authorized_ids": [string] (required),
   },
 },
 "metadata": map (lte=100,dive,keys,lte=50,endkeys,lte=100),
}
```
##### Command peer example:
```bash
peer chaincode invoke -n mycc -c '{"Args":["registerAlgo","{\"key\":\"fd1bb7c3-1f62-244c-0f3a-761cc1688042\",\"name\":\"hog + svm\",\"checksum\":\"fd1bb7c31f62244c0f3a761cc168804227115793d01c270021fe3f7935482dcc\",\"storage_address\":\"https://toto/algo/222/algo\",\"description_checksum\":\"e2dbb7c31f62244c0f3a761cc168804227115793d01c270021fe3f7935482dca\",\"description_storage_address\":\"https://toto/algo/222/description\",\"permissions\":{\"process\":{\"public\":true,\"authorized_ids\":[]}},\"metadata\":null}"]}' -C myc
```
##### Command output:
```json
{
 "key": "fd1bb7c3-1f62-244c-0f3a-761cc1688042"
}
```
#### ------------ Add Train DataSample ------------
Smart contract: `registerDataSample`

##### JSON Inputs:
```go
{
 "keys": [string] (required,dive,len=36),
 "data_manager_keys": [string] (omitempty,dive,len=36),
 "testOnly": string (required,oneof=true false),
}
```
##### Command peer example:
```bash
peer chaincode invoke -n mycc -c '{"Args":["registerDataSample","{\"keys\":[\"aa1bb7c3-1f62-244c-0f3a-761cc1688042\",\"aa2bb7c3-1f62-244c-0f3a-761cc1688042\"],\"data_manager_keys\":[\"da1bb7c3-1f62-244c-0f3a-761cc1688042\"],\"testOnly\":\"false\"}"]}' -C myc
```
##### Command output:
```json
{
 "keys": [
  "aa1bb7c3-1f62-244c-0f3a-761cc1688042",
  "aa2bb7c3-1f62-244c-0f3a-761cc1688042"
 ]
}
```
#### ------------ Query DataManagers ------------
##### Command peer example:
```bash
peer chaincode query -n mycc -c '{"Args":["queryDataManagers"]}' -C myc
```
##### Command output:
```json
{
 "bookmark": "",
 "results": [
  {
   "description": {
    "checksum": "8d4bb7c31f62244c0f3a761cc168804227115793d01c270021fe3f7935482eee",
    "storage_address": "https://toto/dataManager/42234/description"
   },
   "key": "da1bb7c3-1f62-244c-0f3a-761cc1688042",
   "metadata": {},
   "name": "liver slide",
   "objective_key": "5c1d9cd1-c2c1-082d-de09-21b56d11030c",
   "opener": {
    "checksum": "da1bb7c31f62244c0f3a761cc168804227115793d01c270021fe3f7935482dcc",
    "storage_address": "https://toto/dataManager/42234/opener"
   },
   "owner": "SampleOrg",
   "permissions": {
    "process": {
     "authorized_ids": [],
     "public": true
    }
   },
   "type": "images"
  }
 ]
}
```
#### ------------ Query DataSamples ------------
##### Command peer example:
```bash
peer chaincode query -n mycc -c '{"Args":["queryDataSamples"]}' -C myc
```
##### Command output:
```json
{
 "bookmark": "",
 "results": [
  {
   "data_manager_keys": [
    "da1bb7c3-1f62-244c-0f3a-761cc1688042"
   ],
   "key": "aa1bb7c3-1f62-244c-0f3a-761cc1688042",
   "owner": "SampleOrg"
  },
  {
   "data_manager_keys": [
    "da1bb7c3-1f62-244c-0f3a-761cc1688042"
   ],
   "key": "aa2bb7c3-1f62-244c-0f3a-761cc1688042",
   "owner": "SampleOrg"
  },
  {
   "data_manager_keys": [
    "da1bb7c3-1f62-244c-0f3a-761cc1688042"
   ],
   "key": "bb1bb7c3-1f62-244c-0f3a-761cc1688042",
   "owner": "SampleOrg"
  },
  {
   "data_manager_keys": [
    "da1bb7c3-1f62-244c-0f3a-761cc1688042"
   ],
   "key": "bb2bb7c3-1f62-244c-0f3a-761cc1688042",
   "owner": "SampleOrg"
  }
 ]
}
```
#### ------------ Query Objectives ------------
##### Command peer example:
```bash
peer chaincode query -n mycc -c '{"Args":["queryObjectives"]}' -C myc
```
##### Command output:
```json
{
 "bookmark": "",
 "results": [
  {
   "description": {
    "checksum": "5c1d9cd1c2c1082dde0921b56d11030c81f62fbb51932758b58ac2569dd0b379",
    "storage_address": "https://toto/objective/222/description"
   },
   "key": "5c1d9cd1-c2c1-082d-de09-21b56d11030c",
   "metadata": {},
   "metrics": {
    "checksum": "4a1d9cd1c2c1082dde0921b56d11030c81f62fbb51932758b58ac2569dd0b379",
    "name": "accuracy",
    "storage_address": "https://toto/objective/222/metrics"
   },
   "name": "MSI classification",
   "owner": "SampleOrg",
   "permissions": {
    "process": {
     "authorized_ids": [],
     "public": true
    }
   },
   "test_dataset": {
    "data_manager_key": "da1bb7c3-1f62-244c-0f3a-761cc1688042",
    "data_sample_keys": [
     "bb1bb7c3-1f62-244c-0f3a-761cc1688042",
     "bb2bb7c3-1f62-244c-0f3a-761cc1688042"
    ],
    "metadata": {},
    "worker": ""
   }
  }
 ]
}
```
#### ------------ Add Traintuple ------------
Smart contract: `createTraintuple`

##### JSON Inputs:
```go
{
 "key": string (required,len=36),
 "algo_key": string (required,len=36),
 "in_models": [string] (omitempty,dive,len=36),
 "data_manager_key": string (required,len=36),
 "data_sample_keys": [string] (required,unique,gt=0,dive,len=36),
 "compute_plan_key": string (required_with=Rank),
 "rank": string (),
 "tag": string (omitempty,lte=64),
 "metadata": map (lte=100,dive,keys,lte=50,endkeys,lte=100),
}
```
##### Command peer example:
```bash
peer chaincode invoke -n mycc -c '{"Args":["createTraintuple","{\"key\":\"b0289ab8-3a71-f01e-2b72-0259a6452244\",\"algo_key\":\"fd1bb7c3-1f62-244c-0f3a-761cc1688042\",\"in_models\":[],\"data_manager_key\":\"da1bb7c3-1f62-244c-0f3a-761cc1688042\",\"data_sample_keys\":[\"aa1bb7c3-1f62-244c-0f3a-761cc1688042\",\"aa2bb7c3-1f62-244c-0f3a-761cc1688042\"],\"compute_plan_key\":\"\",\"rank\":\"\",\"tag\":\"\",\"metadata\":null}"]}' -C myc
```
##### Command output:
```json
{
 "key": "b0289ab8-3a71-f01e-2b72-0259a6452244"
}
```
#### ------------ Add Traintuple with inModel from previous traintuple ------------
Smart contract: `createTraintuple`

##### JSON Inputs:
```go
{
 "key": string (required,len=36),
 "algo_key": string (required,len=36),
 "in_models": [string] (omitempty,dive,len=36),
 "data_manager_key": string (required,len=36),
 "data_sample_keys": [string] (required,unique,gt=0,dive,len=36),
 "compute_plan_key": string (required_with=Rank),
 "rank": string (),
 "tag": string (omitempty,lte=64),
 "metadata": map (lte=100,dive,keys,lte=50,endkeys,lte=100),
}
```
##### Command peer example:
```bash
peer chaincode invoke -n mycc -c '{"Args":["createTraintuple","{\"key\":\"bbb89ab8-3a71-f01e-2b72-0259a6452244\",\"algo_key\":\"fd1bb7c3-1f62-244c-0f3a-761cc1688042\",\"in_models\":[\"b0289ab8-3a71-f01e-2b72-0259a6452244\"],\"data_manager_key\":\"da1bb7c3-1f62-244c-0f3a-761cc1688042\",\"data_sample_keys\":[\"aa1bb7c3-1f62-244c-0f3a-761cc1688042\",\"aa2bb7c3-1f62-244c-0f3a-761cc1688042\"],\"compute_plan_key\":\"\",\"rank\":\"\",\"tag\":\"\",\"metadata\":null}"]}' -C myc
```
##### Command output:
```json
{
 "key": "bbb89ab8-3a71-f01e-2b72-0259a6452244"
}
```
##### Command output:
```json
{
 "key": "bbb89ab8-3a71-f01e-2b72-0259a6452244"
}
```
#### ------------ Query Traintuples of worker with todo status ------------
Smart contract: `queryFilter`

##### JSON Inputs:
```go
{
 "indexName": string (required),
 "attributes": string (required),
}
```
##### Command peer example:
```bash
peer chaincode invoke -n mycc -c '{"Args":["queryFilter","{\"indexName\":\"traintuple~worker~status\",\"attributes\":\"SampleOrg, todo\"}"]}' -C myc
```
##### Command output:
```json
[
 {
  "algo": {
   "checksum": "fd1bb7c31f62244c0f3a761cc168804227115793d01c270021fe3f7935482dcc",
   "key": "fd1bb7c3-1f62-244c-0f3a-761cc1688042",
   "name": "hog + svm",
   "storage_address": "https://toto/algo/222/algo"
  },
  "compute_plan_key": "",
  "creator": "SampleOrg",
  "dataset": {
   "data_sample_keys": [
    "aa1bb7c3-1f62-244c-0f3a-761cc1688042",
    "aa2bb7c3-1f62-244c-0f3a-761cc1688042"
   ],
   "key": "da1bb7c3-1f62-244c-0f3a-761cc1688042",
   "metadata": {},
   "opener_checksum": "da1bb7c31f62244c0f3a761cc168804227115793d01c270021fe3f7935482dcc",
   "worker": "SampleOrg"
  },
  "in_models": null,
  "key": "b0289ab8-3a71-f01e-2b72-0259a6452244",
  "log": "",
  "metadata": {},
  "out_model": null,
  "permissions": {
   "process": {
    "authorized_ids": [],
    "public": true
   }
  },
  "rank": 0,
  "status": "todo",
  "tag": ""
 }
]
```
#### ------------ Log Start Training ------------
Smart contract: `logStartTrain`

##### JSON Inputs:
```go
{
 "key": string (required,len=36),
}
```
##### Command peer example:
```bash
peer chaincode invoke -n mycc -c '{"Args":["logStartTrain","{\"key\":\"b0289ab8-3a71-f01e-2b72-0259a6452244\"}"]}' -C myc
```
##### Command output:
```json
{
 "algo": {
  "checksum": "fd1bb7c31f62244c0f3a761cc168804227115793d01c270021fe3f7935482dcc",
  "key": "fd1bb7c3-1f62-244c-0f3a-761cc1688042",
  "name": "hog + svm",
  "storage_address": "https://toto/algo/222/algo"
 },
 "compute_plan_key": "",
 "creator": "SampleOrg",
 "dataset": {
  "data_sample_keys": [
   "aa1bb7c3-1f62-244c-0f3a-761cc1688042",
   "aa2bb7c3-1f62-244c-0f3a-761cc1688042"
  ],
  "key": "da1bb7c3-1f62-244c-0f3a-761cc1688042",
  "metadata": {},
  "opener_checksum": "da1bb7c31f62244c0f3a761cc168804227115793d01c270021fe3f7935482dcc",
  "worker": "SampleOrg"
 },
 "in_models": null,
 "key": "b0289ab8-3a71-f01e-2b72-0259a6452244",
 "log": "",
 "metadata": {},
 "out_model": null,
 "permissions": {
  "process": {
   "authorized_ids": [],
   "public": true
  }
 },
 "rank": 0,
 "status": "doing",
 "tag": ""
}
```
#### ------------ Log Success Training ------------
Smart contract: `logSuccessTrain`

##### JSON Inputs:
```go
{
 "key": string (required,len=36),
 "log": string (lte=200),
 "out_model": (required){
   "key": string (required,len=36),
   "checksum": string (required,len=64,hexadecimal),
   "storage_address": string (required),
 },
}
```
##### Command peer example:
```bash
peer chaincode invoke -n mycc -c '{"Args":["logSuccessTrain","{\"key\":\"b0289ab8-3a71-f01e-2b72-0259a6452244\",\"log\":\"no error, ah ah ah\",\"out_model\":{\"key\":\"eedbb7c3-1f62-244c-0f3a-761cc1688042\",\"checksum\":\"eedbb7c31f62244c0f3a761cc168804227115793d01c270021fe3f7935482eed\",\"storage_address\":\"https://substrabac/model/toto\"}}"]}' -C myc
```
##### Command output:
```json
{
 "algo": {
  "checksum": "fd1bb7c31f62244c0f3a761cc168804227115793d01c270021fe3f7935482dcc",
  "key": "fd1bb7c3-1f62-244c-0f3a-761cc1688042",
  "name": "hog + svm",
  "storage_address": "https://toto/algo/222/algo"
 },
 "compute_plan_key": "",
 "creator": "SampleOrg",
 "dataset": {
  "data_sample_keys": [
   "aa1bb7c3-1f62-244c-0f3a-761cc1688042",
   "aa2bb7c3-1f62-244c-0f3a-761cc1688042"
  ],
  "key": "da1bb7c3-1f62-244c-0f3a-761cc1688042",
  "metadata": {},
  "opener_checksum": "da1bb7c31f62244c0f3a761cc168804227115793d01c270021fe3f7935482dcc",
  "worker": "SampleOrg"
 },
 "in_models": null,
 "key": "b0289ab8-3a71-f01e-2b72-0259a6452244",
 "log": "no error, ah ah ah",
 "metadata": {},
 "out_model": {
  "checksum": "eedbb7c31f62244c0f3a761cc168804227115793d01c270021fe3f7935482eed",
  "key": "eedbb7c3-1f62-244c-0f3a-761cc1688042",
  "storage_address": "https://substrabac/model/toto"
 },
 "permissions": {
  "process": {
   "authorized_ids": [],
   "public": true
  }
 },
 "rank": 0,
 "status": "done",
 "tag": ""
}
```
#### ------------ Query Traintuple From key ------------
Smart contract: `queryTraintuple`

##### JSON Inputs:
```go
{
 "key": string (required,len=36),
}
```
##### Command peer example:
```bash
peer chaincode invoke -n mycc -c '{"Args":["queryTraintuple","{\"key\":\"b0289ab8-3a71-f01e-2b72-0259a6452244\"}"]}' -C myc
```
##### Command output:
```json
{
 "algo": {
  "checksum": "fd1bb7c31f62244c0f3a761cc168804227115793d01c270021fe3f7935482dcc",
  "key": "fd1bb7c3-1f62-244c-0f3a-761cc1688042",
  "name": "hog + svm",
  "storage_address": "https://toto/algo/222/algo"
 },
 "compute_plan_key": "",
 "creator": "SampleOrg",
 "dataset": {
  "data_sample_keys": [
   "aa1bb7c3-1f62-244c-0f3a-761cc1688042",
   "aa2bb7c3-1f62-244c-0f3a-761cc1688042"
  ],
  "key": "da1bb7c3-1f62-244c-0f3a-761cc1688042",
  "metadata": {},
  "opener_checksum": "da1bb7c31f62244c0f3a761cc168804227115793d01c270021fe3f7935482dcc",
  "worker": "SampleOrg"
 },
 "in_models": null,
 "key": "b0289ab8-3a71-f01e-2b72-0259a6452244",
 "log": "no error, ah ah ah",
 "metadata": {},
 "out_model": {
  "checksum": "eedbb7c31f62244c0f3a761cc168804227115793d01c270021fe3f7935482eed",
  "key": "eedbb7c3-1f62-244c-0f3a-761cc1688042",
  "storage_address": "https://substrabac/model/toto"
 },
 "permissions": {
  "process": {
   "authorized_ids": [],
   "public": true
  }
 },
 "rank": 0,
 "status": "done",
 "tag": ""
}
```
#### ------------ Add Non-Certified Testtuple ------------
Smart contract: `createTesttuple`

##### JSON Inputs:
```go
{
 "key": string (required,len=36),
 "data_manager_key": string (omitempty,len=36),
 "data_sample_keys": [string] (omitempty,dive,len=36),
 "objective_key": string (required,len=36),
 "tag": string (omitempty,lte=64),
 "metadata": map (omitempty,lte=100,dive,keys,lte=50,endkeys,lte=100),
 "traintuple_key": string (required,len=36),
}
```
##### Command peer example:
```bash
peer chaincode invoke -n mycc -c '{"Args":["createTesttuple","{\"key\":\"dadada11-50f6-26d3-fa86-1bf6387e3896\",\"data_manager_key\":\"da1bb7c3-1f62-244c-0f3a-761cc1688042\",\"data_sample_keys\":[\"aa1bb7c3-1f62-244c-0f3a-761cc1688042\",\"aa2bb7c3-1f62-244c-0f3a-761cc1688042\"],\"objective_key\":\"5c1d9cd1-c2c1-082d-de09-21b56d11030c\",\"tag\":\"\",\"metadata\":null,\"traintuple_key\":\"b0289ab8-3a71-f01e-2b72-0259a6452244\"}"]}' -C myc
```
##### Command output:
```json
{
 "key": "dadada11-50f6-26d3-fa86-1bf6387e3896"
}
```
#### ------------ Add Certified Testtuple ------------
Smart contract: `createTesttuple`

##### JSON Inputs:
```go
{
 "key": string (required,len=36),
 "data_manager_key": string (omitempty,len=36),
 "data_sample_keys": [string] (omitempty,dive,len=36),
 "objective_key": string (required,len=36),
 "tag": string (omitempty,lte=64),
 "metadata": map (omitempty,lte=100,dive,keys,lte=50,endkeys,lte=100),
 "traintuple_key": string (required,len=36),
}
```
##### Command peer example:
```bash
peer chaincode invoke -n mycc -c '{"Args":["createTesttuple","{\"key\":\"bbbada11-50f6-26d3-fa86-1bf6387e3896\",\"data_manager_key\":\"\",\"data_sample_keys\":null,\"objective_key\":\"5c1d9cd1-c2c1-082d-de09-21b56d11030c\",\"tag\":\"\",\"metadata\":null,\"traintuple_key\":\"b0289ab8-3a71-f01e-2b72-0259a6452244\"}"]}' -C myc
```
##### Command output:
```json
{
 "key": "bbbada11-50f6-26d3-fa86-1bf6387e3896"
}
```
#### ------------ Add Testtuple with not done traintuple ------------
Smart contract: `createTesttuple`

##### JSON Inputs:
```go
{
 "key": string (required,len=36),
 "data_manager_key": string (omitempty,len=36),
 "data_sample_keys": [string] (omitempty,dive,len=36),
 "objective_key": string (required,len=36),
 "tag": string (omitempty,lte=64),
 "metadata": map (omitempty,lte=100,dive,keys,lte=50,endkeys,lte=100),
 "traintuple_key": string (required,len=36),
}
```
##### Command peer example:
```bash
peer chaincode invoke -n mycc -c '{"Args":["createTesttuple","{\"key\":\"cccada11-50f6-26d3-fa86-1bf6387e3896\",\"data_manager_key\":\"\",\"data_sample_keys\":null,\"objective_key\":\"5c1d9cd1-c2c1-082d-de09-21b56d11030c\",\"tag\":\"\",\"metadata\":null,\"traintuple_key\":\"bbb89ab8-3a71-f01e-2b72-0259a6452244\"}"]}' -C myc
```
##### Command output:
```json
{
 "key": "cccada11-50f6-26d3-fa86-1bf6387e3896"
}
```
#### ------------ Query Testtuples of worker with todo status ------------
Smart contract: `queryFilter`

##### JSON Inputs:
```go
{
 "indexName": string (required),
 "attributes": string (required),
}
```
##### Command peer example:
```bash
peer chaincode invoke -n mycc -c '{"Args":["queryFilter","{\"indexName\":\"testtuple~worker~status\",\"attributes\":\"SampleOrg, todo\"}"]}' -C myc
```
##### Command output:
```json
[
 {
  "algo": {
   "checksum": "fd1bb7c31f62244c0f3a761cc168804227115793d01c270021fe3f7935482dcc",
   "key": "fd1bb7c3-1f62-244c-0f3a-761cc1688042",
   "name": "hog + svm",
   "storage_address": "https://toto/algo/222/algo"
  },
  "certified": true,
  "compute_plan_key": "",
  "creator": "SampleOrg",
  "dataset": {
   "data_sample_keys": [
    "bb1bb7c3-1f62-244c-0f3a-761cc1688042",
    "bb2bb7c3-1f62-244c-0f3a-761cc1688042"
   ],
   "key": "da1bb7c3-1f62-244c-0f3a-761cc1688042",
   "opener_checksum": "da1bb7c31f62244c0f3a761cc168804227115793d01c270021fe3f7935482dcc",
   "perf": 0,
   "worker": "SampleOrg"
  },
  "key": "bbbada11-50f6-26d3-fa86-1bf6387e3896",
  "log": "",
  "metadata": {},
  "objective": {
   "key": "5c1d9cd1-c2c1-082d-de09-21b56d11030c",
   "metrics": {
    "checksum": "4a1d9cd1c2c1082dde0921b56d11030c81f62fbb51932758b58ac2569dd0b379",
    "storage_address": "https://toto/objective/222/metrics"
   }
  },
  "rank": 0,
  "status": "todo",
  "tag": "",
  "traintuple_key": "b0289ab8-3a71-f01e-2b72-0259a6452244",
  "traintuple_type": "traintuple"
 },
 {
  "algo": {
   "checksum": "fd1bb7c31f62244c0f3a761cc168804227115793d01c270021fe3f7935482dcc",
   "key": "fd1bb7c3-1f62-244c-0f3a-761cc1688042",
   "name": "hog + svm",
   "storage_address": "https://toto/algo/222/algo"
  },
  "certified": false,
  "compute_plan_key": "",
  "creator": "SampleOrg",
  "dataset": {
   "data_sample_keys": [
    "aa1bb7c3-1f62-244c-0f3a-761cc1688042",
    "aa2bb7c3-1f62-244c-0f3a-761cc1688042"
   ],
   "key": "da1bb7c3-1f62-244c-0f3a-761cc1688042",
   "opener_checksum": "da1bb7c31f62244c0f3a761cc168804227115793d01c270021fe3f7935482dcc",
   "perf": 0,
   "worker": "SampleOrg"
  },
  "key": "dadada11-50f6-26d3-fa86-1bf6387e3896",
  "log": "",
  "metadata": {},
  "objective": {
   "key": "5c1d9cd1-c2c1-082d-de09-21b56d11030c",
   "metrics": {
    "checksum": "4a1d9cd1c2c1082dde0921b56d11030c81f62fbb51932758b58ac2569dd0b379",
    "storage_address": "https://toto/objective/222/metrics"
   }
  },
  "rank": 0,
  "status": "todo",
  "tag": "",
  "traintuple_key": "b0289ab8-3a71-f01e-2b72-0259a6452244",
  "traintuple_type": "traintuple"
 }
]
```
#### ------------ Log Start Testing ------------
Smart contract: `logStartTest`

##### JSON Inputs:
```go
{
 "key": string (required,len=36),
}
```
##### Command peer example:
```bash
peer chaincode invoke -n mycc -c '{"Args":["logStartTest","{\"key\":\"bbbada11-50f6-26d3-fa86-1bf6387e3896\"}"]}' -C myc
```
##### Command output:
```json
{
 "algo": {
  "checksum": "fd1bb7c31f62244c0f3a761cc168804227115793d01c270021fe3f7935482dcc",
  "key": "fd1bb7c3-1f62-244c-0f3a-761cc1688042",
  "name": "hog + svm",
  "storage_address": "https://toto/algo/222/algo"
 },
 "certified": true,
 "compute_plan_key": "",
 "creator": "SampleOrg",
 "dataset": {
  "data_sample_keys": [
   "bb1bb7c3-1f62-244c-0f3a-761cc1688042",
   "bb2bb7c3-1f62-244c-0f3a-761cc1688042"
  ],
  "key": "da1bb7c3-1f62-244c-0f3a-761cc1688042",
  "opener_checksum": "da1bb7c31f62244c0f3a761cc168804227115793d01c270021fe3f7935482dcc",
  "perf": 0,
  "worker": "SampleOrg"
 },
 "key": "bbbada11-50f6-26d3-fa86-1bf6387e3896",
 "log": "",
 "metadata": {},
 "objective": {
  "key": "5c1d9cd1-c2c1-082d-de09-21b56d11030c",
  "metrics": {
   "checksum": "4a1d9cd1c2c1082dde0921b56d11030c81f62fbb51932758b58ac2569dd0b379",
   "storage_address": "https://toto/objective/222/metrics"
  }
 },
 "rank": 0,
 "status": "doing",
 "tag": "",
 "traintuple_key": "b0289ab8-3a71-f01e-2b72-0259a6452244",
 "traintuple_type": "traintuple"
}
```
#### ------------ Log Success Testing ------------
Smart contract: `logSuccessTest`

##### JSON Inputs:
```go
{
 "key": string (required,len=36),
 "log": string (lte=200),
 "perf": float32 (omitempty),
}
```
##### Command peer example:
```bash
peer chaincode invoke -n mycc -c '{"Args":["logSuccessTest","{\"key\":\"bbbada11-50f6-26d3-fa86-1bf6387e3896\",\"log\":\"no error, ah ah ah\",\"perf\":0.9}"]}' -C myc
```
##### Command output:
```json
{
 "algo": {
  "checksum": "fd1bb7c31f62244c0f3a761cc168804227115793d01c270021fe3f7935482dcc",
  "key": "fd1bb7c3-1f62-244c-0f3a-761cc1688042",
  "name": "hog + svm",
  "storage_address": "https://toto/algo/222/algo"
 },
 "certified": true,
 "compute_plan_key": "",
 "creator": "SampleOrg",
 "dataset": {
  "data_sample_keys": [
   "bb1bb7c3-1f62-244c-0f3a-761cc1688042",
   "bb2bb7c3-1f62-244c-0f3a-761cc1688042"
  ],
  "key": "da1bb7c3-1f62-244c-0f3a-761cc1688042",
  "opener_checksum": "da1bb7c31f62244c0f3a761cc168804227115793d01c270021fe3f7935482dcc",
  "perf": 0.9,
  "worker": "SampleOrg"
 },
 "key": "bbbada11-50f6-26d3-fa86-1bf6387e3896",
 "log": "no error, ah ah ah",
 "metadata": {},
 "objective": {
  "key": "5c1d9cd1-c2c1-082d-de09-21b56d11030c",
  "metrics": {
   "checksum": "4a1d9cd1c2c1082dde0921b56d11030c81f62fbb51932758b58ac2569dd0b379",
   "storage_address": "https://toto/objective/222/metrics"
  }
 },
 "rank": 0,
 "status": "done",
 "tag": "",
 "traintuple_key": "b0289ab8-3a71-f01e-2b72-0259a6452244",
 "traintuple_type": "traintuple"
}
```
#### ------------ Query Testtuple from its key ------------
Smart contract: `queryTesttuple`

##### JSON Inputs:
```go
{
 "key": string (required,len=36),
}
```
##### Command peer example:
```bash
peer chaincode query -n mycc -c '{"Args":["queryTesttuple","{\"key\":\"bbbada11-50f6-26d3-fa86-1bf6387e3896\"}"]}' -C myc
```
##### Command output:
```json
{
 "algo": {
  "checksum": "fd1bb7c31f62244c0f3a761cc168804227115793d01c270021fe3f7935482dcc",
  "key": "fd1bb7c3-1f62-244c-0f3a-761cc1688042",
  "name": "hog + svm",
  "storage_address": "https://toto/algo/222/algo"
 },
 "certified": true,
 "compute_plan_key": "",
 "creator": "SampleOrg",
 "dataset": {
  "data_sample_keys": [
   "bb1bb7c3-1f62-244c-0f3a-761cc1688042",
   "bb2bb7c3-1f62-244c-0f3a-761cc1688042"
  ],
  "key": "da1bb7c3-1f62-244c-0f3a-761cc1688042",
  "opener_checksum": "da1bb7c31f62244c0f3a761cc168804227115793d01c270021fe3f7935482dcc",
  "perf": 0.9,
  "worker": "SampleOrg"
 },
 "key": "bbbada11-50f6-26d3-fa86-1bf6387e3896",
 "log": "no error, ah ah ah",
 "metadata": {},
 "objective": {
  "key": "5c1d9cd1-c2c1-082d-de09-21b56d11030c",
  "metrics": {
   "checksum": "4a1d9cd1c2c1082dde0921b56d11030c81f62fbb51932758b58ac2569dd0b379",
   "storage_address": "https://toto/objective/222/metrics"
  }
 },
 "rank": 0,
 "status": "done",
 "tag": "",
 "traintuple_key": "b0289ab8-3a71-f01e-2b72-0259a6452244",
 "traintuple_type": "traintuple"
}
```
#### ------------ Query all Testtuples ------------
##### Command peer example:
```bash
peer chaincode query -n mycc -c '{"Args":["queryTesttuples"]}' -C myc
```
##### Command output:
```json
{
 "bookmark": "",
 "results": [
  {
   "algo": {
    "checksum": "fd1bb7c31f62244c0f3a761cc168804227115793d01c270021fe3f7935482dcc",
    "key": "fd1bb7c3-1f62-244c-0f3a-761cc1688042",
    "name": "hog + svm",
    "storage_address": "https://toto/algo/222/algo"
   },
   "certified": false,
   "compute_plan_key": "",
   "creator": "SampleOrg",
   "dataset": {
    "data_sample_keys": [
     "aa1bb7c3-1f62-244c-0f3a-761cc1688042",
     "aa2bb7c3-1f62-244c-0f3a-761cc1688042"
    ],
    "key": "da1bb7c3-1f62-244c-0f3a-761cc1688042",
    "opener_checksum": "da1bb7c31f62244c0f3a761cc168804227115793d01c270021fe3f7935482dcc",
    "perf": 0,
    "worker": "SampleOrg"
   },
   "key": "dadada11-50f6-26d3-fa86-1bf6387e3896",
   "log": "",
   "metadata": {},
   "objective": {
    "key": "5c1d9cd1-c2c1-082d-de09-21b56d11030c",
    "metrics": {
     "checksum": "4a1d9cd1c2c1082dde0921b56d11030c81f62fbb51932758b58ac2569dd0b379",
     "storage_address": "https://toto/objective/222/metrics"
    }
   },
   "rank": 0,
   "status": "todo",
   "tag": "",
   "traintuple_key": "b0289ab8-3a71-f01e-2b72-0259a6452244",
   "traintuple_type": "traintuple"
  },
  {
   "algo": {
    "checksum": "fd1bb7c31f62244c0f3a761cc168804227115793d01c270021fe3f7935482dcc",
    "key": "fd1bb7c3-1f62-244c-0f3a-761cc1688042",
    "name": "hog + svm",
    "storage_address": "https://toto/algo/222/algo"
   },
   "certified": true,
   "compute_plan_key": "",
   "creator": "SampleOrg",
   "dataset": {
    "data_sample_keys": [
     "bb1bb7c3-1f62-244c-0f3a-761cc1688042",
     "bb2bb7c3-1f62-244c-0f3a-761cc1688042"
    ],
    "key": "da1bb7c3-1f62-244c-0f3a-761cc1688042",
    "opener_checksum": "da1bb7c31f62244c0f3a761cc168804227115793d01c270021fe3f7935482dcc",
    "perf": 0.9,
    "worker": "SampleOrg"
   },
   "key": "bbbada11-50f6-26d3-fa86-1bf6387e3896",
   "log": "no error, ah ah ah",
   "metadata": {},
   "objective": {
    "key": "5c1d9cd1-c2c1-082d-de09-21b56d11030c",
    "metrics": {
     "checksum": "4a1d9cd1c2c1082dde0921b56d11030c81f62fbb51932758b58ac2569dd0b379",
     "storage_address": "https://toto/objective/222/metrics"
    }
   },
   "rank": 0,
   "status": "done",
   "tag": "",
   "traintuple_key": "b0289ab8-3a71-f01e-2b72-0259a6452244",
   "traintuple_type": "traintuple"
  },
  {
   "algo": {
    "checksum": "fd1bb7c31f62244c0f3a761cc168804227115793d01c270021fe3f7935482dcc",
    "key": "fd1bb7c3-1f62-244c-0f3a-761cc1688042",
    "name": "hog + svm",
    "storage_address": "https://toto/algo/222/algo"
   },
   "certified": true,
   "compute_plan_key": "",
   "creator": "SampleOrg",
   "dataset": {
    "data_sample_keys": [
     "bb1bb7c3-1f62-244c-0f3a-761cc1688042",
     "bb2bb7c3-1f62-244c-0f3a-761cc1688042"
    ],
    "key": "da1bb7c3-1f62-244c-0f3a-761cc1688042",
    "opener_checksum": "da1bb7c31f62244c0f3a761cc168804227115793d01c270021fe3f7935482dcc",
    "perf": 0,
    "worker": "SampleOrg"
   },
   "key": "cccada11-50f6-26d3-fa86-1bf6387e3896",
   "log": "",
   "metadata": {},
   "objective": {
    "key": "5c1d9cd1-c2c1-082d-de09-21b56d11030c",
    "metrics": {
     "checksum": "4a1d9cd1c2c1082dde0921b56d11030c81f62fbb51932758b58ac2569dd0b379",
     "storage_address": "https://toto/objective/222/metrics"
    }
   },
   "rank": 0,
   "status": "waiting",
   "tag": "",
   "traintuple_key": "bbb89ab8-3a71-f01e-2b72-0259a6452244",
   "traintuple_type": "traintuple"
  }
 ]
}
```
#### ------------ Query details about a model ------------
Smart contract: `queryModelDetails`

##### JSON Inputs:
```go
{
 "key": string (required,len=36),
}
```
##### Command peer example:
```bash
peer chaincode query -n mycc -c '{"Args":["queryModelDetails","{\"key\":\"b0289ab8-3a71-f01e-2b72-0259a6452244\"}"]}' -C myc
```
##### Command output:
```json
{
 "non_certified_testtuples": [
  {
   "algo": {
    "checksum": "fd1bb7c31f62244c0f3a761cc168804227115793d01c270021fe3f7935482dcc",
    "key": "fd1bb7c3-1f62-244c-0f3a-761cc1688042",
    "name": "hog + svm",
    "storage_address": "https://toto/algo/222/algo"
   },
   "certified": false,
   "compute_plan_key": "",
   "creator": "SampleOrg",
   "dataset": {
    "data_sample_keys": [
     "aa1bb7c3-1f62-244c-0f3a-761cc1688042",
     "aa2bb7c3-1f62-244c-0f3a-761cc1688042"
    ],
    "key": "da1bb7c3-1f62-244c-0f3a-761cc1688042",
    "opener_checksum": "da1bb7c31f62244c0f3a761cc168804227115793d01c270021fe3f7935482dcc",
    "perf": 0,
    "worker": "SampleOrg"
   },
   "key": "dadada11-50f6-26d3-fa86-1bf6387e3896",
   "log": "",
   "metadata": {},
   "objective": {
    "key": "5c1d9cd1-c2c1-082d-de09-21b56d11030c",
    "metrics": {
     "checksum": "4a1d9cd1c2c1082dde0921b56d11030c81f62fbb51932758b58ac2569dd0b379",
     "storage_address": "https://toto/objective/222/metrics"
    }
   },
   "rank": 0,
   "status": "todo",
   "tag": "",
   "traintuple_key": "b0289ab8-3a71-f01e-2b72-0259a6452244",
   "traintuple_type": "traintuple"
  }
 ],
 "testtuple": {
  "algo": {
   "checksum": "fd1bb7c31f62244c0f3a761cc168804227115793d01c270021fe3f7935482dcc",
   "key": "fd1bb7c3-1f62-244c-0f3a-761cc1688042",
   "name": "hog + svm",
   "storage_address": "https://toto/algo/222/algo"
  },
  "certified": true,
  "compute_plan_key": "",
  "creator": "SampleOrg",
  "dataset": {
   "data_sample_keys": [
    "bb1bb7c3-1f62-244c-0f3a-761cc1688042",
    "bb2bb7c3-1f62-244c-0f3a-761cc1688042"
   ],
   "key": "da1bb7c3-1f62-244c-0f3a-761cc1688042",
   "opener_checksum": "da1bb7c31f62244c0f3a761cc168804227115793d01c270021fe3f7935482dcc",
   "perf": 0.9,
   "worker": "SampleOrg"
  },
  "key": "bbbada11-50f6-26d3-fa86-1bf6387e3896",
  "log": "no error, ah ah ah",
  "metadata": {},
  "objective": {
   "key": "5c1d9cd1-c2c1-082d-de09-21b56d11030c",
   "metrics": {
    "checksum": "4a1d9cd1c2c1082dde0921b56d11030c81f62fbb51932758b58ac2569dd0b379",
    "storage_address": "https://toto/objective/222/metrics"
   }
  },
  "rank": 0,
  "status": "done",
  "tag": "",
  "traintuple_key": "b0289ab8-3a71-f01e-2b72-0259a6452244",
  "traintuple_type": "traintuple"
 },
 "traintuple": {
  "algo": {
   "checksum": "fd1bb7c31f62244c0f3a761cc168804227115793d01c270021fe3f7935482dcc",
   "key": "fd1bb7c3-1f62-244c-0f3a-761cc1688042",
   "name": "hog + svm",
   "storage_address": "https://toto/algo/222/algo"
  },
  "compute_plan_key": "",
  "creator": "SampleOrg",
  "dataset": {
   "data_sample_keys": [
    "aa1bb7c3-1f62-244c-0f3a-761cc1688042",
    "aa2bb7c3-1f62-244c-0f3a-761cc1688042"
   ],
   "key": "da1bb7c3-1f62-244c-0f3a-761cc1688042",
   "metadata": {},
   "opener_checksum": "da1bb7c31f62244c0f3a761cc168804227115793d01c270021fe3f7935482dcc",
   "worker": "SampleOrg"
  },
  "in_models": null,
  "key": "b0289ab8-3a71-f01e-2b72-0259a6452244",
  "log": "no error, ah ah ah",
  "metadata": {},
  "out_model": {
   "checksum": "eedbb7c31f62244c0f3a761cc168804227115793d01c270021fe3f7935482eed",
   "key": "eedbb7c3-1f62-244c-0f3a-761cc1688042",
   "storage_address": "https://substrabac/model/toto"
  },
  "permissions": {
   "process": {
    "authorized_ids": [],
    "public": true
   }
  },
  "rank": 0,
  "status": "done",
  "tag": ""
 }
}
```
#### ------------ Query all models ------------
##### Command peer example:
```bash
peer chaincode query -n mycc -c '{"Args":["queryModels"]}' -C myc
```
##### Command output:
```json
{
 "bookmark": "{\"traintuple\":\"\",\"composite_traintuple\":\"\",\"aggregatetuple\":\"\"}",
 "results": [
  {
   "traintuple": {
    "algo": {
     "checksum": "fd1bb7c31f62244c0f3a761cc168804227115793d01c270021fe3f7935482dcc",
     "key": "fd1bb7c3-1f62-244c-0f3a-761cc1688042",
     "name": "hog + svm",
     "storage_address": "https://toto/algo/222/algo"
    },
    "compute_plan_key": "",
    "creator": "SampleOrg",
    "dataset": {
     "data_sample_keys": [
      "aa1bb7c3-1f62-244c-0f3a-761cc1688042",
      "aa2bb7c3-1f62-244c-0f3a-761cc1688042"
     ],
     "key": "da1bb7c3-1f62-244c-0f3a-761cc1688042",
     "metadata": {},
     "opener_checksum": "da1bb7c31f62244c0f3a761cc168804227115793d01c270021fe3f7935482dcc",
     "worker": "SampleOrg"
    },
    "in_models": null,
    "key": "b0289ab8-3a71-f01e-2b72-0259a6452244",
    "log": "no error, ah ah ah",
    "metadata": {},
    "out_model": {
     "checksum": "eedbb7c31f62244c0f3a761cc168804227115793d01c270021fe3f7935482eed",
     "key": "eedbb7c3-1f62-244c-0f3a-761cc1688042",
     "storage_address": "https://substrabac/model/toto"
    },
    "permissions": {
     "process": {
      "authorized_ids": [],
      "public": true
     }
    },
    "rank": 0,
    "status": "done",
    "tag": ""
   }
  },
  {
   "traintuple": {
    "algo": {
     "checksum": "fd1bb7c31f62244c0f3a761cc168804227115793d01c270021fe3f7935482dcc",
     "key": "fd1bb7c3-1f62-244c-0f3a-761cc1688042",
     "name": "hog + svm",
     "storage_address": "https://toto/algo/222/algo"
    },
    "compute_plan_key": "",
    "creator": "SampleOrg",
    "dataset": {
     "data_sample_keys": [
      "aa1bb7c3-1f62-244c-0f3a-761cc1688042",
      "aa2bb7c3-1f62-244c-0f3a-761cc1688042"
     ],
     "key": "da1bb7c3-1f62-244c-0f3a-761cc1688042",
     "metadata": {},
     "opener_checksum": "da1bb7c31f62244c0f3a761cc168804227115793d01c270021fe3f7935482dcc",
     "worker": "SampleOrg"
    },
    "in_models": [
     {
      "checksum": "eedbb7c31f62244c0f3a761cc168804227115793d01c270021fe3f7935482eed",
      "key": "eedbb7c3-1f62-244c-0f3a-761cc1688042",
      "storage_address": "https://substrabac/model/toto",
      "traintuple_key": "b0289ab8-3a71-f01e-2b72-0259a6452244"
     }
    ],
    "key": "bbb89ab8-3a71-f01e-2b72-0259a6452244",
    "log": "",
    "metadata": {},
    "out_model": null,
    "permissions": {
     "process": {
      "authorized_ids": [],
      "public": true
     }
    },
    "rank": 0,
    "status": "todo",
    "tag": ""
   }
  }
 ]
}
```
#### ------------ Query model ------------
Smart contract: `queryModel`

##### JSON Inputs:
```go
{
 "key": string (required,len=36),
}
```
##### Command peer example:
```bash
peer chaincode query -n mycc -c '{"Args":["queryModel","{\"key\":\"eedbb7c3-1f62-244c-0f3a-761cc1688042\"}"]}' -C myc
```
##### Command output:
```json
{
 "key": "eedbb7c3-1f62-244c-0f3a-761cc1688042",
 "owner": "SampleOrg",
 "permissions": {
  "download": {
   "authorized_ids": [],
   "public": true
  },
  "process": {
   "authorized_ids": [],
   "public": true
  }
 },
 "storage_address": "https://substrabac/model/toto"
}
```
#### ------------ Query Dataset ------------
Smart contract: `queryDataset`

##### JSON Inputs:
```go
{
 "key": string (required,len=36),
}
```
##### Command peer example:
```bash
peer chaincode query -n mycc -c '{"Args":["queryDataset","{\"key\":\"da1bb7c3-1f62-244c-0f3a-761cc1688042\"}"]}' -C myc
```
##### Command output:
```json
{
 "description": {
  "checksum": "8d4bb7c31f62244c0f3a761cc168804227115793d01c270021fe3f7935482eee",
  "storage_address": "https://toto/dataManager/42234/description"
 },
 "key": "da1bb7c3-1f62-244c-0f3a-761cc1688042",
 "metadata": {},
 "name": "liver slide",
 "objective_key": "5c1d9cd1-c2c1-082d-de09-21b56d11030c",
 "opener": {
  "checksum": "da1bb7c31f62244c0f3a761cc168804227115793d01c270021fe3f7935482dcc",
  "storage_address": "https://toto/dataManager/42234/opener"
 },
 "owner": "SampleOrg",
 "permissions": {
  "process": {
   "authorized_ids": [],
   "public": true
  }
 },
 "test_data_sample_keys": [
  "bb1bb7c3-1f62-244c-0f3a-761cc1688042",
  "bb2bb7c3-1f62-244c-0f3a-761cc1688042"
 ],
 "train_data_sample_keys": [
  "aa1bb7c3-1f62-244c-0f3a-761cc1688042",
  "aa2bb7c3-1f62-244c-0f3a-761cc1688042"
 ],
 "type": "images"
}
```
#### ------------ Query nodes ------------
##### Command peer example:
```bash
peer chaincode query -n mycc -c '{"Args":["queryNodes"]}' -C myc
```
##### Command output:
```json
[
 {
  "id": "SampleOrg"
 }
]
```
#### ------------ Update Data Sample with new data manager ------------
Smart contract: `updateDataSample`

##### JSON Inputs:
```go
{
 "keys": [string] (required,dive,len=36),
 "data_manager_keys": [string] (required,dive,len=36),
}
```
##### Command peer example:
```bash
peer chaincode invoke -n mycc -c '{"Args":["updateDataSample","{\"keys\":[\"aa1bb7c3-1f62-244c-0f3a-761cc1688042\"],\"data_manager_keys\":[\"38a320b2-a67c-8003-cc74-8d6666534f2b\"]}"]}' -C myc
```
##### Command output:
```json
{
 "key": "{\"keys\": [\"aa1bb7c3-1f62-244c-0f3a-761cc1688042\"]}"
}
```
#### ------------ Query the new Dataset ------------
Smart contract: `queryDataset`

##### JSON Inputs:
```go
{
 "key": string (required,len=36),
}
```
##### Command peer example:
```bash
peer chaincode query -n mycc -c '{"Args":["queryDataset","{\"key\":\"38a320b2-a67c-8003-cc74-8d6666534f2b\"}"]}' -C myc
```
##### Command output:
```json
{
 "description": {
  "checksum": "8d4bb7c31f62244c0f3a761cc168804227115793d01c270021fe3f7935482eee",
  "storage_address": "https://toto/dataManager/42234/description"
 },
 "key": "38a320b2-a67c-8003-cc74-8d6666534f2b",
 "metadata": {},
 "name": "liver slide",
 "objective_key": "",
 "opener": {
  "checksum": "da1bb7c31f62244c0f3a761cc168804227115793d01c270021fe3f7935482dcc",
  "storage_address": "https://toto/dataManager/42234/opener"
 },
 "owner": "SampleOrg",
 "permissions": {
  "process": {
   "authorized_ids": [],
   "public": true
  }
 },
 "test_data_sample_keys": [],
 "train_data_sample_keys": [
  "aa1bb7c3-1f62-244c-0f3a-761cc1688042"
 ],
 "type": "images"
}
```
#### ------------ Create a ComputePlan ------------
Smart contract: `createComputePlan`

##### JSON Inputs:
```go
{
 "tag": string (omitempty,lte=64),
 "metadata": map (lte=100,dive,keys,lte=50,endkeys,lte=100),
 "key": string (required,len=36),
 "traintuples": (omitempty) [{
   "key": string (required,len=36),
   "data_manager_key": string (required,len=36),
   "data_sample_keys": [string] (required,dive,len=36),
   "algo_key": string (required,len=36),
   "id": string (required,lte=64),
   "in_models_ids": [string] (omitempty,dive,lte=64),
   "tag": string (omitempty,lte=64),
   "metadata": map (omitempty,lte=100,dive,keys,lte=50,endkeys,lte=100),
 }],
 "aggregatetuples": (omitempty) [{
   "key": string (required,len=36),
   "algo_key": string (required,len=36),
   "id": string (required,lte=64),
   "in_models_ids": [string] (omitempty,dive,lte=64),
   "tag": string (omitempty,lte=64),
   "metadata": map (omitempty,lte=100,dive,keys,lte=50,endkeys,lte=100),
   "worker": string (required),
 }],
 "composite_traintuples": (omitempty) [{
   "key": string (required,len=36),
   "data_manager_key": string (required,len=36),
   "data_sample_keys": [string] (required,dive,len=36),
   "algo_key": string (required,len=36),
   "id": string (required,lte=64),
   "in_head_model_id": string (required_with=InTrunkModelID,omitempty,len=64,hexadecimal),
   "in_trunk_model_id": string (required_with=InHeadModelID,omitempty,len=64,hexadecimal),
   "out_trunk_model_permissions": (required){
     "process": (required){
       "public": bool (required),
       "authorized_ids": [string] (required),
     },
   },
   "tag": string (omitempty,lte=64),
   "metadata": map (omitempty,lte=100,dive,keys,lte=50,endkeys,lte=100),
 }],
 "testtuples": (omitempty) [{
   "key": string (required,len=36),
   "data_manager_key": string (omitempty,len=36),
   "data_sample_keys": [string] (omitempty,dive,len=36),
   "objective_key": string (required,len=36),
   "tag": string (omitempty,lte=64),
   "metadata": map (omitempty,lte=100,dive,keys,lte=50,endkeys,lte=100),
   "traintuple_id": string (required,lte=64),
 }],
}
```
##### Command peer example:
```bash
peer chaincode invoke -n mycc -c '{"Args":["createComputePlan","{\"clean_models\":false,\"tag\":\"a tag is simply a string\",\"metadata\":null,\"key\":\"00000000-50f6-26d3-fa86-1bf6387e3896\",\"traintuples\":[{\"key\":\"11000000-50f6-26d3-fa86-1bf6387e3896\",\"data_manager_key\":\"da1bb7c3-1f62-244c-0f3a-761cc1688042\",\"data_sample_keys\":[\"aa1bb7c3-1f62-244c-0f3a-761cc1688042\"],\"algo_key\":\"fd1bb7c3-1f62-244c-0f3a-761cc1688042\",\"id\":\"firstTraintupleID\",\"in_models_ids\":null,\"tag\":\"\",\"metadata\":null},{\"key\":\"22000000-50f6-26d3-fa86-1bf6387e3896\",\"data_manager_key\":\"da1bb7c3-1f62-244c-0f3a-761cc1688042\",\"data_sample_keys\":[\"aa2bb7c3-1f62-244c-0f3a-761cc1688042\"],\"algo_key\":\"fd1bb7c3-1f62-244c-0f3a-761cc1688042\",\"id\":\"secondTraintupleID\",\"in_models_ids\":[\"firstTraintupleID\"],\"tag\":\"\",\"metadata\":null}],\"aggregatetuples\":null,\"composite_traintuples\":null,\"testtuples\":[{\"key\":\"11000033-50f6-26d3-fa86-1bf6387e3896\",\"data_manager_key\":\"da1bb7c3-1f62-244c-0f3a-761cc1688042\",\"data_sample_keys\":[\"bb1bb7c3-1f62-244c-0f3a-761cc1688042\",\"bb2bb7c3-1f62-244c-0f3a-761cc1688042\"],\"objective_key\":\"5c1d9cd1-c2c1-082d-de09-21b56d11030c\",\"tag\":\"\",\"metadata\":null,\"traintuple_id\":\"secondTraintupleID\"}]}"]}' -C myc
```
##### Command output:
```json
{
 "aggregatetuple_keys": null,
 "clean_models": false,
 "composite_traintuple_keys": null,
 "done_count": 0,
 "id_to_key": {
  "firstTraintupleID": "11000000-50f6-26d3-fa86-1bf6387e3896",
  "secondTraintupleID": "22000000-50f6-26d3-fa86-1bf6387e3896"
 },
 "key": "00000000-50f6-26d3-fa86-1bf6387e3896",
 "metadata": {},
 "status": "todo",
 "tag": "a tag is simply a string",
 "testtuple_keys": [
  "11000033-50f6-26d3-fa86-1bf6387e3896"
 ],
 "traintuple_keys": [
  "11000000-50f6-26d3-fa86-1bf6387e3896",
  "22000000-50f6-26d3-fa86-1bf6387e3896"
 ],
 "tuple_count": 3
}
```
#### ------------ Update a ComputePlan ------------
Smart contract: `updateComputePlan`

##### JSON Inputs:
```go
{
 "key": string (required,len=36),
 "traintuples": (omitempty) [{
   "key": string (required,len=36),
   "data_manager_key": string (required,len=36),
   "data_sample_keys": [string] (required,dive,len=36),
   "algo_key": string (required,len=36),
   "id": string (required,lte=64),
   "in_models_ids": [string] (omitempty,dive,lte=64),
   "tag": string (omitempty,lte=64),
   "metadata": map (omitempty,lte=100,dive,keys,lte=50,endkeys,lte=100),
 }],
 "aggregatetuples": (omitempty) [{
   "key": string (required,len=36),
   "algo_key": string (required,len=36),
   "id": string (required,lte=64),
   "in_models_ids": [string] (omitempty,dive,lte=64),
   "tag": string (omitempty,lte=64),
   "metadata": map (omitempty,lte=100,dive,keys,lte=50,endkeys,lte=100),
   "worker": string (required),
 }],
 "composite_traintuples": (omitempty) [{
   "key": string (required,len=36),
   "data_manager_key": string (required,len=36),
   "data_sample_keys": [string] (required,dive,len=36),
   "algo_key": string (required,len=36),
   "id": string (required,lte=64),
   "in_head_model_id": string (required_with=InTrunkModelID,omitempty,len=64,hexadecimal),
   "in_trunk_model_id": string (required_with=InHeadModelID,omitempty,len=64,hexadecimal),
   "out_trunk_model_permissions": (required){
     "process": (required){
       "public": bool (required),
       "authorized_ids": [string] (required),
     },
   },
   "tag": string (omitempty,lte=64),
   "metadata": map (omitempty,lte=100,dive,keys,lte=50,endkeys,lte=100),
 }],
 "testtuples": (omitempty) [{
   "key": string (required,len=36),
   "data_manager_key": string (omitempty,len=36),
   "data_sample_keys": [string] (omitempty,dive,len=36),
   "objective_key": string (required,len=36),
   "tag": string (omitempty,lte=64),
   "metadata": map (omitempty,lte=100,dive,keys,lte=50,endkeys,lte=100),
   "traintuple_id": string (required,lte=64),
 }],
}
```
##### Command peer example:
```bash
peer chaincode invoke -n mycc -c '{"Args":["updateComputePlan","{\"key\":\"00000000-50f6-26d3-fa86-1bf6387e3896\",\"traintuples\":[{\"key\":\"33000000-50f6-26d3-fa86-1bf6387e3896\",\"data_manager_key\":\"da1bb7c3-1f62-244c-0f3a-761cc1688042\",\"data_sample_keys\":[\"aa1bb7c3-1f62-244c-0f3a-761cc1688042\"],\"algo_key\":\"fd1bb7c3-1f62-244c-0f3a-761cc1688042\",\"id\":\"thirdTraintupleID\",\"in_models_ids\":[\"firstTraintupleID\",\"secondTraintupleID\"],\"tag\":\"\",\"metadata\":null}],\"aggregatetuples\":null,\"composite_traintuples\":null,\"testtuples\":[{\"key\":\"22000033-50f6-26d3-fa86-1bf6387e3896\",\"data_manager_key\":\"da1bb7c3-1f62-244c-0f3a-761cc1688042\",\"data_sample_keys\":[\"bb1bb7c3-1f62-244c-0f3a-761cc1688042\",\"bb2bb7c3-1f62-244c-0f3a-761cc1688042\"],\"objective_key\":\"5c1d9cd1-c2c1-082d-de09-21b56d11030c\",\"tag\":\"\",\"metadata\":null,\"traintuple_id\":\"thirdTraintupleID\"}]}"]}' -C myc
```
##### Command output:
```json
{
 "aggregatetuple_keys": null,
 "clean_models": false,
 "composite_traintuple_keys": null,
 "done_count": 0,
 "id_to_key": {
  "thirdTraintupleID": "33000000-50f6-26d3-fa86-1bf6387e3896"
 },
 "key": "00000000-50f6-26d3-fa86-1bf6387e3896",
 "metadata": {},
 "status": "todo",
 "tag": "a tag is simply a string",
 "testtuple_keys": [
  "11000033-50f6-26d3-fa86-1bf6387e3896",
  "22000033-50f6-26d3-fa86-1bf6387e3896"
 ],
 "traintuple_keys": [
  "11000000-50f6-26d3-fa86-1bf6387e3896",
  "22000000-50f6-26d3-fa86-1bf6387e3896",
  "33000000-50f6-26d3-fa86-1bf6387e3896"
 ],
 "tuple_count": 5
}
```
#### ------------ Query an ObjectiveLeaderboard ------------
Smart contract: `queryObjectiveLeaderboard`

##### JSON Inputs:
```go
{
 "objective_key": string (omitempty,len=36),
 "ascendingOrder": bool (required),
}
```
##### Command peer example:
```bash
peer chaincode invoke -n mycc -c '{"Args":["queryObjectiveLeaderboard","{\"objective_key\":\"5c1d9cd1-c2c1-082d-de09-21b56d11030c\",\"ascendingOrder\":true}"]}' -C myc
```
##### Command output:
```json
{
 "objective": {
  "description": {
   "checksum": "5c1d9cd1c2c1082dde0921b56d11030c81f62fbb51932758b58ac2569dd0b379",
   "storage_address": "https://toto/objective/222/description"
  },
  "key": "5c1d9cd1-c2c1-082d-de09-21b56d11030c",
  "metadata": {},
  "metrics": {
   "checksum": "4a1d9cd1c2c1082dde0921b56d11030c81f62fbb51932758b58ac2569dd0b379",
   "name": "accuracy",
   "storage_address": "https://toto/objective/222/metrics"
  },
  "name": "MSI classification",
  "owner": "SampleOrg",
  "permissions": {
   "process": {
    "authorized_ids": [],
    "public": true
   }
  },
  "test_dataset": {
   "data_manager_key": "da1bb7c3-1f62-244c-0f3a-761cc1688042",
   "data_sample_keys": [
    "bb1bb7c3-1f62-244c-0f3a-761cc1688042",
    "bb2bb7c3-1f62-244c-0f3a-761cc1688042"
   ],
   "metadata": {},
   "worker": ""
  }
 },
 "testtuples": [
  {
   "algo": {
    "checksum": "fd1bb7c31f62244c0f3a761cc168804227115793d01c270021fe3f7935482dcc",
    "key": "fd1bb7c3-1f62-244c-0f3a-761cc1688042",
    "name": "hog + svm",
    "storage_address": "https://toto/algo/222/algo"
   },
   "creator": "SampleOrg",
   "key": "bbbada11-50f6-26d3-fa86-1bf6387e3896",
   "perf": 0.9,
   "tag": "",
   "traintuple_key": "b0289ab8-3a71-f01e-2b72-0259a6452244"
  }
 ]
}
```
#### ------------ Query Compute Plan(s) ------------
Smart contract: `queryComputePlan`

##### JSON Inputs:
```go
{
 "key": string (required,len=36),
}
```
##### Command peer example:
```bash
peer chaincode invoke -n mycc -c '{"Args":["queryComputePlan","{\"key\":\"00000000-50f6-26d3-fa86-1bf6387e3896\"}"]}' -C myc
```
##### Command output:
```json
{
 "aggregatetuple_keys": null,
 "clean_models": false,
 "composite_traintuple_keys": null,
 "done_count": 0,
 "id_to_key": {},
 "key": "00000000-50f6-26d3-fa86-1bf6387e3896",
 "metadata": {},
 "status": "todo",
 "tag": "a tag is simply a string",
 "testtuple_keys": [
  "11000033-50f6-26d3-fa86-1bf6387e3896",
  "22000033-50f6-26d3-fa86-1bf6387e3896"
 ],
 "traintuple_keys": [
  "11000000-50f6-26d3-fa86-1bf6387e3896",
  "22000000-50f6-26d3-fa86-1bf6387e3896",
  "33000000-50f6-26d3-fa86-1bf6387e3896"
 ],
 "tuple_count": 5
}
```
##### Command peer example:
```bash
peer chaincode invoke -n mycc -c '{"Args":["queryComputePlans"]}' -C myc
```
##### Command output:
```json
{
 "bookmark": "",
 "results": [
  {
   "aggregatetuple_keys": null,
   "clean_models": false,
   "composite_traintuple_keys": null,
   "done_count": 0,
   "id_to_key": {},
   "key": "00000000-50f6-26d3-fa86-1bf6387e3896",
   "metadata": {},
   "status": "todo",
   "tag": "a tag is simply a string",
   "testtuple_keys": [
    "11000033-50f6-26d3-fa86-1bf6387e3896",
    "22000033-50f6-26d3-fa86-1bf6387e3896"
   ],
   "traintuple_keys": [
    "11000000-50f6-26d3-fa86-1bf6387e3896",
    "22000000-50f6-26d3-fa86-1bf6387e3896",
    "33000000-50f6-26d3-fa86-1bf6387e3896"
   ],
   "tuple_count": 5
  }
 ]
}
```
#### ------------ Cancel a ComputePlan ------------
Smart contract: `cancelComputePlan`

##### JSON Inputs:
```go
{
 "key": string (required,len=36),
}
```
##### Command peer example:
```bash
peer chaincode invoke -n mycc -c '{"Args":["cancelComputePlan","{\"key\":\"00000000-50f6-26d3-fa86-1bf6387e3896\"}"]}' -C myc
```
##### Command output:
```json
{
 "aggregatetuple_keys": null,
 "clean_models": false,
 "composite_traintuple_keys": null,
 "done_count": 0,
 "id_to_key": {},
 "key": "00000000-50f6-26d3-fa86-1bf6387e3896",
 "metadata": {},
 "status": "canceled",
 "tag": "a tag is simply a string",
 "testtuple_keys": [
  "11000033-50f6-26d3-fa86-1bf6387e3896",
  "22000033-50f6-26d3-fa86-1bf6387e3896"
 ],
 "traintuple_keys": [
  "11000000-50f6-26d3-fa86-1bf6387e3896",
  "22000000-50f6-26d3-fa86-1bf6387e3896",
  "33000000-50f6-26d3-fa86-1bf6387e3896"
 ],
 "tuple_count": 5
}
```
