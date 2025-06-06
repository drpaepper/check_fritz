package fritz

// SoapData is the data structure for a SOAP request including the response
type SoapData struct {
	Username     []byte
	Password     []byte
	URL          string
	URLPath      string
	Action       string
	Service      string
	ServiceIndex string
	XMLVariable  SoapDataVariable
}

type FritzServerCredentials struct {
	Username  string
	Password  string
	URL       string
	TR064Port string
}

// SoapDataVariable is the data structure for a variable that can injected in the SOAP request
type SoapDataVariable struct {
	Name  string
	Value string
}

// AddSoapDataVariable adds a SoapRequestVariable to a SoapRequest
func (soapData *SoapData) AddSoapDataVariable(soapDataVariable SoapDataVariable) {
	soapData.XMLVariable = soapDataVariable
}

// CreateNewSoapData creates a new FritzSoapRequest structure
func CreateNewSoapData(username string, password string, hostname string, port string, e Endpoint) SoapData {
	var fSR SoapData

	fSR.URL = "https://" + hostname + ":" + port + e.url

	fSR.URLPath = e.url
	fSR.Username = []byte(username)
	fSR.Password = []byte(password)
	fSR.Action = e.action
	fSR.Service = e.service
	fSR.ServiceIndex = e.serviceIndex

	return fSR
}

// CreateNewSoapVariable creates a new SoapRequestVariable
func CreateNewSoapVariable(name string, value string) SoapDataVariable {
	var soapDataVariable SoapDataVariable

	soapDataVariable.Name = name
	soapDataVariable.Value = value

	return soapDataVariable
}

func GetDataFromEndpoint(s FritzServerCredentials, e Endpoint) [][]byte {
	resps := make(chan []byte)
	errs := make(chan error)

	soapReq := CreateNewSoapData(s.Username, s.Password, s.URL, s.TR064Port, e)
	go DoSoapRequest(&soapReq, resps, errs, false)

	resp, _ := ProcessSoapResponse(resps, errs, 1, 30)
	return resp
}
