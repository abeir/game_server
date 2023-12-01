package conf

type IConfigFinder interface {
	Find() (string, bool)
}
