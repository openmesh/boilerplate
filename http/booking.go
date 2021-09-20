package http

import (
	"context"
	"net/http"

	"github.com/openmesh/booking"

	httptransport "github.com/go-kit/kit/transport/http"
	"github.com/go-kit/transport"
	"github.com/gorilla/mux"
	"github.com/openmesh/booking/endpoint"
)

func (s *Server) registerBookingRoutes(r *mux.Router) {
	e := endpoint.MakeBookingEndpoints(s.BookingService)

	options := []httptransport.ServerOption{
		httptransport.ServerErrorHandler(transport.NewLogErrorHandler(s.logger)),
		httptransport.ServerErrorEncoder(encodeError),
	}
	
	r.Methods("insert_method_here").Path("insert_path_here").Handler(httptransport.NewServer(
		e.FindBookingByIDEndpoint,
		decodeFindBookingByIDRequest,
		encodeResponse,
		options...,
	))

	r.Methods("insert_method_here").Path("insert_path_here").Handler(httptransport.NewServer(
		e.FindBookingsEndpoint,
		decodeFindBookingsRequest,
		encodeResponse,
		options...,
	))

	r.Methods("insert_method_here").Path("insert_path_here").Handler(httptransport.NewServer(
		e.CreateBookingEndpoint,
		decodeCreateBookingRequest,
		encodeResponse,
		options...,
	))

	r.Methods("insert_method_here").Path("insert_path_here").Handler(httptransport.NewServer(
		e.UpdateBookingEndpoint,
		decodeUpdateBookingRequest,
		encodeResponse,
		options...,
	))

	r.Methods("insert_method_here").Path("insert_path_here").Handler(httptransport.NewServer(
		e.DeleteBookingEndpoint,
		decodeDeleteBookingRequest,
		encodeResponse,
		options...,
	))
}

func decodeFindBookingByIDRequest(_ context.Context, r *http.Request) (interface{}, error) {
	return booking.FindBookingByIDRequest{}, nil
}

func decodeFindBookingsRequest(_ context.Context, r *http.Request) (interface{}, error) {
	return booking.FindBookingsRequest{}, nil
}

func decodeCreateBookingRequest(_ context.Context, r *http.Request) (interface{}, error) {
	return booking.CreateBookingRequest{}, nil
}

func decodeUpdateBookingRequest(_ context.Context, r *http.Request) (interface{}, error) {
	return booking.UpdateBookingRequest{}, nil
}

func decodeDeleteBookingRequest(_ context.Context, r *http.Request) (interface{}, error) {
	return booking.DeleteBookingRequest{}, nil
}
