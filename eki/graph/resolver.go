package graph

import (
	"context"
	"database/sql"
	"errors"
	"github.com/laqiiz/graphql-go-learning/eki/graph/model"
	"github.com/laqiiz/graphql-go-learning/eki/models"
	_ "github.com/lib/pq"
)

// This file will not be regenerated automatically.
//
// It serves as dependency injection for your app, add any dependencies you require here.

type Resolver struct{}

var db *sql.DB

func init() {
	conn, err := sql.Open("postgres", "user=postgres password=postgres sslmode=disable")
	if err != nil {
		panic(err)
	}
	db = conn
}

func (r *Resolver) getStationByName(ctx context.Context, name *string) ([]*model.Station, error) {
	sts, err := models.StationsByStationName(db, *name)
	if err != nil {
		return nil, err
	}

	resp := make([]*model.Station, 0, len(sts))
	for _, v := range sts {
		resp = append(resp, &model.Station{
			StationCd:   v.StationCd,
			StationGcd:  v.StationGCd,
			StationName: v.StationName,
			LineCd:      v.LineCd,
			LineName:    &v.LineName,
			Address:     &v.Address,
		})
	}

	return resp, nil
}

func (r *Resolver) getStation(ctx context.Context, stationCd *int) (*model.StationConn, error) {
	stations, err := models.StationConnsByStationCD(db, *stationCd)
	if err != nil {
		return nil, err
	}
	if len(stations) == 0 {
		return nil, errors.New("not found")
	}

	first := stations[0]

	var beforeStation *model.Station
	if first.BeforeStationName != "" {
		beforeStation = &model.Station{
			StationCd:   first.BeforeStationCd,
			StationGcd:  0,
			StationName: first.BeforeStationName,
			LineCd:      first.LineCd,
			LineName:    &first.LineName,
			Address:     nil,
		}
	}

	var afterStation *model.Station
	if first.AfterStationName != "" {
		afterStation = &model.Station{
			StationCd:   first.AfterStationCd,
			StationGcd:  0,
			StationName: first.AfterStationName,
			LineCd:      first.LineCd,
			LineName:    &first.LineName,
			Address:     nil,
		}
	}

	transfers := make([]*model.Station, 0, len(stations))
	for _, v := range stations {
		if v.TransferStationName == "" {
			continue
		}
		transfers = append(transfers, &model.Station{
			StationCd:   v.TransferStationCd,
			StationGcd:  v.StationGCd,
			StationName: v.TransferStationName,
			LineCd:      0,
			LineName:    &v.TransferLineName,
			Address:     nil,
		})
	}

	return &model.StationConn{
		Station: &model.Station{
			StationCd:   first.StationCd,
			StationGcd:  first.StationGCd,
			StationName: first.StationName,
			LineCd:      first.LineCd,
			LineName:    &first.LineName,
			Address:     &first.Address,
		},
		TransferStation: transfers,
		BeforeStation:   beforeStation,
		AfterStation:    afterStation,
	}, nil
}
