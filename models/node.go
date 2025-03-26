package models

type Node struct {
    ID          int    `json:"id"`
    DisplayName string `json:"display_name"`
    ParentID    *int   `json:"parent_id"` // Pointer for null support
    NodeLeft    int    `json:"node_left"`
    NodeRight   int    `json:"node_right"`
}