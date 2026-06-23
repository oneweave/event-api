package controller

const (
	Prefix = "ctl"

	ControllerUpdateRequestedV1Type = "controller.update.requested.v1"
	ControllerUpdateSucceededV1Type = "controller.update.succeeded.v1"
	ControllerUpdateRejectedV1Type  = "controller.update.rejected.v1"
	ControllerUpdateFailedV1Type    = "controller.update.failed.v1"
)

type ControllerUpdateBaseData struct{}
