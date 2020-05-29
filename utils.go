package main

import (
	"encoding/json"
	"fmt"
)

func initAnomaly() AnomalyDetector {
	svc := anomalyDetector{}
	return svc
}

type nostalgiaResponse struct {
	Result []App `json:"result"`
}

func (ret *nostalgiaResponse) getNostalgiaResponse() {
	var jsonStr = []byte(`{"result":[{"date":"2020-05-18T00:00:00Z","app":"24403717","dau":265,"requests":680,"responses":683},{"date":"2020-05-18T00:00:00Z","app":"79542943","dau":226,"requests":5956,"responses":5946},{"date":"2020-05-18T00:00:00Z","app":"10269465","dau":166,"requests":820},{"date":"2020-05-18T00:00:00Z","app":"10677633","dau":455,"requests":2809,"responses":2811},{"date":"2020-05-18T00:00:00Z","app":"40288377","dau":1,"requests":1,"responses":1},{"date":"2020-05-18T00:00:00Z","app":"16431510","dau":2188,"requests":30,"responses":28},{"date":"2020-05-18T00:00:00Z"},{"date":"2020-05-18T00:00:00Z","app":"43346372","dau":168604,"requests":373322,"responses":368639},{"date":"2020-05-18T00:00:00Z","app":"93472839","dau":2267,"requests":18664,"responses":4029},{"date":"2020-05-18T00:00:00Z","app":"66583952","dau":634,"requests":143,"responses":143},{"date":"2020-05-18T00:00:00Z","app":"79446414","dau":4403,"requests":39694,"responses":16433},{"date":"2020-05-18T00:00:00Z","app":"58898147","dau":116,"requests":371,"responses":371},{"date":"2020-05-18T00:00:00Z","app":"11294382","dau":4,"requests":3},{"date":"2020-05-18T00:00:00Z","app":"17994617","dau":9,"requests":142,"responses":141},{"date":"2020-05-18T00:00:00Z","app":"11538209","dau":1210,"requests":12238},{"date":"2020-05-18T00:00:00Z","app":"12693001","dau":11,"requests":33},{"date":"2020-05-18T00:00:00Z","app":"43480563","dau":2260,"requests":31216,"responses":31113},{"date":"2020-05-18T00:00:00Z","app":"29917280","dau":1897,"requests":2912,"responses":2900},{"date":"2020-05-18T00:00:00Z","app":"13391173","dau":26624,"requests":631850,"responses":414655},{"date":"2020-05-18T00:00:00Z","app":"14259489","dau":271,"requests":1}]}`)

	err := json.Unmarshal(jsonStr, &ret)
	if err != nil {
		fmt.Println(err)
	}

}
