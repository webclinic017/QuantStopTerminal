package engine

import (
	"fmt"
	"github.com/quantstop/quantstopterminal/internal/log"
	"reflect"
	"sync"
)

// ServiceRegistry provides a useful pattern for managing services.
// It allows for ease of dependency management and ensures services
// dependent on others use the same references in memory.
type ServiceRegistry struct {
	services     map[reflect.Type]iSubsystem // map of types to iSubsystem.
	serviceTypes []reflect.Type              // keep an ordered slice of registered service types.
}

// NewServiceRegistry starts a registry instance for convenience
func NewServiceRegistry() *ServiceRegistry {
	return &ServiceRegistry{
		services: make(map[reflect.Type]iSubsystem),
	}
}

// RegisterService appends a service constructor function to the service
// registry.
func (s *ServiceRegistry) RegisterService(service iSubsystem) error {
	kind := reflect.TypeOf(service)
	if _, exists := s.services[kind]; exists {
		return fmt.Errorf("subsystem already exists: %v", kind)
	}
	s.services[kind] = service
	s.serviceTypes = append(s.serviceTypes, kind)
	return nil
}

// StartAll initialized each service in order of registration.
func (s *ServiceRegistry) StartAll(wg *sync.WaitGroup) {
	log.Debugf(log.SubsystemLogger, "Starting %d subsystems: %v", len(s.serviceTypes), s.serviceTypes)

	// Loop through all services
	for _, kind := range s.serviceTypes {
		if s.services[kind].isEnabled() && s.services[kind].isInitialized() {
			// Make sure service is enabled, and initialized, then try starting
			log.Debugf(log.SubsystemLogger, "Starting subsystem type %v", kind)
			//if err := StartSubsystem(s.services[kind], wg); err != nil {
			if err := s.services[kind].start(wg); err != nil {
				log.Errorf(log.SubsystemLogger, "Unable to start subsystem %v : %v", kind, err)
			}
		}

	}
}

// Start takes in a command string, and wait group, attempts to find the service by its name, and then start
func (s *ServiceRegistry) Start(name string, wg *sync.WaitGroup) error {

	found := false

	for _, kind := range s.serviceTypes {
		if s.services[kind].getName() == name {
			found = true
			if !s.services[kind].isEnabled() {
				return ErrSubsystemNotEnabled
			}
			if !s.services[kind].isInitialized() {
				return ErrSubsystemNotInitialized
			}
			//if err := StartSubsystem(s.services[kind], wg); err != nil {
			if err := s.services[kind].start(wg); err != nil {
				return err
			}

			return nil
		}
	}

	if !found {
		return fmt.Errorf("unknown service: %s", name)
	}

	return nil
}

// StopAll ends every service in reverse order of registration, logging a
// panic if any of them fail to stop.
func (s *ServiceRegistry) StopAll() {
	for i := len(s.serviceTypes) - 1; i >= 0; i-- {
		kind := s.serviceTypes[i]
		service := s.services[kind]
		//if err := StopSubsystem(service); err != nil {
		if err := service.stop(); err != nil {
			log.Errorf(log.SubsystemLogger, "Could not stop the following service: %v, %v", kind, err)
		}
	}
}

// FetchService takes in a struct pointer and sets the value of that pointer
// to a service currently stored in the service registry. This ensures the input argument is
// set to the right pointer that refers to the originally registered service.
func (s *ServiceRegistry) FetchService(service interface{}) error {
	if reflect.TypeOf(service).Kind() != reflect.Ptr {
		return fmt.Errorf("input must be of pointer type, received value type instead: %T", service)
	}
	element := reflect.ValueOf(service).Elem()
	if running, ok := s.services[element.Type()]; ok {
		element.Set(reflect.ValueOf(running))
		return nil
	}
	return fmt.Errorf("unknown service: %T", service)
}
