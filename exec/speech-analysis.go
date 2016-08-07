package exec

import (
	//	log "github.com/inconshreveable/log15"
	"bytes"
	"encoding/json"
	"gopkg.in/pipe.v2"
	"io"
	"os"
	"time"
)

type SAArguments struct {
type SA_Arguments struct {
	FileOUTPATH string
	SAScriptLOC string
	FileOUT     bool
	SpeechREQ   io.ReadCloser
}

func returnSpeechAnalysis(a *SA_Arguments) (result string, err error) {
	buffer := new(bytes.Buffer)
	filename := fmt.Sprintf("%s/speechoutput_%s.json", a.FileOUTPATH,
		time.Now().Format(time.RFC3339))

	file, err := os.Create(filename)
	if err != nil {
		return "", err
	}
	defer file.Close()

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

// SAReturnASYS takes the SA_Arguments struct, and returns the analysis
// from the speech analysis program
func SAReturnASYS(a *SA_Arguments) (result string, err error) {

	result = ""
	err = nil

	if a.FileOUT == true {
		result, err = returnSpeechAnalysis(a)
	}

	return result, err

}
