package models

import(
	"github.com/jinzhou/gorm"
	"github.com/te70/stunning-octo-tribble/pkg/config"
)

var db *gorm.DB

type Book struct(
	gorm.model
	Name string `gorm:""json:"name"`
	Author string `json:"author"`
	Publication string `json:"pulication"`
)

func init(){
	config.Connect()
	db = config.GetDB()
	db.AutoMigrate(&book{})
}

func (b *Book) CreateBook() *Book{
	db.NewRecord(b)
	db.Create(&b)
	return b
}

func GetAllBooks() []Book(
	var Books []Book
	db.Find(&Books)
	return Books
)

func GetBookById(Id int64) (*Book, *gorm.DB){
	var getBook book
	db:=db.Where("ID=?", Id).Find(&getBook)
	return &getBook, db
}

func DeleteBook(ID int64) Book(
	var book Book
	db.where("ID=?", ID).Delete(book)
	return book
)