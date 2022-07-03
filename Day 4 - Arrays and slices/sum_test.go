package main

import (
	"reflect"
	"testing"
)

func TestSum(t *testing.T){
	
	
	t.Run("Collection of 6 numbers", func(t *testing.T){
		value := [6]int{1,2,3,4,5,6}
		got := Sum(value)
		want := 21
		if got != want {
			t.Errorf("got %d want %d given, %v", got, want, value)
		}
	})

	t.Run("Collection of any size numbers", func(t *testing.T){
		value := []int{1,2,3,4,5,6}
		got := SumUnknown(value)
		want := 21

		if got != want {
			t.Errorf("got %d want %d given, %v", got, want, value)
		}
	})

}

func TestSumAll(t *testing.T){
	got := SumAll([]int{1,2}, []int{0,9})
	want := []int{3,9}

	if !reflect.DeepEqual(got,want) {
		t.Errorf("got %v want %v", got,want)
	}
}

func TestSumAllTails(t *testing.T){

	checkSums := func(t testing.TB, got, want []int){
		t.Helper()
		if !reflect.DeepEqual(got, want){
			t.Errorf("got %v want %v", got , want)
		}
	}
	


	t.Run("make the sum of some slices", func(t* testing.T){
		got := SumAllTails([]int{1,2}, []int{0,9})
		want := []int{2,9}
	
		checkSums(t,got,want)
	})

	t.Run("Safely sum empty slices", func(t* testing.T){
		got := SumAllTails([]int{}, []int{0,9})
		want := []int{0,9}

		checkSums(t,got,want)
	})

}