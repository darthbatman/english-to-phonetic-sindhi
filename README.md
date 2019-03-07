# english-to-phonetic-sindhi
Tool that converts user-inputted English text to its Phonetic Sindhi translation displayed with English letters.

## Usage

`go build app.go`

`./app "This is some text I want translated."`

This tool uses the Google Cloud Translation API.

Make sure you have set up `gcloud` and set the path to the Google Cloud API Credentials using `export GOOGLE_APPLICATION_CREDENTIALS="PATH"`