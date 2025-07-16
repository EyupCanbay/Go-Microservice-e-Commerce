package types

import (
	"time"

	"go.mongodb.org/mongo-driver/bson/primitive"
)

// Address, bir kullanıcının coğrafi adres bilgilerini içerir.
// Address contains the geographical address information for a user.
type Address struct {
	ID        string `bson:"id" json:"id"`               // Adres için benzersiz kimlik / Unique identifier for the address
	Type      string `bson:"type" json:"type"`           // Adres Tipi (örn: 'shipping', 'billing') / Address Type (e.g., 'shipping', 'billing')
	Street    string `bson:"street" json:"street"`       // Cadde ve Sokak Bilgisi / Street Information
	Apartment string `bson:"apartment" json:"apartment"` // Daire, Kapı Numarası / Apartment, Suite, etc.
	District  string `bson:"district" json:"district"`   // İlçe / District
	City      string `bson:"city" json:"city"`           // Şehir / City
	State     string `bson:"state" json:"state"`         // Eyalet veya İl / State or Province
	ZipCode   string `bson:"zip_code" json:"zip_code"`   // Posta Kodu / Zip Code
	Country   string `bson:"country" json:"country"`     // Ülke / Country
	IsActive  bool   `json:"is_active"`                  // Adresin aktif olup olmadığını belirtir / Indicates if the address is active
}

// Customer, veritabanındaki müşteri kaydının tamamını temsil eder.
// Customer represents the entire customer record in the database.
type Customer struct {
	Id                primitive.ObjectID `bson:"_id" json:"id"`                                  // Benzersiz Müşteri Kimliği / Unique Customer Identifier
	FirstName         string             `bson:"first_name" json:"first_name"`                   // Müşterinin Adı / Customer's First Name
	LastName          string             `bson:"last_name" json:"last_name"`                     // Müşterinin Soyadı / Customer's Last Name
	Email             string             `bson:"email" json:"email"`                             // Müşterinin E-posta Adresi / Customer's Email Address
	PasswordHash      string             `bson:"password_hash" json:"-"`                         // Parolanın şifrelenmiş hali / Hashed password
	Phone             string             `bson:"phone" json:"phone"`                             // Müşterinin Telefon Numarası / Customer's Phone Number
	Addresses         []Address          `bson:"addresses" json:"addresses"`                     // Müşterinin kaydedilmiş adres listesi / A list of the customer's saved addresses
	IsActive          bool               `bson:"is_active" json:"is_active"`                     // Hesabın aktif olup olmadığını belirtir / Indicates if the account is active
	DateOfBirth       time.Time          `bson:"date_of_birth" json:"date_of_birth"`             // Müşterinin Doğum Tarihi / Customer's Date of Birth
	Gender            string             `bson:"gender" json:"gender"`                           // Cinsiyet / Gender
	Notes             []string           `bson:"notes" json:"notes"`                             // Müşteriyle ilgili notlar dizisi / Array of notes about the customer
	Preferences       map[string]string  `bson:"preferences" json:"preferences"`                 // Kullanıcı tercihleri (örn: tema, dil) / User preferences (e.g., theme, language)
	ProfilePictureURL string             `bson:"profile_picture_url" json:"profile_picture_url"` // Profil Resmi URL'si / Profile Picture URL
	CreatedAt         time.Time          `bson:"created_at" json:"created_at"`                   // Kayıt Oluşturulma Zamanı / Record Creation Time
	UpdatedAt         time.Time          `bson:"updated_at" json:"updated_at"`                   // Kayıt Güncellenme Zamanı / Record Last Update Time
}
