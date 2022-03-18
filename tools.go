package ggol

// Check if two CellLiveStatusMaps are equal.
func AreCellLiveStatusMapsEqual(a CellLiveStatusMap, b CellLiveStatusMap) bool {
	for i := 0; i < len(a); i++ {
		for j := 0; j < len(a[i]); j++ {
			if a[i][j] != b[i][j] {
				return false
			}
		}
	}
	return true
}
