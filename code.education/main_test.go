package main

import (
	"testing"
)

func TestGreeting(t *testing.T){
	got := greeting("Code.Education Rocks");
	want := "<b>Code.Education Rocks</b>";

	if got != want{
		t.Errorf("greeting(Code.Education Rocks) \n got: %v \n want: %v", got, want);
	}
}