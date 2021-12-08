package exchangeVersion

// Enum defines the each available Exchange release version
type Enum int

func (e Enum) String() string {
	switch e {
	case Exchange2007SP1:
		return "Exchange2007_SP1"
	case Exchange2010:
		return "Exchange2010"
	case Exchange2010SP1:
		return "Exchange2010_SP1"
	case Exchange2010SP2:
		return "Exchange2010_SP2"
	case Exchange2013:
		return "Exchange2013"
	case Exchange2013SP1:
		return "Exchange2013_SP1"
	case Exchange2015:
		return "Exchange2015"
	case Exchange2016:
		return "Exchange2016"
	case V20151005:
		return "V20151005"
	}
	return ""
}

const (
	// Exchange2007SP1 defines Microsoft Exchange 2007, Service Pack 1
	Exchange2007SP1 Enum = iota
	// Exchange2010 defines Microsoft Exchange 2010
	Exchange2010
	// Exchange2010SP1 defines Microsoft Exchange 2010, Service Pack 1
	Exchange2010SP1
	// Exchange2010SP2 defines Microsoft Exchange 2010, Service Pack 2
	Exchange2010SP2
	// Exchange2013 defines Microsoft Exchange 2013
	Exchange2013
	// Exchange2013SP1 defines Microsoft Exchange 2013 SP1
	Exchange2013SP1
	// Exchange2015 defines Microsoft Exchange 2015 (aka Exchange 2016)
	Exchange2015
	// Exchange2016 defines Microsoft Exchange 2016
	Exchange2016
	// V20151005 defines Functionality starting 10/05/2015
	V20151005
)
