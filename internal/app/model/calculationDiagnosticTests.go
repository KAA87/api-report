package model

import (
	"reflect"
	"strconv"
)

// CalculationDiagnosticTests ...
type CalculationDiagnosticTests struct {
	ID                int
	DiagnosticTestsID int
	DeviceUsersID     int
	PatientID         int
	Device            string
	Value1            int
	Fall1             int
	Value2            int
	Fall2             int
	Value3            int
	Fall3             int
	Value4            int
	Fall4             int
	Value5            int
	Fall5             int
	Value6            int
	Fall6             int
	Value7            int
	Fall7             int
	Value8            int
	Fall8             int
	Value9            int
	Fall9             int
	Value10           int
	Fall10            int
	Value11           int
	Fall11            int
	Value12           int
	Fall12            int
	Value13           int
	Fall13            int
	Value14           int
	Fall14            int
	Value15           int
	Fall15            int
	Value16           int
	Fall16            int
	Value17           int
	Fall17            int
	Value18           int
	Fall18            int
	Value19           int
	Fall19            int
	Value20           int
	Fall20            int
	Value21           int
	Fall21            int
	Value22           int
	Fall22            int
	Value23           int
	Fall23            int
	Value24           int
	Fall24            int
	Value25           int
	Fall25            int
	Value26           int
	Fall26            int
	Value27           int
	Fall27            int
	Value28           int
	Fall28            int
	Value29           int
	Fall29            int
	Value30           int
	Fall30            int
	Value31           int
	Fall31            int
	Value32           int
	Fall32            int
	Value33           int
	Fall33            int
	Value34           int
	Fall34            int
	Value35           int
	Fall35            int
	Value36           int
	Fall36            int
	Value37           int
	Fall37            int
	Value38           int
	Fall38            int
	Value39           int
	Fall39            int
	Value40           int
	Fall40            int
	Value41           int
	Fall41            int
	Value42           int
	Fall42            int
	Value43           int
	Fall43            int
	Value44           int
	Fall44            int
	Value45           int
	Fall45            int
	Value46           int
	Fall46            int
	Value47           int
	Fall47            int
	Countitempw       int
	Countitemac       int
	Countitemwa       int
	Createdate        string
}

// GetValues
func (t *CalculationDiagnosticTests) GetValues() (map[string]int) {
	values := make(map[string]int)

	r := reflect.ValueOf(t)
	for i := 1; i <= 47; i++ {
		key := "Value" + strconv.Itoa(i)

		v := reflect.Indirect(r).FieldByName(key)

		values[key] = int(v.Int())
	}

	return values
}

func (t *CalculationDiagnosticTests) AvgValues() int {

	sumValues := 0

	for _, item := range t.GetValues() {
		sumValues += item
	}

	avgValues := sumValues/len(t.GetValues())

	return avgValues
}