package model

import "github.com/globalxtreme/gobaseconf/model"

type RabbitMQAsyncWorkflowStep struct {
	model.RabbitMQBaseModel
	WorkflowId     uint                           `gorm:"column:workflowId;type:bigint"`
	Service        string                         `gorm:"column:service;type:varchar(100);not null"`
	Queue          string                         `gorm:"column:queue;type:varchar(200);not null"`
	StepOrder      int                            `gorm:"column:stepOrder;type:int"`
	StatusId       int                            `gorm:"column:statusId;type:tinyint"`
	Description    string                         `gorm:"column:description;type:text;null"`
	Payload        *model.MapInterfaceColumn      `gorm:"column:payload;type:json;default:null"`
	ForwardPayload *model.MapInterfaceColumn      `gorm:"column:forwardPayload;type:json;default:null"`
	Errors         *model.ArrayMapInterfaceColumn `gorm:"column:errors;type:json;default:null"`
	Response       *model.MapInterfaceColumn      `gorm:"column:response;type:json;default:null"`
	Reprocessed    float64                        `gorm:"column:reprocessed;type:int;default:0"`
}

func (RabbitMQAsyncWorkflowStep) TableName() string {
	return "async_workflow_steps"
}
