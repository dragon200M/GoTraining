package stringutil

/*Rune literals are just an integer value (as you've written).
They are "mapped" to their unicode codepoint. For example the rule literal 'a' is in reality the number 97.
*/
func reverseTwo(s string) string {
	r := []rune(s)

	for i, j := 0, len(r)-1; i < len(r)/2; i, j = i+1, j-1 {
		r[i], r[j] = r[j], r[i]
	}

	return string(r)

}
