/*******************************************************************************
 * Copyright 2019 Samsung Electronics All Rights Reserved.
 *
 * Licensed under the Apache License, Version 2.0 (the "License");
 * you may not use this file except in compliance with the License.
 * You may obtain a copy of the License at
 *
 * http://www.apache.org/licenses/LICENSE-2.0
 *
 * Unless required by applicable law or agreed to in writing, software
 * distributed under the License is distributed on an "AS IS" BASIS,
 * WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
 * See the License for the specific language governing permissions and
 * limitations under the License.
 *
 *******************************************************************************/

//Package discoverymgr provides functions to register local device to network and find other orchestration devices
package discoverymgr

import (
	wrapper "controller/discoverymgr/wrapper"
	"sync"
)

const logPrefix = "[discoverymgr]"

// OrchestrationInformation is the struct to handle orchestration
type OrchestrationInformation struct {
	Platform      string `json:"Platform"`
	ExecutionType string `json:"ExecutionType"`

	//interface-ip 형태의 구조체 리스트로.
	IPv4 []string `json:"IPv4"`
	// IPv6     []string   `json:"IPv6"`
	ServiceList []string `json:"ServiceList"`
}

// OrchestrationMap has full Orchestration Information
type OrchestrationMap map[string]*OrchestrationInformation

// ExportDeviceMap gives device info map for discoverymgr user
type ExportDeviceMap map[string]OrchestrationInformation

type discoverymgrInformation struct {
	deviceID         string
	platform         string
	executionType    string
	orchestrationMap OrchestrationMap
	mapMTX           sync.Mutex
	wrapperIns       wrapper.ZeroconfInterface
	shutdownChan     chan struct{}
}
