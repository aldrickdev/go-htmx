package routes

import (
	"fmt"
	"net/http"

	"github.com/aldrickdev/go-htmx/cmd/app/utils"
	"github.com/gin-gonic/gin"
	"github.com/golang-jwt/jwt/v5"
)

func Configure(c *gin.Context) {
	ghpat, found := c.GetPostForm("ghpat")
	if !found {
		c.HTML(200, "setup", nil)
		return
	}
	// Abstract
	claims := &jwt.RegisteredClaims{
		Issuer:  "User",
		Subject: ghpat,
	}
	token := jwt.NewWithClaims(jwt.SigningMethodHS512, claims)
	signedToken, err := token.SignedString(utils.Env.JWT_SIGNING_KEY)
	if err != nil {
		fmt.Println("Error signing the token: ", err)
		c.HTML(http.StatusUnprocessableEntity, "setup", nil)
		return
	}
	//

	c.SetCookie("github_issues_app", signedToken, 0, "/", "", true, true)
	c.Header("HX-Redirect", "/")
}
