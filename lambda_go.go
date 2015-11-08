package main

import (
	"encoding/json"
	"github.com/jasonmoo/lambda_proc"
)

// Struct for the Intent Schema
type AlexaSkillEvent struct {
	Session struct {
		Sessionid   string `json:"sessionId"`
		Application struct {
			Applicationid string `json:"applicationId"`
		} `json:"application"`
		Attributes interface{} `json:"attributes"`
		User       struct {
			Userid      string      `json:"userId"`
			Accesstoken interface{} `json:"accessToken"`
		} `json:"user"`
		New bool `json:"new"`
	} `json:"session"`
	Request struct {
		Type      string `json:"type"`
		Requestid string `json:"requestId"`
		Timestamp int64  `json:"timestamp"`
		Intent    struct {
			Name  string `json:"name"`
			Slots struct {
				Action struct {
					Name  string `json:"name"`
					Value string `json:"value"`
				} `json:"Action"`
				Name struct {
					Name  string `json:"name"`
					Value string `json:"value"`
				} `json:"Name"`
			} `json:"slots"`
		} `json:"intent"`
		Reason interface{} `json:"reason"`
	} `json:"request"`
}

// Struct for the resonse to Alexa
type AlexaResponse struct {
	Version  string `json:"version,omitempty"`
	Response struct {
		Outputspeech struct {
			Type string `json:"type"`
			Text string `json:"text"`
		} `json:"outputSpeech"`
		Reprompt         interface{} `json:"reprompt,omitempty"`
		Shouldendsession bool        `json:"shouldEndSession"`
	} `json:"response"`
	Sessionattributes struct {
	} `json:"sessionAttributes,omitempty"`
}

func main() {
	// Handle the incoming request from Alexa on Lambda from Node.js
	lambda_proc.Run(func(context *lambda_proc.Context, eventJSON json.RawMessage) (interface{}, error) {
		event := &AlexaSkillEvent{}
		json.Unmarshal(eventJSON, event)
		return processRequest(event), nil
	})
}

func processRequest(event *AlexaSkillEvent) *AlexaResponse {
	text := ""
	switch event.Request.Intent.Name {
	case "Command":
		action := event.Request.Intent.Slots.Action.Value
		name := event.Request.Intent.Slots.Name.Value
		text = "You told " + name + " to " + action
	default:
		text = "I have no idea what you just said, could you speak more clearly already"
	}
	return generateAlexaResponse(text)
}

func generateAlexaResponse(text string) *AlexaResponse {
	response := &AlexaResponse{}
	response.Version = "1.0"
	response.Response.Outputspeech.Type = "PlainText"
	response.Response.Shouldendsession = true
	response.Response.Outputspeech.Text = text
	return response
}
