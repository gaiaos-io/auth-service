package ports

type TokenHasher interface {
	Hash(token []byte) []byte
	Verify(hash, token []byte) (bool, error)
}
