package presentation

import (
	"bookManagement/logic"
	"log"
	"net/http"

	"github.com/gin-gonic/gin"
)

func (sm *ServerManager) SignUp(c *gin.Context) {
	var newUser logic.User
	err := c.BindJSON(&newUser)
	if err != nil {
		log.Println(err.Error())
		c.Status(http.StatusBadRequest)
		return
	}

	if err = logic.AddUser(newUser, sm.Db); err != nil {
		log.Println(err)
		c.IndentedJSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	c.IndentedJSON(http.StatusOK, gin.H{"message": "Signed Up successfully"})
}

func (sm *ServerManager) SignInByCredintials(c *gin.Context) {
	var signInCre struct {
		UserName string `json:"userName"`
		Password string `json:"password"`
	}
	if err := c.BindJSON(&signInCre); err != nil {
		log.Println(err.Error())
		c.Status(http.StatusBadRequest)
		return
	}
	token, err := logic.UserSignIn(signInCre.UserName, signInCre.Password, sm.Db, sm.JwtManager)
	if err != nil {
		log.Println(err)
		c.IndentedJSON(http.StatusUnauthorized, gin.H{"error": err.Error()})
		return
	}
	c.Header("Authorization", token)
	c.IndentedJSON(http.StatusOK, gin.H{"message": "Signed in successfully"})
}

func (sm *ServerManager) SignInByToken(c *gin.Context) {
	token := c.GetHeader("Authorization")
	if token == "" {
		log.Println("Token is not provided")
		c.IndentedJSON(http.StatusForbidden, gin.H{"error": "Token is not provided"})
		return
	}
	user, err := logic.UserSignInByToken(token, sm.Db, sm.JwtManager)
	if err != nil {
		log.Println(err.Error())
		c.IndentedJSON(http.StatusForbidden, gin.H{"error": err.Error()})
		return
	}
	c.IndentedJSON(http.StatusOK, gin.H{
		"userName":     user.UserName,
		"emailAddress": user.EmailAddress,
		"phoneNumber":  user.PhoneNumber,
		"firstName":    user.FirstName,
		"lastName":     user.LastName,
	})
}

func (sm *ServerManager) CreateBook(c *gin.Context) {
	token := c.GetHeader("Authorization")
	if token == "" {
		log.Println("Token is not provided")
		c.IndentedJSON(http.StatusForbidden, gin.H{"error": "Token is not provided"})
		return
	}
	var book logic.Book
	if err := c.BindJSON(&book); err != nil {
		log.Println(err.Error())
		c.Status(http.StatusBadRequest)
		return
	}

	if err := logic.AddUserBook(token, &book, sm.Db, sm.JwtManager); err != nil {
		c.IndentedJSON(http.StatusForbidden, gin.H{"error": err.Error()})
		return
	}
	c.IndentedJSON(http.StatusOK, gin.H{"message": "The book added successfully"})
}

func (sm *ServerManager) GetAllBooks(c *gin.Context) {
	token := c.GetHeader("Authorization")
	if token == "" {
		log.Println("Token is not provided")
		c.IndentedJSON(http.StatusForbidden, gin.H{"error": "Token is not provided"})
		return
	}
	books, err := logic.GetAllBooks(token, sm.Db, sm.JwtManager)
	if err != nil {
		c.IndentedJSON(http.StatusForbidden, gin.H{"error": err.Error()})
		return
	}
	c.IndentedJSON(http.StatusOK, gin.H{"books": books})
}

func (sm *ServerManager) GetBook(c *gin.Context) {
	token := c.GetHeader("Authorization")
	if token == "" {
		log.Println("Token is not provided")
		c.IndentedJSON(http.StatusForbidden, gin.H{"error": "Token is not provided"})
		return
	}

	id := c.Param("id")
	if id == "" {
		c.IndentedJSON(http.StatusBadRequest, gin.H{"error": "Id of the book is not provided"})
		return
	}

	book, err := logic.GetUserBookById(token, id, sm.Db, sm.JwtManager)
	if err != nil {
		c.IndentedJSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	c.IndentedJSON(http.StatusOK, book)
}

func (sm *ServerManager) UpdateBook(c *gin.Context) {
	token := c.GetHeader("Authorization")
	if token == "" {
		log.Println("Token is not provided")
		c.IndentedJSON(http.StatusForbidden, gin.H{"error": "Token is not provided"})
		return
	}

	id := c.Param("id")
	if id == "" {
		c.IndentedJSON(http.StatusBadRequest, gin.H{"error": "Id of the book is not provided"})
		return
	}

	var newBook logic.Book
	if err := c.BindJSON(&newBook); err != nil {
		log.Println(err.Error())
		c.Status(http.StatusBadRequest)
		return
	}

	err := logic.UpdateUserBook(token, id, newBook, sm.Db, sm.JwtManager)
	if err != nil {
		c.IndentedJSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	c.IndentedJSON(http.StatusOK, gin.H{"message": "The Book info Updated successfully"})
}

func (sm *ServerManager) DeleteBook(c *gin.Context) {
	token := c.GetHeader("Authorization")
	if token == "" {
		log.Println("Token is not provided")
		c.IndentedJSON(http.StatusForbidden, gin.H{"error": "Token is not provided"})
		return
	}

	id := c.Param("id")
	if id == "" {
		c.IndentedJSON(http.StatusBadRequest, gin.H{"error": "Id of the book is not provided"})
		return
	}

	err := logic.DeleteUserBookById(token, id, sm.Db, sm.JwtManager)
	if err != nil {
		c.IndentedJSON(http.StatusForbidden, gin.H{"error": err.Error()})
		return
	}

	c.IndentedJSON(http.StatusOK, gin.H{"message": "The Book deleted successfully"})
}
