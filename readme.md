## Visa Transaction Controls API integrations sample codes in Go
### Installation
#### 1. Clone the repository:  
    git clone https://github.com/NjiruClinton/tectonic_cards.git
    cd tectonic_cards
#### 2. Initialize the Go module:  
    go mod tidy
### Configuration
#### 1. Place your certificates in the root directory
- caCertificate './cacert.pem'
- clientCertificate './cert.pem'
- clientKey './key.pem'

#### Create a .env file with your Visa Developer credentials:
```env
USER_ID=your_visa_user_id
PASSWORD=your_visa_password
```

## Package Functions
#### Register a Card
```go
registered, err := registercard.RegisterCard(panPrefix)
```
#### Toggle Card Status
```go
response, err := registercard.ToggleCard(documentID, status)
```
#### Perform a Card Transaction
```go
response, err := transactions.PerformCardTransaction(panPrefix, dateTimeLocal, howPresented, isDomestic)
```
#### Delete a Card
```go
response, err := registercard.DeleteCard(documentID)
```
#### Retrieve Controls
```go
response, err := transactions.RetrieveControls(panPrefix)
```
#### Create a Customer
```go
response, err := customer.CreateCustomer(panPrefix, email, firstName, lastName)
```

for development all functions and examples can be initialized from `main.go`  
project is in active maintenance and development

### Disclaimer
This project is not affiliated with, endorsed by, or certified by Visa. It is intended solely for educational and illustrative purposes to demonstrate how to work with the Visa Developer Platform APIs. Actual implementation requires adherence to Visa's legal, operational, and confidentiality obligations, including:
- Ensuring appropriate licenses, consents, and compliance with applicable regulations. 
- Consulting with legal and compliance advisors to meet specific business requirements.

### License
This repository is licensed under the [MIT License](https://github.com/NjiruClinton/tectonic_cards/blob/main/LICENSE), covering only the code developed within this repository. The Visa Transaction Controls API and any related documentation remain the intellectual property of Visa and are governed by their respective terms and conditions.