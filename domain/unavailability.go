package booking

import (
	"context"
	"fmt"
	"time"
)

type UnavailabilityService interface {
	FindUnavailabilityByID(ctx context.Context, req FindUnavailabilityByIDRequest) FindUnavailabilityByIDResponse
	FindUnavailabilities(ctx context.Context, req FindUnavailabilitiesRequest) FindUnavailabilitiesResponse
	CreateUnavailability(ctx context.Context, req CreateUnavailabilityRequest) CreateUnavailabilityResponse
	UpdateUnavailability(ctx context.Context, req UpdateUnavailabilityRequest) UpdateUnavailabilityResponse
	DeleteUnavailability(ctx context.Context, req DeleteUnavailabilityRequest) DeleteUnavailabilityResponse
}

// FindUnavailabilityByIDRequest represents a payload used by the FindUnavailabilityByID method of a UnavailabilityService
type FindUnavailabilityByIDRequest struct {
	// insert request properties here
}

// Validate a FindUnavailabilityByID. Returns a ValidationError for each requirement that fails.
func (r FindUnavailabilityByIDRequest) Validate() []ValidationError {
	// insert validation logic here
	return nil
}

// FindUnavailabilityByIDResponse represents a response returned by the FindUnavailabilityByID method of a UnavailabilityService.
type FindUnavailabilityByIDResponse struct {
	// insert response properties here
}

// Error implements the errorer interface. Returns property Err from the response.
func (r FindUnavailabilityByIDResponse) Error() error { return r.Err }

// FindUnavailabilitiesRequest represents a payload used by the FindUnavailabilities method of a UnavailabilityService
type FindUnavailabilitiesRequest struct {
	// insert request properties here
}

// Validate a FindUnavailabilities. Returns a ValidationError for each requirement that fails.
func (r FindUnavailabilitiesRequest) Validate() []ValidationError {
	// insert validation logic here
	return nil
}

// FindUnavailabilitiesResponse represents a response returned by the FindUnavailabilities method of a UnavailabilityService.
type FindUnavailabilitiesResponse struct {
	// insert response properties here
}

// Error implements the errorer interface. Returns property Err from the response.
func (r FindUnavailabilitiesResponse) Error() error { return r.Err }

// CreateUnavailabilityRequest represents a payload used by the CreateUnavailability method of a UnavailabilityService
type CreateUnavailabilityRequest struct {
	// insert request properties here
}

// Validate a CreateUnavailability. Returns a ValidationError for each requirement that fails.
func (r CreateUnavailabilityRequest) Validate() []ValidationError {
	// insert validation logic here
	return nil
}

// CreateUnavailabilityResponse represents a response returned by the CreateUnavailability method of a UnavailabilityService.
type CreateUnavailabilityResponse struct {
	// insert response properties here
}

// Error implements the errorer interface. Returns property Err from the response.
func (r CreateUnavailabilityResponse) Error() error { return r.Err }

// UpdateUnavailabilityRequest represents a payload used by the UpdateUnavailability method of a UnavailabilityService
type UpdateUnavailabilityRequest struct {
	// insert request properties here
}

// Validate a UpdateUnavailability. Returns a ValidationError for each requirement that fails.
func (r UpdateUnavailabilityRequest) Validate() []ValidationError {
	// insert validation logic here
	return nil
}

// UpdateUnavailabilityResponse represents a response returned by the UpdateUnavailability method of a UnavailabilityService.
type UpdateUnavailabilityResponse struct {
	// insert response properties here
}

// Error implements the errorer interface. Returns property Err from the response.
func (r UpdateUnavailabilityResponse) Error() error { return r.Err }

// DeleteUnavailabilityRequest represents a payload used by the DeleteUnavailability method of a UnavailabilityService
type DeleteUnavailabilityRequest struct {
	// insert request properties here
}

// Validate a DeleteUnavailability. Returns a ValidationError for each requirement that fails.
func (r DeleteUnavailabilityRequest) Validate() []ValidationError {
	// insert validation logic here
	return nil
}

// DeleteUnavailabilityResponse represents a response returned by the DeleteUnavailability method of a UnavailabilityService.
type DeleteUnavailabilityResponse struct {
	// insert response properties here
}

// Error implements the errorer interface. Returns property Err from the response.
func (r DeleteUnavailabilityResponse) Error() error { return r.Err }

// UnavailabilityServiceMiddleware defines a middleware for UnavailabilityService
type UnavailabilityServiceMiddleware func(service UnavailabilityService) UnavailabilityService

// UnavailabilityValidationMiddleware returns a middleware for validating requests made to a UnavailabilityService
func UnavailabilityValidationMiddleware() UnavailabilityServiceMiddleware {
	return func(next UnavailabilityService) UnavailabilityService {
		return unavailabilityValidationMiddleware{next}
	}
}

// FindUnavailabilityByID validates a FindUnavailabilityByIDRequest. Returns a domain error if any requirements fail and invokes the next middleware otherwise.
func (mw unavailabilityValidationMiddlware) FindUnavailabilityByID(ctx context.Context, req FindUnavailabilityByIDRequest) FindUnavailabilityByIDResponse {
	errs := req.Validate()
	if len(errs) > 0 {
		return FindUnavailabilityByIDResponse{Err: wrapValidationErrors(errs)}
	}
	return mw.UnavailabilityService.FindUnavailabilityByID(ctx, req)
}

// FindUnavailabilities validates a FindUnavailabilitiesRequest. Returns a domain error if any requirements fail and invokes the next middleware otherwise.
func (mw unavailabilityValidationMiddlware) FindUnavailabilities(ctx context.Context, req FindUnavailabilitiesRequest) FindUnavailabilitiesResponse {
	errs := req.Validate()
	if len(errs) > 0 {
		return FindUnavailabilitiesResponse{Err: wrapValidationErrors(errs)}
	}
	return mw.UnavailabilityService.FindUnavailabilities(ctx, req)
}

// CreateUnavailability validates a CreateUnavailabilityRequest. Returns a domain error if any requirements fail and invokes the next middleware otherwise.
func (mw unavailabilityValidationMiddlware) CreateUnavailability(ctx context.Context, req CreateUnavailabilityRequest) CreateUnavailabilityResponse {
	errs := req.Validate()
	if len(errs) > 0 {
		return CreateUnavailabilityResponse{Err: wrapValidationErrors(errs)}
	}
	return mw.UnavailabilityService.CreateUnavailability(ctx, req)
}

// UpdateUnavailability validates a UpdateUnavailabilityRequest. Returns a domain error if any requirements fail and invokes the next middleware otherwise.
func (mw unavailabilityValidationMiddlware) UpdateUnavailability(ctx context.Context, req UpdateUnavailabilityRequest) UpdateUnavailabilityResponse {
	errs := req.Validate()
	if len(errs) > 0 {
		return UpdateUnavailabilityResponse{Err: wrapValidationErrors(errs)}
	}
	return mw.UnavailabilityService.UpdateUnavailability(ctx, req)
}

// DeleteUnavailability validates a DeleteUnavailabilityRequest. Returns a domain error if any requirements fail and invokes the next middleware otherwise.
func (mw unavailabilityValidationMiddlware) DeleteUnavailability(ctx context.Context, req DeleteUnavailabilityRequest) DeleteUnavailabilityResponse {
	errs := req.Validate()
	if len(errs) > 0 {
		return DeleteUnavailabilityResponse{Err: wrapValidationErrors(errs)}
	}
	return mw.UnavailabilityService.DeleteUnavailability(ctx, req)
}
