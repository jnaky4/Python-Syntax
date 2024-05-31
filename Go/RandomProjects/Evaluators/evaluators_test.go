// package main
package main

import (
	"testing"
)

type testEvaluation struct {
	resource      ComputeResource
	expectedScore int
}

//	var testResources = []testEvaluation{
//		{ComputeResource{"hv1", 1, 50, "scp", .75, .2}, 25},
//		{ComputeResource{"hv2", 3, 5000, "scp", .4, .2}, 60},
//		{ComputeResource{"hv3", 2, 500, "scp", .25, .7}, 75},
//		{ComputeResource{"hvLowCore", 0, 5000, "scp", .4, .2}, 0},
//		{ComputeResource{"hvLowRam", 3, 0, "scp", .4, .2},0},
//		{ComputeResource{"hvNoTag", 2, 500, "", .1, .1},0},
//	}
var testResources = []testEvaluation{
	{ComputeResource{"hv1", 1, 50, "scp", .75, .2, -1}, 25},
	{ComputeResource{"hv2", 3, 5000, "scp", .4, .2, -1}, 60},
	{ComputeResource{"hv3", 2, 500, "scp", .25, .7, -1}, 75},
	{ComputeResource{"hvLowCore", 0, 5000, "scp", .4, .2, -1}, 0},
	{ComputeResource{"hvLowRam", 3, 0, "scp", .4, .2, -1}, 0},
	{ComputeResource{"hvNoTag", 2, 500, "", .1, .1, -1}, 0},
}

//for i := 0; i < 10000; i++{
//testResources = append (testResources, ComputeResource{
//fmt.Sprintf("hv%d", i), rand.Intn(5), rand.Intn(500), "scp", rand.Float64(), rand.Float64()})
//}

func TestEvaluateResource(t *testing.T) {
	expectedChoice := testEvaluation{ComputeResource{"expected", 2, 500, "scp", .15, .7, -1}, 85}

	testResources = append(testResources, expectedChoice)

	requested := ComputeResource{"requested", 1, 50, "scp", 0, 0, -1}
	for available := range testResources {
		evaluateResourcePointer(&requested, &testResources[available].resource)
	}

	highest := ComputeResource{Score: -1}
	for _, resource := range testResources {

		if resource.resource.Score > highest.Score {
			highest = resource.resource
		}
	}

	if highest.Id != expectedChoice.resource.Id {
		t.Errorf("evaluateResourcePointer: expected %v, actual %v\n", expectedChoice, highest)
	}
}

func TestEachEvaluateResourceScore(t *testing.T) {

	requested := ComputeResource{"requested", 1, 50, "scp", 0, 0, -1}

	for available := range testResources {
		evaluateResourcePointer(&requested, &testResources[available].resource)
	}

	for _, results := range testResources {

		if results.resource.Score != results.expectedScore {
			t.Errorf("evaluate Score %s: expected %d, actual %d\n", results.resource.Id, results.resource.Score, results.expectedScore)
		}

	}
}

func TestEvaluateResources(t *testing.T) {

	var testResourcesArr []ComputeResource
	for i := 0; i < len(testResources); i++ {
		testResourcesArr = append(testResourcesArr, testResources[i].resource)
	}

	expectedChoice := testEvaluation{ComputeResource{"expected", 2, 500, "scp", .1, .7, -1}, 90}
	testResourcesArr = append(testResourcesArr, expectedChoice.resource)

	scoreChannel := make(chan ComputeResource, len(testResources))
	requested := ComputeResource{"requested", 1, 50, "scp", 0, 0, -1}

	go evaluateResourcesPointer(scoreChannel, &requested, &testResourcesArr)

	highest := <-scoreChannel

	if highest.Id != expectedChoice.resource.Id {
		t.Errorf("evaluateResourcePointer: expected %v, actual %v\n", expectedChoice, highest)
	}
}

//func TestEvaluateResource(t *testing.T) {
//	expectedChoice := testEvaluation{ComputeResource{"expected", 2, 500, "scp", .15, .7}, 85}
//
//	testResources = append(testResources, expectedChoice)
//
//	scoreChannel := make(chan Score, len(testResources))
//	requested := ComputeResource{"requested", 1, 50, "scp", 0, 0}
//	for available := range testResources {
//		go evaluateResourcePointer(scoreChannel, &requested, &testResources[available].resource)
//	}
//	highest := Score{Score: 0}
//	for range testResources {
//		csv := <- scoreChannel
//		if csv.Score > highest.Score {
//			highest = csv
//		}
//	}
//
//	if highest.Id != expectedChoice.resource.Id{
//		t.Errorf("evaluateResourcePointer: expected %v, actual %v\n", expectedChoice, highest)
//	}
//}
//
//func TestEachEvaluateResourceScore(t *testing.T){
//	scoreChannel := make(chan Score, len(testResources))
//
//	requested := ComputeResource{"requested", 1, 50, "scp", 0, 0}
//
//	for available := range testResources {
//		go evaluateResourcePointer(scoreChannel, &requested, &testResources[available].resource)
//	}
//
//	for range testResources {
//		results := <- scoreChannel
//		for _, available := range testResources {
//			if available.resource.Id == results.Id && available.expectedScore != results.Score{
//				t.Errorf("evaluate Score %s: expected %d, actual %d\n", results.Id ,available.expectedScore, results.Score)
//			}
//		}
//	}
//}
//
//func TestEvaluateResources(t *testing.T){
//
//	var testResourcesArr []ComputeResource
//	for i := 0; i < len(testResources); i++ {
//		testResourcesArr = append(testResourcesArr, testResources[i].resource)
//	}
//
//	expectedChoice := testEvaluation{ComputeResource{"expected", 2, 500, "scp", .1, .7}, 90}
//	testResourcesArr = append(testResourcesArr, expectedChoice.resource)
//
//	scoreChannel := make(chan Score, len(testResources))
//	requested := ComputeResource{"requested", 1, 50, "scp", 0, 0}
//
//	go evaluateResourcesPointer(scoreChannel, &requested, &testResourcesArr)
//
//	highest := <- scoreChannel
//
//	if highest.Id != expectedChoice.resource.Id{
//		t.Errorf("evaluateResourcePointer: expected %v, actual %v\n", expectedChoice, highest)
//	}
//}

//func BenchmarkEvaluateResource(b *testing.B){
//	scoreChannel := make(chan Score, len(testResources))
//
//	requested := ComputeResource{"requested", 1, 50, "scp", 0, 0}
//
//	for i := 0; i < b.N; i++ {
//		evaluateResourcePointer(scoreChannel, &requested, &testResources[0].resource)
//	}
//	for range testResources {
//		_ = <-scoreChannel
//	}
//}

//todo Compute Evaluator

// todo compute evaluator

//defer wg.Done()
//for job := range jobs{
//	EvaluateResources(scoreChannel, &job.ComputeRequest, &availableResources)
//
//}
//}

// Todo Constants
