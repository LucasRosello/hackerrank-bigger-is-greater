package main

func biggerIsGood(word string) string {

	wordEnd := word[len(word)-2:]

	for i := 3; i <= len(word); i++ {
		if len(word) >= 4 && wordEnd > word[len(word)-i:len(word)-2] {
			word = word[:len(word)-i] + wordEnd + word[len(word)-i:len(word)-2]
			return word
		}
	}

	return "no answer"
}
