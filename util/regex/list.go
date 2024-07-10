package regex

// MatchStringInList will return true if item is contained in list. If
// exactMatch is set to false, list may contain regex to be matched.

// Check if NameSpace is Equal with direct match
// If not Equal, and Exact Match is false, and does not comply with RFC 1123, then check if it matches the pattern
// In this context, if it isn't RFC 1123, it is probably regex, otherwise un-intentional substring matching might occur.
func MatchStringInList(list []string, item string, exactMatch bool) bool {
	for _, ll := range list {
		if item == ll || (!exactMatch && Match(ll, item) && !Match("(?=\\A[-a-z0-9]{1,63}\\Z)\\A[a-z0-9]+(-[a-z0-9]+)*\\Z", ll)) {
			return true
		}
	}
	return false
}
