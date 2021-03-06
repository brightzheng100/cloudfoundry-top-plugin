// Copyright (c) 2016 ECS Team, Inc. - All Rights Reserved
// https://github.com/ECSTeam/cloudfoundry-top-plugin
//
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
//
// http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License.

package eventrouting

import (
	"sync/atomic"
	"time"

	"github.com/cloudfoundry/sonde-go/events"
	"github.com/ecsteam/cloudfoundry-top-plugin/eventdata"
)

type EventRouter struct {
	eventCount uint64
	startTime  time.Time
	processor  *eventdata.EventProcessor
}

func NewEventRouter(processor *eventdata.EventProcessor) *EventRouter {
	return &EventRouter{
		processor: processor,
		startTime: time.Now(),
	}
}

func (er *EventRouter) GetProcessor() *eventdata.EventProcessor {
	return er.processor
}

func (er *EventRouter) GetEventCount() uint64 {
	return atomic.LoadUint64(&er.eventCount)
}

func (er *EventRouter) GetStartTime() time.Time {
	return er.startTime
}

func (er *EventRouter) Clear() {
	atomic.StoreUint64(&er.eventCount, 0)
	er.startTime = time.Now()
	er.processor.ClearStats()
}

func (er *EventRouter) Route(instanceId int, msg *events.Envelope) {
	atomic.AddUint64(&er.eventCount, 1)
	er.processor.Process(instanceId, msg)
}
