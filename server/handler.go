package server

import "reflect"

type GRPCHandler struct {
	name    string
	handler interface{}
}

func NewGPRCHandler(handler interface{}, opts ...Option) *GRPCHandler {
	options := Options{
		Name: DefaultName,
	}

	//typ := reflect.TypeOf(handler)
	hdlr := reflect.ValueOf(handler)
	name := reflect.Indirect(hdlr).Type().Name()

	for _, o := range opts {
		o(&options)
	}

	// for m := 0; m < typ.NumMethod(); m++ {
	// 	if e := extractEndpoint(typ.Method(m)); e != nil {
	// 		e.Name = name + "." + e.Name

	// 		for k, v := range options.Metadata[e.Name] {
	// 			e.Metadata[k] = v
	// 		}

	// 		endpoints = append(endpoints, e)
	// 	}
	// }

	return &GRPCHandler{
		name:    name,
		handler: handler,
	}
}
