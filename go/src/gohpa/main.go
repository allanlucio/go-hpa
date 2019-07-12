package main

import (
	"fmt"
	"net/http"
	"math"
)
func Greeting(nome string) string {
	return "<b>" +nome+ "</b>"
}

func SqrtLooping(valor float64) float64 {
	sqrtLoopingResult :=0.0;
	for i := 0.0; i < valor; i++ {
		sqrtLoopingResult += math.Sqrt(i)
	}

	return sqrtLoopingResult;

}
func main() {

	http.HandleFunc("/", func (w http.ResponseWriter, r *http.Request) {
		
		// s := fmt.Sprintf("%f", SqrtLooping())
		// fmt.Fprintf(w, Greeting(s))
		SqrtLooping(100000000.0)
		fmt.Fprintf(w, Greeting("Code.education Rocks!"))
	})

	fs := http.FileServer(http.Dir("static/"))
	http.Handle("/static/", http.StripPrefix("/static/", fs))

	http.ListenAndServe(":8000", nil)
}