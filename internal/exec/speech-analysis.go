package exec

import (
	"bytes"
	"encoding/json"
	"fmt"
	log "github.com/inconshreveable/log15"
	pipe "gopkg.in/pipe.v2"
	"strings"
	"time"
)

const defaultSpeechAnalyseScriptLocation = "/docker/analyse"

type SpeechAnalysisArgs struct {
	input          string // Inputted speech for script.
	scriptLocation string // To be populated via `config` module.
	debugMode      bool   // Output speech analysis to `stdout`. (TODO: See L51)
}

type AnalyzedSpeech struct {
	analysis    string        // Returned JSON from Ben's script.
	processTime time.Duration // Used for metrics on status page.
}

func AnalyzeSpeech(args *SpeechAnalysisArgs) (AnalyzedSpeech, error) {
	buf := new(bytes.Buffer)
	var err error = nil
	var scriptLocation string

	// Check if script path is defined.
	if len(strings.TrimSpace(args.scriptLocation)) == 0 {
		// Assign default script path.
		scriptLocation = defaultSpeechAnalyseScriptLocation
	} else {
		// Otherwise trust args.
		scriptLocation = args.scriptLocation
	}

	// Start timing.
	start := time.Now()

	// Begin construction of pipe.
	pipeline := pipe.Line(
		pipe.Read(strings.NewReader(args.input)),
		pipe.Exec(scriptLocation),
		pipe.Tee(buf))

	// Run pipe, assign to declared `err` var
	err = pipe.Run(pipeline)
	if err != nil {
		// Should we let the controller handle errors, and return a 5xx code?
		log.Error("Could not run speech analysis pipeline.",
			log.Ctx{"error": err.Error()})
		return AnalyzedSpeech{}, err
	}

	if args.debugMode { // TODO: Refactor to remove `debugMode`, use logging level arg in ENV. Output slog-style.
		fmt.Printf("JSON output from speech analysis request: %s\n", buf.String())
	}

	analyzedSpeech := new(AnalyzedSpeech)
	compacted := new(bytes.Buffer)
	if err := json.Compact(compacted, buf.Bytes()); err != nil {
		return AnalyzedSpeech{}, err
	}
	analyzedSpeech.analysis = compacted.String()
	analyzedSpeech.processTime = time.Since(start) // TODO: Maybe directly report to metrics, and associate with request ID?

	return *analyzedSpeech, nil
}
