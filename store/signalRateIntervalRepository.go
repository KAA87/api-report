package store

import (
	"fmt"
	"github.com/KAA87/api-report.git/internal/app/model"
)

// SignalRateIntervalRepository ...
type SignalRateIntervalRepository struct {
	store *Store
}

func (r *SignalRateIntervalRepository) All() ([]*model.SignalRateInterval, error) {

	rows, err := r.store.db.Query("SELECT * FROM signal_rate_interval WHERE")
	if err != nil {
		return nil, err
	}

	list := make([]*model.SignalRateInterval, 0)
	for rows.Next() {
		t := &model.SignalRateInterval{}
		err := rows.Scan(&t.ID, &t.IntervalNormMin, &t.IntervalNormMax)
		if err != nil {
			return nil, err
		}

		list = append(list, t)
	}

	fmt.Print("\n -------------- \n")
	fmt.Print("test2")
	fmt.Print("\n -------------- \n")

	return list, nil
}


func (r *SignalRateIntervalRepository) AvgBorderValues() map[string]int {

	avgBorderValues := make(map[string]int)

	avgBorderValues["min"] = 0
	avgBorderValues["max"] = 0

	r.All()



	//rateIntervals,err := r.All()
	//if err != nil {
	//	return avgBorderValues
	//}
	//
	//for _, item := range rateIntervals {
		//avgBorderValues["min"] += item.IntervalNormMin
		//avgBorderValues["max"] += item.IntervalNormMax
	//}
	//
	//avgBorderValues["min"] = avgBorderValues["min"] / len(rateIntervals)
	//avgBorderValues["max"] = avgBorderValues["max"] / len(rateIntervals)


	return avgBorderValues
}