// Code generated by OpenAPI Generator. DO NOT EDIT.
/*
 * Routing API v8
 *
 * A location service providing customizable route calculations for a variety of vehicle types as well as pedestrian modes.
 *
 * API version: 8.3.1
 * Generated by: OpenAPI Generator (https://openapi-generator.tech)
 */

package routingv8
// RouterSection One of the possible sections that can be part of the Router route.  `TransitSection` is only used for ferries and car shuttles. 
type RouterSection struct {
	// Unique identifier of the section
	Id string `json:"id"`
	// Section type used by the client to identify what extension to the BaseSection are available.
	Type string `json:"type"`
	// Actions that must be done prior to `departure`.
	PreActions []BaseAction `json:"preActions,omitempty"`
	// Actions that must be done during the travel portion of the section, i.e., between `departure` and `arrival`.
	Actions []OffsetAction `json:"actions,omitempty"`
	// Language of the localized strings in the section, if any, in BCP47 format.
	Language string `json:"language,omitempty"`
	// Actions that must be done after `arrival`.
	PostActions []BaseAction `json:"postActions,omitempty"`
	// Actions for turn by turn guidance during the travel portion of the section, i.e., between `departure` and `arrival`.
	TurnByTurnActions []OffsetAction `json:"turnByTurnActions,omitempty"`
	Departure TransitDeparture `json:"departure"`
	Arrival TransitDeparture `json:"arrival"`
	// Total value of key attributes (e.g., duration, length) summed up for the entire section, including `preActions`, `postActions`, and the travel portion of the section. 
	Summary BaseSummary `json:"summary,omitempty"`
	// Total value of key attributes (e.g., duration, length) summed up for just the travel portion of the section, between `departure` and `arrival`. `preActions` and `postActions` are excluded. 
	TravelSummary BaseSummary `json:"travelSummary,omitempty"`
	// Line string in [Flexible Polyline](https://github.com/heremaps/flexible-polyline) format
	Polyline string `json:"polyline,omitempty"`
	// Contains a list of issues related to this section of the route.  Follows a list of possible notice codes:  * `noSchedule`: A timetable schedule is not available for the transit line in this section,   and only the run frequency is available. As a result, departure/arrival times are approximated. * `noIntermediate`: Information about intermediate stops is not available for this transit line. * `unwantedMode`: This section contains a transport mode that was explictly disabled.   Mode filtering is not available in this area. * `scheduledTimes`: The times returned on this section are the scheduled times even though delay information are available. * `simplePolyline`: An accurate polyline is not available for this transit line.   The returned polyline has been generated from departure, arrival and intermediate stop sequence. 
	Notices []Notice `json:"notices,omitempty"`
	Transport TransitTransport `json:"transport"`
	// Span attached to a `Section` describing transit content. 
	Spans []TransitSpan `json:"spans,omitempty"`
	// Intermediate stops between departure and destination of the transit line. It can be missing if this information is not available or not requested. 
	IntermediateStops []TransitStop `json:"intermediateStops,omitempty"`
	Agency Agency `json:"agency,omitempty"`
	// List of required attributions to display.
	Attributions []Attribution `json:"attributions,omitempty"`
	// List of tickets to pay for this section of the route
	Fares []Fare `json:"fares,omitempty"`
}
