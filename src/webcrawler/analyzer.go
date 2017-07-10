package main
import (

)
type ParseResponse func(httpResp *http.Response,respDepth uint32)([]Data,[]error)
type Analyzer interface{
	Module
	RespParsers()[]ParseResponse
	Analyze(resp *Response)([]Data,[]error)
}



















