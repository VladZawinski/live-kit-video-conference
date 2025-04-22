package sdk

type TokenSdkService interface {
	GenerateJoinToken(identity string, roomName string, isPublisher bool) (string, error)
}

type tokenServiceImpl struct {
	apiKey    string
	secretKey string
}

func NewTokenSdkService(apiKey, secretKey string) TokenSdkService {
	return &tokenServiceImpl{
		apiKey:    apiKey,
		secretKey: secretKey,
	}
}

func (s *tokenServiceImpl) GenerateJoinToken(identity string, roomName string, isPublisher bool) (string, error) {
	return "", nil
}
