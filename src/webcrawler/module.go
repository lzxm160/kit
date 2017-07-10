package main
import (
	"net/http"
)
type MID string
type CalculateScore func(counts Counts)uint64
type SummaryStruct struct{
	ID MID `json:"id"`
	Called uint64 `json:"called"`
	Accepted uint64 `json:"accepted"`
	Completed uint64 `json:"completed"`
	Handling uint64 `json:"handling"`
	Extra interface{} `json:"extra,omitempty"`
}
type Module interface{
	ID() MID
	Addr()string
	Score()uint64
	SetScore(score uint64)
	ScoreCalculator()CalculateScore
	CalledCount()uint64
	AcceptedCount()uint64
	CompletedCount()uint64
	HandlingNumber()uint64
	Counts()Counts
	Summary()SummaryStruct
}
var midTemplate="%s%d|%s"
type Type string
const(
	TYPE_DOWNLOADER Type="downloader"
	TYPE_ANALYZER Type="analyzer"
	TYPE_PIPELINE Type="pipeline"
)
var legalTypeLetterMap=map[Type]string{
	TYPE_DOWNLOADER:"D",
	TYPE_ANALYZER:"A",
	TYPE_PIPELINE:"P",
}
type SNGenerator interface{
	Start()uint64
	Max()uint64
	Next()uint64
	CycleCount()uint64
	Get()uint64
}
type Registeror interface{
	Register(module Module)(bool,error)
	Unregister(mid MID)(bool,error)
	Get(moduleType Type)(Module,error)
	GetAllByType(moduleType Type)(map[MID]Module,error)
	GetAll()map[MID]Module
	Clear()
}

type ModuleInternal interface{
	Module
	IncrCalledCount()
	IncrAcceptedCount()
	IncrCompletedCount()
	IncrHandlingNumber()
	DecrHandlingNumber()
	Clear()
}

type myModule struct{
	mid MID
	addr string
	score uint64
	scoreCalculator CalculateScore
	calledCount uint64
	acceptedCount uint64
	completedCount uint64
	handlingNumber uint64
}
func NewModuleInternal(mid MID,scoreCalculator CalculateScore)(ModuleInternal,error) {
	parts,err:=SplitMID(mid)
	if err!=nil{
		return nil,errors.New(fmt.Sprintf("illegal ID %q:%s",mid,err))
	}
	return &myModule{mid:mid,addr:parts[2],ScoreCalculator:ScoreCalculator},nil
}
















