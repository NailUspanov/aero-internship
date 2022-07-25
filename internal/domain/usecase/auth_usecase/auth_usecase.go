package auth_usecase

import (
	"aero-internship/internal/adapters/postgres"
	"aero-internship/internal/domain/entity/tokens"
	users2 "aero-internship/internal/domain/entity/users"
	"aero-internship/pkg/config"
	"context"
	"errors"
	"fmt"
	jwt "github.com/golang-jwt/jwt/v4"
	"github.com/google/uuid"
	"golang.org/x/crypto/bcrypt"
	"google.golang.org/grpc"
	"google.golang.org/grpc/metadata"
	"log"
	"strconv"
	"strings"
	"time"
)

type contextKey int

const (
	clientIDKey contextKey = iota
)

type AuthService struct {
	postgres.Repository
	secret []byte
	method jwt.SigningMethod
	ttl    time.Duration
	cfg    *config.Config
}

func NewAuthService(cfg *config.Config, repo postgres.Repository) *AuthService {
	jwt_ttl, _ := time.ParseDuration(cfg.GetJWTttl())
	return &AuthService{
		secret:     []byte(cfg.GetJWTSecret()),
		method:     jwt.SigningMethodHS256,
		ttl:        jwt_ttl,
		cfg:        cfg,
		Repository: repo,
	}
}

func (tm *AuthService) GenerateTokens(tknDTO *tokens.TokenDTO) (*Tokens, error) {

	tkn := jwt.NewWithClaims(jwt.SigningMethodHS256, jwt.MapClaims{
		"sub":  tknDTO.UserId,
		"role": tknDTO.IsAdmin,
		"exp":  fmt.Sprintf("%d", time.Now().Add(tm.ttl).Unix()),
	})

	token, err := tkn.SignedString(tm.secret)
	if err != nil {
		return nil, err
	}

	refreshToken := uuid.New().String()

	err = tm.Repository.MakeNewSession(tm.cfg, tknDTO.UserId, refreshToken)
	if err != nil {
		return nil, err
	}

	return &Tokens{
		AccessToken:  token,
		RefreshToken: refreshToken,
	}, nil
}

func (tm *AuthService) GenerateTokensFromUserDTO(userDTO *users2.UserDTO) (*Tokens, error) {

	log.Printf("ok, we have UserDTO...\n")

	tokenDTO, err := tm.Repository.GetTokenDTOFromUserDTO(tm.cfg, userDTO)
	if err != nil {
		return nil, err
	}

	log.Printf("now we have TokenDTO( userId: %s, role: %s)\n", tokenDTO.UserId, tokenDTO.IsAdmin)

	log.Printf("let's generate tokens from this DTO!\n")

	tokens, err := tm.GenerateTokens(tokenDTO)
	if err != nil {
		return nil, err
	}
	return tokens, nil
}

func (tm *AuthService) ParseToken(tkn string) (TokenDTO, error) {
	token, err := jwt.Parse(tkn, func(token *jwt.Token) (interface{}, error) {
		if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
			return nil, fmt.Errorf("tnexpected signing method: %v", token.Header["alg"])
		}
		claims, ok := token.Claims.(jwt.MapClaims)
		exp, _ := strconv.Atoi(claims["exp"].(string))
		if !ok || int64(exp) < time.Now().Unix() {
			return nil, errors.New("token is expired")
		}
		return tm.secret, nil
	})
	if err.Error() == "token is expired" {
		return TokenDTO{}, err
	}
	claims, _ := token.Claims.(jwt.MapClaims)
	if claims["sub"] == nil || claims["role"] == nil {
		return TokenDTO{}, errors.New("claims is not valid")
	}
	return TokenDTO{
		UserId:  strconv.Itoa(int(claims["sub"].(float64))),
		IsAdmin: claims["role"].(bool),
	}, nil
}

func (tm *AuthService) RegisterUser(userDTO *users2.UserDTO) (*Tokens, error) {

	hashedPassword, err := hashPassword(userDTO.Password)
	if err != nil {
		return nil, err
	}

	log.Printf("Let's make a txn!\n")

	err = tm.Users.MakeRegistrationTxn(tm.cfg, users2.UserDTO{
		Name:     userDTO.Name,
		Email:    userDTO.Email,
		Password: hashedPassword,
		IsAdmin:  userDTO.IsAdmin,
	})
	if err != nil {
		return nil, fmt.Errorf("error while making txn: %v", err)
	}

	log.Printf("txn commited!\n")

	log.Printf("let's generate some tokens...\n")

	tokens, err := tm.GenerateTokensFromUserDTO(userDTO)
	if err != nil {
		return nil, fmt.Errorf("error while creating tokens: %v", err)
	}

	log.Printf("nice, tokens %s:%s", tokens.AccessToken, tokens.RefreshToken)
	return tokens, nil
}

func (tm *AuthService) SignIn(signInDTO *SignInDTO) (*Tokens, error) {

	user, err := tm.Users.GetUserDTObyEmail(tm.cfg, signInDTO.Email)
	if err != nil {
		return nil, err
	}

	userDTO := &users2.UserDTO{
		Name:     user.Name,
		Email:    user.Email,
		Password: user.Password,
		IsAdmin:  user.IsAdmin,
	}

	log.Printf("nice, now we have userDTO for %s\n", user.Name)

	ok := doPasswordsMatch(userDTO.Password, signInDTO.Password)

	log.Printf("password check: %v\n", ok)

	if !ok {
		return nil, err
	}
	tokens, err := tm.GenerateTokensFromUserDTO(userDTO)
	if err != nil {
		return nil, err
	}

	log.Printf("we have tokens now\n")

	return tokens, nil
}

func (tm *AuthService) UnaryInterceptor(ctx context.Context, req interface{}, info *grpc.UnaryServerInfo, handler grpc.UnaryHandler) (interface{}, error) {

	var clientID string

	if md, ok := metadata.FromIncomingContext(ctx); ok {
		if md.Get("authorization") != nil {
			token, err := tm.ParseToken(md.Get("authorization")[0])
			if err != nil {
				return nil, err
			}
			clientID = token.UserId
		} else {
			id, err := tm.authenticateClient(ctx, tm.cfg)
			if err != nil {
				return nil, err
			}
			clientID = id
		}
	}

	ctx = context.WithValue(ctx, clientIDKey, clientID)

	return handler(ctx, req)
}

func (tm *AuthService) authenticateClient(ctx context.Context, cfg *config.Config) (string, error) {
	if md, ok := metadata.FromIncomingContext(ctx); ok {

		clientLogin := strings.Join(md["login"], "")

		user, err := tm.Users.GetUserDTObyEmail(cfg, clientLogin)
		if err != nil {
			return "", err
		}

		log.Printf("authenticated client: %s", clientLogin)

		return strconv.Itoa(user.Id), nil
	}
	return "", fmt.Errorf("missing credentials")
}

func hashPassword(pas string) (string, error) {
	hashedPasswordBytes, err := bcrypt.GenerateFromPassword([]byte(pas), bcrypt.MinCost)
	if err != nil {
		return "", err
	}
	return string(hashedPasswordBytes), nil
}

func doPasswordsMatch(hashedPassword, currPassword string) bool {
	err := bcrypt.CompareHashAndPassword(
		[]byte(hashedPassword), []byte(currPassword))
	return err == nil
}
