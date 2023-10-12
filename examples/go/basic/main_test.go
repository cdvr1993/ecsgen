package main

import (
	"runtime"
	"testing"
	"time"
)

func BenchmarkGenerated(b *testing.B) {
	input := Base{
		AtTimestamp: time.Now(),
		Labels: map[string]interface{}{
			"foo": "bar",
		},
		Message: "message",
		Tags:    []string{"foo", "bar"},
		Agent: Agent{
			EphemeralID: "ephemeral_id",
			ID:          "id",
			Name:        "name",
			Type:        "Type",
			Version:     "version",
		},
		Client: Client{
			Address: "address",
		},
		Cloud: Cloud{
			Account: CloudAccount{
				ID: "id",
			},
		},
		Container: Container{
			ID: "id",
		},
		Destination: Destination{
			Address: "address",
		},
		DLL:           DLL{},
		DNS:           DNS{},
		ECS:           ECS{},
		Error:         Error{},
		Event:         Event{},
		File:          File{},
		Group:         Group{},
		Hash:          Hash{},
		Host:          Host{},
		HTTP:          HTTP{},
		Log:           Log{},
		Network:       Network{},
		Observer:      Observer{},
		Organization:  Organization{},
		Package:       Package{},
		Process:       Process{},
		Registry:      Registry{},
		Related:       Related{},
		Rule:          Rule{},
		Server:        Server{},
		Service:       Service{},
		Source:        Source{},
		Threat:        Threat{},
		TLS:           TLS{},
		Trace:         Trace{},
		Transaction:   Transaction{},
		URL:           URL{},
		User:          User{},
		UserAgent:     UserAgent{},
		Vulnerability: Vulnerability{},
	}
	b.ResetTimer()
	var in []byte
	var err error
	for i := 0; i < b.N; i++ {
		in, err = input.MarshalJSON()
		if err != nil {
			b.Fatal(err)
		}
	}
	//b.Logf("input: %s", in)
	runtime.KeepAlive(in)
}
