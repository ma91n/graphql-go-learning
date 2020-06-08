package graph

import (
	"context"
	"database/sql"
	"errors"
	"github.com/laqiiz/graphql-go-learning/eki2/graph/model"
	"github.com/laqiiz/graphql-go-learning/eki2/models"
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
	sts, err := models.StationByNamesByStationName(db, *name)
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
	stations, err := models.StationByCDsByStationCD(db, *stationCd)
	if err != nil {
		return nil, err
	}
	if len(stations) == 0 {
		return nil, errors.New("not found")
	}
	first := stations[0]

	return &model.StationConn{
		Station: &model.Station{
			StationCd:   first.StationCd,
			StationGcd:  first.StationGCd,
			StationName: first.StationName,
			LineCd:      first.LineCd,
			LineName:    &first.LineName,
			Address:     &first.Address,
		},
	}, nil
}

func (r *stationConnResolver) transferStation(ctx context.Context, obj *model.StationConn) ([]*model.Station, error) {
	stationGroupCD := obj.Station.StationGcd

	records, err := models.TransfersByStationCD(db, stationGroupCD)
	if err != nil {
		return nil, err
	}

	resp := make([]*model.Station, 0, len(records))
	for _, v := range records {
		if v.TransferStationName == "" {
			continue
		}
		resp = append(resp, &model.Station{
			StationCd:   v.TransferStationCd,
			StationGcd:  v.StationGCd,
			StationName: v.TransferStationName,
			LineCd:      v.TransferLineCd,
			LineName:    &v.TransferLineName,
			Address:     &v.TransferAddress,
		})
	}

	return resp, nil
}

func (r *stationConnResolver) beforeStation(ctx context.Context, obj *model.StationConn) (*model.Station, error) {
	stationCD := obj.Station.StationCd
	records, err := models.BeforesByStationCD(db, stationCD)
	if err != nil {
		return nil, err
	}

	if len(records) == 0 {
		return nil, nil
	}

	if records[0].BeforeStationName == "" {
		return nil, nil
	}
	
	return &model.Station{
		StationCd:   records[0].StationCd,
		StationGcd:  records[0].BeforeStationGCd,
		StationName: records[0].BeforeStationName,
		LineCd:      records[0].LineCd,
		LineName:    &records[0].LineName,
		Address:     &records[0].BeforeStationAddress,
	}, nil
}

func (r *stationConnResolver) afterStation(ctx context.Context, obj *model.StationConn) (*model.Station, error) {
	stationCD := obj.Station.StationCd
	records, err := models.AftersByStationCD(db, stationCD)
	if err != nil {
		return nil, err
	}

	if len(records) == 0 {
		return nil, nil
	}

	if records[0].AfterStationName == "" {
		return nil, nil
	}

	return &model.Station{
		StationCd:   records[0].AfterStationCd,
		StationGcd:  records[0].AfterStationGCd,
		StationName: records[0].AfterStationName,
		LineCd:      records[0].LineCd,
		LineName:    &records[0].LineName,
		Address:     &records[0].AfterStationAddress,
	}, nil

}
