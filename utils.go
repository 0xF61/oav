package main

import (
	"fmt"
	"log"
	"os"
	"strings"

	"github.com/AlienVault-OTX/OTX-Go-SDK/src/otxapi"
	"github.com/VirusTotal/vt-go"
)

// Look for file hashe's on vt and update indicator description
func CheckPulseDetail(pd *otxapi.PulseDetail) {
	indics := pd.Indicators

	// checkedVTs := 0
	vtKEY := os.Getenv("VT_API_KEY")
	if vtKEY == "" {
		return
	}

	for i, k := range indics {
		if !strings.Contains(*k.Type, "FileHash") {
			continue
		}

		// Generally good people have some type of duplicate hashes here
		// Don't waste our limited API calls for found file with other 2 hashes
		// SHA256 of ********************************
		// Sample Description is like above for those ones, we can skip I guess
		if strings.Contains(*k.Description, " of ") {
			continue
		}

		vtclient := vt.NewClient(vtKEY)
		file, err := vtclient.GetObject(vt.URL("files/" + *k.Indicator))
		if err != nil {
			log.Println(err)
			if err.Error() == "Quota exceeded" {
				*pd.Indicators[i].Description = err.Error()
				return
			}
			continue
		}

		lsd, err := file.GetTime("last_submission_date")
		if err != nil {
			fmt.Printf("lsd err: %v\n", err)
		}

		*pd.Indicators[i].Description = lsd.String()

		// checkedVTs++
	}
}
