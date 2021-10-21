package enumint

const (
	SEX_MAN     = 1
	SEX_WOMAN    = 0
)

var SexEnum = Enum{
	maps: map[int]string{
		SEX_MAN:   "man",
		SEX_WOMAN: "woman",
	},
}
