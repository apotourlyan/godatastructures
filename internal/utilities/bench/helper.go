package bench

func ToKiloBytes(totalSize int, unitSize int) float64 {
	return float64(totalSize) * float64(unitSize) / 1024
}
