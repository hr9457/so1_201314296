package models

// struct data memory ram
type MemoryRam struct {
	TotalMemory    int     `json:"total_memory"`
	ConsumedMemory int     `json:"consumed_memory"`
	UseMemory      float64 `json:"use_memory"`
}

type ModuleRam struct {
	Total  int `json:"total_memory"`
	Free   int `json:"free_memory"`
	Buffer int `json:"buffer_memory"`
}
