package ga

import (
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
func CreateIndividual(nodes) *Individual {
	var flattench []int
	var chromosome [][]int
	customersIDList = nodes.CusotmersIDList()
	copy(flattench, customersIDList)
	chromosome = shapeFlatToVehicles(flattench)
	indv = &Individual{chromosome: chromosome}
	return indv
}
