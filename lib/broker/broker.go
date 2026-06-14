package broker

const (
	BrokerUpdateRequestedV1Type = "broker.update.requested.v1"
	BrokerUpdateSucceededV1Type = "broker.update.succeeded.v1"
	BrokerUpdateRejectedV1Type  = "broker.update.rejected.v1"
	BrokerUpdateFailedV1Type    = "broker.update.failed.v1"
)

type BrokerUpdateEventData struct {
	BrokerID           string `json:"brokerId" bson:"broker_id" validate:"required,uuid"`
	ServiceID          string `json:"serviceId" bson:"service_id" validate:"required,uuid"`
	ServiceVersionHash string `json:"serviceVersionHash" bson:"service_version_hash" validate:"required,sha256"`
}

func NewBrokerUpdateEventData() BrokerUpdateEventData {
	return BrokerUpdateEventData{}
}
