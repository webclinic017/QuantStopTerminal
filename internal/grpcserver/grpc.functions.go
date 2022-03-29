package grpcserver

import (
	"context"
	"fmt"
	"github.com/quantstop/quantstopterminal/internal/database/models"
	"github.com/quantstop/quantstopterminal/internal/grpcserver/auth"
	"github.com/quantstop/quantstopterminal/internal/log"
)

// GetInfo returns info about the current session
func (s *GRPCServer) GetInfo(_ context.Context, _ *GetInfoRequest) (*GetInfoResponse, error) {

	return &GetInfoResponse{
		Uptime:          s.GetUptime(),
		Version:         s.GetVersion(),
		SubsystemStatus: s.GetSubsystemsStatus(),
	}, nil

}

// GetSubsystems returns a list of subsystems and their status
func (s *GRPCServer) GetSubsystems(_ context.Context, _ *GetSubsystemsRequest) (*GetSusbsytemsResponse, error) {
	return &GetSusbsytemsResponse{SubsystemsStatus: s.GetSubsystemsStatus()}, nil
}

// EnableSubsystem enables a engine subsystem
func (s *GRPCServer) EnableSubsystem(ctx context.Context, r *GenericSubsystemRequest) (*GenericResponse, error) {

	// auth endpoint
	if _, err := auth.CheckAuth(ctx); err != nil {
		return &GenericResponse{Status: "404", Data: ""}, nil
	}

	// set subsystem enabled
	err := s.SetSubsystem(r.Subsystem, true)
	if err != nil {
		return nil, err
	}
	return &GenericResponse{Status: "success",
		Data: fmt.Sprintf("subsystem %s enabled", r.Subsystem)}, nil

}

// DisableSubsystem disables a engine subsystem
func (s *GRPCServer) DisableSubsystem(ctx context.Context, r *GenericSubsystemRequest) (*GenericResponse, error) {

	// auth endpoint
	if _, err := auth.CheckAuth(ctx); err != nil {
		return &GenericResponse{Status: "404", Data: ""}, nil
	}

	// set subsystem disabled
	err := s.SetSubsystem(r.Subsystem, false)
	if err != nil {
		return nil, err
	}
	return &GenericResponse{Status: "success",
		Data: fmt.Sprintf("subsystem %s disabled", r.Subsystem)}, nil

}

// Login authenticates user to endpoints using jwt
func (s *GRPCServer) Login(_ context.Context, in *LoginRequest) (*LoginReply, error) {

	// check if user exists in database
	user := models.User{}
	if err := user.GetUserByUsername(s.GetSQL(), in.Username); err != nil {
		log.Errorf(log.GRPClog, "Error authenticating client, could not find user: %v", err)
		return &LoginReply{Status: "403", Token: ""}, nil
	}

	// check that supplied password matches
	if in.Password != user.Password {
		log.Errorf(log.GRPClog, "Error authenticating client, invalid password supplied.")
		return &LoginReply{Status: "403", Token: ""}, nil
	}

	// create and return token
	tokenString, err := auth.CreateToken(in.Username)
	if err != nil {
		log.Errorf(log.GRPClog, "Error authenticating client, cannot create token.")
		return &LoginReply{Status: "403", Token: ""}, nil
	}
	return &LoginReply{Status: "200", Token: tokenString}, nil

}
