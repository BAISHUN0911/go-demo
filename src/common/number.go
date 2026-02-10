package common

import (
	"bytes"
	"encoding/json"
	"fmt"
	"math"
)

func Number_demo() {
	var a int = 10
	const Pi float64 = 3.1415926

	sum1 := float64(a) + Pi
	fmt.Println(sum1)
}

func PointScoringSystem25() {
	var score1 int
	var score2 int
	var score3 int
	score1 = 80
	score2 = 90
	score3 = 71
	floatScore := float64(score1 + score2 + score3)
	finalScore := (floatScore / 3) * 0.25
	fmt.Println(math.Round(finalScore*100) / 100)
}

func Deserialization() {
	var m map[string]interface{}
	json.Unmarshal([]byte(`{"score": 1.23}`), &m)
	fmt.Println(m)
	fmt.Printf("Type of score: %T\n", m["score"]) // Type of score: float64

	var m2 map[string]interface{}
	dec := json.NewDecoder(bytes.NewReader([]byte(`{"score": 1.23}`)))
	dec.UseNumber()
	dec.Decode(&m2)
	fmt.Println(m2)
	fmt.Printf("Type of score: %T\n", m2["score"]) // Type of score: json.Number
}
