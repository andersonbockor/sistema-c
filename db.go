// db.go

package main

import (
	
  "github.com/aws/aws-sdk-go/aws"
  "github.com/aws/aws-sdk-go/aws/session"
  "github.com/aws/aws-sdk-go/service/dynamodb"
  "github.com/aws/aws-sdk-go/service/dynamodb/dynamodbattribute"
  "github.com/aws/aws-sdk-go/service/dynamodb/expression"
  
  "fmt"
)
const AWS_REGION = "sa-east-1"
const TABLE_NAME = "status"

var db = dynamodb.New(session.New(), aws.NewConfig().WithRegion(AWS_REGION));


func getStatus(cpf string) ([]Status, error) {

	filt := expression.Name("cpf").Equal(expression.Value(cpf))
	//proj := expression.NamesList(expression.Name("SongTitle"), expression.Name("AlbumTitle"))
	//expr, err := expression.NewBuilder().WithFilter(filt).WithProjection(proj).Build()
	expr, err := expression.NewBuilder().WithFilter(filt).Build()
	if err != nil {
  		fmt.Println(err)
	}

  	input := &dynamodb.ScanInput{
  		ExpressionAttributeNames:  expr.Names(),
  		ExpressionAttributeValues: expr.Values(),
  		FilterExpression:          expr.Filter(),
  		//ProjectionExpression:      expr.Projection(),
    	TableName: aws.String(TABLE_NAME),
  	}
  
  	result, err := db.Scan(input)
  	if err != nil {
    	return []Status{}, err
  	}
  	if len(result.Items) == 0 {
    	return []Status{}, nil
  	}

  	var status[]Status
  	err = dynamodbattribute.UnmarshalListOfMaps(result.Items, &status)
  	if err != nil {
    	return []Status{}, err
  	}

  return status, nil
}

// CreateUser inserts a new User item to the table.
func insertStatusFake(status Status) error {

  // Generates a new random ID
  //uuid, err := uuid.NewV4()
  //if err != nil {
  //  return err
  //}
 // err := 0


 //av, 
 /*err := dynamodbattribute.MarshalMap(status)
if err != nil {
    fmt.Println("Got error marshalling new movie item:")
    fmt.Println(err.Error())
    os.Exit(1)
}*/
fmt.Println("3")
fmt.Println("44" + status.Cpf)
input := &dynamodb.PutItemInput{
    //Item:      av,
    TableName: aws.String(TABLE_NAME),
    Item: map[string]*dynamodb.AttributeValue{
      "cpf": {
        S: aws.String(status.Cpf),
      },
      "nome": {
        S: aws.String(status.Nome),
      },
    },
}

  /*// Creates the item that's going to be inserted
  input := &dynamodb.PutItemInput{
    TableName: aws.String(TABLE_NAME),
    Item: map[string]*dynamodb.AttributeValue{
      "id": {
        S: aws.String(strconv.Itoa(status.id)),
      },
      "cpd": {
        S: aws.String(status.cpf),
      },
      "nome": {
        S: aws.String(status.nome),
      },
      /*"dataUltimaConsulta": {
        S: aws.String(status.Name),
      },
      "valorUltimaConsulta": {
        S: aws.String(fmt.Sprintf("%v")),
      },
      "estabelecimentoUltimaConsulta": {
        S: aws.String(user.Name),
      },
      "bureauUltimaConsulta": {
        S: aws.String(fmt.Sprintf("%v")),
      },
      "dataUltimaCompraCartao": {
        S: aws.String(user.Name),
      },
      "valorUltimaCompraCartao": {
        S: aws.String(fmt.Sprintf("%v")),
      },
      "estabelecimentoUltimaCompraCartao": {
        S: aws.String(user.Name),
      },*/
 //   },
  //}

  _, err := db.PutItem(input)
  return err
}