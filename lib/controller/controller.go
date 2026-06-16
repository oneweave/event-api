package controller

const (
	ControllerUpdateRequestedV1Type = "controller.update.requested.v1"
	ControllerUpdateSucceededV1Type = "controller.update.succeeded.v1"
	ControllerUpdateRejectedV1Type  = "controller.update.rejected.v1"
	ControllerUpdateFailedV1Type    = "controller.update.failed.v1"
)

type ControllerUpdatedEventBaseData struct {
	ControllerID string `json:"controllerId" bson:"controller_id" validate:"required,uuid"`
	ServiceID    string `json:"serviceId" bson:"service_id" validate:"required,uuid"`
}
