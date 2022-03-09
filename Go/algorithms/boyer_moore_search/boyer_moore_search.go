package boyer_moore_search

import (
	"Go/time_completion"
	"strings"
	"unicode/utf8"
)

//Worst-case performance		Θ(m) preprocessing + O(mn) matching
//Best-case performance			Θ(m) preprocessing + Ω(n/m) matching
//Worst-case space complexity	Θ(k)

/*
Pseudocode
precomputeShifts(){
	int lengthOfPattern = this.pattern.length()
	for index := 0; index < lengthOfPattern; index++{
		char actualCharacter = this.pattern.chartAt(intex)
		maxShift = math.max(1, lenghtOfPatter - index 01)
		this.mismatchShiftTable.put(actualCharacter, maxShift)
	}
}


for i := 0; i < lengthOfText - lengthOfPattern; i += numOfSkips{
	numOfSkips = 0;
	for j := lengthOfPattern -1; j >= 0; j--{
		if pattern.charAt(j) != text.charAt(i + j){
			if this.mismatchShiftTable.get(text.charAt(i + j) != null{
				numOfSkip = this.mismatchSiftsTable.get(text.charAt(i+j))
				break
			} else {
				numOfSkips = lengthOfPattern
				break
			}
		}

	}
	if numOfSkips == 0{
		return i
	}
}
 */

//Explanation: https://www.youtube.com/watch?v=3Ft3HMizsCk
//Original Code: https://github.com/cubicdaiya/bms/blob/master/bms.go

func BuildReadableSkipTable(s string) (badMatchMap map[string]int){
	defer time_completion.FunctionTimer(BuildReadableSkipTable)()

	if len(s) == 0{return nil}

	badMatchMap = make(map[string]int)

	for i := 0; i < len(s) -1; i++{
		//max(1, lengthOfPattern - indexOfChar -1)
		if len(s) - i - 1 < 1 {
			badMatchMap[string(s[i])] = 1
		} else {
			badMatchMap[string(s[i])] = len(s) - i - 1
		}

	}
	badMatchMap["*"] = len(s)

	return badMatchMap
}

// build skip table of needle for Boyer-Moore search.
func BuildSkipMap(searchTerm string) (badMatchMap map[rune]int) {
	defer time_completion.FunctionTimer(BuildSkipMap)()


	l := utf8.RuneCountInString(searchTerm)
	runes := []rune(searchTerm)

	badMatchMap = make(map[rune]int)

	for i := 0; i < l-1; i++ {
		j := runes[i]

		//max(1, lengthOfPattern - indexOfChar -1)
		if l - i - 1 < 1 {
			badMatchMap[j] = 1
		} else {
			badMatchMap[j] = l - i - 1
		}

		//originally just
		//badMatchMap[j] = l - i - 1
	}
	return badMatchMap
}

// search a needle in haystack and return count of needle.
// table is build by BuildSkipMap.
func SearchBySkipTable(text, searchTerm string, skipMap map[rune]int) (searchTermCount int, locations[]int) {
	defer time_completion.FunctionTimer(SearchBySkipTable)()

	locations = make([]int, 0)
	i, count := 0, 0
	textRunes := []rune(text)
	searchTermRunes := []rune(searchTerm)
	textLen := utf8.RuneCountInString(text)
	sTermLen := utf8.RuneCountInString(searchTerm)

	//search string, term length 0 return
	//search string < search term return
	if textLen == 0 || sTermLen == 0 || textLen < sTermLen {
		return 0, locations
	}

	//search string == search len return count 1
	if textLen == sTermLen && text == searchTerm {
		return 1, append(locations, 0)
	}

loop:
	//while index + search term length < text length
	for i+sTermLen <= textLen {
		//start at last char of searchTerm going backwards
		for j := sTermLen - 1; j >= 0; j-- {
			//if char @ (text location + sTerm char @ j) != sTerm char @ j, jump
			if textRunes[i+j] != searchTermRunes[j] {
				//if char @ text location not in skipMap
				if _, value := skipMap[textRunes[i+j]]; !value {
					//if 0 matches occurred
					if j == sTermLen-1 {
						//jump len of sTerm string
						i += sTermLen
					} else {
						//jump len of remaining sTerm string
						i += sTermLen - j - 1
					}
					//char @ text location in skipMap
				} else {

					i += skipMap[textRunes[i+j]]// - (sTermLen - j - 1)
					//todo? why check n here?
					//n := skipMap[textRunes[i+j]] - (sTermLen - j - 1)
					////if skip value < 0 skip 1
					//if n <= 0 {
					//	i++
					////else skip len in skipMap
					//} else {
					//	i += n
					//}
				}
				//match not found at top of if, go back to loop
				goto loop
			}
			//match found, compare next char in reverse order ie j--
		}
		//complete match found, jump to
		//if jump to end of searchTerm finds character in skipMap
		if _, value := skipMap[textRunes[i+sTermLen-1]]; value {
			//todo ? why no check like n here?
			//jump to skipMap value
			i += skipMap[textRunes[i+sTermLen-1]]
			//not in skipMap, just jump to end of searchTerm
		} else {
			//jump to end of searchTerm
			i += sTermLen
		}

		count++
		//current location is end of searchTerm String, set to index of beginning of searchTerm in text
		locations = append(locations, i-(sTermLen-1))
	}

	return count, locations
}

// Boyer Moore Search
func Search(text, search string) (searchTermCount int, locations []int) {
	defer time_completion.FunctionTimer(Search)()
	text = strings.ToLower(text)
	search = strings.ToLower(search)
	table := BuildSkipMap(search)
	return SearchBySkipTable(text, search, table)
}
