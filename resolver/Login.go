package resolver

import (
	"context"
	helper "eaciit/gopher/helper"
	model "eaciit/gopher/model"
	types "eaciit/gopher/types"
	"github.com/dgrijalva/jwt-go"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
	"gopkg.in/go-playground/validator.v9"
	"os"
	"time"
)

// Login dkfjdsf
func (r *RootResolver) Login(args types.RegisterUserArgs) (*UserResolver, error) {
	// Load your secret key from a safe place and reuse it across multiple
	// Seal/Open calls. (Obviously don't use this example key for anything
	// real.) If you want to convert a passphrase to a key, use a suitable
	// package like bcrypt or scrypt.
	// When decoded the key should be 16 bytes (AES-128) or 32 (AES-256).
	validate := validator.New()
	user := &types.User{
		Email:    args.Email,
		Password: args.Password,
	}

	errs := validate.Struct(user)
	if errs != nil {
		return &UserResolver{user}, &types.CommonError{Code: "Error", Message: errs.Error()}
	}

	opts := options.FindOne().SetSort(bson.D{{"Email", 1}})
	var result bson.M
	usercol := model.User{}
	err := usercol.Collection().FindOne(context.TODO(), bson.D{{"Email", user.Email}}, opts).Decode(&result)
	if err != nil {
		// ErrNoDocuments means that the filter did not match any documents in the collection
		if err == mongo.ErrNoDocuments {
			return &UserResolver{user}, &types.CommonError{Code: "Error", Message: "Email not found"}
		}
	}
	password := result["Password"].(string)
	decrypt, err := helper.Decrypt(password)
	if errs != nil {
		return &UserResolver{user}, &types.CommonError{Code: "Error", Message: "Decrypt Error"}
	}
	if decrypt != args.Password {
		return &UserResolver{user}, &types.CommonError{Code: "Error", Message: "Password invalid"}
	}

	mySigningKey := []byte(os.Getenv("SIGNING_KEY"))

	// Create the Claims
	ExpiresAt := time.Now().Add(time.Minute * 5).Unix()
	claims := types.TokenClaim{
		result["Email"].(string),
		jwt.StandardClaims{
			ExpiresAt: ExpiresAt,
			Issuer:    "test",
		},
	}

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	ss, err := token.SignedString(mySigningKey)

	if err != nil {
		return &UserResolver{user}, &types.CommonError{Code: "Error", Message: "Error creating token"}
	}
	user.Token = ss
	user.ExpiresAt = string(ExpiresAt)
	return &UserResolver{user}, nil
}
