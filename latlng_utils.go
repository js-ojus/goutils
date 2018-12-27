package goutils

import (
	"errors"
	"regexp"
	"strconv"
)

var (
	re1 = regexp.MustCompile(`([NEWS])(.+)`)
	re2 = regexp.MustCompile(`(.+?)([NEWS])`)
	reD = regexp.MustCompile(`(\d{1,3})°?(.*)`)
	reM = regexp.MustCompile(`(\d{1,2})'?(.*)`)
	reS = regexp.MustCompile(`(\d{1,2}\.?\d*)"?`)
)

// DMSStr2Float takes either a latitude or a longitude in one of the two
// following (string) formats:
//
//   [NEWS]\d{1,3}°\d{1,2}'\d{1,2}(\.\d+?")?
//   \d{1,3}°\d{1,2}'\d{1,2}(\.\d+?)?"[NEWS]
//
// and converts it into a decimal value.
//
// *N.B.* This does NOT perform range checks on latitudes and longitudes.
func DMSStr2Float(dms string) (float64, error) {
	d, m, s := 0.0, 0.0, 0.0
	var dir, rest string
	var err error

	gs := re1.FindStringSubmatch(dms)
	if gs != nil {
		dir = gs[1]
		rest = gs[2]
	} else {
		if gs = re2.FindStringSubmatch(dms); gs != nil {
			rest = gs[1]
			dir = gs[2]
		} else {
			return 0.0, errors.New("unknown coordinate format")
		}
	}
	sign := 1.0
	if dir == "W" || dir == "S" {
		sign = -1.0
	}

	gs = reD.FindStringSubmatch(rest)
	if gs == nil {
		return 0.0, errors.New("unknown coordinate format")
	}
	d, err = strconv.ParseFloat(gs[1], 64)
	if err != nil {
		return 0.0, errors.New(gs[1] + " : d : " + err.Error())
	}
	if gs[2] == "" {
		return sign * d, nil
	}
	rest = gs[2]

	gs = reM.FindStringSubmatch(rest)
	if gs == nil {
		return 0.0, errors.New("unknown coordinate format")
	}
	m, err = strconv.ParseFloat(gs[1], 64)
	if err != nil {
		return 0.0, errors.New(gs[1] + " : m : " + err.Error())
	}
	if gs[2] == "" {
		return sign * (d + m/60), nil
	}
	rest = gs[2]

	gs = reS.FindStringSubmatch(rest)
	if gs == nil {
		return 0.0, errors.New("unknown coordinate format")
	}
	s, err = strconv.ParseFloat(gs[1], 64)
	if err != nil {
		return 0.0, errors.New(gs[1] + " : s : " + err.Error())
	}
	return sign * (d + (m*60+s)/3600), nil
}
