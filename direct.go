package direct

import (
	"fmt"

	"github.com/dgrijalva/jwt-go"
)

// Hello uses jwt
func Hello() {
	fmt.Println(jwt.GetSigningMethod(""))
}
