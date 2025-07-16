package types

import "time"

// CustomerRequestModel, istemciden gelen müşteri oluşturma/güncelleme verilerini tanımlar.
// CustomerRequestModel defines the data for creating/updating a customer from the client.
type CustomerRequestModel struct {
	FirstName   string            `json:"first_name"`    // Müşterinin Adı / Customer's First Name
	LastName    string            `json:"last_name"`     // Müşterinin Soyadı / Customer's Last Name
	Email       string            `json:"email"`         // Müşterinin E-posta Adresi / Customer's Email Address
	Password    string            `json:"password"`      // Müşterinin Parolası (düz metin) / Customer's Password (plain text)
	Phone       string            `json:"phone"`         // Müşterinin Telefon Numarası / Customer's Phone Number
	Addresses   []Address         `json:"addresses"`     // Kullanıcının adres listesi / The user's list of addresses
	DateOfBirth time.Time         `json:"date_of_birth"` // Müşterinin Doğum Tarihi / Customer's Date of Birth
	Gender      string            `json:"gender"`        // Cinsiyet / Gender
	Preferences map[string]string `json:"preferences"`   // Kullanıcı tercihleri (örn: dil, tema) / User preferences (e.g., language, theme)
}

// CustomerResponseModel, API'nin istemciye döndüreceği müşteri verilerini tanımlar.
// CustomerResponseModel defines the customer data that the API will return to the client.
type CustomerResponseModel struct {
	Id                       string    `json:"id"`                          // Benzersiz Müşteri Kimliği / Unique Customer Identifier
	FirstName                string    `json:"first_name"`                  // Müşterinin Adı / Customer's First Name
	LastName                 string    `json:"last_name"`                   // Müşterinin Soyadı / Customer's Last Name
	Email                    string    `json:"email"`                       // Müşterinin E-posta Adresi / Customer's Email Address
	Phone                    string    `json:"phone"`                       // Müşterinin Telefon Numarası / Customer's Phone Number
	Addresses                []Address `json:"addresses"`                   // Müşterinin kayıtlı adresleri / The customer's registered addresses
	DefaultBillingAddressID  string    `json:"default_billing_address_id"`  // Varsayılan Fatura Adresi Kimliği / Default Billing Address ID
	DefaultShippingAddressID string    `json:"default_shipping_address_id"` // Varsayılan Teslimat Adresi Kimliği / Default Shipping Address ID
	DateOfBirth              time.Time `json:"date_of_birth"`               // Müşterinin Doğum Tarihi / Customer's Date of Birth
	Gender                   string    `json:"gender"`                      // Cinsiyet / Gender
	Notes                    []string  `json:"notes"`                       // Müşteriyle ilgili notlar / Notes about the customer
	ProfilePictureURL        string    `json:"profile_picture_url"`         // Profil Resmi URL'si / Profile Picture URL
}
