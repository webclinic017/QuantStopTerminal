package grpcserver

import (
	"context"
	"fmt"
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
