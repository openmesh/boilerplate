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

func (s *Server) registerUnavailabilityRoutes(r *mux.Router) {
	e := endpoint.MakeUnavailabilityEndpoints(s.UnavailabilityService)

	options := []httptransport.ServerOption{
		httptransport.ServerErrorHandler(transport.NewLogErrorHandler(s.logger)),
		httptransport.ServerErrorEncoder(encodeError),
	}
	
	r.Methods("insert_method_here").Path("insert_path_here").Handler(httptransport.NewServer(
		e.FindUnavailabilityByIDEndpoint,
		decodeFindUnavailabilityByIDRequest,
		encodeResponse,
		options...,
	))

	r.Methods("insert_method_here").Path("insert_path_here").Handler(httptransport.NewServer(
		e.FindUnavailabilitiesEndpoint,
		decodeFindUnavailabilitiesRequest,
		encodeResponse,
		options...,
	))

	r.Methods("insert_method_here").Path("insert_path_here").Handler(httptransport.NewServer(
		e.CreateUnavailabilityEndpoint,
		decodeCreateUnavailabilityRequest,
		encodeResponse,
		options...,
	))

	r.Methods("insert_method_here").Path("insert_path_here").Handler(httptransport.NewServer(
		e.UpdateUnavailabilityEndpoint,
		decodeUpdateUnavailabilityRequest,
		encodeResponse,
		options...,
	))

	r.Methods("insert_method_here").Path("insert_path_here").Handler(httptransport.NewServer(
		e.DeleteUnavailabilityEndpoint,
		decodeDeleteUnavailabilityRequest,
		encodeResponse,
		options...,
	))
}

func decodeFindUnavailabilityByIDRequest(_ context.Context, r *http.Request) (interface{}, error) {
	return booking.FindUnavailabilityByIDRequest{}, nil
}

func decodeFindUnavailabilitiesRequest(_ context.Context, r *http.Request) (interface{}, error) {
	return booking.FindUnavailabilitiesRequest{}, nil
}

func decodeCreateUnavailabilityRequest(_ context.Context, r *http.Request) (interface{}, error) {
	return booking.CreateUnavailabilityRequest{}, nil
}

func decodeUpdateUnavailabilityRequest(_ context.Context, r *http.Request) (interface{}, error) {
	return booking.UpdateUnavailabilityRequest{}, nil
}

func decodeDeleteUnavailabilityRequest(_ context.Context, r *http.Request) (interface{}, error) {
	return booking.DeleteUnavailabilityRequest{}, nil
}
