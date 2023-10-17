package builtin

import (
	"encoding/json"
	"fmt"
	"os"
	"reflect"
)

/*	名词解释： 【定义中json可以换成其他可存储形式】
	序列化（encoding）： 对象 =》 json
	反序列化（decoding） ： json =》 对象
*/

type response1 struct {
	Page   int
	Fruits []string
}

type response2 struct {
	Page   int      `json:"Page"`
	Fruits []string `json:"fruits"`
}

func IJSON() {
	//基础数据类型 序列化
	// json.Marshal 返回的是byte数组
	bolB, _ := json.Marshal(true)
	fmt.Println(string(bolB))

	intB, _ := json.Marshal(1)
	fmt.Println(string(intB))

	fltB, _ := json.Marshal(2.34)
	fmt.Println(string(fltB))

	strB, _ := json.Marshal("gopher")
	fmt.Println(string(strB))

	slcD := []string{"apple", "peahch", "pear"}
	slcB, _ := json.Marshal(slcD)
	fmt.Println(string(slcB))

	//struct to json
	res2D := &response2{
		Page:   1,
		Fruits: []string{"apple", "peahch", "pear"},
	}
	res2B, _ := json.Marshal(res2D)
	fmt.Println(string(res2B))

	//反序列化 Unmarshal
	byt := []byte(`{"num":6.13,"strs":["a","b"]}`)
	var dat map[string]interface{}
	if err := json.Unmarshal(byt, &dat); err != nil {
		panic(err)
	}
	fmt.Println(dat)
	fmt.Println(reflect.TypeOf(dat))

	enc := json.NewEncoder(os.Stdout)
	d := map[string]int{"apple": 5, "lettuce": 7}
	enc.Encode(d)

}
