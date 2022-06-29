package service

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
	"sort"
)

type CountW struct {
	word  string
	count int
}

type CountW_ struct {
	Word  string `json:"word"`
	Count int    `json:"count"`
}

type CountWs []CountW_

func PostRequest(w http.ResponseWriter, r *http.Request) {
	content, err := ioutil.ReadAll(r.Body)
	if err != nil {
		fmt.Println("eroor in requesting body", err)

	}
	s := []string{string(content)}
	new_ := WcCount(s)
	WC := make([]CountW, 0, len(new_))
	for key, val := range new_ {
		WC = append(WC, CountW{word: key, count: val})
	}

	sort.Slice(WC, func(i, j int) bool {
		return WC[i].count > WC[j].count
	})
	for i := 0; i < len(WC) && i < 10; i++ {
		count_ := CountWs{CountW_{WC[i].word, WC[i].count}}
		json.Marshal(count_)
		json.NewEncoder(w).Encode(count_)
	}
}
