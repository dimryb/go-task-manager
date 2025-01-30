package utils

import (
	"fmt"
	"strings"
	"time"
)

type JSONTime time.Time

func (jt *JSONTime) UnmarshalJSON(b []byte) error {
	str := strings.Trim(string(b), "\"")
	fmt.Println("Парсим дату:", str)
	if str == "null" || str == "" {
		return nil
	}

	t, err := time.Parse(time.RFC3339, str)
	if err != nil {
		return fmt.Errorf("ошибка парсинга времени: %w", err)
	}

	*jt = JSONTime(t)
	return nil
}

func (jt JSONTime) Time() time.Time {
	return time.Time(jt)
}
