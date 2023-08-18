package db

import (
	"bookManagement/db/models"
	"errors"
)

func (dbStruct *Db) CreateUser(user *models.User) error {
	result := dbStruct.db.Create(user)
	if result.Error != nil {
		return result.Error
	}
	return nil
}

func (dbStruct *Db) FindUserByUserName(userName string) (*models.User, error) {
	var foundUser models.User
	result := dbStruct.db.Where(&models.User{UserName: userName}).First(&foundUser)
	if result.Error != nil {
		return nil, errors.New("user not found")
	}
	return &foundUser, nil
}

func (dbStruct *Db) UserNameExistence(userName string) bool {
	var foundUser models.User
	result := dbStruct.db.Where(&models.User{UserName: userName}).First(&foundUser)
	return result.Error == nil
}

func (dbStruct *Db) EmialAddressExistence(emailAddress string) bool {
	var foundUser models.User
	result := dbStruct.db.Where(&models.User{EmailAddress: emailAddress}).First(&foundUser)
	return result.Error == nil
}

func (dbStruct *Db) PhoneNumberExistence(phoneNumber string) bool {
	var foundUser models.User
	result := dbStruct.db.Where(&models.User{PhoneNumber: phoneNumber}).First(&foundUser)
	return result.Error == nil
}

func (dbStruct *Db) AddUserBook(book *models.Book, userName string) error {
	user, err := dbStruct.FindUserByUserName(userName)
	if err != nil {
		return errors.New("user not found")
	}

	user.Books = append(user.Books, *book)

	err = dbStruct.db.Updates(&user).Error
	if err != nil {
		return errors.New("unable to add the book")
	}
	return nil
}

func (dbStruct *Db) GetAllBooks() ([]models.Book, error) {
	var books []models.Book
	result := dbStruct.db.Model(&models.Book{}).Preload("Author").Find(&books)
	if result.Error != nil {
		return nil, errors.New("unable to get all books")
	}
	return books, nil
}

func (dbStruct *Db) GetUserBooks(userName string) ([]models.Book, error) {
	var books []models.Book
	user, err := dbStruct.FindUserByUserName(userName)
	if err != nil {
		return nil, errors.New("user not found")
	}
	err = dbStruct.db.Model(&user).Association("Books").Find(&books)
	if err != nil {
		return nil, errors.New("unable to get user Books")
	}
	return books, nil
}

func (dbStruct *Db) GetUserBookById(userName string, bookId uint) (*models.Book, error) {
	var user models.User
	err := dbStruct.db.Where(&models.User{UserName: userName}).Preload("Books").First(&user).Error
	if err != nil {
		return nil, errors.New("user not found")
	}

	for _, book := range user.Books {
		if book.ID == bookId {
			err := dbStruct.db.Where(&models.Book{}).Preload("Author").First(&book, "id = ?", book.ID).Error
			if err != nil {
				return nil, errors.New("book not found")
			}
			return &book, nil
		}
	}

	return nil, errors.New("book not found")
}

func (dbStruct *Db) UpdateUserBook(userName string, newBook models.Book, bookId uint) error {
	oldBook, err := dbStruct.GetUserBookById(userName, bookId)
	if err != nil {
		return err
	}
	oldBook.Name = newBook.Name
	oldBook.Author.FirstName = newBook.Author.FirstName
	oldBook.Author.LastName = newBook.Author.LastName
	oldBook.Author.Birthday = newBook.Author.Birthday
	oldBook.Author.Nationality = newBook.Author.Nationality
	oldBook.Category = newBook.Category
	oldBook.Volume = newBook.Volume
	oldBook.PublishedAt = newBook.PublishedAt
	oldBook.Summary = newBook.Summary
	oldBook.TableOfContents = newBook.TableOfContents
	oldBook.Publisher = newBook.Publisher
	if err := dbStruct.db.Updates(&oldBook).Error; err != nil {
		return errors.New("unable to update the user's book")
	}
	return nil
}

func (dbStruct *Db) DeleteUserBookById(bookId uint) error {
	result := dbStruct.db.Delete(&models.Book{}, bookId)
	if result.Error != nil {
		return errors.New("unable to delete the Book")
	}
	if result.RowsAffected == 0 {
		return errors.New("book not found")
	}
	return nil
}
