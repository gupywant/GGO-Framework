// models/user.go
package models

import (
	"GGO/utils"

    "errors"
    "gorm.io/gorm"
    "crypto/md5"
    "encoding/hex"
	"github.com/google/uuid"
)

type User struct {
    ID       uuid.UUID  `gorm:"type:uuid;default:uuid_generate_v4();primaryKey"`
    Username string `gorm:"unique;not null"`
    Password string `gorm:"not null"`
}

func (User) TableName() string {
	return "user" 
}

// FindUserByUsername fetches a user from the database by username
func FindUserByUsername(db *gorm.DB, username string) (*User, error) {
    var user User
	utils.ODserver(user)
    if err := db.Where("username = ?", username).First(&user).Error; err != nil {
        return nil, err
    }

	// utils.DDserver(w, user)

    return &user, nil
}



func (u *User) Authenticate(password string) error {
    // Hash the provided password using MD5
    hasher := md5.New()
    hasher.Write([]byte(password))
    hashedPassword := hex.EncodeToString(hasher.Sum(nil))

    // Compare hashed password with stored password
    if u.Password != hashedPassword {
        return errors.New("invalid password")
    }
    return nil
}