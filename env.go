package uausgosdk

import "os"

var (
	URL        = os.Getenv("UAUS_URL")
	SERVICE_ID = os.Getenv("SERVICE_ID")
	API_KEY    = os.Getenv("API_KEY")
)
