package ga

import (
	"../node"
	"reflect"
)

type Individual struct {
	Chromosome [][]int
	Distance   float64
	Fitness    float64
}

//***********************//
// Methods of Individual //
//***********************//
func (indv *Individual) IsEqual(counterpart *Individual) bool {
	return reflect.DeepEqual(indv.Chromosome, counterpart.Chromosome)
}

func (indv *Individual) NVehicle() int {
	nvehicle := len(indv.Chromosome)
	return nvehicle
}

//***********//
// Functions //
//***********//
func CreateIndividual(nodes *node.NodeList) *Individual {
	flattench := make([]int, 0, len(nodes.CusotmersIDList()))
	chromosome := make([][]int, 0)
	customersIDList := nodes.CusotmersIDList()
	copy(flattench, customersIDList)
	chromosome = shapeFlatToVehicles(nodes, flattench)
	indv := &Individual{Chromosome: chromosome}
	return indv
}
