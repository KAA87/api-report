package calculations

import (
	"fmt"
	"github.com/KAA87/api-report.git/internal/app/model"
	"github.com/KAA87/api-report.git/store"
)

type rGeneralState struct {
	GeneralState               map[string]string
	UnbalanceFactor            string
	DegreeRiskForMajorSystems  string
	GeneralStability           string
	TypeAutonomicNervousSystem string
	CentralNervousSystem       string
}

func CalcGeneralState(t *model.CalculationDiagnosticTests) (*rGeneralState, error) {
	c := &rGeneralState{}

	c.calcGeneralState(t)

	return c, nil
}

func (r *rGeneralState) calcGeneralState(t *model.CalculationDiagnosticTests){
	r.GeneralState = make(map[string]string)

	avgValue := t.AvgValues()

	avgBordersSignals := (&store.SignalRateIntervalRepository{}).AvgBorderValues()
	fmt.Print("\n -------------- \n")
	fmt.Print(avgBordersSignals)
	fmt.Print("\n -------------- \n")


	avgBorderMinSignal := avgBordersSignals["min"]
	avgBorderMaxSignal := avgBordersSignals["max"]

	avgBorderMinSignalItem := avgBorderMinSignal/2
	avgBorderMaxSignalItem := (100 - avgBorderMaxSignal)/2

	if avgValue < avgBorderMinSignalItem {
		r.GeneralState["color"] = "purple"
		r.GeneralState["description"] = "unsatisfactory"
	} else if avgValue >= avgBorderMinSignalItem && avgValue < avgBorderMinSignal {
		r.GeneralState["color"] = "blue"
		r.GeneralState["description"] = "satisfactory"
	} else if avgValue >= avgBorderMaxSignal && avgValue < (avgBorderMaxSignal + avgBorderMaxSignalItem) {
		r.GeneralState["color"] = "orange"
		r.GeneralState["description"] = "satisfactory"
	} else if avgValue >= (avgBorderMaxSignal + avgBorderMaxSignalItem) {
		r.GeneralState["color"] = "red"
		r.GeneralState["description"] = "unsatisfactory"
	} else {
		r.GeneralState["color"] = "green"
		r.GeneralState["description"] = "good"
	}

}
