package global

var (
	// withKingpin 使用原生kingpin.v2
	withKingpin = false
)

// Kingpin ...
func Kingpin() bool {
	return withKingpin
}

// WitchKingpin ...
func WitchKingpin(v bool) {
	withKingpin = v
}
