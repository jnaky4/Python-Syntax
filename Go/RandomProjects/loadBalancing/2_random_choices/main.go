package main

import (
	"fmt"
	"math"
	"math/rand"
)

type ComputeResource struct {
	Id    string
	Cores uint8
	Ram   uint8
}

type Node struct {
	Resource []ComputeResource
	sumCores int
	sumCPU   int
	Id       int
}

func RandUint8() uint8 {
	return uint8(rand.Intn(math.MaxUint8 + 1))
}
func get2DifferentNodes(n int) (int, int) {
	index := rand.Intn(n)
	index2 := rand.Intn(n)

	for { //re-roll if node is same
		if index != index2 {
			break
		}
		index2 = rand.Intn(n)
	}
	return index, index2
}

func main() {
	r := 20000
	n := 5000

	var resources []ComputeResource
	for i := 0; i < r; i++ { // create r resources
		resources = append(resources, ComputeResource{
			fmt.Sprintf("hv%d", i), RandUint8(), RandUint8()})
	}

	NodeMap2Choice := map[int]Node{}

	for i := 0; i < n; i++ { // create n nodes
		NodeMap2Choice[i] = Node{Id: i, Resource: []ComputeResource{}}
	}

	for _, v := range resources { //place r resources on n nodes
		index, index2 := get2DifferentNodes(n)

		nmc := NodeMap2Choice
		ramSum1 := 0
		cpuSum1 := 0
		for _, i := range nmc[index].Resource {
			ramSum1 += int(i.Ram)
			cpuSum1 += int(i.Cores)
		}

		ramSum2 := 0
		cpuSum2 := 0
		for _, i := range nmc[index2].Resource {
			ramSum2 += int(i.Ram)
			cpuSum2 += int(i.Cores)
		}

		if ramSum1 >= ramSum2 {
			//println("Picked 2")
			index = index2
		}

		//fmt.Printf("ID: %d , len(%d)\n", n.Id, len(n.Resource))
		//fmt.Printf("Cores: %d , Ram: %d\n", cpu_sum, ram_sum)

		if entry, ok := NodeMap2Choice[index]; ok {
			entry.Resource = append(entry.Resource, v)
			NodeMap2Choice[index] = entry
		}

	}

	//println("NodeMap")
	//for _, n := range NodeMap{
	//	ram_sum := 0
	//	cpu_sum := 0
	//	for _, r := range n.Resource{
	//		ram_sum += r.Ram
	//		cpu_sum += r.Cores
	//	}
	//	fmt.Printf("ID: %d , len(%d)\n", n.Id, len(n.Resource))
	//	fmt.Printf("Cores: %d , Ram: %d\n", cpu_sum, ram_sum)
	//}

	println()
	println("NodeMap 2 random choices")
	for _, n := range NodeMap2Choice {
		ramSum, cpuSum := 0, 0
		ramMean, cpuMean := 0.0, 0.0
		for _, r := range n.Resource {
			ramSum += r.Ram
			cpuSum += r.Cores
		}
		ramMean = float64(ramSum / len(n.Resource))
		cpuMean = float64(cpuSum / len(n.Resource))
		fmt.Printf("ID: %d , len(%d)\n", n.Id, len(n.Resource))
		fmt.Printf("Cores: %d , Ram: %d\n", cpuSum, ramSum)
		fmt.Printf("Mean Cores: %f , Mean Ram: %f\n", cpuMean, ramMean)
	}

}
