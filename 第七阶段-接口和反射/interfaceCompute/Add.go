package main

type AddOperator interface {
	 Add() interface{}
}


type  IntOarams struct {
	P1 int
	P2 int
}

// 实现一个方法
func (params *IntOarams) Add() interface{} {
	return params.P1 + params.P2
}

