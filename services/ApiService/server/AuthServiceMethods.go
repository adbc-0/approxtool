package server

import (
	"apisrv/protos/auth"
	"apisrv/protos/user"
	"context"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

// Login Not Secured
func (s *Server) Login(ctx context.Context, request *auth.LoginRequest) (*auth.LoginResponse, error) {
	verifiedUser, err := s.UserClient.VerifyPassword(context.Background(), &user.VerifyPasswordRequest{
		Email:    request.Email,
		Password: request.Password,
	})
	if err != nil {
		return nil, err
	}

	resp, err := s.AuthClient.Login(context.Background(), &auth.InternalLoginRequest{
		UserId:   verifiedUser.UserId,
		UserRole: verifiedUser.UserRole,
	})
	if err != nil {
		return nil, err
	}

	return resp, nil
}

// Logout Not Secured
func (s *Server) Logout(ctx context.Context, request *auth.LogoutRequest) (*auth.LogoutResponse, error) {
	resp, err := s.AuthClient.Logout(context.Background(), request)
	if err != nil {
		return nil, err
	}
	return resp, nil
}

// RefreshToken Not Secured
func (s Server) RefreshToken(ctx context.Context, request *auth.RefreshRequest) (*auth.RefreshResponse, error) {
	tokens, err := s.AuthClient.RefreshToken(context.Background(), request)
	if err != nil {
		return nil, status.Error(codes.Unauthenticated, "error refreshing tokens")
	}

	return &auth.RefreshResponse{
		AccessToken:  tokens.AccessToken,
		RefreshToken: tokens.RefreshToken,
	}, err
}