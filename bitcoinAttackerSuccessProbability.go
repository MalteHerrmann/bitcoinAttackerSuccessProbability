package main

import (
	"embed"
	"encoding/json"
	"io"
	"log"
	"math"
	"net/http"
)

// Declare filesystem
var (
	//go:embed chart.js index.html
	staticFS embed.FS
)

// Define table struct with entities necessary for computations
type Table struct {
	q             []float64
	z             []int
	probabilities [][]float64
}

// Calculates the success probability for an attacker trying to harm
// the blockchain, depending on the probability q and number of transactions z
// in the chain.
func AttackerSuccessProbability(q float64, z int) float64 {
	p := 1.0 - q
	lambda := float64(z) * (q / p)
	sum := 1.0

	var poisson float64

	for k := 0; k <= z; k++ {
		poisson = poissonDensity(lambda, k)
		sum = sum - poisson*(1-math.Pow(q/p, float64(z-k)))
	}

	return sum
}

// Calculates the poisson density for the expected value lambda
// and the number of events k.
func poissonDensity(lambda float64, k int) float64 {
	poisson := math.Exp(-lambda)
	for i := 1; i <= k; i++ {
		poisson = poisson * (lambda / float64(i))
	}

	return poisson
}

// Calculates the values to be displayed
func calculateValues() (Table, error) {
	var q []float64
	var z []int
	var prob [][]float64

	for i_q := 0; i_q <= 49; i_q++ {
		var probRow []float64
		q_i := float64(i_q) / 100.0
		for z_i := 1; z_i < 100; z_i++ {
			res := AttackerSuccessProbability(q_i, z_i)
			probRow = append(probRow, res)
		}
		prob = append(prob, probRow)
	}

	var table Table
	table.q = q
	table.z = z
	table.probabilities = prob

	return table, nil
}

// tableJSON writes table data as JSON into w
func tableJSON(table Table, w io.Writer) error {
	reply := map[string]interface{}{
		"data": []map[string]interface{}{
			{
				"x":    table.q,
				"y":    table.z,
				"z":    table.probabilities,
				"type": "surface",
			},
		},
		"layout": map[string]interface{}{
			"title":    "Attacker Success Probability",
			"autosize": false,
			"width":    1500,
			"height":   1500,
		},
	}

	return json.NewEncoder(w).Encode(reply)
}

// Function to handle the http requests and responses
func dataHandler(w http.ResponseWriter, r *http.Request) {
	table, err := calculateValues()
	if err != nil {
		log.Printf("Error while calculating the values!")
	}

	if err := tableJSON(table, w); err != nil {
		log.Printf("table: %s", err)
	}
}

// Main function, that serves http on localhost and sends the
// calculated data to the html document.
func main() {
	http.Handle("/", http.FileServer(http.FS(staticFS)))
	http.HandleFunc("/data", dataHandler)

	if err := http.ListenAndServe(":8080", nil); err != nil {
		log.Fatal(err)
	}
}
