package inventory

import (
	"encoding/json"
	"fmt"
	"net/http"

	"github.com/GolangPractice/gRPC2/proto"
	"google.golang.org/grpc"

	"io/ioutil"

	"github.com/gin-gonic/gin"
)

func Connection() proto.AvailabilityServiceClient {
	conn, err := grpc.Dial("localhost:8060", grpc.WithInsecure()) //connect to GRPC server
	if err != nil {
		panic(err)
	}
	return proto.NewAvailabilityServiceClient(conn)
}

func AddItemHandler(ctx *gin.Context) {
	client := Connection()
	newItem := proto.AvailableRequest{}              //create empty message
	reqBody, err := ioutil.ReadAll(ctx.Request.Body) //read request body
	if err != nil {
		panic(err)
	}
	result := json.Unmarshal(reqBody, &newItem) //Unmarshal json to message object
	fmt.Println(result)
	if response, err := client.AddItem(ctx, &newItem); err == nil {
		ctx.JSON(http.StatusOK, gin.H{
			"result": fmt.Sprint(response.Status),
		})
	} else {
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
	}

}

func UpdateItemHandler(ctx *gin.Context) {
	client := Connection()
	updItem := proto.AvailableRequest{}
	reqBody, err := ioutil.ReadAll(ctx.Copy().Request.Body)
	if err != nil {
		panic(err)
	}
	result := json.Unmarshal(reqBody, &updItem)
	fmt.Println(result)
	if response, err := client.UpdateInventory(ctx, &updItem); err == nil {
		ctx.JSON(http.StatusOK, gin.H{
			"result": fmt.Sprint(response.Status),
		})
	} else {
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
	}
}

func GetDetailsHandler(ctx *gin.Context) {
	client := Connection()
	id := ctx.Param("id") //get ItemID

	req := &proto.AvailableRequest{ItemId: id}                    //create message object
	if response, err := client.GetDetails(ctx, req); err == nil { //Call Server method
		ctx.JSON(http.StatusOK, gin.H{
			"result": response,
		})
	} else {
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
	}
}

func CheckInventoryHandler(ctx *gin.Context) {
	client := Connection()
	param := ctx.Param("status")
	fmt.Println(param)
	req := &proto.StatusResponse{Status: param}
	if response, err := client.CheckInventory(ctx, req); err == nil {
		ctx.JSON(http.StatusOK, gin.H{
			"result": response,
		})
	} else {
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
	}
}

func RemoveItemHandler(ctx *gin.Context) {
	client := Connection()
	id := ctx.Param("id")

	req := &proto.AvailableRequest{ItemId: id}
	if response, err := client.RemoveItem(ctx, req); err == nil {
		ctx.JSON(http.StatusOK, gin.H{
			"result": response,
		})
	} else {
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
	}
}
