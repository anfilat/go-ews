package exchangeVersion

// Enum defines the each available Exchange release version
type Enum int

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
