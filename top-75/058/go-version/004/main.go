package main

func countSubstrings(s string) int {
    /*
    Walk the string
    For each caracter, see the ones in the sides if they are palindromes
    If so, keep going, if not, drop it.

    1 - 3 - 5
    2 - 4 - 6 (try it with the -1 substring) and do the same logic
    */

    palindromeCount := func(i, j int) int {
        count := 0

        for i >= 0 && j < len(s) && s[i] == s[j]{
            count++
            i--
            j++
        }

        return count
    }

    res := 0
    for i := 0; i < len(s); i++ {
        res += palindromeCount(i, i)
    }

    for i := 1; i < len(s); i++ {
        res += palindromeCount(i-1, i)
    }

    return res
}