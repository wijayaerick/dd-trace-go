// Unless explicitly stated otherwise all files in this repository are licensed
// under the Apache License Version 2.0.
// This product includes software developed at Datadog (https://www.datadoghq.com/).
// Copyright 2025 Datadog, Inc.

// Package telemetrytest provides a mock implementation of the telemetry client for testing purposes
package telemetrytest

import (
	"gopkg.in/DataDog/dd-trace-go.v1/internal/telemetry"

	"github.com/stretchr/testify/mock"
)

// MockClient implements Client and is used for testing purposes outside the telemetry package,
// e.g. the tracer and profiler.
type MockClient struct {
	mock.Mock
}

func (m *MockClient) Close() error {
	return nil
}

type MockMetricHandle struct {
	mock.Mock
}

func (m *MockMetricHandle) Submit(value float64) {
	m.Called(value)
}

func (m *MockMetricHandle) Get() float64 {
	return m.Called().Get(0).(float64)
}

func (m *MockClient) Count(namespace telemetry.Namespace, name string, tags []string) telemetry.MetricHandle {
	return m.Called(namespace, name, tags).Get(0).(telemetry.MetricHandle)
}

func (m *MockClient) Rate(namespace telemetry.Namespace, name string, tags []string) telemetry.MetricHandle {
	return m.Called(namespace, name, tags).Get(0).(telemetry.MetricHandle)
}

func (m *MockClient) Gauge(namespace telemetry.Namespace, name string, tags []string) telemetry.MetricHandle {
	return m.Called(namespace, name, tags).Get(0).(telemetry.MetricHandle)
}

func (m *MockClient) Distribution(namespace telemetry.Namespace, name string, tags []string) telemetry.MetricHandle {
	return m.Called(namespace, name, tags).Get(0).(telemetry.MetricHandle)
}

func (m *MockClient) Log(level telemetry.LogLevel, text string, options ...telemetry.LogOption) {
	m.Called(level, text, options)
}

func (m *MockClient) ProductStarted(product telemetry.Namespace) {
	m.Called(product)
}

func (m *MockClient) ProductStopped(product telemetry.Namespace) {
	m.Called(product)
}

func (m *MockClient) ProductStartError(product telemetry.Namespace, err error) {
	m.Called(product, err)
}

func (m *MockClient) RegisterAppConfig(key string, value any, origin telemetry.Origin) {
	m.Called(key, value, origin)
}

func (m *MockClient) RegisterAppConfigs(kvs ...telemetry.Configuration) {
	m.Called(kvs)
}

func (m *MockClient) MarkIntegrationAsLoaded(integration telemetry.Integration) {
	m.Called(integration)
}

func (m *MockClient) Flush() {
	m.Called()
}

func (m *MockClient) AppStart() {
	m.Called()
}

func (m *MockClient) AppStop() {
	m.Called()
}

var _ telemetry.Client = (*MockClient)(nil)
