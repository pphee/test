package bmi

import "time"

type BMI struct {
	ID              int64     `json:"id"`
	Uuid            string    `json:"uuid"`
	Name            string    `json:"name"`
	Age             int64     `json:"age"`
	Height          float64   `json:"height"`
	Weight          float64   `json:"weight"`
	Value           float64   `json:"value"`
	Category        string    `json:"category,omitempty"`
	Risk            string    `json:"risk,omitempty"`
	Vector          []float32 `json:"vector,omitempty"`
	CreatedAt       time.Time `json:"created_at"`
	RecordNumber    int       `json:"-"`
	VectorGenerated bool      `json:"-"`
}

type BMIIV struct {
	Uuid      string    `json:"uuid"`
	Name      string    `json:"name"`
	Age       int64     `json:"age"`
	Height    float64   `json:"height"`
	Weight    float64   `json:"weight"`
	Value     float64   `json:"value"`
	Category  string    `json:"category,omitempty"`
	Risk      string    `json:"risk,omitempty"`
	Vector    []float32 `json:"vector,omitempty"`
	CreatedAt time.Time `json:"created_at"`
}

type BMICalculationRequest struct {
	Height float64 `json:"height" validate:"required,gt=0"`
	Weight float64 `json:"weight" validate:"required,gt=0"`
}

type PointStruct struct {
	ID      string                 `json:"id"`
	Vector  []float32              `json:"vector"`
	Payload map[string]interface{} `json:"payload"`
}
