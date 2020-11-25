package models

type statusTracking struct {
	s0 string
	s1 string
	s2 string
	s3 string
}

func StatusTracking() *statusTracking {
	return &statusTracking{
		s0: "",
		s1: "",
		s2: "",
		s3: "",
	}
}
