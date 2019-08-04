/**
 * 
 * @author wangchen
 * @version 2019-07-15 10:28
 */
package goform

import (
	"fmt"
	"reflect"
	"testing"
)


type SubTargetA struct{
	SubAfieldA string
	SubAfieldB int
}
type SubTargetB struct{
	SubBfieldA string
	SubBfieldB int
}

type ResponseTarget struct{
	FieldA string `type:"text" title:"FieldA"`
	FieldB int	`type:"text" title:"FieldA"`
	FieldC interface{}	`type:"json" title:"FieldC"`
	SubA []*SubTargetA `type:"table" title:"SubA"`
	SubB []*SubTargetB `type:"table" title:"SubB"`
}

func TestMarshalResponse(t *testing.T)  {
	listA := make([]*SubTargetA,0)
	listB := make([]*SubTargetB,0)
	listA = append(listA, &SubTargetA{
		"SubAfieldA1",
		111111111,
	})
	listA = append(listA, &SubTargetA{
		"SubAfieldA2",
		222222222,
	})
	listA = append(listA, &SubTargetA{
		"SubAfieldA3",
		333333333,
	})
	listB = append(listB, &SubTargetB{
		"SubTargetB1",
		111111111,
	})
	listB = append(listB, &SubTargetB{
		"SubTargetB2",
		222222222,
	})
	listB = append(listB, &SubTargetB{
		"SubTargetB3",
		333333333,
	})
	target := ResponseTarget{
		FieldA:"aaaaa",
		FieldB:123,
		FieldC:SubTargetA{
			"JsonA",
			12312312,
		},
		SubA:listA,
		SubB:listB,
	}
	s,_ := MarshalResponse(target)
	if s !=`[{"data":"aaaaa","title":"FieldA","type":"text"},{"data":"123","title":"FieldA","type":"text"},{"data":{"SubAfieldA":"JsonA","SubAfieldB":12312312},"title":"FieldC","type":"json"},{"data":{"columnNames":[{"key":"SubAfieldA","title":"SubAfieldA"},{"key":"SubAfieldB","title":"SubAfieldB"}],"tableDatas":[{"SubAfieldA":"SubAfieldA1","SubAfieldB":111111111},{"SubAfieldA":"SubAfieldA2","SubAfieldB":222222222},{"SubAfieldA":"SubAfieldA3","SubAfieldB":333333333}]},"title":"SubA","type":"table"},{"data":{"columnNames":[{"key":"SubBfieldA","title":"SubBfieldA"},{"key":"SubBfieldB","title":"SubBfieldB"}],"tableDatas":[{"SubBfieldA":"SubTargetB1","SubBfieldB":111111111},{"SubBfieldA":"SubTargetB2","SubBfieldB":222222222},{"SubBfieldA":"SubTargetB3","SubBfieldB":333333333}]},"title":"SubB","type":"table"}]`{
		fmt.Println(s)
		t.Fail()
	}
}


type FormTarget struct{
	FieldA string `type:"string" title:"FieldA" json:"field_a"`
	FieldB string	`type:"string" title:"FieldB" json:"field_b"`
	FieldC string	`type:"string" title:"FieldC" enum:"a,b,c" json:"field_c"`

}

func TestMarshalForm(t *testing.T) {
	s,_ := MarshalForm(reflect.TypeOf(FormTarget{}))
	if s!=`{"schema":{"FieldA":{"title":"FieldA","type":"string"},"FieldB":{"title":"FieldB","type":"string"},"FieldC":{"enum":["a","b","c"],"title":"FieldC","type":"string"}}}`{
		fmt.Println(s)
		t.Fail()
	}

}
