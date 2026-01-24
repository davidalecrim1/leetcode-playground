package main

type Solution struct{}

// Join the strings with a default splitter that allows us
// to decode but not interfer on the string text.
// e.g. \n

func (s *Solution) Encode(strs []string) string {
	encodedBytes := make([]byte, 0, 200*len(strs)+len(strs)*1)
	for i := range strs {
		encodedString := []byte(strs[i])
		encodedString = append(encodedString, '\n')
		encodedBytes = append(encodedBytes, encodedString...)
	}

	return string(encodedBytes)

}

func (s *Solution) Decode(encoded string) []string {
	decodedStrings := make([]string, 0, len(encoded))

	left := 0
	encodedBytes := []byte(encoded)
	for right := 0; right < len(encoded); right++ {
		if encodedBytes[right] != '\n' {
			continue
		}

		decodedStrings = append(decodedStrings, string(encodedBytes[left:right]))
		left = right + 1 // remove the \n
	}

	return decodedStrings
}
