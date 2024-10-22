package code

// Result struct to hold the duplicate elements and the number of duplicates
type Result[T any] struct {
	DuplicatedElements []T
	NumberOfDuplicates uint
}

// FindDuplicates function to find the duplicated elements in  a slice
func FindDuplicates[T comparable](elements []T) Result[T] {
	// Map to store the count of eachelement
	elementCount := make(map[T]int)

	// Slice to store the unique duplicated elemts
	var duplicates []T

	// Count occurences of each element
	for _, element := range elements {
		elementCount[element]++
		// If the element occurs exactly tiwice, add it to the duplicates slice
		if elementCount[element] == 2 {
			duplicates = append(duplicates, element)
		}
	}

	// Return the result with the number of duplicates and the unique duplicated elements
	return Result[T]{
		DuplicatedElements: duplicates,
		NumberOfDuplicates: uint(len(duplicates)),
	}
}
