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
func (s *GRPCServer) EnableSubsystem(_ context.Context, r *GenericSubsystemRequest) (*GenericResponse, error) {
	err := s.SetSubsystem(r.Subsystem, true)
	if err != nil {
		return nil, err
	}
	return &GenericResponse{Status: "success",
		Data: fmt.Sprintf("subsystem %s enabled", r.Subsystem)}, nil

}

// DisableSubsystem disables a engine subsystem
func (s *GRPCServer) DisableSubsystem(_ context.Context, r *GenericSubsystemRequest) (*GenericResponse, error) {
	err := s.SetSubsystem(r.Subsystem, false)
	if err != nil {
		return nil, err
	}
	return &GenericResponse{Status: "success",
		Data: fmt.Sprintf("subsystem %s disabled", r.Subsystem)}, nil

}

func (s *GRPCServer) Login(ctx context.Context, in *LoginRequest) (*LoginReply, error) {
	//fmt.Println("Loginrequest: ", in.Username)

	// check if user exists in database
	user := models.User{}
	if err := user.GetUserByUsername(s.GetSQL(), in.Username); err != nil {
		log.Errorf(log.DatabaseLogger, "Error authenticating client, could not find user: %v", err)
		return &LoginReply{Status: "403", Token: ""}, nil
	}

	// check that supplied password matches
	if in.Password != user.Password {
		log.Errorf(log.DatabaseLogger, "Error authenticating client, invalid password supplied.")
		return &LoginReply{Status: "403", Token: ""}, nil
	}

	// create and return token
	tokenString := auth.CreateToken(in.Username)
	return &LoginReply{Status: "200", Token: tokenString}, nil

	/*if in.Username == "gavin" && in.Password == "gavin" {
		tokenString := auth.CreateToken(in.Username)
		return &LoginReply{Status: "200", Token: tokenString}, nil

	} else {
		return &LoginReply{Status: "403", Token: ""}, nil
	}*/

}
