// Package census simulates a system used to collect census data.
package census

// Resident represents a resident in this city.
type Resident struct {
	Name    string
	Age     int
	Address map[string]string
}

// NewResident registers a new resident in this city.
func NewResident(name string, age int, address map[string]string) *Resident {
	result := Resident{Name: name, Age: age, Address: address}

	return &result
}

// HasRequiredInfo determines if a given resident has all of the required information.
func (r *Resident) HasRequiredInfo() bool {
	hasAddress := r.Address != nil && r.Address["street"] != ""

	return hasAddress && r.Name != ""
}

// Delete deletes a resident's information.
func (r *Resident) Delete() {
	r.Address = nil
	r.Age = 0
	r.Name = ""
}

// Count counts all residents that have provided the required information.
func Count(residents []*Resident) int {
	count := 0

	for _, resident := range residents {
		if resident.HasRequiredInfo() {
			count++
		}
	}

	return count
}
