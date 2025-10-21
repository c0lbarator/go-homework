package api

type DecodeRequest struct {
	InputString string `json:"inputString"`
}

type DecodeResponse struct {
	OutputString string `json:"outputString"`
}

type VersionResponse struct {
	Version string `json:"version"`
}
