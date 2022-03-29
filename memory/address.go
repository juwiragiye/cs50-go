package memory

func Address() (*int) {
	n := 50
	p := &n
	return p
}