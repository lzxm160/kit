package main
import (

)
type ProcessItem func(item Item)(result Item,err error)

type Pipeline interface{
	Module
	ItemProcessors()[]ProcessItem
	Send(item Item)[]error
	FailFast()bool
	setFailFast(failFast bool)
}



















