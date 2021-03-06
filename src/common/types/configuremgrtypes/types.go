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

// Package configuremgrtypes defines types for configuremgr
package configuremgrtypes

import (
	"common/resourceutil"
)

// ServiceInfo for configuremgr
type ServiceInfo struct {
	ServiceName string
	//	ScoringFunc    func(cpu float32, mem float32, net float32, score *float32) int
	//ScoringFunc func(func(string) float64) float64
	ScoringFunc    resourceutil.Command
	IntervalTimeMs int
}
