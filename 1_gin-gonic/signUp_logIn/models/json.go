package models

import (
	"github.com/dgrijalva/jwt-go"
)

//
type UserAccount struct {
	UserName string `json:"userName"`
	Email    string `json:"email"`
}

// royxatdan muvofaqiyatli otgan foydalanuvchi
// to'ldiragan va unga berilgan ma'lumotlarni korishligi uchun struct
type SignedUser struct {
	UserId   uint   `json:"userId"`
	UserName string `json:"userName"`
	Email    string `json:"email"`
	Password string `json:"PassWord"`
}

// royxatdan otgan foydalanuvchi sayt imkonyatlaridan foydalana olshiligi
// uchun beriladigan tokenga kerak bolga bazi malumotlar uchun struct
type Payload struct {
	UserName string
	Role     int
	jwt.StandardClaims
}

// ro'yxatdan o'tgan foydalanuvchi token olishligi uchun email accountiga
// kelgan kodni kiritishi uchun struct
type PostUserPassword struct {
	UserName string `json:"userName"`
	Password string `json:"pasword"`
}

//foydalanuvchi oziga berilgan tokkeni ko'rishi uchun struct
type Token struct {
	Token       string      `json:"token"`
	UserAccount UserAccount `json:"userAccount"`
}

// foydalanuvchi berilgan parolni o'zgartitirishi uchun struct
type NewPassword struct {
	UserName        string `json:"userName"`
	CurrentPassword string `json:"currentPassword"`
	NewPassword     string `json:"newPassword"`
}

//websocker

type Header struct {
	Name string `json:"userName"`
}
