package version

const (
	Major = "0"
	Minor = "1"
	Patch = "0"
)

// Version 返回完整的版本号
func Version() string {
	return Major + "." + Minor + "." + Patch
}
