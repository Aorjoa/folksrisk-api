package model

import (
	"github.com/jinzhu/gorm"
)

// Personal stores name of risk person
type Personal struct {
	gorm.Model
	Identity     string        `json:"identity"`
	Name         string        `json:"name"`
	Email        string        `json:"email"`
	Point        Point         `json:"point"`
	Image        string        `json:"image"`
	BankAccounts []BankAccount `json:"bank_accounts"`
	PhoneNumbers []PhoneNumber `json:"phone_numbers"`
	Evidences    []Evidence    `json:"evidences"`
	Files        []File        `json:"files"`
}

//BankAccount store bank account associate with person
type BankAccount struct {
	PersonalID  uint
	BankAccount string
}

//PhoneNumber store phone number associate with person
type PhoneNumber struct {
	PersonalID  uint
	PhoneNumber string
}

//Point store point of person
type Point struct {
	PersonalID        uint `json:"personalid"`
	Risk              int8
	PersinalIDVerify  bool
	BankAccountVerify bool
	PhoneNumberVerify bool
	EmailActivated    bool
	Sponsed           bool
}

// Evidence stores description of user
type Evidence struct {
	gorm.Model
	PersonalID  uint   `json:"personalid"`
	Description string `json:"description"`
}

// File stores files of evidence
type File struct {
	gorm.Model
	PersonalID    uint   `json:"personalid"`
	FileName      string `json:"filename"`
	FilePath      string `json:"filepath"`
	FileExtension string `json:"fileextension"`
}
