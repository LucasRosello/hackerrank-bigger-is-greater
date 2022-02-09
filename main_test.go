package main

import (
	"bufio"
	"log"
	"os"
	"testing"
)

func TestNoAnswer(t *testing.T) {

	var word, expected, result string

	// Test 1
	word = "bb"
	expected = "no answer"
	result = biggerIsGreater(word)

	if expected != result {
		t.Errorf("Esperado: %v, obtenido %v", expected, result)
	}

	// Test 2
	word = "dcba"
	expected = "no answer"
	result = biggerIsGreater(word)

	if expected != result {
		t.Errorf("Esperado: %v, obtenido %v", expected, result)
	}

	// Test 3
	word = "dcbb"
	expected = "no answer"
	result = biggerIsGreater(word)

	if expected != result {
		t.Errorf("Esperado: %v, obtenido %v", expected, result)
	}
}

func TestGreaterWord(t *testing.T) {

	var word, expected, result string

	// Test 1
	word = "ab"
	expected = "ba"
	result = biggerIsGreater(word)

	if expected != result {
		t.Errorf("Esperado: %v, obtenido %v", expected, result)
	}

	// Test 2
	word = "hefg"
	expected = "hegf"
	result = biggerIsGreater(word)

	if expected != result {
		t.Errorf("Esperado: %v, obtenido %v", expected, result)
	}

	// Test 3
	word = "dhck"
	expected = "dhkc"
	result = biggerIsGreater(word)

	if expected != result {
		t.Errorf("Esperado: %v, obtenido %v", expected, result)
	}

	// Test 4
	word = "lmno"
	expected = "lmon"
	result = biggerIsGreater(word)

	if expected != result {
		t.Errorf("Esperado: %v, obtenido %v", expected, result)
	}

	// Test 5
	word = "lmno"
	expected = "lmon"
	result = biggerIsGreater(word)

	if expected != result {
		t.Errorf("Esperado: %v, obtenido %v", expected, result)
	}

	// Test 6
	word = "abcd"
	expected = "abdc"
	result = biggerIsGreater(word)

	if expected != result {
		t.Errorf("Esperado: %v, obtenido %v", expected, result)
	}

	// Test 7
	word = "fedcbabcd"
	expected = "fedcbabdc"
	result = biggerIsGreater(word)

	if expected != result {
		t.Errorf("Esperado: %v, obtenido %v", expected, result)
	}

	// Test 8
	word = "dkhc"
	expected = "hcdk"
	result = biggerIsGreater(word)

	if expected != result {
		t.Errorf("Esperado: %v, obtenido %v", expected, result)
	}

	// Test 9
	word = "abdc"
	expected = "acbd"
	result = biggerIsGreater(word)

	if expected != result {
		t.Errorf("Esperado: %v, obtenido %v", expected, result)
	}

	// Test 10
	word = "imllmmcslslkyoegymoa"
	expected = "imllmmcslslkyoegyoam"
	result = biggerIsGreater(word)

	if expected != result {
		t.Errorf("Esperado: %v, obtenido %v", expected, result)
	}
}

func TestHackerRank(t *testing.T) {

	word := readFile("input01")
	expected := readFile("output01")

	for i, w := range word {
		result := biggerIsGreater(w)
		if expected[i] != result {
			t.Errorf("Esperado: %v, recibido %v", expected[i], result)
			return
		}
	}
}

func readFile(fileName string) (list []string) {
	file, err := os.Open("./rawTests/" + fileName + ".txt")
	if err != nil {
		log.Fatal(err)
	}

	defer file.Close()

	scanner := bufio.NewScanner(file)
	// optionally, resize scanner's capacity for lines over 64K, see next example

	for scanner.Scan() {
		list = append(list, scanner.Text())

	}

	if err := scanner.Err(); err != nil {
		log.Fatal(err)
	}

	return
}
