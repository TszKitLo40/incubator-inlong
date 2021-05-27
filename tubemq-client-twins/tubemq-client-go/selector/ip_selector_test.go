/**
 * Licensed to the Apache Software Foundation (ASF) under one or more
 * contributor license agreements.  See the NOTICE file distributed with
 * this work for additional information regarding copyright ownership.
 * The ASF licenses this file to You under the Apache License, Version 2.0
 * (the "License"); you may not use this file except in compliance with
 * the License.  You may obtain a copy of the License at
 * <p>
 * http://www.apache.org/licenses/LICENSE-2.0
 * <p>
 * Unless required by applicable law or agreed to in writing, software
 * distributed under the License is distributed on an "AS IS" BASIS,
 * WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
 * See the License for the specific language governing permissions and
 * limitations under the License.
 */

package selector

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestSingleIP(t *testing.T) {
	serviceName := "192.168.0.1:9092"
	selector := Get("ip")
	node, err := selector.Select(serviceName)
	assert.Nil(t, err)
	assert.Equal(t, node.HasNext, false)
	assert.Equal(t, node.Address, "192.168.0.1:9092")
	assert.Equal(t, node.ServiceName, "192.168.0.1:9092")
}

func TestSingleDNS(t *testing.T) {
	serviceName := "tubemq:8081"
	selector := Get("dns")
	node, err := selector.Select(serviceName)
	assert.Nil(t, err)
	assert.Equal(t, node.HasNext, false)
	assert.Equal(t, node.Address, "tubemq:8081")
	assert.Equal(t, node.ServiceName, "tubemq:8081")
}

func TestMultipleIP(t *testing.T) {
	serviceName := "192.168.0.1:9091,192.168.0.1:9092,192.168.0.1:9093,192.168.0.1:9094"
	selector := Get("dns")
	node, err := selector.Select(serviceName)
	assert.Nil(t, err)
	assert.Equal(t, true, node.HasNext)
	assert.Equal(t, "192.168.0.1:9091", node.Address)
	assert.Equal(t, "192.168.0.1:9091,192.168.0.1:9092,192.168.0.1:9093,192.168.0.1:9094", node.ServiceName)

	node, err = selector.Select(serviceName)
	assert.Equal(t, true, node.HasNext)
	assert.Equal(t, "192.168.0.1:9092", node.Address)
	assert.Equal(t, "192.168.0.1:9091,192.168.0.1:9092,192.168.0.1:9093,192.168.0.1:9094", node.ServiceName)

	node, err = selector.Select(serviceName)
	assert.Equal(t, true, node.HasNext)
	assert.Equal(t, "192.168.0.1:9093", node.Address)
	assert.Equal(t, "192.168.0.1:9091,192.168.0.1:9092,192.168.0.1:9093,192.168.0.1:9094", node.ServiceName)

	node, err = selector.Select(serviceName)
	assert.Equal(t, false, node.HasNext)
	assert.Equal(t, "192.168.0.1:9094", node.Address)
	assert.Equal(t, "192.168.0.1:9091,192.168.0.1:9092,192.168.0.1:9093,192.168.0.1:9094", node.ServiceName)
}

func TestMultipleDNS(t *testing.T) {
	serviceName := "tubemq:8081,tubemq:8082,tubemq:8083,tubemq:8084"
	selector := Get("dns")
	node, err := selector.Select(serviceName)
	assert.Nil(t, err)
	assert.Equal(t, true, node.HasNext)
	assert.Equal(t, "tubemq:8081", node.Address)
	assert.Equal(t, "tubemq:8081,tubemq:8082,tubemq:8083,tubemq:8084", node.ServiceName)

	node, err = selector.Select(serviceName)
	assert.Equal(t, true, node.HasNext)
	assert.Equal(t, "tubemq:8082", node.Address)
	assert.Equal(t, "tubemq:8081,tubemq:8082,tubemq:8083,tubemq:8084", node.ServiceName)

	node, err = selector.Select(serviceName)
	assert.Equal(t, true, node.HasNext)
	assert.Equal(t, "tubemq:8083", node.Address)
	assert.Equal(t, "tubemq:8081,tubemq:8082,tubemq:8083,tubemq:8084", node.ServiceName)

	node, err = selector.Select(serviceName)
	assert.Equal(t, false, node.HasNext)
	assert.Equal(t, "tubemq:8084", node.Address)
	assert.Equal(t, "tubemq:8081,tubemq:8082,tubemq:8083,tubemq:8084", node.ServiceName)
}


func TestEmptyService(t *testing.T) {
	serviceName := ""
	selector := Get("ip")
	_, err := selector.Select(serviceName)
	assert.Error(t, err)
}