package infrastructure

import (
	jwt_pkg "github.com/bookmark-project-learn/bookmark-common-libs/pkg/jwt"
	"github.com/bookmark-project-learn/user-service/constant"
)

func CreateJwtProvider() (jwt_pkg.JwtGenerator, jwt_pkg.JwtValidator) {
	jwtGenerator := jwt_pkg.NewJWTGenerator(constant.PrivateKeyPath)
	jwtValidator := jwt_pkg.NewJWTValidator(constant.PublicKeyPath)
	return jwtGenerator, jwtValidator
}
