package booking

import (
	"context"
	"fmt"
	"time"
)

type BookingService interface {
	FindBookingByID(ctx context.Context, req FindBookingByIDRequest) FindBookingByIDResponse
	FindBookings(ctx context.Context, req FindBookingsRequest) FindBookingsResponse
	CreateBooking(ctx context.Context, req CreateBookingRequest) CreateBookingResponse
	UpdateBooking(ctx context.Context, req UpdateBookingRequest) UpdateBookingResponse
	DeleteBooking(ctx context.Context, req DeleteBookingRequest) DeleteBookingResponse
}

// FindBookingByIDRequest represents a payload used by the FindBookingByID method of a BookingService
type FindBookingByIDRequest struct {
	// insert request properties here
}

// Validate a FindBookingByID. Returns a ValidationError for each requirement that fails.
func (r FindBookingByIDRequest) Validate() []ValidationError {
	// insert validation logic here
	return nil
}

// FindBookingByIDResponse represents a response returned by the FindBookingByID method of a BookingService.
type FindBookingByIDResponse struct {
	// insert response properties here
}

// Error implements the errorer interface. Returns property Err from the response.
func (r FindBookingByIDResponse) Error() error { return r.Err }

// FindBookingsRequest represents a payload used by the FindBookings method of a BookingService
type FindBookingsRequest struct {
	// insert request properties here
}

// Validate a FindBookings. Returns a ValidationError for each requirement that fails.
func (r FindBookingsRequest) Validate() []ValidationError {
	// insert validation logic here
	return nil
}

// FindBookingsResponse represents a response returned by the FindBookings method of a BookingService.
type FindBookingsResponse struct {
	// insert response properties here
}

// Error implements the errorer interface. Returns property Err from the response.
func (r FindBookingsResponse) Error() error { return r.Err }

// CreateBookingRequest represents a payload used by the CreateBooking method of a BookingService
type CreateBookingRequest struct {
	// insert request properties here
}

// Validate a CreateBooking. Returns a ValidationError for each requirement that fails.
func (r CreateBookingRequest) Validate() []ValidationError {
	// insert validation logic here
	return nil
}

// CreateBookingResponse represents a response returned by the CreateBooking method of a BookingService.
type CreateBookingResponse struct {
	// insert response properties here
}

// Error implements the errorer interface. Returns property Err from the response.
func (r CreateBookingResponse) Error() error { return r.Err }

// UpdateBookingRequest represents a payload used by the UpdateBooking method of a BookingService
type UpdateBookingRequest struct {
	// insert request properties here
}

// Validate a UpdateBooking. Returns a ValidationError for each requirement that fails.
func (r UpdateBookingRequest) Validate() []ValidationError {
	// insert validation logic here
	return nil
}

// UpdateBookingResponse represents a response returned by the UpdateBooking method of a BookingService.
type UpdateBookingResponse struct {
	// insert response properties here
}

// Error implements the errorer interface. Returns property Err from the response.
func (r UpdateBookingResponse) Error() error { return r.Err }

// DeleteBookingRequest represents a payload used by the DeleteBooking method of a BookingService
type DeleteBookingRequest struct {
	// insert request properties here
}

// Validate a DeleteBooking. Returns a ValidationError for each requirement that fails.
func (r DeleteBookingRequest) Validate() []ValidationError {
	// insert validation logic here
	return nil
}

// DeleteBookingResponse represents a response returned by the DeleteBooking method of a BookingService.
type DeleteBookingResponse struct {
	// insert response properties here
}

// Error implements the errorer interface. Returns property Err from the response.
func (r DeleteBookingResponse) Error() error { return r.Err }

// BookingServiceMiddleware defines a middleware for BookingService
type BookingServiceMiddleware func(service BookingService) BookingService

// BookingValidationMiddleware returns a middleware for validating requests made to a BookingService
func BookingValidationMiddleware() BookingServiceMiddleware {
	return func(next BookingService) BookingService {
		return bookingValidationMiddleware{next}
	}
}

// FindBookingByID validates a FindBookingByIDRequest. Returns a domain error if any requirements fail and invokes the next middleware otherwise.
func (mw bookingValidationMiddlware) FindBookingByID(ctx context.Context, req FindBookingByIDRequest) FindBookingByIDResponse {
	errs := req.Validate()
	if len(errs) > 0 {
		return FindBookingByIDResponse{Err: wrapValidationErrors(errs)}
	}
	return mw.BookingService.FindBookingByID(ctx, req)
}

// FindBookings validates a FindBookingsRequest. Returns a domain error if any requirements fail and invokes the next middleware otherwise.
func (mw bookingValidationMiddlware) FindBookings(ctx context.Context, req FindBookingsRequest) FindBookingsResponse {
	errs := req.Validate()
	if len(errs) > 0 {
		return FindBookingsResponse{Err: wrapValidationErrors(errs)}
	}
	return mw.BookingService.FindBookings(ctx, req)
}

// CreateBooking validates a CreateBookingRequest. Returns a domain error if any requirements fail and invokes the next middleware otherwise.
func (mw bookingValidationMiddlware) CreateBooking(ctx context.Context, req CreateBookingRequest) CreateBookingResponse {
	errs := req.Validate()
	if len(errs) > 0 {
		return CreateBookingResponse{Err: wrapValidationErrors(errs)}
	}
	return mw.BookingService.CreateBooking(ctx, req)
}

// UpdateBooking validates a UpdateBookingRequest. Returns a domain error if any requirements fail and invokes the next middleware otherwise.
func (mw bookingValidationMiddlware) UpdateBooking(ctx context.Context, req UpdateBookingRequest) UpdateBookingResponse {
	errs := req.Validate()
	if len(errs) > 0 {
		return UpdateBookingResponse{Err: wrapValidationErrors(errs)}
	}
	return mw.BookingService.UpdateBooking(ctx, req)
}

// DeleteBooking validates a DeleteBookingRequest. Returns a domain error if any requirements fail and invokes the next middleware otherwise.
func (mw bookingValidationMiddlware) DeleteBooking(ctx context.Context, req DeleteBookingRequest) DeleteBookingResponse {
	errs := req.Validate()
	if len(errs) > 0 {
		return DeleteBookingResponse{Err: wrapValidationErrors(errs)}
	}
	return mw.BookingService.DeleteBooking(ctx, req)
}
