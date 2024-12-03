package hash

func djb2(s string) int64 {
	var hash int64 = 5381

	for _, c := range s {
		hash = ((hash << 5) + hash) + int64(c)
	}

	return hash
}
