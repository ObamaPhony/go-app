package exec

import (
	"bytes"
	"encoding/json"
	"fmt"

	"io"
	"io/ioutil"
	"os"
	"time"

	log "github.com/inconshreveable/log15"

	pipe "gopkg.in/pipe.v2"
)

type SAArguments struct {
	FileOUTPATH string
	SAScriptLOC string
	FileOUT     bool
	SpeechREQ   io.ReadCloser
}

func returnSpeechAnalysis(a *SAArguments) (result string, err error) {
	log.Debug("New bytes buffer created!")
	buffer := new(bytes.Buffer)

	tempdir, err := ioutil.TempDir("", "speechoutput_")
	if err != nil {
		return "", err
	}

	filename := fmt.Sprintf("%s/speechoutput_%s.json",
		tempdir,
		time.Now().Format(time.RFC3339))

	file, err := os.Create(filename)
	defer file.Close()
	if err != nil {
		return "", err
	}

	pi := pipe.Line(
		pipe.Read(a.SpeechREQ),
		pipe.Exec(a.SAScriptLOC),
		pipe.Tee(buffer),
	)

	err = pipe.Run(pi)
	if err != nil {
		return "", err
	}

	output := &bytes.Buffer{}
	if err := json.Compact(output, buffer.Bytes()); err != nil {
		return "", err
	}

	return output.String(), nil

}

// SAReturnASYS takes the SAArguments struct, and returns the analysis
// from the speech analysis program
func SAReturnASYS(a *SAArguments) (result string, err error) {
	if a.FileOUT == true {
		result, err := returnSpeechAnalysis(a)
		return result, err
	}

	return "", nil
}
