package password

type Service interface {
	GenerateHash(passwd string) (string, error)
	CheckHash(plain, hash string) bool
}
