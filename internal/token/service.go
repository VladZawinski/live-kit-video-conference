package token

type TokenService interface {
	GenerateJoinToken(identity string, roomName string, isPublisher bool) (string, error)
}
