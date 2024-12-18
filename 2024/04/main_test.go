package main

import "testing"

func TestFindsWordLeftToRight(t *testing.T) {
	puzzle := [][]rune{
		{'X', 'M', 'A', 'S'},
		{' ', ' ', ' ', ' '},
		{' ', ' ', ' ', ' '},
		{' ', ' ', ' ', ' '},
	}
	count := searchWord(puzzle, "XMAS")
	if count != 1 {
		t.Errorf("Expected 1, got %d", count)
	}
}

func TestFindsWordRightToLeft(t *testing.T) {
	puzzle := [][]rune{
		{' ', ' ', ' ', ' '},
		{'S', 'A', 'M', 'X'},
		{' ', ' ', ' ', ' '},
		{' ', ' ', ' ', ' '},
	}
	count := searchWord(puzzle, "XMAS")
	if count != 1 {
		t.Errorf("Expected 1, got %d", count)
	}
}

func TestFindsWordUpToDown(t *testing.T) {
	puzzle := [][]rune{
		{' ', ' ', 'X', ' '},
		{' ', ' ', 'M', ' '},
		{' ', ' ', 'A', ' '},
		{' ', ' ', 'S', ' '},
	}
	count := searchWord(puzzle, "XMAS")
	if count != 1 {
		t.Errorf("Expected 1, got %d", count)
	}
}

func TestFindsWordDownToUp(t *testing.T) {
	puzzle := [][]rune{
		{' ', ' ', 'S', ' '},
		{' ', ' ', 'A', ' '},
		{' ', ' ', 'M', ' '},
		{' ', ' ', 'X', ' '},
	}
	count := searchWord(puzzle, "XMAS")
	if count != 1 {
		t.Errorf("Expected 1, got %d", count)
	}
}

func TestFindsWordDiagonalTopLeftToBottomRight(t *testing.T) {
	puzzle := [][]rune{
		{'X', ' ', ' ', ' '},
		{' ', 'M', ' ', ' '},
		{' ', ' ', 'A', ' '},
		{' ', ' ', ' ', 'S'},
	}
	count := searchWord(puzzle, "XMAS")
	if count != 1 {
		t.Errorf("Expected 1, got %d", count)
	}
}

func TestFindsWordDiagonalTopRightToBottomLeft(t *testing.T) {
	puzzle := [][]rune{
		{' ', ' ', ' ', 'X'},
		{' ', ' ', 'M', ' '},
		{' ', 'A', ' ', ' '},
		{'S', ' ', ' ', ' '},
	}
	count := searchWord(puzzle, "XMAS")
	if count != 1 {
		t.Errorf("Expected 1, got %d", count)
	}
}
