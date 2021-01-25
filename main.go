// main.go

package main

import (
  	"net/http"
  	"github.com/aws/aws-lambda-go/events"
  	"github.com/aws/aws-lambda-go/lambda"
 	"encoding/json"
 	"fmt"
)

func handleGetStatus(req events.APIGatewayProxyRequest) (events.APIGatewayProxyResponse, error) {
	
cpf, found := req.QueryStringParameters["cpf"]
if !found {
	return events.APIGatewayProxyResponse{
      StatusCode: http.StatusInternalServerError,
      Body: http.StatusText(http.StatusInternalServerError),
    }, nil
}

fmt.Println("rawParam1" + cpf)

  status, err := getStatus(cpf)
  if err != nil {
    return events.APIGatewayProxyResponse{
      StatusCode: http.StatusInternalServerError,
      Body: http.StatusText(http.StatusInternalServerError),
    }, nil
  }

  js, err := json.Marshal(status)
  if err != nil {
    return events.APIGatewayProxyResponse{
      StatusCode: http.StatusInternalServerError,
      Body: http.StatusText(http.StatusInternalServerError),
    }, nil
  }

  return events.APIGatewayProxyResponse{
    StatusCode: http.StatusOK,
    Body: string(js),
  }, nil
}

func handleInsertStatus(req events.APIGatewayProxyRequest) (events.APIGatewayProxyResponse, error) {

  var status Status
  err := json.Unmarshal([]byte(req.Body), &status)
  if err != nil {
    return events.APIGatewayProxyResponse{
      StatusCode: http.StatusInternalServerError,
      Body:       err.Error(),
    }, nil
  }

  fmt.Println("55" + req.Body)
  fmt.Println("66" + status.Cpf)

  err = insertStatusFake(status)
  if err != nil {
    return events.APIGatewayProxyResponse{
      StatusCode: http.StatusInternalServerError,
      Body:       err.Error(),
    }, nil
  }
fmt.Println("9")
  return events.APIGatewayProxyResponse{
    StatusCode: http.StatusCreated,
    Body:       "Status Fake Inserido",
  }, nil
}







func router(req events.APIGatewayProxyRequest) (events.APIGatewayProxyResponse, error) {

fmt.Println("1.1")

  if req.Path == "/status" {
  	fmt.Println("1.2")
    if req.HTTPMethod == "GET" {
      return handleGetStatus(req)
      //return "hjwgdjsjghd"
    }
    if req.HTTPMethod == "POST" {
    	fmt.Println("2")
      return handleInsertStatus(req)
    }
  }
  /*if req.Path == "/movimentacao" {
    if req.HTTPMethod == "GET" {
      return handleGetMovimentacao(req)
    }
    if req.HTTPMethod == "POST" {
      return handleInsertMovimentacao(req)
    }
  }*/
fmt.Println("1.f")
  return events.APIGatewayProxyResponse{
    StatusCode: http.StatusMethodNotAllowed,
    Body:       http.StatusText(http.StatusMethodNotAllowed),
  }, nil
}

func main() {
	fmt.Println("1")
        lambda.Start(router)
        fmt.Println("F")
}