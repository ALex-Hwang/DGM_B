package models

import (
	"gopkg.in/mgo.v2/bson"
	"time"
)

const (
	// CollectionCustomer holds the name of the customer collection
	CollectionUser = "user"
)

// User model
type User struct {
	ID               	bson.ObjectId   `json:"_id,omitempty" bson:"_id,omitempty"`
	AccountName      	string          `json:"account_name" form:"account_name" query:"account_name"`
	CreatedAt        	time.Time       `json:"created_at,omitempty" form:"created_at,omitempty" query:"created_at,omitempty"`//`json:"createdAt" bson:"createdAt"`
	Email			 	string			`json:"email" form:"email" query:"email"`
	Username		 	string			`json:"username" form:"username" query:"username"`
	Password		 	string			`json:"password" form:"password" query:"password"`
	PhoneNumber			string 			`json:"phone_number" form:"phone_number" query:"phone_number"`
	Avatar				string 			`json:"avatar" form:"avatar" query:"avatar"`
	Intro				string 			`json:"intro" form:"intro" query:"intro"`
	//Description      string          `json:"description,omitempty" bson:"description,omitempty"`
	//Address          string          `json:"address,omitempty" bson:"address,omitempty"`
	//Contact          string          `json:"contact,omitempty" bson:"contact,omitempty"`
	//Phone            string          `json:"phone,omitempty" bson:"phone,omitempty"`
	//OrganizationID   bson.ObjectId   `json:"organizationId" bson:"organizationId"`
	//ProductID        []bson.ObjectId `json:"productId,omitempty" bson:"productId,omitempty"`
	//CreatedBy        bson.ObjectId   `json:"createdBy" bson:"createdBy"`
	//OrganizationName string          `json:"organizationName" bson:"organizationName"`
	//ProductName      string          `json:"productName" bson:"productName"`
	//CreatedName      string          `json:"createdName" bson:"createdName"`

	//MemberCount      int             `json:"memberCount,omitempty" bson:"memberCount,omitempty"`
	//ProductCount     int             `json:"productCount,omitempty" bson:"productCount,omitempty"`
	//DeviceCount      int             `json:"deviceCount,omitempty" bson:"deviceCount,omitempty"`
}

// Login param
type LoginRequest struct {
	AccountName 	string //`json:"email"`
	Password 		string //`json:"password"`
}
