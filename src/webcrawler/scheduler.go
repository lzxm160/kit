package main
import (

)
type Scheduler interface{
	Init(requestArgs RequestArgs,dataArgs DataArgs,moduleArgs ModuleArgs)error
	Start(firstHTTPReq *http.Request)error
	Stop()error
	Status()Status
	ErrorChan()<-chan error
	Idle()bool
	Summary()SchedSummary
}

type RequestArgs struct{
	AcceptedDomains []string `json:"accepted_primary_domains`
	MaxDepth uint32 `json:"max_depth"`
}
type DataArgs struct{
	ReqBufferCap uint32 `json:"req_buffer_cap"`
	ReqMaxBufferCap uint32 `json:"req_max_buffer_number"`

	RespBufferCap uint32 `json:"resp_buffer_cap"`
	RespMaxBufferCap uint32 `json:"resp_max_buffer_number"`

	ItemBufferCap uint32 `json:"item_buffer_cap"`
	ItemMaxBufferCap uint32 `json:"item_max_buffer_number"`

	ErrorBufferCap uint32 `json:"error_buffer_cap"`
	ErrorMaxBufferCap uint32 `json:"error_max_buffer_number"`
}
type ModuleArgs struct{
	Downloaders []module.Downloaders
	Analyzers []module.Analyzers
	Pipelines []module.Pipelines
}
type Args interface{
	Check()error
}
type Status uint8
const(
	SCHED_STATUS_UNINITIALIZED Status=0
	SCHED_STATUS_INITIALIZING Status=1
	SCHED_STATUS_INITIALIZED Status=2
	SCHED_STATUS_STARTING Status=3
	SCHED_STATUS_STARTED Status=4
	SCHED_STATUS_STOPPING Status=5
	SCHED_STATUS_STOPPED Status=6
)
type SchedSummary interface{
	Struct() SummaryStruct
	String() string
}
type SummaryStruct struct{
	RequestArgs RequestArgs `json:"request_args"`
	DataArgs DataArgs `json:"data_args"`
	ModuleArgs ModuleArgsSummary `json:"module_args"`
	Status string `json:"status"`
	Downloaders []module.SummaryStruct `json:"downloaders"`
	Analyzers []module.SummaryStruct `json:"analyzers"`
	Pipelines []module.SummaryStruct `json:"pipelines"`
	ReqBufferPool BufferPoolSummaryStruct `json:request_buffer_pool`
	RespBufferPool BufferPoolSummaryStruct `json:response_buffer_pool`
	ItemBufferPool BufferPoolSummaryStruct `json:item_buffer_pool`
	ErrorBufferPool BufferPoolSummaryStruct `json:error_buffer_pool`
	NumURL uint64 `json:url_number`
}













