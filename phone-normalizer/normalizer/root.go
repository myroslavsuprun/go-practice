package normalizer

func Normalize(phoneNumber string) string {
	var result []int32

	for _, char := range phoneNumber {
		if char >= '0' && char <= '9' {
			result = append(result, char)
		}
	}

	return string(result)
}
