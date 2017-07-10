package main
import (
	"net/http"
	"fmt"
)
type Request struct{
	httpReq *http.Request
	depth uint32
}
func NewRequest(httpReq *http.Request,depth uint32)*Request {
	return &Request{httpReq:httpReq,depth:depth}
}
func (this *Request)HTTPReq()*http.Request {
	return this.httpReq
}
func (this *Request)Depth()uint32 {
	return this.depth
}
//////////////////////////////////////////////////
type Response struct{
	httpResp *http.Response
	depth uint32
}
func NewResponse(httpResp *http.Response,depth uint32)*Response {
	return &Response{httpResp:httpResp,depth:depth}
}
func (this *Response)HTTPResp()*http.Response {
	return this.httpResp
}
func (this *Response)Depth()uint32 {
	return this.depth
}
///////////////////////////////////////////////////
type Item map[string]interface{}
//////////////////////////////////////////////////
type Data interface{
	Valid()bool
}
func (this *Request)Valid()bool {
	return this.httpReq!=nil&&this.httpReq.URL!=nil
}
func (this *Response)Valid()bool {
	return this.httpResp!=nil&&this.httpResp.Body!=nil
}
func (this *Item)Valid()bool {
	return this!=nil
}
///////////////////////////////////////////////////
type ErrorType string
const(
	ERROR_TYPE_DOWNLOADER ErrorType="downloader error"
	ERROR_TYPE_ANALYZER ErrorType="analyzer error"
	ERROR_TYPE_PIPELINE ErrorType="pipeline error"
	ERROR_TYPE_SCHEDULER ErrorType="scheduler error"
)
type CrawlerError interface{
	Type() ErrorType
	Error()string
}
type myCrawlerError struct{
	errType ErrorType
	errMsg string
	fullErrMsg string
}
func NewCrawlerError(errType ErrorType,errMsg string)CrawlerError {
	return &myCrawlerError{errType:errType,errMsg:errMsg}
}
func (this *myCrawlerError)Type()ErrorType {
	return this.errType
}
func (this *myCrawlerError)Error()string {
	if this.fullErrMsg==""{
		this.getFullErrMsg()
	}
	return this.fullErrMsg
}
func (this *myCrawlerError)getFullErrMsg() {
	this.fullErrMsg=fmt.Sprintf("CrawlerError:%s:%s",this.errType,this.errMsg)
}















