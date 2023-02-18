package quora

import (
	"time"

	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/mysql"
)

// Question struct to hold question data
type Question struct {
	ID        uint `gorm:"primary_key"`
	Title     string
	Body      string
	UserID    uint
	Timestamp time.Time
	Answers   []Answer
	Votes     []Vote
	Comments  []Comment
}

// Answer struct to hold answer data
type Answer struct {
	ID         uint `gorm:"primary_key"`
	Body       string
	UserID     uint
	QuestionID uint
	Timestamp  time.Time
	Votes      []Vote
	Comments   []Comment
}

// User struct to hold user data
type User struct {
	ID       uint `gorm:"primary_key"`
	Username string
	Email    string
	Password string
}

// Comment struct to hold comment data
type Comment struct {
	ID         uint `gorm:"primary_key"`
	Body       string
	UserID     uint
	QuestionID uint
	AnswerID   uint
	Timestamp  time.Time
}

// Topic struct to hold topic data
type Topic struct {
	ID          uint `gorm:"primary_key"`
	Name        string
	Description string
	Questions   []Question
}
type Vote struct {
	ID         uint `gorm:"primary_key"`
	UserID     uint
	QuestionID uint
	AnswerID   uint
	Value      int
}

// Quora struct to hold the database connection
type Quora struct {
	DB *gorm.DB
}

// NewQuora returns a new Quora struct with a database connection
func NewQuora(connString string) (*Quora, error) {
	db, err := gorm.Open("mysql", connString)
	if err != nil {
		return nil, err

	}
	return &Quora{DB: db}, nil

}

// AskQuestion adds a new question to the platform
func (q *Quora) AskQuestion(userID uint, title, body string) (*Question, error) {
	question := &Question{
		Title:  title,
		Body:   body,
		UserID: userID,
	}
	err := q.DB.Create(question).Error
	if err != nil {
		return nil, err

	}
	return question, nil

}

// AnswerQuestion adds a new answer to a question
func (q *Quora) AnswerQuestion(userID, questionID uint, body string) (*Answer, error) {
	answer := &Answer{
		Body:       body,
		UserID:     userID,
		QuestionID: questionID,
	}
	err := q.DB.Create(answer).Error
	if err != nil {
		return nil, err

	}
	return answer, nil

}

// UpvoteQuestion upvotes a question
func (q *Quora) UpvoteQuestion(userID, questionID uint) error {
	vote := &Vote{
		UserID:     userID,
		QuestionID: questionID,
		Value:      1,
	}
	return q.DB.Create(vote).Error

}

// DownvoteQuestion downvotes a question
func (q *Quora) DownvoteQuestion(userID, questionID uint) error {
	vote := &Vote{
		UserID:     userID,
		QuestionID: questionID,
		Value:      -1,
	}
	return q.DB.Create(vote).Error

}

// CommentOnQuestion adds a comment on a question
func (q *Quora) CommentOnQuestion(userID, questionID uint, body string) (*Comment, error) {
	comment := &Comment{
		Body:       body,
		UserID:     userID,
		QuestionID: questionID,
	}
	err := q.DB.Create(comment).Error
	if err != nil {
		return nil, err

	}
	return comment, nil

}

// SearchQuestions searches for questions by keyword
func (q *Quora) SearchQuestions(keyword string) ([]Question, error) {
	var questions []Question
	err := q.DB.Where("title LIKE ? OR body LIKE ?", "%"+keyword+"%", "%"+keyword+"%").Find(&questions).Error
	if err != nil {
		return nil, err

	}
	return questions, nil

}

/*A similar service in Go could have the following features:

Users can ask and answer questions on a wide variety of topics
Users can upvote or downvote questions and answers
Users can comment on questions and answers
Users can search for questions and answers by keywords
Users can follow specific topics or users to receive updates when new questions or answers are posted
Here's an example of how you might structure the main components of the service:

Question struct: Represents a question that is asked on the platform. It could have fields such as ID, Title, Body, UserID, and Timestamp.
Answer struct: Represents an answer to a question. It could have fields such as ID, Body, UserID, QuestionID, and Timestamp.
User struct: Represents a user on the platform. It could have fields such as ID, Username, Email, and Password.
Comment struct: Represents a comment on a question or answer. It could have fields such as ID, Body, UserID, QuestionID, AnswerID, and Timestamp.
Topic struct: Represents a topic on the platform. It could have fields such as ID, Name, Description, and Questions (a slice of Question structs).
Vote struct: Represents a vote on a question or answer. It could have fields such as ID, UserID, QuestionID, AnswerID, and Value (1 for an upvote, -1 for a downvote).
A database is also necessary to store all of the data.

These are just examples of how you might structure the main components of the service, depending on the use case you might need to add more fields or change the current fields to suit your needs.

You could also use external libraries such as Gorm(ORM) to interact with the database and add functionalities like pagination, filtering and ordering.

You can also use external libraries such as JWT (JSON Web Token) or OAuth2 to handle authentication and authorization.

It's a complex service and the implementation details can be complex, but with the right knowledge and resources, it's definitely possible to create a similar service using golang.*/
