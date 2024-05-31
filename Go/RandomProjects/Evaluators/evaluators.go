package main

//package Evaluators

import (
	"fmt"
	"math"
	"time"
)

type ComputeResource struct {
	Id        string
	Cores     int
	Ram       float64
	AvgCpu15m float64
	Score float64
}

type Score struct {
	Id    string
	Score int
}

func Timer() func() {
	start := time.Now()
	return func() {
		fmt.Printf("elapsed: %s\n", time.Since(start))
	}
}

//func evaluateResource(c chan Score, requested *ComputeResource, available *ComputeResource){
//	overallScore := Score{available.Id, 1}
//	if requested.Cores > available.Cores{
//		//fmt.Printf("not enough cores %v\n", available)
//		overallScore.Score = 0
//	} else if requested.Ram > available.Ram{
//		//fmt.Printf("not enough ram %v\n", available)
//		// Ram is partitioned so no overallocation occurs, just need the requested amount
//		overallScore.Score = 0
//	} else if requested.Tag != "" && requested.Tag != available.Tag{
//		//fmt.Printf("wrong tag %v\n", available)
//		overallScore.Score = 0
//	} else {
//		//TODO is avgCPU available or used avg?
//		// Target over allocates cores 16:1 so cpu avg must be considered
//		overallScore.Score *= int((1 - available.AvgCpu15m) * 100)
//		//fmt.Printf("Score %v\n", overallScore)
//
//	}
//
//	c <- overallScore
//}

func evaluateResources(c chan Score, requested *ComputeResource, availableResources *[]ComputeResource) {
	highest := Score{"null", -1}
	overallScore := Score{}
	for _, available := range *availableResources {
		overallScore.Id = available.Id
		overallScore.Score = 1

		if requested.Cores > available.Cores {
			//fmt.Printf("not enough cores %v\n", available)
			overallScore.Score = 0
		} else if requested.Ram > available.Ram {
			//fmt.Printf("not enough ram %v\n", available)
			// Ram is partitioned so no overallocation occurs, just need the requested amount
			overallScore.Score = 0
		} else {
			//TODO is avgCPU available or used avg?
			// Target over allocates cores 16:1 so cpu avg must be considered
			overallScore.Score *= int((1 - available.AvgCpu15m) * 100)
			//fmt.Printf("Score %v\n", overallScore)
		}
		if overallScore.Score > highest.Score {
			highest = overallScore
		}
	}

	c <- highest
}

//func evaluateResourcePointer(requested *ComputeResource, available *ComputeResource) {
//	available.Score = 1
//	if requested.Cores > available.Cores {
//		available.Score = 0
//	} else if requested.Ram > available.Ram {
//		// Ram is partitioned so no overallocation occurs, just need the requested amount
//		available.Score = 0
//	} else {
//		//TODO is avgCPU available or used avg?
//		// Target over allocates cores 16:1 so cpu avg must be considered
//		available.Score *= int((1 - available.AvgCpu15m) * 100)
//	}
//
//}

//func evaluateResourcesPointer(c chan ComputeResource, requested *ComputeResource, availableResources *[]ComputeResource) {
//	highest := ComputeResource{Score: -1}
//
//	for _, available := range *availableResources {
//		go evaluateResourcePointer(requested, &available)
//
//		if available.Score > highest.Score {
//			highest = available
//		}
//	}
//
//	c <- highest
//}

//func evaluateResourceReflect(c chan Score, requested *ComputeResource, available *ComputeResource){
//	overallScore := Score{available.Id, 1}
//
//	rFields := reflect.TypeOf(*requested)
//	rValues := reflect.ValueOf(*requested)
//
//	aFields := reflect.TypeOf(*available)
//	aValues := reflect.ValueOf(*available)
//
//
//	var aValue reflect.value
//	var rValue reflect.value
//
//
//
//	num := rFields.NumField()
//	if num != aFields.NumField(){
//		println("Error, different length structs")
//	}
//
//	for i := 0; i < num; i++ {
//		aValue = aValues.Field(i)
//		rValue = rValues.Field(i)
//		rName := rValues.Type().Field(i).Name
//
//		switch rValue.Kind() {
//		case reflect.String:
//			rv := rValue.String()
//			av := aValue.String()
//
//			if rv != "" && rValue.Type().Name() == "Tag" && rv != av {
//				fmt.Printf("wrong tag %v\n", available)
//				overallScore.Score = 0
//			}
//			break
//		case reflect.Int, reflect.Int32, reflect.Int64:
//			//rv := strconv.FormatInt(rValue.Int(), 10)
//			rv := strconv.FormatInt(rValue.Int(), 10)
//			av := strconv.FormatInt(aValue.Int(), 10)
//			if  rv < av {
//				overallScore.Score = 0
//			}
//			break
//
//		case reflect.Float32, reflect.Float64:
//			//rv := rValue.Float()
//			av := aValue.Float()
//
//			if rName == "AvgCpu15m" {
//				overallScore.Score *= int((1 - av) * 100)
//			}
//			break
//
//		}
//
//	}
//
//	//fmt.Printf("Type: %s\tValue: %v\n", typeOfS., v.Field(i).)
//	//fmt.Printf("Field: %s\tValue: %v\n", typeOfS.Field(i).Id, v.Field(i).Interface())
//
//	c <- overallScore
//}

//func evaluateResourceArr(c chan Score, requested *ComputeResource, available *[]interface{}){
//
//}

func main() {
	testNormalizeWeights()
	//var resources []ComputeResource
	//resources = append(resources,
	//	ComputeResource{"hv1", 1, .5, .75, -1},
	//	ComputeResource{"hv2", 2, .3, .25, -1},
	//	ComputeResource{"hv3", 3, .7, .4, -1},
	//	ComputeResource{"hvNoCore", 0, .4, .4, -1},
	//	ComputeResource{"hvNoRam", 3, 1.0, .4, -1},
	//	ComputeResource{"hvNoTag", 2, .4, .1, -1},
	//	ComputeResource{"BAD", 1, .1, .1, -1},
	//)
	//
	//chosen := ComputeResource{}
	//for i,v := range resources{
	//	resources[i].Score = ScoreFitBinPack(v.Ram, v.AvgCpu15m)
	//	if resources[i].Score > chosen.Score{
	//		chosen = resources[i]
	//	}
	//}
	//
	//fmt.Printf("Chosen %v\n", chosen)
	//fmt.Printf("%v\n", resources)

	//for i := 0; i < 500000; i++ {
	//	resources = append(resources, ComputeResource{
	//		fmt.Sprintf("hv%d", i), rand.Intn(5), rand.Intn(500), rand.Float64(), -1})
	//}

	//testPointer(resources) 		  	//@ 5000 150-180 µs @ 500000 7.5 ms
	//testEvaluateResources(resources)	//@ 5000 200-300 µs @ 500000 3-6 ms
	//testEvaluateResourcesPointer(resources) //@ 5000 2.2 ms :(
	//test(resources) //@ 5000 3 ms :(
	//testReflect(resources) //@ 5000 5 ms :(

}

//func test(resources []ComputeResource){
//	defer Timer()()
//	scoreChannel := make(chan Score, len(resources))
//	requested := ComputeResource{"requested", 1, 50, "scp", 0, 0, -1}
//	for available := range resources{
//		go evaluateResource(scoreChannel, &requested, &resources[available])
//	}
//	highest := Score{Score: 0}
//	for range resources{
//		csv := <- scoreChannel
//		if csv.Score > highest.Score {
//			highest = csv
//		}
//	}
//	fmt.Printf("Highest Score Normal %v\n", highest)
//}

//func testPointer(resources []ComputeResource) {
//	defer Timer()()
//
//	requested := ComputeResource{"requested", 1, 50, 0,  -1}
//	for available := range resources {
//		evaluateResourcePointer(&requested, &resources[available])
//	}
//	highest := ComputeResource{Score: -1}
//	for _, resource := range resources {
//		if resource.Score > highest.Score {
//			highest = resource
//		}
//	}
//
//	fmt.Printf("Highest Score Normal Pointer %v\n", highest)
//}

func testEvaluateResources(resources []ComputeResource) {
	defer Timer()()
	chunks := 100

	scoreChannel := make(chan Score, len(resources))
	requested := ComputeResource{"requested", 1, 50, 0, -1}

	for i := 0; i < len(resources); i += chunks {
		if i+chunks >= len(resources) {
			resourcesRemaining := resources[i : len(resources)-1]
			go evaluateResources(scoreChannel, &requested, &resourcesRemaining)
		} else {
			resources5 := resources[i : i+chunks]
			go evaluateResources(scoreChannel, &requested, &resources5)
		}

	}

	highest := Score{}
	for i := 0; i < (len(resources) / chunks); i++ {
		temp := <-scoreChannel
		if temp.Score > highest.Score {
			highest = temp
		}
	}

	//for range resources{
	//	csv := <- scoreChannel
	//	if csv.Score > highest.Score {
	//		highest = csv
	//	}
	//}
	fmt.Printf("Highest Score Chunks %v\n", highest)
}

//func testReflect(resources []ComputeResource){
//	defer Timer()()
//	scoreChannel := make(chan Score, len(resources))
//	requested := ComputeResource{"requested", 1, 50, "scp", 0, 0, -1}
//	for available := range resources{
//		go evaluateResourceReflect(scoreChannel, &requested, &resources[available])
//	}
//	highest := Score{Score: 0}
//	for range resources{
//		csv := <- scoreChannel
//		if csv.Score > highest.Score {
//			highest = csv
//		}
//	}
//	fmt.Printf("Highest Score Reflect %v\n", highest)
//}

//func testEvaluateResourcesPointer(resources []ComputeResource) {
//	defer Timer()()
//	chunks := 100
//
//	scoreChannel := make(chan ComputeResource, len(resources))
//	requested := ComputeResource{"requested", 1, 50, 0, -1}
//
//	for i := 0; i < len(resources); i += chunks {
//		if i+chunks >= len(resources) {
//			resourcesRemaining := resources[i : len(resources)-1]
//			go evaluateResourcesPointer(scoreChannel, &requested, &resourcesRemaining)
//		} else {
//			resources5 := resources[i : i+chunks]
//			go evaluateResourcesPointer(scoreChannel, &requested, &resources5)
//		}
//
//	}
//
//	highest := ComputeResource{Score: -1}
//	for i := 0; i < (len(resources) / chunks); i++ {
//		csv := <-scoreChannel
//		if csv.Score > highest.Score {
//			highest = csv
//		}
//	}
//
//	//for range resources{
//	//	csv := <- scoreChannel
//	//	if csv.Score > highest.Score {
//	//		highest = csv
//	//	}
//	//}
//	fmt.Printf("Highest Score Chunks Pointer %v\n", highest)
//}


//func computeFreePercentage(node *Node, util *ComparableResources) (freePctCpu, freePctRam float64) {
//	// COMPAT(0.11): Remove in 0.11
//	reserved := node.ComparableReservedResources()
//	res := node.ComparableResources()
//
//	// Determine the node availability
//	nodeCpu := float64(res.Flattened.Cpu.CpuShares)
//	nodeMem := float64(res.Flattened.Memory.MemoryMB)
//	if reserved != nil {
//		nodeCpu -= float64(reserved.Flattened.Cpu.CpuShares)
//		nodeMem -= float64(reserved.Flattened.Memory.MemoryMB)
//	}
//
//	// Compute the free percentage
//	freePctCpu = 1 - (float64(util.Flattened.Cpu.CpuShares) / nodeCpu)
//	freePctRam = 1 - (float64(util.Flattened.Memory.MemoryMB) / nodeMem)
//	return freePctCpu, freePctRam
//}


func ScoreFitBinPack(freePctCpu float64, freePctRam float64) float64 {
	//freePctCpu, freePctRam := computeFreePercentage(node, util)

	// Total will be "maximized" the smaller the value is.
	// At 100% utilization, the total is 2, while at 0% util it is 20.
	//V3: score(i,j) = 10 free_ram_pct(i) + 10 free_cpu_pct(i)
	total := math.Pow(10, freePctCpu) + math.Pow(10, freePctRam)
	println("total: ", int(total))
	// Invert so that the "maximized" total represents a high-value
	// score. Because the floor is 20, we simply use that as an anchor.
	// This means at a perfect fit, we return 18 as the score.
	score := 20.0 - total

	// Bound the score, just in case
	// If the score is over 18, that means we've overfit the node.
	if score > 18.0 {
		score = 18.0
	} else if score < 0 {
		score = 0
	}
	return score
}

func testNormalizeWeights(){
	cpuAvg := 10.0
	requestRam := 200.0
	availableRam := 400.0
	maxCPUScore, minCPUScore := 40.0, 20.0

	minRAMScore, maxRamScore := requestRam, requestRam * 2

	//Normalize scores: x - xmin / xmax - xmin
	cpuScore := (cpuAvg - minCPUScore) / (maxCPUScore - minCPUScore) * 100
	fmt.Printf("%f\n",cpuScore)
	ramScore := (availableRam - minRAMScore) / (maxRamScore - minRAMScore) * 100
	fmt.Printf("%f\n", ramScore)

	//Weights
	//cpuW := .8
	//ramW := .2

	//total := (cpuW * cpuScore) + (ramW * ramScore)
	total := cpuScore
	fmt.Printf("%f\n", total)

}