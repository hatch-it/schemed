package main

import (
	"reflect"
	"strings"
	"github.com/gin-gonic/gin"
)

type Route struct {
	Name		string
	Method		string
	Pattern		string
	HandlerFunc gin.HandlerFunc
}

type Routes []Route

func GenerateRoutesForService(s Service, routes *Routes) {
	var name string

	if t := reflect.TypeOf(s); t.Kind() == reflect.Ptr {
		name = t.Elem().Name()
	} else {
		name = t.Name()
	}

	pattern := "/" + strings.ToLower(name)

	*routes = append(*routes, Routes{
		Route{
			name + " Get",
			"GET",
			pattern + ":id",
			GetService,
		},
		Route{
			name + " Fetch",
			"GET",
			pattern,
			FetchService,
		},
		Route{
			name + " Create",
			"POST",
			pattern,
			CreateService,
		},
		Route{
			name + " Update",
			"POST",
			pattern + ":id",
			UpdateService,
		},
		Route{
			name + " Delete",
			"DELETE",
			pattern + ":id",
			DeleteService,
		},
	}...)
}

var routes = make(Routes, 10)