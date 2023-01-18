package extra

var ERRORS = map[string]string{
	"CORE:001": "ERROR: Wrong OP COUNT in Simulation mode.",
	"CORE:002": "ERROR: Can not process OP as STACK length is less than 2.",
	"CORE:003": "ERROR: Are you seriously trying to divide on 0?",
}	

func ToThrow(errCode string) string {
	err := ERRORS[errCode]
	
	return err
}