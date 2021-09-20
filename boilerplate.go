package boilerplate

import (
	"bytes"
	"fmt"
	"os"
	"regexp"
	"strings"
	"text/template"
)

type Service struct {
	Name        string
	Methods     []string
	PackageName string
}

func Generate(s Service) {
	generateServiceDefinition(s)
}

func generateServiceDefinition(s Service) {
	generateBoilerplate(serviceTemplate, fmt.Sprintf("domain/%s.go", strings.ToLower(s.Name)), s)
	generateBoilerplate(endpointTemplate, fmt.Sprintf("endpoint/%s.go", strings.ToLower(s.Name)), s)
	generateBoilerplate(logTemplate, fmt.Sprintf("log/%s.go", strings.ToLower(s.Name)), s)
	generateBoilerplate(metricsTemplate, fmt.Sprintf("metrics/%s.go", strings.ToLower(s.Name)), s)
	generateBoilerplate(httpTemplate, fmt.Sprintf("http/%s.go", strings.ToLower(s.Name)), s)
}

func generateBoilerplate(templateStr, output string, service Service) {
	funcMap := template.FuncMap{
		"ToLower":     strings.ToLower,
		"ToSnakeCase": ToSnakeCase,
	}
	templ := template.New(output).Funcs(funcMap)
	templ, err := templ.Parse(templateStr)
	if err != nil {
		panic(err)
	}
	buf := new(bytes.Buffer)
	templ.Execute(buf, service)

	// Write buffer to file
	err = os.WriteFile(output, buf.Bytes(), 0644)
	if err != nil {
		panic(err)
	}
}

var (
	matchFirstCap = regexp.MustCompile("([A-Z])([A-Z][a-z])")
	matchAllCap   = regexp.MustCompile("([a-z0-9])([A-Z])")
)

// ToSnakeCase converts the provided string to snake_case.
// Based on https://gist.github.com/stoewer/fbe273b711e6a06315d19552dd4d33e6
func ToSnakeCase(input string) string {
	output := matchFirstCap.ReplaceAllString(input, "${1}_${2}")
	output = matchAllCap.ReplaceAllString(output, "${1}_${2}")
	output = strings.ReplaceAll(output, "-", "_")
	return strings.ToLower(output)
}

var serviceTemplate = `package {{.PackageName}}

import (
	"context"
	"fmt"
	"time"
)

type {{.Name}}Service interface {{"{"}}{{range $method := .Methods}}
	{{$method}}(ctx context.Context, req {{$method}}Request) {{$method}}Response{{end}}
}
{{range $method := .Methods}}
// {{$method}}Request represents a payload used by the {{$method}} method of a {{$.Name}}Service
type {{$method}}Request struct {
	// insert request properties here
}

// Validate a {{$method}}. Returns a ValidationError for each requirement that fails.
func (r {{$method}}Request) Validate() []ValidationError {
	// insert validation logic here
	return nil
}

// {{$method}}Response represents a response returned by the {{$method}} method of a {{$.Name}}Service.
type {{$method}}Response struct {
	// insert response properties here
}

// Error implements the errorer interface. Returns property Err from the response.
func (r {{$method}}Response) Error() error { return r.Err }
{{end}}
// {{.Name}}ServiceMiddleware defines a middleware for {{.Name}}Service
type {{.Name}}ServiceMiddleware func(service {{.Name}}Service) {{.Name}}Service

// {{.Name}}ValidationMiddleware returns a middleware for validating requests made to a {{.Name}}Service
func {{.Name}}ValidationMiddleware() {{.Name}}ServiceMiddleware {
	return func(next {{.Name}}Service) {{.Name}}Service {
		return {{.Name | ToLower}}ValidationMiddleware{next}
	}
}
{{range $method := .Methods}}
// {{$method}} validates a {{$method}}Request. Returns a domain error if any requirements fail and invokes the next middleware otherwise.
func (mw {{$.Name | ToLower}}ValidationMiddlware) {{$method}}(ctx context.Context, req {{$method}}Request) {{$method}}Response {
	errs := req.Validate()
	if len(errs) > 0 {
		return {{$method}}Response{Err: wrapValidationErrors(errs)}
	}
	return mw.{{$.Name}}Service.{{$method}}(ctx, req)
}
{{end}}`

var endpointTemplate = `package endpoint

import (
	"context"

	"github.com/go-kit/kit/endpoint"
	"github.com/openmesh/{{.PackageName}}"
)

// {{.Name}}Endpoints collects all the endpoints that compose a {{.PackageName}}.{{.Name}}Service.
// It's used as a helper struct, to collect all the endpoints into a single
// parameter.
type {{.Name}}Endpoints struct {{"{"}}{{range $method := .Methods}}
	{{$method}}Endpoint endpoint.Endpoint{{end}}
}

// Make{{.Name}}Endpoints returns a {{.Name}}Endpoints struct where each endpoint
// invokes the corresponding method on the provided service.
func Make{{.Name}}Endpoints(s {{.PackageName}}.{{.Name}}Service) {{.Name}}Endpoints {
	return {{.Name}}Endpoints{{"{"}}{{range $method := .Methods}}
		{{$method}}Endpoint: Make{{$method}}Endpoint(s),{{end}}
	}
}
{{range $method := .Methods}}
// Make{{$method}}Endpoint returns an endpoint via the passed service.
func Make{{$method}}Endpoint(s {{$.PackageName}}.{{$.Name}}Service) endpoint.Endpoint {
	return func(ctx context.Context, r interface{}) (interface{}, error) {
		return s.{{$method}}(ctx, r.({{$.PackageName}}.{{$method}}Request)), nil
	}
}{{end}}`

var logTemplate = `package log

import (
	"context"
	"time"

	"github.com/go-kit/kit/log"
	"github.com/openmesh/{{.PackageName}}"
)

func {{.Name}}LoggingMiddleware(logger log.Logger) {{.PackageName}}.{{.Name}}ServiceMiddleware {
	return func(next {{.PackageName}}.{{.Name}}Service) {{.PackageName}}.{{.Name}}Service {
		return {{.Name | ToLower}}LoggingMiddleware{logger, next}
	}
}

type {{.Name | ToLower}}LoggingMiddleware struct {
	logger log.Logger
	{{.PackageName}}.{{.Name}}Service
}
{{range $method := .Methods}}
func (mw {{$.Name | ToLower}}LoggingMiddleware) {{$method}}(ctx context.Context, req {{$.PackageName}}.{{$method}}Request) (res {{$.PackageName}}.{{$method}}Response) {
	defer func(begin time.Time) {
		_ = mw.logger.Log(
			"method", "{{$method | ToSnakeCase}}",
			"request", req,
			"response", res,
			"took", time.Since(begin),
		)
	}(time.Now())
	res = mw.{{$.Name}}Service.{{$method}}(ctx, req)
	return
}
{{end}}`

var metricsTemplate = `package metrics

import (
	"context"
	"time"

	"github.com/go-kit/kit/metrics"
	"github.com/openmesh/{{.PackageName}}"
)

func {{.Name}}MetricsMiddleware(
	requestCount metrics.Counter,
	errorCount metrics.Counter,
	requestDuration metrics.Histogram,
) {{.PackageName}}.{{.Name}}ServiceMiddleware {
	return func(next {{.PackageName}}.{{.Name}}Service) {{.PackageName}}.{{.Name}}Service {
		return {{.Name | ToLower}}MetricsMiddleware{requestCount, errorCount, requestDuration, next}
	}
}

type {{.Name | ToLower}}MetricsMiddleware struct {
	requestCount    metrics.Counter
	errorCount      metrics.Counter
	requestDuration metrics.Histogram
	{{.PackageName}}.{{.Name}}Service
}

{{range $method := .Methods}}
func (mw {{$.Name | ToLower}}MetricsMiddleware) {{$method}}(ctx context.Context, req {{$.PackageName}}.{{$method}}Request) (res {{$.PackageName}}.{{$method}}Response) {
	defer func(begin time.Time) {
		lvs := []string{"method", "{{$method | ToSnakeCase}}"}
		mw.requestCount.With(lvs...).Add(1)
		mw.requestDuration.With(lvs...).Observe(time.Since(begin).Seconds())
		if res.Err != nil {
			mw.errorCount.With(lvs...).Add(1)
		}
	}(time.Now())
	res = mw.{{$.Name}}Service.{{$method}}(ctx, req)
	return
}
{{end}}`

var httpTemplate = `package http

import (
	"context"
	"net/http"

	"github.com/openmesh/{{.PackageName}}"

	httptransport "github.com/go-kit/kit/transport/http"
	"github.com/go-kit/transport"
	"github.com/gorilla/mux"
	"github.com/openmesh/{{.PackageName}}/endpoint"
)

func (s *Server) register{{.Name}}Routes(r *mux.Router) {
	e := endpoint.Make{{.Name}}Endpoints(s.{{.Name}}Service)

	options := []httptransport.ServerOption{
		httptransport.ServerErrorHandler(transport.NewLogErrorHandler(s.logger)),
		httptransport.ServerErrorEncoder(encodeError),
	}
	{{range $method := .Methods}}
	r.Methods("insert_method_here").Path("insert_path_here").Handler(httptransport.NewServer(
		e.{{$method}}Endpoint,
		decode{{$method}}Request,
		encodeResponse,
		options...,
	))
{{end}}}
{{range $method := .Methods}}
func decode{{$method}}Request(_ context.Context, r *http.Request) (interface{}, error) {
	return {{$.PackageName}}.{{$method}}Request{}, nil
}
{{end}}`
