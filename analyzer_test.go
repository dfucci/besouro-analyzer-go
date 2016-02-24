package main

import "testing"

func TestInitialTimestamp(t *testing.T) {
	val := InitialAction("actions.txt", true)
	if val != 1412845504561 {
		t.Fatalf("Expected 1141284550456 but got %d", val)
	}
}

func TestFistEpisodeDuration(t *testing.T){
  val := FirstEpisodeDuration(1412845504561)
  if val != 847216 {
    t.Fatalf("Expected 12345, but got %d", val)
  }
}
