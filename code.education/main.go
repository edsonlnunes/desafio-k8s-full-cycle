package main

import (
	"fmt"
	"net/http"
);

func greeting(phrase string) string{
	return "<b>" + phrase +"</b>";
}

func main(){
	
	phrase := greeting("Code.Education Rocks");

	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		fmt.Fprintf(w, phrase);
	});

	http.ListenAndServe(":8000", nil);
}