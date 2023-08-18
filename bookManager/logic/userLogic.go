package logic

import (
	"bookManagement/db"
	"bookManagement/db/models"
	"errors"
	"log"
	"strconv"
	"strings"

	"golang.org/x/crypto/bcrypt"
)

func hashPassword(password string) (string, error) {
	bytes, err := bcrypt.GenerateFromPassword([]byte(password), 14)
	return string(bytes), err
}

func checkPasswordHash(password, hash string) bool {
	err := bcrypt.CompareHashAndPassword([]byte(hash), []byte(password))
	return err == nil
}

func AddUser(newUser User, dbStruct *db.Db) error {
	if !UserNameValidation(newUser.UserName) || !PasswordValidation(newUser.Password) || !EmailValidation(newUser.EmailAddress) || !PhoneNumberValidation(newUser.PhoneNumber) || !NameValidation(newUser.FirstName, newUser.LastName) {
		return errors.New("credintials are invalid")
	}

	str := ""
	if dbStruct.EmialAddressExistence(newUser.EmailAddress) {
		str += "emailAddress "
	}
	if dbStruct.UserNameExistence(newUser.UserName) {
		str += "userName "
	}

	if dbStruct.PhoneNumberExistence(newUser.PhoneNumber) {
		str += "phoneNumber "
	}
	if len(str) != 0 {
		str += "already exist(s)"
		return errors.New(str)
	}

	hashedPassword, err := hashPassword(newUser.Password)
	if err != nil {
		log.Println("Couldn't hash the password")
		return errors.New("unable to signUp")
	}

	dbNewUser := models.User{
		UserName:     newUser.UserName,
		Password:     hashedPassword,
		EmailAddress: newUser.EmailAddress,
		PhoneNumber:  newUser.PhoneNumber,
		Gender:       newUser.Gender,
		FirstName:    newUser.FirstName,
		LastName:     newUser.LastName,
	}

	err = dbStruct.CreateUser(&dbNewUser)
	if err != nil {
		return errors.New("the User was not added successfully")
	}
	return nil
}

func UserSignIn(userName, password string, dbStruct *db.Db, jwtManager *JwtManager) (string, error) {
	if !UserNameValidation(userName) || !PasswordValidation(password) {
		return "", errors.New("credintials are invalid")
	}

	foundUser, err := dbStruct.FindUserByUserName(userName)
	if err != nil {
		return "", errors.New("user not found")
	}

	if !checkPasswordHash(password, foundUser.Password) {
		return "", errors.New("password is incorrect")
	}

	token, err := jwtManager.createNewToken(userName)
	if err != nil {
		return "", err
	}
	return token, nil
}

func UserSignInByToken(token string, dbStruct *db.Db, jwtManager *JwtManager) (*User, error) {
	userName, err := jwtManager.verifyToken(token, dbStruct)
	if err != nil {
		return nil, err
	}
	temp, err := dbStruct.FindUserByUserName(userName)
	if err != nil {
		return nil, err
	}
	return &User{
		UserName:     temp.UserName,
		Password:     temp.Password,
		EmailAddress: temp.EmailAddress,
		PhoneNumber:  temp.PhoneNumber,
		Gender:       temp.Gender,
		FirstName:    temp.FirstName,
		LastName:     temp.LastName,
	}, nil
}

func AddUserBook(token string, book *Book, dbStruct *db.Db, jwtManager *JwtManager) error {
	userName, err := jwtManager.verifyToken(token, dbStruct)
	if err != nil {
		return err
	}

	if !validateBookInfo(book) {
		return errors.New("book information is not enough")
	}

	userBooks, err := dbStruct.GetUserBooks(userName)
	if err != nil {
		return err
	}

	for _, userBook := range userBooks {
		if userBook.Name == book.Name {
			return errors.New("the user has this book already")
		}
	}

	err = dbStruct.AddUserBook(&models.Book{
		Name: book.Name,
		Author: models.Author{
			FirstName:   book.Author.FirstName,
			LastName:    book.Author.LastName,
			Birthday:    book.Author.Birthday,
			Nationality: book.Author.Nationality,
		},
		Category:        book.Category,
		Volume:          book.Volume,
		PublishedAt:     book.PublishedAt,
		Summary:         book.Summary,
		TableOfContents: strings.Join(book.TableOfContents, "#"),
		Publisher:       book.Publisher,
	}, userName)
	if err != nil {
		return err
	}
	return nil
}

type tempBook map[string]any

func GetAllBooks(token string, dbStruct *db.Db, jwtManager *JwtManager) ([]tempBook, error) {
	_, err := jwtManager.verifyToken(token, dbStruct)
	if err != nil {
		return nil, err
	}

	dbBooks, err := dbStruct.GetAllBooks()
	if err != nil {
		return nil, err
	}

	books := make([]tempBook, 0, len(dbBooks))
	for _, dbBook := range dbBooks {
		books = append(books, tempBook{
			"name":        dbBook.Name,
			"author":      dbBook.Author.FirstName + " " + dbBook.Author.LastName,
			"category":    dbBook.Category,
			"volume":      dbBook.Volume,
			"publishedAt": dbBook.PublishedAt,
			"summary":     dbBook.Summary,
			"publisher":   dbBook.Publisher,
		})
	}
	return books, nil

}

func GetUserBookById(token string, id string, dbStruct *db.Db, jwtManager *JwtManager) (*Book, error) {
	userName, err := jwtManager.verifyToken(token, dbStruct)
	if err != nil {
		return nil, err
	}
	idInt, err := strconv.Atoi(id)
	if err != nil {
		return nil, errors.New("id is not in valid format")
	}

	book, err := dbStruct.GetUserBookById(userName, uint(idInt))
	if err != nil {
		return nil, err
	}
	return &Book{
		Name: book.Name,
		Author: author{
			FirstName:   book.Author.FirstName,
			LastName:    book.Author.LastName,
			Birthday:    book.Author.Birthday,
			Nationality: book.Author.Nationality,
		},
		Category:        book.Category,
		Volume:          book.Volume,
		PublishedAt:     book.PublishedAt,
		Summary:         book.Summary,
		TableOfContents: strings.Split(book.TableOfContents, "#"),
		Publisher:       book.Publisher,
	}, nil
}

func completeNewBook(newBook *Book, oldBook models.Book) {
	if newBook.Name == "" {
		newBook.Name = oldBook.Name
	}
	if newBook.Author.FirstName == "" {
		newBook.Author.FirstName = oldBook.Author.FirstName
	}
	if newBook.Author.LastName == "" {
		newBook.Author.LastName = oldBook.Author.LastName
	}
	if newBook.Author.Birthday.IsZero() {
		newBook.Author.Birthday = oldBook.Author.Birthday
	}
	if newBook.Author.Nationality == "" {
		newBook.Author.Nationality = oldBook.Author.Nationality
	}
	if newBook.Category == "" {
		newBook.Category = oldBook.Category
	}
	if newBook.Volume == 0 {
		newBook.Volume = oldBook.Volume
	}
	if newBook.PublishedAt.IsZero() {
		newBook.PublishedAt = oldBook.PublishedAt
	}
	if newBook.Summary == "" {
		newBook.Summary = oldBook.Summary
	}
	if newBook.TableOfContents == nil || len(newBook.TableOfContents) == 0 {
		newBook.TableOfContents = strings.Split(oldBook.TableOfContents, "#")
	}
	if newBook.Publisher == "" {
		newBook.Publisher = oldBook.Publisher
	}
}

func UpdateUserBook(token string, id string, newBook Book, dbStruct *db.Db, jwtManager *JwtManager) error {
	userName, err := jwtManager.verifyToken(token, dbStruct)
	if err != nil {
		return err
	}
	idInt, err := strconv.Atoi(id)
	if err != nil {
		return errors.New("id is not in valid format")
	}

	oldBook, err := dbStruct.GetUserBookById(userName, uint(idInt))
	if err != nil {
		return err
	}

	// Filling the fields wich weren't provided and are in their default values
	completeNewBook(&newBook, *oldBook)

	err = dbStruct.UpdateUserBook(userName, models.Book{
		Name: newBook.Name,
		Author: models.Author{
			FirstName:   newBook.Author.FirstName,
			LastName:    newBook.Author.LastName,
			Birthday:    newBook.Author.Birthday,
			Nationality: newBook.Author.Nationality,
		},
		Category:        newBook.Category,
		Volume:          newBook.Volume,
		PublishedAt:     newBook.PublishedAt,
		Summary:         newBook.Summary,
		TableOfContents: strings.Join(newBook.TableOfContents, "#"),
		Publisher:       newBook.Publisher,
	}, uint(idInt))
	if err != nil {
		return err
	}

	return nil
}

func DeleteUserBookById(token string, id string, dbStruct *db.Db, jwtManager *JwtManager) error {
	_, err := jwtManager.verifyToken(token, dbStruct)
	if err != nil {
		return err
	}
	idInt, err := strconv.Atoi(id)
	if err != nil {
		return errors.New("id is not in valid format")
	}

	err = dbStruct.DeleteUserBookById(uint(idInt))
	if err != nil {
		return err
	}

	return nil
}
