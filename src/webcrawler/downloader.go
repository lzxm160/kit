package main
import (

)
type Downloader interface{
	Module
	Download(req *Request)(*Response,error)
}



















