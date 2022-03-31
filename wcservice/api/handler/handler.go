package handler

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"regexp"
	"sort"
	"strconv"
	"strings"
)

var inputdatafinal []string

type WordCount struct {
	word  string
	count int
}

func InputHandler(w http.ResponseWriter, r *http.Request) {

	body, err := ioutil.ReadAll(r.Body)

	if err != nil {
		log.Fatal(err)
	}
	getdata := string(body)

	re, err := regexp.Compile(`[^\w]`)
	if err != nil {
		log.Fatal(err)
	}

	getdata = re.ReplaceAllString(getdata, " ")
	getdata = strings.ToLower(getdata)

	str, err := json.Marshal(getdata)
	if err != nil {
		fmt.Printf("Error: %s", err.Error())
	}

	userinput := string(str[:])

	words := strings.Fields(userinput)

	m := make(map[string]int)
	for _, word := range words {

		_, ok := m[word]
		if !ok {
			m[word] = 1
		} else {
			m[word]++
		}
	}

	// create and fill slice of word-count pairs for sorting by count
	wordCounts := make([]WordCount, 0, len(m))
	for key, val := range m {
		wordCounts = append(wordCounts, WordCount{word: key, count: val})
	}

	// sort wordCount slice by decreasing count number
	sort.Slice(wordCounts, func(i, j int) bool {
		return wordCounts[i].count > wordCounts[j].count
	})

	// display the three most frequent words
	for i := 0; i < len(wordCounts) && i < 10; i++ {
		// fmt.Println(wordCounts[i].word, ":", wordCounts[i].count)
		// fmt.Fprintf(w, "The count is: ", wordCounts[i].word, ":", wordCounts[i].count)
		inputdatafinal = append(inputdatafinal, " word: ", wordCounts[i].word, " occurs: ", strconv.Itoa(wordCounts[i].count))
	}
	fmt.Fprintf(w, "The count is: %s", inputdatafinal)
}
