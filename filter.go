package goxi

import (
	"net/url"
	"reflect"
)

// FilterInterface allows for future support of new objects the Nagios API may introduce
type FilterInterface interface {
	build() string
}

// HostFilter struct allows filtering of results from the API
// A filter can be defined to only return a specific subset of data; an empty filter will return results for all hosts
type HostFilter struct {
	HostID               []string `param:"host_id"`
	Name                 []string `param:"name"`
	DisplayName          []string `param:"display_name"`
	Address              []string `param:"address"`
	Alias                []string `param:"alias"`
	CurrentState         []string `param:"current_state"`
	AcknowledgementType  []string `param:"acknowledgement_type"`
	StateType            []string `param:"state_type"`
	ProblemAcknowledged  []string `param:"problem_acknowledged"`
	PassiveChecksEnabled []string `param:"passive_checks_enabled"`
	ActiveChecksEnabled  []string `param:"active_checks_enabled"`
	FlapDetectionEnabled []string `param:"flap_detection_enabled"`
	IsFlapping           []string `param:"is_flapping"`
	Records              string   `param:"records"`
}

// ServiceFilter struct allows filtering of results from the API
// A filter can be defined to only return a specific subset of data; an empty filter will return results for all hosts
type ServiceFilter struct {
	InstanceID           []string `param:"instance_id"`
	ServiceID            []string `param:"service_id"`
	HostID               []string `param:"host_id"`
	HostName             []string `param:"host_name"`
	HostAlias            []string `param:"host_alias"`
	Name                 []string `param:"name"`
	HostDisplayName      []string `param:"host_display_name"`
	HostAddress          []string `param:"host_address"`
	DisplayName          []string `param:"display_name"`
	CurrentState         []string `param:"current_state"`
	AcknowledgementType  []string `param:"acknowledgement_type"`
	CheckCommand         []string `param:"check_command"`
	StateType            []string `param:"state_type"`
	NotificationsEnabled []string `param:"notifications_enabled"`
	ProblemAcknowledged  []string `param:"problem_acknowledged"`
	PassiveChecksEnabled []string `param:"passive_checks_enabled"`
	ActiveChecksEnabled  []string `param:"active_checks_enabled"`
	IsFlapping           []string `param:"is_flapping"`
	Records              string   `param:"records"`
}

func (h HostFilter) build() string {
	return toParam(h)
}

func (s ServiceFilter) build() string {
	return toParam(s)
}

// toParam takes in a filter and returns a string containing the parameters for the URL
func toParam(f FilterInterface) string {

	var fs string
	t := reflect.TypeOf(f)  //type of filter
	v := reflect.ValueOf(f) //values of filter

	// Loop through each field in the struct
	for i := 0; i < t.NumField(); i++ {

		// For each field, get: entire field; tag for field; value of field
		field := t.Field(i)
		tag := field.Tag.Get("param")
		value := v.Field(i)

		// If field is a non-empty slice, URL structured with an `in` query for slice elements
		// Else, separate each slice element with an ampersand
		if value.Kind() == reflect.Slice && value.Len() != 0 {
			fs += "&" + tag + "=in:"
			for i := 0; i < value.Len(); i++ {
				fs += url.QueryEscape(value.Index(i).String())
				if i != value.Len()-1 {
					fs += ","
				}
			}

		} else if value.Len() > 0 {
			fs += "&" + tag + "=" + url.QueryEscape(value.String())
		}

	}
	return fs
}
