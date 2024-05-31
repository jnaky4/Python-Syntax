package main

import (
	"fmt"
	k8s "github.com/cilium/cilium/pkg/k8s/slim/k8s/apis/meta/v1"
	"github.com/cilium/cilium/pkg/labels"
	cpa "github.com/cilium/cilium/pkg/policy/api"
)

type CiliumPolicy struct {
	Labels           `json:"labels"`
	EndpointSelector EndpointSelector `json:"endpointSelector"`
	Egress           []Egress         `json:"egress"`
}

type Egress struct {
	ToCIDRSet ToCIDRSet `json:"toCIDRSet,omitempty"`
	ToCIDR    []string  `json:"toCIDR,omitempty"`
}

type ToCIDRSet []struct {
	Cidr   string   `json:"cidr"`
	Except []string `json:"except,omitempty"`
}
type EndpointSelector struct {
	MatchLabels map[string]string `json:"matchLabels"`
}
type Labels []struct {
	Key   string `json:"key"`
	Value string `json:"value"`
}

func newPolicy(key string, value string, setFor []string, outBoundCIDR string, exceptions []string) []CiliumPolicy {
	cp := []CiliumPolicy{{
		Labels: Labels{{Key: key, Value: value}},
		EndpointSelector: EndpointSelector{
			MatchLabels: map[string]string{},
		},
	}}

	cp[0].EndpointSelector.MatchLabels[key] = value

	//egress := []struct {
	//	ToCIDRSet ToCIDRSet `json:"toCIDRSet,omitempty"`
	//	ToCIDR    []string  `json:"toCIDR,omitempty"`
	//}{
	//	{
	//
	//		ToCIDR: setFor,
	//		ToCIDRSet: ToCIDRSet{
	//			{
	//				Cidr: outBoundCIDR,
	//				Except: exceptions,
	//			},
	//		},
	//	},
	//}
	//cp[0].Egress[0] = egress

	return cp
	//cp[0].Egress = append(cp.Egress, egress...)

	//cp.Egress[0].ToCIDR = setFor
	//cp.Egress[0].ToCIDRSet[0].Cidr = outBoundCIDR
	//cp.Egress[0].ToCIDRSet[0].Except = exceptions
}

func NewCiliumPolicy() []CiliumPolicy {
	return []CiliumPolicy{{
		Labels: Labels{{Key: "test", Value: "test"}},
		EndpointSelector: EndpointSelector{
			MatchLabels: map[string]string{},
		},
	}}
}

//[{
//  "labels":[{"key":"app","value":"myService"}],
//  "endpointSelector":{"matchLabels":{"app":"myService"}},
//  "egress":[
//    {"toCIDRSet":[{"cidr":"10.0.0.0/8","except":["10.96.0.0/12"]}],
//      "toCIDR":["20.1.1.1/32"]
//    }
//  ]
//}]

func main() {
	rule := cpa.Rule{
		Egress: []cpa.EgressRule{{
			EgressCommonRule: cpa.EgressCommonRule{
				ToCIDR: cpa.CIDRSlice{
					"20.1.1.1/32",
				},
			},
		},
			{
				EgressCommonRule: cpa.EgressCommonRule{
					ToCIDRSet: []cpa.CIDRRule{{
						Cidr:        "10.0.0.0/8",
						ExceptCIDRs: []cpa.CIDR{"10.96.0.0/12"},
					}},
				},
			},
		},
		Labels: labels.LabelArray{
			labels.NewLabel("testKey", "testValue", ""),
		},
		Description: "test",
		EndpointSelector: cpa.EndpointSelector{
			LabelSelector: &k8s.LabelSelector{
				MatchLabels: map[string]k8s.MatchLabelsValue{
					"app": "myService",
				},
			},
		},
		Ingress: []cpa.IngressRule{{
			IngressCommonRule: cpa.IngressCommonRule{},
		}},
		IngressDeny:  []cpa.IngressDenyRule{},
		EgressDeny:   []cpa.EgressDenyRule{},
		NodeSelector: cpa.EndpointSelector{},
	}

	marshalJSON, err := rule.MarshalJSON()
	if err != nil {
		return
	}

	fmt.Printf("%s", marshalJSON)

	//T := "[{\"labels\": [{\"key\": \"app\", \"value\": \"myService\"}],\"endpointSelector\": {\"matchLabels\":{\"app\":\"myService\"}},\"egress\": [{\"toCIDR\": [\"20.1.1.1/32\"]}, {\"toCIDRSet\": [{\"cidr\": \"10.0.0.0/8\",\"except\": [\"10.96.0.0/12\"]}]}]}]"

	//cp := newPolicy("app", "myService",[]string{"20.1.1.1/32"}, "10.0.0.0/8", []string{"10.96.0.0/12"})
	//marshaling, err := json.Marshal(cp)
	//if err != nil {
	//	return
	//}
	//fmt.Printf("%s\n",marshaling)
	//
	//var cpt []T
	//T := "[{\"labels\": [{\"key\": \"app\", \"value\": \"myService\"}],\"endpointSelector\": {\"matchLabels\":{\"app\":\"myService\"}},\"egress\": [{\"toCIDR\": [\"20.1.1.1/32\"]}, {\"toCIDRSet\": [{\"cidr\": \"10.0.0.0/8\",\"except\": [\"10.96.0.0/12\"]}]}]}]"
	//err = json.Unmarshal([]byte(T), &cpt)
	//if err != nil {
	//	println(err.Error())
	//}
	//marshaling, err = json.Marshal(cpt)
	//if err != nil {
	//	return
	//}
	//fmt.Printf("%s\n",marshaling)
	//fmt.Printf("%+v\n", cpt)

}
