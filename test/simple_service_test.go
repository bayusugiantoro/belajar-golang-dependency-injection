package test

import (
	"belajar-golang-restful-api/simple"
	"fmt"
	"testing"
)

func TestSimpleService(t *testing.T) {
	simpleService := simple.InitializedService()
	fmt.Println(simpleService.SimpleRepository)
}