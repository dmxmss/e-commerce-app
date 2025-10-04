package dto

import "time"

type Date struct {
	*time.Time
}

func (d *Date) UnmarshalParam(param string) error { // function to parse time as query parameter
	if param == "" {
		d.Time = nil
		return nil
	}

	t, err := time.Parse("2006-01-02T15:04:05.999999Z", param)
	if err != nil { 
		return err
	}

	d.Time = &t

	return nil
}
