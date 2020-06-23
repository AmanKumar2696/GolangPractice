package main

import (
	"context"
	"fmt"
	"net"

	"github.com/GolangPractice/gRPC2/db"
	"github.com/GolangPractice/gRPC2/proto"
	"google.golang.org/grpc"
	"google.golang.org/grpc/reflection"
)

type Server struct{}

func main() {
	listener, err := net.Listen("tcp", ":8060") //Listen on port 8060
	if err != nil {
		panic(err)
	}
	srv := grpc.NewServer()
	proto.RegisterAvailabilityServiceServer(srv, &Server{}) //Bind service to the server
	reflection.Register(srv)

	if er := srv.Serve(listener); er != nil {
		panic(er)
	}
	fmt.Println("Working")
}

/*----------------Add Item--------------------------*/
func (s *Server) AddItem(ctx context.Context, request *proto.AvailableRequest) (*proto.StatusResponse, error) {
	db := db.Conn()  //Connect to Database
	defer db.Close() //Close db connection before function returns
	insert, err := db.Query("INSERT INTO inventory VALUES(?, ?, ?, ?)",
		request.GetItemId(), request.GetItemName(), request.GetPrice(), request.GetQuantity())
	if err != nil {
		panic(err)
	}
	defer insert.Close()

	return &proto.StatusResponse{Status: "Item Added"}, nil

}

/*----------------Get Details of 1 Item-------------------*/
func (s *Server) GetDetails(ctx context.Context, request *proto.AvailableRequest) (*proto.AvailableResponse, error) {
	db := db.Conn()
	defer db.Close()
	result := proto.AvailableResponse{}
	err := db.QueryRow("SELECT * FROM inventory WHERE ItemID = ?", request.GetItemId()).Scan(&result.ItemId,
		&result.ItemName, &result.Price, &result.Quantity)
	fmt.Println(err)
	return &result, nil

}

/*-----------------Get complete inventory-----------*/
func (s *Server) CheckInventory(ctx context.Context, request *proto.StatusResponse) (*proto.AllResponse, error) {
	db := db.Conn()
	defer db.Close()
	row, err := db.QueryContext(ctx, "SELECT * FROM inventory")
	if err != nil {
		fmt.Println(err)
	}
	var result []*proto.AvailableResponse
	for row.Next() {
		res := &proto.AvailableResponse{}
		rowscan := row.Scan(&res.ItemId, &res.ItemName, &res.Price, &res.Quantity)
		if rowscan != nil {
			panic(rowscan.Error())
		}
		result = append(result, res)
	}
	return &proto.AllResponse{Inventory: result}, nil

}

/*--------------Remove Item-----------------*/

func (s *Server) RemoveItem(ctx context.Context, request *proto.AvailableRequest) (*proto.StatusResponse, error) {
	db := db.Conn()
	defer db.Close()
	del, err := db.Query("DELETE FROM inventory WHERE ItemId = ?", request.GetItemId())
	if err != nil {
		fmt.Println(err)
	}
	del.Close()
	return &proto.StatusResponse{Status: "Item Removed"}, nil
}

/*---------------Update-----------------*/

func (s *Server) UpdateInventory(ctx context.Context, request *proto.AvailableRequest) (*proto.StatusResponse, error) {
	db := db.Conn()
	defer db.Close()
	update, err := db.Query("UPDATE inventory SET ItemName = ?, Price = ?, Quantity = ? WHERE ItemId = ?",
		request.GetItemName(), request.GetPrice(), request.GetQuantity(), request.GetItemId())
	if err != nil {
		fmt.Println(err)
	}
	update.Close()
	return &proto.StatusResponse{Status: "Item Updated"}, nil
}
