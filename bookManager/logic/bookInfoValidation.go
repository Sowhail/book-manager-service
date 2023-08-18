package logic

func validateBookInfo(book *Book) bool {
	switch true {
	case book.Name == "":
		return false
	case book.Author.FirstName == "":
		return false
	case book.Author.LastName == "":
		return false
	case book.Author.Birthday.IsZero():
		return false
	case book.Author.Nationality == "":
		return false
	case book.Category == "":
		return false
	case book.Volume == 0:
		return false
	case book.PublishedAt.IsZero():
		return false
	case book.Summary == "":
		return false
	case book.TableOfContents == nil || len(book.TableOfContents) == 0:
		return false
	case book.Publisher == "":
		return false
	default:
		return true
	}
}
