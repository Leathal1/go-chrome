package chrome

import (
	"app/chrome/protocol"
	"encoding/json"

	log "github.com/Sirupsen/logrus"
)

/*
HeapProfiler - https://chromedevtools.github.io/devtools-protocol/tot/HeapProfiler/
EXPERIMENTAL
*/
type HeapProfiler struct{}

/*
AddInspectedHeapObject enables console to refer to the node with given id via $x (see Command Line
API for more details $x functions).
*/
func (HeapProfiler) AddInspectedHeapObject(
	socket *Socket,
	params *heap_profiler.AddInspectedHeapObjectParams,
) (nil, error) {
	command := &protocol.Command{
		Method: "HeapProfiler.addInspectedHeapObject",
		Params: params,
	}
	socket.SendCommand(command)
	return nil, command.Err
}

/*
CollectGarbage EXPERIMENTAL
*/
func (HeapProfiler) CollectGarbage(
	socket *Socket,
) (nil, error) {
	command := &protocol.Command{
		Method: "HeapProfiler.collectGarbage",
	}
	socket.SendCommand(command)
	return nil, command.Err
}

/*
Disable disables the HeapProfiler.
*/
func (HeapProfiler) Disable(
	socket *Socket,
) (nil, error) {
	command := &protocol.Command{
		Method: "HeapProfiler.disable",
	}
	socket.SendCommand(command)
	return nil, command.Err
}

/*
Enable enables the HeapProfiler.
*/
func (HeapProfiler) Enable(
	socket *Socket,
) (nil, error) {
	command := &protocol.Command{
		Method: "HeapProfiler.enable",
	}
	socket.SendCommand(command)
	return nil, command.Err
}

/*
GetHeapObjectID EXPERIMENTAL
*/
func (HeapProfiler) GetHeapObjectID(
	socket *Socket,
	params *heap_profiler.GetHeapObjectIDParams,
) (heap_profiler.GetHeapObjectIDResult, error) {
	command := &protocol.Command{
		Method: "HeapProfiler.getHeapObjectID",
		Params: params,
	}
	socket.SendCommand(command)
	return command.Result.(heap_profiler.GetHeapObjectIDResult), command.Err
}

/*
GetObjectByHeapObjectID EXPERIMENTAL
*/
func (HeapProfiler) GetObjectByHeapObjectID(
	socket *Socket,
	params *heap_profiler.GetObjectByHeapObjectIDParams,
) (heap_profiler.GetObjectByHeapObjectIDResult, error) {
	command := &protocol.Command{
		Method: "HeapProfiler.getObjectByHeapObjectId",
		Params: params,
	}
	socket.SendCommand(command)
	return command.Result.(heap_profiler.GetObjectByHeapObjectIDResult), command.Err
}

/*
GetSamplingProfile EXPERIMENTAL
*/
func (HeapProfiler) GetSamplingProfile(
	socket *Socket,
	params *heap_profiler.GetSamplingProfileParams,
) (nil, error) {
	command := &protocol.Command{
		Method: "HeapProfiler.getSamplingProfile",
		Params: params,
	}
	socket.SendCommand(command)
	return nil, command.Err
}

/*
StartSampling EXPERIMENTAL
*/
func (HeapProfiler) StartSampling(
	socket *Socket,
	params *heap_profiler.StartSamplingParams,
) (nil, error) {
	command := &protocol.Command{
		Method: "HeapProfiler.startSampling",
		Params: params,
	}
	socket.SendCommand(command)
	return nil, command.Err
}

/*
StartTrackingHeapObjects EXPERIMENTAL
*/
func (HeapProfiler) StartTrackingHeapObjects(
	socket *Socket,
	params *heap_profiler.StartTrackingHeapObjectsParams,
) (nil, error) {
	command := &protocol.Command{
		Method: "HeapProfiler.startTrackingHeapObjects",
		Params: params,
	}
	socket.SendCommand(command)
	return nil, command.Err
}

/*
StopSampling EXPERIMENTAL
*/
func (HeapProfiler) StopSampling(
	socket *Socket,
	params *heap_profiler.StopSamplingParams,
) (nil, error) {
	command := &protocol.Command{
		Method: "HeapProfiler.stopSampling",
		Params: params,
	}
	socket.SendCommand(command)
	return nil, command.Err
}

/*
StopTrackingHeapObjects EXPERIMENTAL
*/
func (HeapProfiler) StopTrackingHeapObjects(
	socket *Socket,
	params *heap_profiler.StopTrackingHeapObjectsParams,
) (nil, error) {
	command := &protocol.Command{
		Method: "HeapProfiler.stopTrackingHeapObjects",
		Params: params,
	}
	socket.SendCommand(command)
	return nil, command.Err
}

/*
TakeHeapSnapshot EXPERIMENTAL
*/
func (HeapProfiler) TakeHeapSnapshot(
	socket *Socket,
	params *heap_profiler.TakeHeapSnapshotParams,
) (nil, error) {
	command := &protocol.Command{
		Method: "HeapProfiler.takeHeapSnapshot",
		Params: params,
	}
	socket.SendCommand(command)
	return nil, command.Err
}

/*
OnAddHeapSnapshotChunk adds a handler to the HeapProfiler.AddHeapSnapshotChunk event.
EXPERIMENTAL
*/
func (HeapProfiler) OnAddHeapSnapshotChunk(
	socket *Socket,
	callback func(event *heap_profiler.AddHeapSnapshotChunkEvent),
) {
	handler := protocol.NewEventHandler(
		"HeapProfiler.addHeapSnapshotChunk",
		func(name string, params []byte) {
			event := &heap_profiler.AddHeapSnapshotChunkEvent{}
			if err := json.Unmarshal(params, event); err != nil {
				log.Error(err)
			} else {
				callback(event)
			}
		},
	)
	socket.AddEventHandler(handler)
}

/*
OnHeapStatsUpdate adds a handler to the DOM.heapStatsUpdate event. DOM.heapStatsUpdate fires if heap
objects tracking has been started then backend may send update for one or more fragments.
*/
func (HeapProfiler) OnHeapStatsUpdate(
	socket *Socket,
	callback func(event *heap_profiler.HeapStatsUpdateEvent),
) {
	handler := protocol.NewEventHandler(
		"HeapProfiler.heapStatsUpdate",
		func(name string, params []byte) {
			event := &heap_profiler.HeapStatsUpdateEvent{}
			if err := json.Unmarshal(params, event); err != nil {
				log.Error(err)
			} else {
				callback(event)
			}
		},
	)
	socket.AddEventHandler(handler)
}

/*
OnLastSeenObjectID adds a handler to the DOM.LastSeenObjectID event. DOM.LastSeenObjectID fires if
heap objects tracking has been started then backend regularly sends a current value for last seen
object id and corresponding timestamp. If the were changes in the heap since last event then one or
more heapStatsUpdate events will be sent before a new lastSeenObjectId event.
*/
func (HeapProfiler) OnLastSeenObjectID(
	socket *Socket,
	callback func(event *heap_profiler.LastSeenObjectIDEvent),
) {
	handler := protocol.NewEventHandler(
		"HeapProfiler.lastSeenObjectID",
		func(name string, params []byte) {
			event := &heap_profiler.LastSeenObjectIDEvent{}
			if err := json.Unmarshal(params, event); err != nil {
				log.Error(err)
			} else {
				callback(event)
			}
		},
	)
	socket.AddEventHandler(handler)
}

/*
OnReportHeapSnapshotProgress adds a handler to the DOM.ReportHeapSnapshotProgress event.
EXPERIMENTAL
*/
func (HeapProfiler) OnReportHeapSnapshotProgress(
	socket *Socket,
	callback func(event *heap_profiler.ReportHeapSnapshotProgressEvent),
) {
	handler := protocol.NewEventHandler(
		"HeapProfiler.reportHeapSnapshotProgress",
		func(name string, params []byte) {
			event := &heap_profiler.ReportHeapSnapshotProgressEvent{}
			if err := json.Unmarshal(params, event); err != nil {
				log.Error(err)
			} else {
				callback(event)
			}
		},
	)
	socket.AddEventHandler(handler)
}

/*
OnResetProfiles adds a handler to the HeapProfiler.ResetProfiles event. EXPERIMENTAL
*/
func (HeapProfiler) OnResetProfiles(
	socket *Socket,
	callback func(event *heap_profiler.ResetProfilesEvent),
) {
	handler := protocol.NewEventHandler(
		"HeapProfiler.resetProfiles",
		func(name string, params []byte) {
			event := &heap_profiler.ResetProfilesEvent{}
			if err := json.Unmarshal(params, event); err != nil {
				log.Error(err)
			} else {
				callback(event)
			}
		},
	)
	socket.AddEventHandler(handler)
}
