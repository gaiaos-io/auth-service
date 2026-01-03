package session

type RefreshToken struct {
	hash [32]byte
}

func NewRefreshToken(hash [32]byte) *RefreshToken {
	return &RefreshToken{hash: hash}
}

func (token RefreshToken) Bytes() []byte {
	bytes := make([]byte, 32)
	copy(bytes, token.hash[:])
	return bytes
}

func (token RefreshToken) Equal(other RefreshToken) bool {
	return token.hash == other.hash
}
