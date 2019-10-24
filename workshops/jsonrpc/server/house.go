package main

import (
	"fmt"
	"net"
	"net/rpc"
	"net/rpc/jsonrpc"
)

type HouseDAO struct {
	Id         int
	Name       string
	Region     string
	CoatOfArms string
	Words      string
}

var houses = []HouseDAO{
	HouseDAO{
		Id:         1,
		Name:       "House Algood",
		Region:     "The Westerlands",
		CoatOfArms: "A golden wreath, on a blue field with a gold border(Azure, a garland of laurel within a bordure or)",
		Words:      "",
	},
	HouseDAO{
		Id:         2,
		Name:       "House Allyrion of Godsgrace",
		Region:     "Dorne",
		CoatOfArms: "Gyronny Gules and Sable, a hand couped Or",
		Words:      "No Foe May Pass",
	},
	HouseDAO{
		Id:         3,
		Name:       "House Amber",
		Region:     "The North",
		CoatOfArms: "",
		Words:      "",
	},
}

type Args struct {
	Id int
}

type House int

func (t *House) GetHouse(args Args, reply *HouseDAO) error {
	fmt.Println("GetHouse")
	*reply = houses[args.Id-1]
	return nil
}

func (t *House) GetHouses(args Args, reply *[]HouseDAO) error {
	fmt.Println("GetHouses")
	*reply = houses
	return nil
}

func main() {
	house := new(House)

	server := rpc.NewServer()
	server.Register(house)

	server.HandleHTTP(rpc.DefaultRPCPath, rpc.DefaultDebugPath)

	listenning, e := net.Listen("tcp", ":1234")
	if e != nil {
		fmt.Println(e.Error())
	}
	defer listenning.Close() // Lancé à la fin de la fonction
	for {
		connection, e := listenning.Accept()
		if e != nil {
			fmt.Println(e.Error())
		}
		lol := jsonrpc.NewServerCodec(connection)
		server.ServeCodec(lol)
	}

}
