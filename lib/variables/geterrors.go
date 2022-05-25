package variables


func GetErrors(Package string) ([]byte, int) {
	if Package == "all" {
		return JsonMarshal(Errors), 200
	}
	var Logs []Log
	for _, Log := range Errors {
		if Log.Package == Package {
			Logs = append(Logs, Log)
		}
	}
	return JsonMarshal(Logs), 200
}