package goxi

import (
	"fmt"
	"testing"
)

const hostJSON string = `{ "recordcount": "2", "hoststatus": [ { "@attributes": { "id": "123" }, "instance_id": "1", "host_id": "5154", "name": "server-a", "display_name": "server-a", "address": "server-a.example", "alias": "Load Balancer A", "status_update_time": "2020-01-04 09:49:20", "status_text": "PING OK - Packet loss = 0%, RTA = 6.98 ms", "status_text_long":{}, "current_state": "0", "icon_image":{}, "icon_image_alt":{}, "performance_data": "rta=6.984000ms;3000.000000;5000.000000;0.000000 pl=0%;80;100;0", "should_be_scheduled": "1", "check_type": "0", "last_state_change": "2019-11-10 14:38:37", "last_hard_state_change": "2019-11-10 14:38:37", "last_hard_state": "0", "last_time_up": "2020-01-04 09:49:16", "last_time_down": "2019-11-10 14:38:38", "last_time_unreachable": "1970-01-01 01:00:00", "last_notification": "2019-11-10 14:38:41", "next_notification": "1970-01-01 01:00:00", "no_more_notifications": "0", "acknowledgement_type": "0", "current_notification_number": "18", "event_handler_enabled": "1", "process_performance_data": "1", "obsess_over_host": "1", "modified_host_attributes": "0", "event_handler":{}, "check_command": "check-host-alive", "normal_check_interval": "5", "retry_check_interval": "2", "check_timeperiod_id": "131", "has_been_checked": "1", "current_check_attempt": "1", "max_check_attempts": "10", "last_check": "2020-01-04 09:49:16", "next_check": "2020-01-04 09:54:16", "state_type": "1", "notifications_enabled": "1", "problem_acknowledged": "0", "passive_checks_enabled": "1", "active_checks_enabled": "1", "flap_detection_enabled": "0", "is_flapping": "0", "percent_state_change": "0", "latency": "0.00146", "execution_time": "4.02664", "scheduled_downtime_depth": "0", "notes_url":{}, "action_url":{} }, { "@attributes": { "id": "1234" }, "instance_id": "1", "host_id": "5153", "name": "server-b", "display_name": "server-b", "address": "server-b.example", "alias": "Load Balancer B", "status_update_time": "2020-01-04 09:49:44", "status_text": "PING OK - Packet loss = 0%, RTA = 8.52 ms", "status_text_long":{}, "current_state": "0", "icon_image":{}, "icon_image_alt":{}, "performance_data": "rta=8.516000ms;3000.000000;5000.000000;0.000000 pl=0%;80;100;0", "should_be_scheduled": "1", "check_type": "0", "last_state_change": "2019-11-29 12:43:17", "last_hard_state_change": "2019-05-07 06:51:17", "last_hard_state": "0", "last_time_up": "2020-01-04 09:49:40", "last_time_down": "2019-11-29 12:43:17", "last_time_unreachable": "1970-01-01 01:00:00", "last_notification": "1970-01-01 01:00:00", "next_notification": "1970-01-01 01:00:00", "no_more_notifications": "0", "acknowledgement_type": "0", "current_notification_number": "7", "event_handler_enabled": "1", "process_performance_data": "1", "obsess_over_host": "1", "modified_host_attributes": "0", "event_handler":{}, "check_command": "check-host-alive", "normal_check_interval": "5", "retry_check_interval": "2", "check_timeperiod_id": "131", "has_been_checked": "1", "current_check_attempt": "1", "max_check_attempts": "10", "last_check": "2020-01-04 09:49:40", "next_check": "2020-01-04 09:54:40", "state_type": "1", "notifications_enabled": "1", "problem_acknowledged": "0", "passive_checks_enabled": "1", "active_checks_enabled": "1", "flap_detection_enabled": "0", "is_flapping": "0", "percent_state_change": "0", "latency": "0", "execution_time": "4.02012", "scheduled_downtime_depth": "0", "notes_url":{}, "action_url":{} } ] }`
const serviceJSON = `{"recordcount":"2","servicestatus":[{"@attributes":{"id":"2518769"},"instance_id":"1","service_id":"123","host_id":"5154","host_name":"server-a","host_alias":"Load Balancer A","name":"Linux Disk Space","host_display_name":{},"host_address":"server-a.example","display_name":"Linux Disk Space","status_update_time":"2020-01-04 10:32:48","status_text":"\/sys: 0%used(0MB\/0MB) \/: 78%used(20733MB\/26450MB) \/sys\/kernel\/debug: 0%used(0MB\/0MB) \/boot: 47%used(48MB\/102MB) (<91%) : OK","status_text_long":{},"current_state":"0","performance_data":{},"should_be_scheduled":"1","check_type":"0","last_state_change":"2020-01-02 13:32:01","last_hard_state_change":"2019-10-24 02:21:26","last_hard_state":"3","last_time_ok":"2020-01-04 10:32:47","last_time_warning":"1970-01-01 01:00:00","last_time_critical":"2018-03-20 23:28:41","last_time_unknown":"2020-01-02 13:32:01","last_notification":"1970-01-01 01:00:00","next_notification":"1970-01-01 01:00:00","no_more_notifications":"0","acknowledgement_type":"0","current_notification_number":"1","process_performance_data":"1","obsess_over_service":"1","event_handler_enabled":"1","modified_service_attributes":"0","event_handler":{},"check_command":"check_snmp_storage","normal_check_interval":"15","retry_check_interval":"5","check_timeperiod_id":"131","icon_image":{},"icon_image_alt":{},"has_been_checked":"1","current_check_attempt":"1","max_check_attempts":"4","last_check":"2020-01-04 10:32:47","next_check":"2020-01-04 10:47:47","state_type":"1","notifications_enabled":"1","problem_acknowledged":"0","flap_detection_enabled":"0","is_flapping":"0","percent_state_change":"0","latency":"0","execution_time":"0.22216","scheduled_downtime_depth":"0","passive_checks_enabled":"1","active_checks_enabled":"1","notes_url":{},"action_url":{}},{"@attributes":{"id":"2518770"},"instance_id":"1","service_id":"1234","host_id":"5154","host_name":"server-b","host_alias":"Load Balancer B","name":"Linux Swap Usage","host_display_name":{},"host_address":"server-b.example","display_name":"Linux swap usage","status_update_time":"2020-01-04 10:33:44","status_text":"Swap Space: 0%used(0MB\/2055MB) (<60%) : OK","status_text_long":{},"current_state":"0","performance_data":{},"should_be_scheduled":"1","check_type":"0","last_state_change":"2019-12-15 22:29:31","last_hard_state_change":"2019-10-06 23:47:35","last_hard_state":"3","last_time_ok":"2020-01-04 10:33:44","last_time_warning":"1970-01-01 01:00:00","last_time_critical":"2018-03-20 23:32:03","last_time_unknown":"2019-12-15 22:29:31","last_notification":"1970-01-01 01:00:00","next_notification":"1970-01-01 01:00:00","no_more_notifications":"0","acknowledgement_type":"0","current_notification_number":"0","process_performance_data":"1","obsess_over_service":"1","event_handler_enabled":"1","modified_service_attributes":"0","event_handler":{},"check_command":"check_snmp_storage","normal_check_interval":"15","retry_check_interval":"5","check_timeperiod_id":"131","icon_image":{},"icon_image_alt":{},"has_been_checked":"1","current_check_attempt":"1","max_check_attempts":"4","last_check":"2020-01-04 10:33:44","next_check":"2020-01-04 10:48:44","state_type":"1","notifications_enabled":"1","problem_acknowledged":"0","flap_detection_enabled":"0","is_flapping":"0","percent_state_change":"0","latency":"0","execution_time":"0.28911","scheduled_downtime_depth":"0","passive_checks_enabled":"1","active_checks_enabled":"1","notes_url":{},"action_url":{}}]}`

func TestGetHosts(t *testing.T) {
	mock := []byte(hostJSON)

	resp := hostResponse{}
	host := host{}

	resp.unmarshal(&mock)
	host.unmarshal(&resp)

	if resp.RecordCount != 2 {
		t.Error(fmt.Sprintf("Expected record count to be %d but instead got %d", 2, resp.RecordCount))
	}

	if host[0].Name != "server-a" {
		t.Error(fmt.Sprintf("Expected host name to be %s but instead got %s", "server-a", host[0].Name))
	}
}

func TestGetServices(t *testing.T) {
	mock := []byte(serviceJSON)

	resp := serviceResponse{}
	service := service{}

	resp.unmarshal(&mock)
	service.unmarshal(&resp)

	if resp.RecordCount != 2 {
		t.Error(fmt.Sprintf("Expected record count to be %d but instead got %d", 2, resp.RecordCount))
	}

	if service[0].Name != "Linux Disk Space" {
		t.Error(fmt.Sprintf("Expected service name be %s but instead got %s", "Linux Disk Space", service[0].Name))
	}
}
