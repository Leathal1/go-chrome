package chrome

import (
	"encoding/json"

	headless_experimental "github.com/mkenney/go-chrome/headless_experimental"
	"github.com/mkenney/go-chrome/protocol"

	log "github.com/Sirupsen/logrus"
)

/*
HeadlessExperimental - https://chromedevtools.github.io/devtools-protocol/tot/HeadlessExperimental/
Provides experimental commands only supported in headless mode. EXPERIMENTAL
*/
type HeadlessExperimental struct{}

/*
BeginFrame sends a BeginFrame to the target and returns when the frame was completed. Optionally
captures a screenshot from the resulting frame. Requires that the target was created with enabled
BeginFrameControl.
*/
func (HeadlessExperimental) BeginFrame(
	socket *Socket,
	params *headless_experimental.BeginFrameParams,
) (headless_experimental.BeginFrameResult, error) {
	command := &protocol.Command{
		Method: "HeadlessExperimental.beginFrame",
		Params: params,
	}
	result := headless_experimental.BeginFrameResult{}
	socket.SendCommand(command)

	if nil != command.Err {
		return result, command.Err
	}

	if nil != command.Result {
		resultData, err := json.Marshal(command.Result)
		if nil != err {
			return result, err
		}

		err = json.Unmarshal(resultData, &result)
		if nil != err {
			return result, err
		}
	}

	return result, command.Err
}

/*
Disable disables headless events for the target.
*/
func (HeadlessExperimental) Disable(
	socket *Socket,
) error {
	command := &protocol.Command{
		Method: "HeadlessExperimental.disable",
	}
	socket.SendCommand(command)
	return command.Err
}

/*
Enable enables headless events for the target.
*/
func (HeadlessExperimental) Enable(
	socket *Socket,
) error {
	command := &protocol.Command{
		Method: "HeadlessExperimental.enable",
	}
	socket.SendCommand(command)
	return command.Err
}

/*
OnNeedsBeginFramesChanged adds a handler to the HeadlessExperimental.needsBeginFramesChanged event.
HeadlessExperimental.needsBeginFramesChanged fires when the target starts or stops needing
BeginFrames.
*/
func (HeadlessExperimental) OnNeedsBeginFramesChanged(
	socket *Socket,
	callback func(event *headless_experimental.NeedsBeginFramesChangedEvent),
) {
	handler := protocol.NewEventHandler(
		"HeadlessExperimental.needsBeginFramesChanged",
		func(name string, params []byte) {
			event := &headless_experimental.NeedsBeginFramesChangedEvent{}
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
OnMainFrameReadyForScreenshots adds a handler to the
HeadlessExperimental.mainFrameReadyForScreenshots event.
HeadlessExperimental.mainFrameReadyForScreenshots fires when the main frame has first submitted a
frame to the browser. May only be fired while a BeginFrame is in flight. Before this event,
screenshotting requests may fail.
*/
func (HeadlessExperimental) OnMainFrameReadyForScreenshots(
	socket *Socket,
	callback func(event *headless_experimental.MainFrameReadyForScreenshotsEvent),
) {
	handler := protocol.NewEventHandler(
		"HeadlessExperimental.mainFrameReadyForScreenshots",
		func(name string, params []byte) {
			event := &headless_experimental.MainFrameReadyForScreenshotsEvent{}
			if err := json.Unmarshal(params, event); err != nil {
				log.Error(err)
			} else {
				callback(event)
			}
		},
	)
	socket.AddEventHandler(handler)
}
