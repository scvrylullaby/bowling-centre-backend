package models

type DashboardState struct {
	ActiveLanes map[int]int `json:"active_lanes"`
	Stats       Stats       `json:"stats"`
}
