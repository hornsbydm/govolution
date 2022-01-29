package mqtt

type UnitofMeasure uint8

const (
	DEGF UnitofMeasure = iota
	DEGC
	PERC
	CFM
	INH2O
	RPM
	OTHER
)

func (u UnitofMeasure) String() string {

	switch u {
	case DEGF:
		return "Degrees Farenheight (F)"
	case DEGC:
		return "Degrees Celcius (C)"
	case PERC:
		return "Percent (%)"
	case CFM:
		return "Cubic Feet Per Minute (CFM)"
	case INH2O:
		return "Inches of Water Column (inH2O)"
	case RPM:
		return "Rotations per Minute (RPM)"
	}

	return ""
}
