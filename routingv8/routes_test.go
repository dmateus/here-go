package routingv8_test

import (
	"bytes"
	"context"
	"encoding/json"
	"errors"
	"io"
	"net/http"
	"testing"

	"go.einride.tech/here/routingv8"
	"gotest.tools/v3/assert"
)

type RoutesMock struct {
	responseStatus int
	responseBody   routingv8.RoutesResponse
	error          *routingv8.HereErrorResponse
}

func (c *RoutesMock) Do(_ *http.Request) (*http.Response, error) {
	headers := http.Header{}
	headers.Add("Content-Type", "application/json")
	b, err := json.Marshal(c.responseBody)
	if err != nil {
		return nil, err
	}
	if c.error != nil {
		b, err = json.Marshal(c.error)
		if err != nil {
			return nil, err
		}
	}
	r := bytes.NewReader(b)
	return &http.Response{
		StatusCode:    c.responseStatus,
		Header:        headers,
		Body:          io.NopCloser(r),
		ContentLength: int64(len(b)),
	}, nil
}

func TestRoutingervice_Routes(t *testing.T) {
	t.Parallel()
	ctx := context.Background()

	// Einride Gothenburg.
	origin := routingv8.GeoWaypoint{
		Lat:  57.707752,
		Long: 11.949767,
	}
	// Einride Stockholm.
	destination := routingv8.GeoWaypoint{
		Lat:  59.337492,
		Long: 18.063672,
	}

	exp := routingv8.RoutesResponse{
		Routes: []routingv8.Route{
			{
				ID: "route-1",
				Sections: []routingv8.Section{
					{
						ID:   "section-1",
						Type: "veicle",
						Departure: routingv8.Place{
							Type:             "place",
							Location:         origin,
							OriginalLocation: origin,
						},
						Arrival: routingv8.Place{
							Type:             "place",
							Location:         destination,
							OriginalLocation: destination,
						},
						Summary: routingv8.Summary{
							Duration:     243,
							Length:       1206,
							BaseDuration: 136,
						},
					},
				},
			},
		},
		ErrorCodes: routingv8.ErrorCodes{routingv8.ErrorCodeSuccess},
	}
	httpClient := RoutesMock{responseBody: exp, responseStatus: 200}
	routingClient := routingv8.NewClient(&httpClient)

	got, err := routingClient.Routing.Routes(ctx, &routingv8.RoutesRequest{
		Origin:        origin,
		Destination:   destination,
		TransportMode: routingv8.TransportModeCar,
	})
	assert.NilError(t, err)
	assert.DeepEqual(t, &exp, got)
}

func TestRoutingervice_Routes_Error(t *testing.T) {
	t.Parallel()
	ctx := context.Background()

	// Einride Gothenburg.
	origin := routingv8.GeoWaypoint{
		Lat:  57.707752,
		Long: 11.949767,
	}
	// Einride Stockholm.
	destination := routingv8.GeoWaypoint{
		Lat:  59.337492,
		Long: 18.063672,
	}

	exp := routingv8.HereErrorResponse{
		Title:  "Mocked Error",
		Status: 400,
	}

	httpClient := RoutesMock{responseStatus: 400, error: &exp}
	routingClient := routingv8.NewClient(&httpClient)

	_, err := routingClient.Routing.Routes(ctx, &routingv8.RoutesRequest{
		Origin:        origin,
		Destination:   destination,
		TransportMode: routingv8.TransportModeCar,
	})
	var responseError *routingv8.ResponseError
	assert.Check(t, errors.As(err, &responseError))
	assert.DeepEqual(t, responseError.Response, &exp)
	assert.Check(t, responseError.HTTPBody != "")
}
