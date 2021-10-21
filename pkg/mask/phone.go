package mask

// FormatMobileStar 手机号中间4位替换为*号
func FormatMobileStar(mobile string) string {
	if len(mobile) <= 10 {
		return mobile
	}
	return mobile[:3] + "****" + mobile[7:]
}
