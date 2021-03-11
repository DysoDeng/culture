package service

import (
	"culture/cloud/base/internal/model"
	"culture/cloud/base/internal/support/api"
	"fmt"
)

type DemoService struct {}

func NewDemoService() *DemoService {
	return &DemoService{}
}
func (demo DemoService) Test(params string) (model.Demo, Error) {
	fmt.Println(params)
	return model.Demo{TestField: "test"}, Error{Code: api.CodeOk}
}
