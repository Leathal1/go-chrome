package chrome

import (
	"app/chrome/protocol"
	"encoding/json"

	log "github.com/Sirupsen/logrus"
)

/*
Debugger - https://chromedevtools.github.io/devtools-protocol/tot/Debugger/
Exposes JavaScript debugging capabilities. It allows setting and removing breakpoints, stepping
through execution, exploring stack traces, etc.
*/
type Debugger struct{}

/*
ContinueToLocation continues execution until specific location is reached.
*/
func (Debugger) ContinueToLocation(
	socket *Socket,
	params *debugger.ContinueToLocationParams,
) (nil, error) {
	command := &protocol.Command{
		method: "Debugger.continueToLocation",
		params: params,
	}
	socket.SendCommand(command)
	return nil, command.Err
}

/*
Disable disables debugger for given page.
*/
func (Debugger) Disable(
	socket *Socket,
) (nil, error) {
	command := &protocol.Command{
		method: "Debugger.disable",
	}
	socket.SendCommand(command)
	return nil, command.Err
}

/*
Enable enables debugger for the given page. Clients should not assume that the debugging has been
enabled until the result for this command is received.
*/
func (Debugger) Enable(
	socket *Socket,
) (nil, error) {
	command := &protocol.Command{
		method: "Debugger.enable",
	}
	socket.SendCommand(command)
	return nil, command.Err
}

/*
EvaluateOnCallFrame evaluates expression on a given call frame.
*/
func (Debugger) EvaluateOnCallFrame(
	socket *Socket,
	params *debugger.EvaluateOnCallFrameParams,
) (debugger.EvaluateOnCallFrameResult, error) {
	command := &protocol.Command{
		method: "Debugger.evaluateOnCallFrame",
		params: params,
	}
	socket.SendCommand(command)
	return command.Result.(debugger.EvaluateOnCallFrameResult), command.Err
}

/*
GetPossibleBreakpoints returns possible locations for breakpoint. scriptId in start and end range
locations should be the same.
*/
func (Debugger) GetPossibleBreakpoints(
	socket *Socket,
	params *debugger.GetPossibleBreakpointsParams,
) (debugger.GetPossibleBreakpointsResult, error) {
	command := &protocol.Command{
		method: "Debugger.getPossibleBreakpoints",
		params: params,
	}
	socket.SendCommand(command)
	return command.Result.(debugger.GetPossibleBreakpointsResult), command.Err
}

/*
GetScriptSource returns source for the script with given id.
*/
func (Debugger) GetScriptSource(
	socket *Socket,
	params *debugger.GetScriptSourceParams,
) (debugger.GetScriptSourceResult, error) {
	command := &protocol.Command{
		method: "Debugger.getScriptSource",
		params: params,
	}
	socket.SendCommand(command)
	return command.Result.(debugger.GetScriptSourceResult), command.Err
}

/*
GetStackTrace returns stack trace with given stackTraceId. EXPERIMENTAL
*/
func (Debugger) GetStackTrace(
	socket *Socket,
	params *debugger.GetStackTraceParams,
) (debugger.GetStackTraceResult, error) {
	command := &protocol.Command{
		method: "Debugger.getStackTrace",
		params: params,
	}
	socket.SendCommand(command)
	return command.Result.(debugger.GetStackTraceResult), command.Err
}

/*
Pause stops on the next JavaScript statement.
*/
func (Debugger) Pause(
	socket *Socket,
) (nil, error) {
	command := &protocol.Command{
		method: "Debugger.pause",
	}
	socket.SendCommand(command)
	return nil, command.Err
}

/*
PauseOnAsyncCall EXPERIMENTAL
*/
func (Debugger) PauseOnAsyncCall(
	socket *Socket,
	params *debugger.PauseOnAsyncCallParams,
) (nil, error) {
	command := &protocol.Command{
		method: "Debugger.pauseOnAsyncCall",
		params: params,
	}
	socket.SendCommand(command)
	return nil, command.Err
}

/*
RemoveBreakpoint removes JavaScript breakpoint.
*/
func (Debugger) RemoveBreakpoint(
	socket *Socket,
	params *debugger.RemoveBreakpointParams,
) (nil, error) {
	command := &protocol.Command{
		method: "Debugger.removeBreakpoint",
		params: params,
	}
	socket.SendCommand(command)
	return nil, command.Err
}

/*
RestartFrame restarts particular call frame from the beginning.
*/
func (Debugger) RestartFrame(
	socket *Socket,
	params *debugger.RestartFrameParams,
) (debugger.RestartFrameResult, error) {
	command := &protocol.Command{
		method: "Debugger.restartFrame",
		params: params,
	}
	socket.SendCommand(command)
	return command.Result.(debugger.RestartFrameResult), command.Err
}

/*
Resume resumes JavaScript execution.
*/
func (Debugger) Resume(
	socket *Socket,
) (nil, error) {
	command := &protocol.Command{
		method: "Debugger.resume",
	}
	socket.SendCommand(command)
	return nil, command.Err
}

/*
ScheduleStepIntoAsync is deprecated - use Debugger.stepInto with breakOnAsyncCall and
Debugger.pauseOnAsyncTask instead. Steps into next scheduled async task if any is scheduled before
next pause. Returns success when async task is actually scheduled, returns error if no task were
scheduled or another scheduleStepIntoAsync was called. EXPERIMENTAL
*/
func (Debugger) ScheduleStepIntoAsync(
	socket *Socket,
) (nil, error) {
	command := &protocol.Command{
		method: "Debugger.scheduleStepIntoAsync",
	}
	socket.SendCommand(command)
	return nil, command.Err
}

/*
SearchInContent searches for given string in script content.
*/
func (Debugger) SearchInContent(
	socket *Socket,
	params *debugger.SearchInContentParams,
) (debugger.SearchInContentResult, error) {
	command := &protocol.Command{
		method: "Debugger.searchInContent",
		params: params,
	}
	socket.SendCommand(command)
	return command.Result.(debugger.SearchInContentResult), command.Err
}

/*
SetAsyncCallStackDepth enables or disables async call stacks tracking.
*/
func (Debugger) SetAsyncCallStackDepth(
	socket *Socket,
	params *debugger.SetAsyncCallStackDepthParams,
) (nil, error) {
	command := &protocol.Command{
		method: "Debugger.setAsyncCallStackDepth",
		params: params,
	}
	socket.SendCommand(command)
	return nil, command.Err
}

/*
SetBlackboxPatterns replaces previous blackbox patterns with passed ones. Forces backend to skip
stepping/pausing in scripts with url matching one of the patterns. VM will try to leave blackboxed
script by performing 'step in' several times, finally resorting to 'step out' if unsuccessful.
EXPERIMENTAL
*/
func (Debugger) SetBlackboxPatterns(
	socket *Socket,
	params *debugger.SetBlackboxPatternsParams,
) (nil, error) {
	command := &protocol.Command{
		method: "Debugger.setBlackboxPatterns",
		params: params,
	}
	socket.SendCommand(command)
	return nil, command.Err
}

/*
SetBlackboxedRanges makes backend skip steps in the script in blackboxed ranges. VM will try leave
blacklisted scripts by performing 'step in' several times, finally resorting to 'step out' if
unsuccessful. Positions array contains positions where blackbox state is changed. First interval
isn't blackboxed. Array should be sorted.
*/
func (Debugger) SetBlackboxedRanges(
	socket *Socket,
	params *debugger.SetBlackboxedRangesParams,
) (nil, error) {
	command := &protocol.Command{
		method: "Debugger.setBlackboxedRanges",
		params: params,
	}
	socket.SendCommand(command)
	return nil, command.Err
}

/*
SetBreakpoint sets JavaScript breakpoint at a given location.
*/
func (Debugger) SetBreakpoint(
	socket *Socket,
	params *debugger.SetBreakpointParams,
) (debugger.SetBreakpointResult, error) {
	command := &protocol.Command{
		method: "Debugger.setBreakpoint",
		params: params,
	}
	socket.SendCommand(command)
	return command.Result.(debugger.SetBreakpointResult), command.Err
}

/*
SetBreakpointByURL sets JavaScript breakpoint at given location specified either by URL or URL
regex. Once this command is issued, all existing parsed scripts will have breakpoints resolved and
returned in locations property. Further matching script parsing will result in subsequent
breakpointResolved events issued. This logical breakpoint will survive page reloads.
*/
func (Debugger) SetBreakpointByURL(
	socket *Socket,
	params *debugger.SetBreakpointByURLParams,
) (debugger.SetBreakpointByURLResult, error) {
	command := &protocol.Command{
		method: "Debugger.setBreakpointByUrl",
		params: params,
	}
	socket.SendCommand(command)
	return command.Result.(debugger.SetBreakpointByURLResult), command.Err
}

/*
SetBreakpointsActive activates / deactivates all breakpoints on the page.
*/
func (Debugger) SetBreakpointsActive(
	socket *Socket,
	params *debugger.SetBreakpointsActiveParams,
) (nil, error) {
	command := &protocol.Command{
		method: "Debugger.setBreakpointsActive",
		params: params,
	}
	socket.SendCommand(command)
	return nil, command.Err
}

/*
SetPauseOnExceptions defines the pause on exceptions state. Can be set to stop on all exceptions,
uncaught exceptions or no exceptions. Initial pause on exceptions state is none.
*/
func (Debugger) SetPauseOnExceptions(
	socket *Socket,
	params *debugger.SetPauseOnExceptionsParams,
) (nil, error) {
	command := &protocol.Command{
		method: "Debugger.setPauseOnExceptions",
		params: params,
	}
	socket.SendCommand(command)
	return nil, command.Err
}

/*
SetReturnValue changes return value in top frame. Available only at return break position.
EXPERIMENTAL
*/
func (Debugger) SetReturnValue(
	socket *Socket,
	params *debugger.SetReturnValueParams,
) (nil, error) {
	command := &protocol.Command{
		method: "Debugger.setReturnValue",
		params: params,
	}
	socket.SendCommand(command)
	return nil, command.Err
}

/*
SetScriptSource edits JavaScript source live.
*/
func (Debugger) SetScriptSource(
	socket *Socket,
	params *debugger.SetScriptSourceParams,
) (debugger.SetScriptSourceResult, error) {
	command := &protocol.Command{
		method: "Debugger.setScriptSource",
		params: params,
	}
	socket.SendCommand(command)
	return command.Result.(debugger.SetScriptSourceResult), command.Err
}

/*
SetSkipAllPauses makes page not interrupt on any pauses (breakpoint, exception, dom exception etc).
*/
func (Debugger) SetSkipAllPauses(
	socket *Socket,
	params *debugger.SetSkipAllPausesParams,
) (nil, error) {
	command := &protocol.Command{
		method: "Debugger.setSkipAllPauses",
		params: params,
	}
	socket.SendCommand(command)
	return nil, command.Err
}

/*
SetVariableValue changes value of variable in a callframe. Object-based scopes are not supported and
must be mutated manually.
*/
func (Debugger) SetVariableValue(
	socket *Socket,
	params *debugger.SetVariableValueParams,
) (nil, error) {
	command := &protocol.Command{
		method: "Debugger.setVariableValue",
		params: params,
	}
	socket.SendCommand(command)
	return nil, command.Err
}

/*
StepInto steps into the function call.
*/
func (Debugger) StepInto(
	socket *Socket,
	params *debugger.StepIntoParams,
) (nil, error) {
	command := &protocol.Command{
		method: "Debugger.stepInto",
		params: params,
	}
	socket.SendCommand(command)
	return nil, command.Err
}

/*
StepOut steps out of the function call.
*/
func (Debugger) StepOut(
	socket *Socket,
) (nil, error) {
	command := &protocol.Command{
		method: "Debugger.stepOut",
	}
	socket.SendCommand(command)
	return nil, command.Err
}

/*
StepOver steps over the statement.
*/
func (Debugger) StepOver(
	socket *Socket,
) (nil, error) {
	command := &protocol.Command{
		method: "Debugger.stepOver",
	}
	socket.SendCommand(command)
	return nil, command.Err
}

/*
OnBreakpointResolved adds a handler to the Debugger.breakpointResolved event.
Debugger.breakpointResolved fires when breakpoint is resolved to an actual script and location.
*/
func (Debugger) OnBreakpointResolved(
	socket *Socket,
	callback func(event *debugger.BreakpointResolvedEvent),
) {
	handler := protocol.NewEventHandler(
		"Debugger.breakpointResolved",
		func(name string, params []byte) {
			event := &debugger.BreakpointResolvedEvent{}
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
OnPaused adds a handler to the Debugger.paused event. Debugger.paused fires when the virtual machine
stopped on breakpoint or exception or any other stop criteria.
*/
func (Debugger) OnPaused(
	socket *Socket,
	callback func(event *debugger.PausedEvent),
) {
	handler := protocol.NewEventHandler(
		"Debugger.paused",
		func(name string, params []byte) {
			event := &debugger.PausedEvent{}
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
OnResumed adds a handler to the Debugger.resumed event. Debugger.resumed fires when the virtual
machine resumes execution.
*/
func (Debugger) OnResumed(
	socket *Socket,
	callback func(event *debugger.ResumedEvent),
) {
	handler := protocol.NewEventHandler(
		"Debugger.resumed",
		func(name string, params []byte) {
			event := &debugger.ResumedEvent{}
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
OnScriptFailedToParse adds a handler to the Debugger.scriptFailedToParse event.
Debugger.scriptFailedToParse fires when the virtual machine fails to parse the script.
*/
func (Debugger) OnScriptFailedToParse(
	socket *Socket,
	callback func(event *debugger.ScriptFailedToParseEvent),
) {
	handler := protocol.NewEventHandler(
		"Debugger.scriptFailedToParse",
		func(name string, params []byte) {
			event := &debugger.ScriptFailedToParseEvent{}
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
OnScriptParsed adds a handler to the Debugger.ScriptParsed event. Debugger.ScriptParsed fires when
virtual machine parses script. This event is also fired for all known and uncollected scripts upon
enabling debugger.
*/
func (Debugger) OnScriptParsed(
	socket *Socket,
	callback func(event *debugger.ScriptParsedEvent),
) {
	handler := protocol.NewEventHandler(
		"Debugger.scriptParsed",
		func(name string, params []byte) {
			event := &debugger.ScriptParsedEvent{}
			if err := json.Unmarshal(params, event); err != nil {
				log.Error(err)
			} else {
				callback(event)
			}
		},
	)
	socket.AddEventHandler(handler)
}