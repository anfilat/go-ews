package ewsType

import "github.com/anfilat/go-ews/wsdl"

type GetUserAvailabilityResults struct {
	FreeBusyResponseArray *wsdl.ArrayOfFreeBusyResponse
	SuggestionsResponse   *wsdl.SuggestionsResponseType
}
