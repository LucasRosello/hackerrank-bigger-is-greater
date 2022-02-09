package main

func biggerIsGreater(word string) string {

	/*
		Recorre cada grupo de letras dentro de mi palabra
	*/
	for cantidadDeLetras := 1; cantidadDeLetras < len(word); cantidadDeLetras++ {

		wordEnd := string(word[len(word)-cantidadDeLetras:])

		/*
			Compara el final de mi palabra (wordEnd) con las distintas partes de la
			palabra para tratar de encontrar una letra que sea inferior en orden alfabetico
		*/
		for i := len(word) - 1 - cantidadDeLetras; i >= 0; i-- {
			letter := word[i]
			if wordEnd > string(letter) {

				/*
					Si encuentra esa letra menor en orden alfabetico intercambia las
					posiciones de mi wordEnd con la letra encontrada
				*/
				return word[:i] + wordEnd + word[i:len(word)-cantidadDeLetras]

			}
		}
	}

	return "no answer"
}
