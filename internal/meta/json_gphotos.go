package meta

import (
	"encoding/json"
	"fmt"
	"time"

	"gopkg.in/ugjka/go-tz.v2/tz"
)

type GPhoto struct {
	Title       string `json:"title"`
	Description string `json:"description"`
	Views       int    `json:"imageViews,string"`
	Album       GAlbum `json:"albumData"`
	Geo         GGeo   `json:"geoData"`
	TakenAt     GTime  `json:"photoTakenTime"`
	CreatedAt   GTime  `json:"creationTime"`
	UpdatedAt   GTime  `json:"modificationTime"`
}

func (m GPhoto) SanitizedTitle() string {
	return SanitizeTitle(m.Title)
}

func (m GPhoto) SanitizedDescription() string {
	return SanitizeDescription(m.Description)
}

type GAlbum struct {
	Title       string `json:"title"`
	Description string `json:"description"`
	Access      string `json:"access"`
	Location    string `json:"location"`
	Date        GTime  `json:"date"`
	Geo         GGeo   `json:"geoData"`
}

func (m GAlbum) Exists() bool {
	return m.Title != ""
}

type GGeo struct {
	Lat      float64 `json:"latitude"`
	Lng      float64 `json:"longitude"`
	Altitude float64 `json:"altitude"`
}

func (m GGeo) Exists() bool {
	return m.Lat != 0.0 && m.Lng != 0.0
}

type GTime struct {
	Unix      int64  `json:"timestamp,string"`
	Formatted string `json:"formatted"`
}

func (m GTime) Exists() bool {
	return m.Unix > 0
}

func (m GTime) Time() time.Time {
	return time.Unix(m.Unix, 0).UTC()
}

// Parses JSON sidecar data as created by Google Photos.
func (data *Data) GPhotos(jsonData []byte) (err error) {
	defer func() {
		if e := recover(); e != nil {
			err = fmt.Errorf("metadata: %s (gphotos panic)", e)
		}
	}()

	p := GPhoto{}

	if err := json.Unmarshal(jsonData, &p); err != nil {
		return err
	}

	if s := p.SanitizedTitle(); s != "" {
		data.Title = s
	}

	if s := p.SanitizedDescription(); s != "" {
		data.Description = s
	}

	if p.Views > 0 {
		data.Views = p.Views
	}

	if p.TakenAt.Exists() {
		data.TakenAt = p.TakenAt.Time()
		data.TakenAtLocal = p.TakenAt.Time()
	}

	if p.Geo.Exists() {
		data.Lat = float32(p.Geo.Lat)
		data.Lng = float32(p.Geo.Lng)
		data.Altitude = int(p.Geo.Altitude)
	}

	if p.Album.Exists() {
		data.Albums = append(data.Albums, p.Album.Title)
	}

	// Set time zone and calculate UTC time.
	if data.Lat != 0 && data.Lng != 0 {
		zones, err := tz.GetZone(tz.Point{
			Lat: float64(data.Lat),
			Lon: float64(data.Lng),
		})

		if err == nil && len(zones) > 0 {
			data.TimeZone = zones[0]
		}

		if !data.TakenAtLocal.IsZero() {
			if loc, err := time.LoadLocation(data.TimeZone); err != nil {
				log.Warnf("metadata: unknown time zone %s (gphotos)", data.TimeZone)
			} else if tl, err := time.ParseInLocation("2006:01:02 15:04:05", data.TakenAtLocal.Format("2006:01:02 15:04:05"), loc); err == nil {
				data.TakenAt = tl.Round(time.Second).UTC()
			} else {
				log.Errorf("metadata: %s (gphotos)", err.Error()) // this should never happen
			}
		}
	}

	return nil
}
