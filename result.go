package zpay

type Result interface {
	Ok() bool
	Error() error
}
