package main

import (
	"fmt"
	"net"
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

func main() {
	connection, erreur := net.Dial("tcp", "localhost:1234")
	if erreur != nil {
		fmt.Println(erreur.Error())
	}
	defer connection.Close()

	args := Args{1}
	var reply []HouseDAO

	client := jsonrpc.NewClient(connection)

	for i := 0; i < 1; i++ {
		erreur = client.Call("House.GetHouses", args, &reply)
		if erreur != nil {
			fmt.Println(erreur.Error())
		}
		fmt.Printf("House: ")
		fmt.Println(reply)

	}
}
