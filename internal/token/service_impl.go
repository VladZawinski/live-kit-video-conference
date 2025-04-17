package token

type tokenServiceImpl struct {
	apiKey    string
	secretKey string
}

func NewTokenService(apiKey, secretKey string) TokenService {
	return &tokenServiceImpl{
		apiKey:    apiKey,
		secretKey: secretKey,
	}
}

func (s *tokenServiceImpl) GenerateJoinToken(identity string, roomName string, isPublisher bool) (string, error) {
	return "", nil
}
