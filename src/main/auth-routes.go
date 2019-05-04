package main

import (
	"errors"
	"fmt"
	"github.com/dgrijalva/jwt-go"
	"github.com/gin-gonic/gin"
	"golang.org/x/crypto/bcrypt"
	"log"
	"net/http"
	"strconv"
	"strings"
	"time"
)

func CheckJWTToken(token string, secret string) (*jwt.Token, error) {
	return jwt.Parse(token, func(token *jwt.Token) (interface{}, error) { // Verify token authenticity
		if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
			return nil, fmt.Errorf("unexpected signing method: %v", token.Header["alg"])
		}
		return []byte(secret), nil
	})
}

// Get claims of JWT token

func GetJWTClaims(token string, secret string) (jwt.MapClaims, error) {
	tok, err := CheckJWTToken(token, secret)
	if err != nil {
		return nil, err
	}
	if !tok.Valid {
		return nil, errors.New("invalid token")
	}
	claims, ok := tok.Claims.(jwt.MapClaims)
	if !ok {
		return nil, errors.New("not ok")
	}
	return claims, nil
}

func AuthRequired() gin.HandlerFunc {
	return func(c *gin.Context) {
		sess, err := c.Cookie("GOSESSID")
		if err != nil {
			c.Redirect(301, "/login")
			c.Abort()
			return
		}

		claims, err := GetJWTClaims(sess, config.Secret)
		if err != nil {
			c.Redirect(301, "/login")
			c.Abort()
			return
		}

		exp, err := strconv.ParseInt(claims["expires"].(string), 10, 64)
		if err != nil {
			log.Println(err.Error())
			c.HTML(http.StatusInternalServerError, "500.html", gin.H{})
			c.Abort()
			return // invalid token
		}

		if exp > time.Now().Unix() {
			c.Redirect(301, "/login")
			c.Abort()
			return
		}

		var user *IUser

		for _, userf := range db.Users {
			if userf.UserName == claims["id"].(string) {
				user = userf
			}
		}

		if user == nil {
			c.Redirect(301, "/login")
			c.Abort()
			return
		}

		c.Set("user", user)
		c.Next()
	}
}

func AuthRoutes() {
	router.GET("/login", func(c *gin.Context) {
		c.HTML(http.StatusOK, "login.html", gin.H{
			"incorrectPassword": false,
			"noUser":            false,
		})
	})

	router.POST("/login", func(c *gin.Context) {
		username := strings.TrimSpace(c.PostForm("username"))
		password := c.PostForm("password")

		var userF *IUser

		for _, user := range db.Users {
			if user.UserName == username {
				userF = user
			}
		}

		if userF == nil {
			c.HTML(http.StatusOK, "login.html", gin.H{
				"incorrectPassword": false,
				"noUser":            true,
			})
			return
		}

		err := bcrypt.CompareHashAndPassword([]byte(userF.Password), []byte(password))

		if err != nil {
			c.HTML(http.StatusOK, "login.html", gin.H{
				"incorrectPassword": true,
				"noUser":            false,
			})
			return
		}

		token := jwt.New(jwt.SigningMethodHS256)
		claims := make(jwt.MapClaims)
		claims["id"] = userF.UserName
		claims["expires"] = time.Now().Add(time.Second * time.Duration(86400)).Unix()
		token.Claims = claims
		tokenString, err := token.SignedString([]byte(config.Secret))

		if err != nil {
			c.HTML(http.StatusInternalServerError, "500.html", gin.H{})
			return
		}

		c.SetCookie("GOSESSID", tokenString, 86400, "/", config.Domain, true, false)

		// successful
		c.HTML(http.StatusOK, "index.html", gin.H{})
	})

	router.GET("/register", func(c *gin.Context) {
		c.HTML(http.StatusOK, "register.html", gin.H{
			"createSuccess": false,
		})
	})

	router.POST("/register", func(c *gin.Context) {
		username := c.PostForm("username")
		password := c.PostForm("password")

		pass, err := bcrypt.GenerateFromPassword([]byte(password), 11)
		if err != nil {
			log.Println(err)
		}

		user := IUser{
			DisplayName:    username,
			UserName:       username,
			Password:       string(pass),
			TimeRegistered: time.Now().Unix(),
			Posts:          []string{},
			Icon:           "",
			Description:    "",
		}

		db.Users = append(db.Users, &user)

		go StoreDB()

		c.HTML(http.StatusOK, "login.html", gin.H{
			"createSuccess": true,
		})
	})
}
