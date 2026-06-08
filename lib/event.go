package lib

type BaseEvent struct {
	SpecVersion     string `json:"specversion" bson:"spec_version" validate:"required,eq=1.0"`
	ID              string `json:"id" bson:"id" validate:"required,uuid"`
	Source          string `json:"source" bson:"source" validate:"required"`
	Subject         string `json:"subject" bson:"subject" validate:"required"`
	Time            string `json:"time" bson:"time" validate:"required,datetime=2006-01-02T15:04:05Z07:00"`
	DataContentType string `json:"datacontenttype" bson:"data_content_type" validate:"required,eq=application/json"`
	Dataschema      string `json:"dataschema" bson:"data_schema" validate:"required"`
	CorrelationID   string `json:"correlationid" bson:"correlation_id" validate:"required,uuid"`
	CausationID     string `json:"causationid" bson:"causation_id" validate:"required,uuid"`
}

func NewBaseEvent() BaseEvent {
	return BaseEvent{
		SpecVersion:     "1.0",
		DataContentType: "application/json",
	}
}
