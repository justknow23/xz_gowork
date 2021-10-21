package enumstring

const (
	VALID   = "valid"
	DELETED = "deleted"
)

var EstateEnum = Enum{
	maps: map[string]string{
		VALID:   "有效",
		DELETED: "失效",
	},
}
