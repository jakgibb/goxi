package goxi

import (
	"bytes"
	"encoding/json"
	"fmt"
	"log"
)

// host struct holds the unmarshaled host results
type host []struct {
	HostID                    string `json:"host_id"`
	Name                      string `json:"name"`
	DisplayName               string `json:"display_name"`
	Address                   string `json:"address"`
	Alias                     string `json:"alias"`
	StatusUpdateTime          string `json:"status_update_time"`
	StatusText                string `json:"status_text"`
	StatusTextLong            string `json:"status_text_long"`
	CurrentState              string `json:"current_state"`
	IconImage                 string `json:"icon_image"`
	IconImageAlt              string `json:"icon_image_alt"`
	PerformanceData           string `json:"performance_data"`
	ShouldBeScheduled         string `json:"should_be_scheduled"`
	CheckType                 string `json:"check_type"`
	LastStateChange           string `json:"last_state_change"`
	LastHardStateChange       string `json:"last_hard_state_change"`
	LastHardState             string `json:"last_hard_state"`
	LastTimeUp                string `json:"last_time_up"`
	LastTimeDown              string `json:"last_time_down"`
	LastTimeUnreachable       string `json:"last_time_unreachable"`
	LastNotification          string `json:"last_notification"`
	NextNotification          string `json:"next_notification"`
	NoMoreNotifications       string `json:"no_more_notifications"`
	AcknowledgementType       string `json:"acknowledgement_type"`
	CurrentNotificationNumber string `json:"current_notification_number"`
	EventHandlerEnabled       string `json:"event_handler_enabled"`
	ProcessPerformanceData    string `json:"process_performance_data"`
	ObsessOverHost            string `json:"obsess_over_host"`
	ModifiedHostAttributes    string `json:"modified_host_attributes"`
	EventHandler              string `json:"event_handler"`
	CheckCommand              string `json:"check_command"`
	NormalCheckInterval       string `json:"normal_check_interval"`
	RetryCheckInterval        string `json:"retry_check_interval"`
	CheckTimeperiodID         string `json:"check_timeperiod_id"`
	HasBeenChecked            string `json:"has_been_checked"`
	CurrentCheckAttempt       string `json:"current_check_attempt"`
	MaxCheckAttempts          string `json:"max_check_attempts"`
	LastCheck                 string `json:"last_check"`
	NextCheck                 string `json:"next_check"`
	StateType                 string `json:"state_type"`
	NotificationsEnabled      string `json:"notifications_enabled"`
	ProblemAcknowledged       string `json:"problem_acknowledged"`
	PassiveChecksEnabled      string `json:"passive_checks_enabled"`
	ActiveChecksEnabled       string `json:"active_checks_enabled"`
	FlapDetectionEnabled      string `json:"flap_detection_enabled"`
	IsFlapping                string `json:"is_flapping"`
	PercentStateChange        string `json:"percent_state_change"`
	Latency                   string `json:"latency"`
	ExecutionTime             string `json:"execution_time"`
	ScheduledDowntimeDepth    string `json:"scheduled_downtime_depth"`
	NotesURL                  string `json:"notes_url"`
	ActionURL                 string `json:"action_url"`
}

//hostResponse struct holds the raw response from the API which will be analysed and unmarshaled into the `host` struct
type hostResponse struct {
	RecordCount int             `json:"recordcount,string"`
	RawStatus   json.RawMessage `json:"hoststatus"`
}

// unmarshal (hostResponse) is a custom unmarshaler which takes the raw JSON result and populates the `hostResponse` struct
func (r *hostResponse) unmarshal(jsonResp *[]byte) {

	// Bug (#1 of #2) exists in the Nagios API <=5.6 which returns null/empty values as an empty object `{}`
	// Reported and will be fixed in v5.7: https://support.nagios.com/forum/viewtopic.php?f=20&t=56741
	// Workaround: remove any empty objects `{}` from the JSON
	*jsonResp = bytes.ReplaceAll(*jsonResp, []byte(":{}"), []byte(":\"\""))

	if err := json.Unmarshal(*jsonResp, r); err != nil {
		log.Fatal(err)
	}
}

// unmarshal (host) is a custom unmarshaler to populate the various fields of the struct as held in `hostResponse`
func (r *host) unmarshal(resp *hostResponse) {
	if resp.RecordCount > 0 {
		// Bug (#2 of #2) in the Nagios API <=5.6 which returns inconsistent JSON types depending on the number of
		// results returned, preventing unmasrhalling into slice of structs
		// Reported and will be fixed in v5.7: https://support.nagios.com/forum/viewtopic.php?f=20&t=56741
		// Workaround: if a single result is returned, create a temporary slice of Hosts (tmp), unmarshal, assign tmp to *r
		if resp.RecordCount == 1 {

			tmp := make(host, 1, 1)

			if err := json.Unmarshal(resp.RawStatus, &tmp[0]); err != nil {
				fmt.Println(err)
			}

			*r = tmp
			return
		}

		if err := json.Unmarshal(resp.RawStatus, r); err != nil {
			fmt.Println(err)
		}
	}
}
