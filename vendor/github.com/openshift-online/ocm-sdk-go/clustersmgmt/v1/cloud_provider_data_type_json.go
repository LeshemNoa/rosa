/*
Copyright (c) 2020 Red Hat, Inc.

Licensed under the Apache License, Version 2.0 (the "License");
you may not use this file except in compliance with the License.
You may obtain a copy of the License at

  http://www.apache.org/licenses/LICENSE-2.0

Unless required by applicable law or agreed to in writing, software
distributed under the License is distributed on an "AS IS" BASIS,
WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
See the License for the specific language governing permissions and
limitations under the License.
*/

// IMPORTANT: This file has been generated automatically, refrain from modifying it manually as all
// your changes will be lost when the file is generated again.

package v1 // github.com/openshift-online/ocm-sdk-go/clustersmgmt/v1

import (
	"io"
	"net/http"

	jsoniter "github.com/json-iterator/go"
	"github.com/openshift-online/ocm-sdk-go/helpers"
)

// MarshalCloudProviderData writes a value of the 'cloud_provider_data' type to the given writer.
func MarshalCloudProviderData(object *CloudProviderData, writer io.Writer) error {
	stream := helpers.NewStream(writer)
	writeCloudProviderData(object, stream)
	err := stream.Flush()
	if err != nil {
		return err
	}
	return stream.Error
}

// writeCloudProviderData writes a value of the 'cloud_provider_data' type to the given stream.
func writeCloudProviderData(object *CloudProviderData, stream *jsoniter.Stream) {
	count := 0
	stream.WriteObjectStart()
	var present_ bool
	present_ = object.bitmap_&1 != 0 && object.aws != nil
	if present_ {
		if count > 0 {
			stream.WriteMore()
		}
		stream.WriteObjectField("aws")
		writeAWS(object.aws, stream)
		count++
	}
	present_ = object.bitmap_&2 != 0 && object.gcp != nil
	if present_ {
		if count > 0 {
			stream.WriteMore()
		}
		stream.WriteObjectField("gcp")
		writeGCP(object.gcp, stream)
		count++
	}
	present_ = object.bitmap_&4 != 0
	if present_ {
		if count > 0 {
			stream.WriteMore()
		}
		stream.WriteObjectField("key_location")
		stream.WriteString(object.keyLocation)
		count++
	}
	present_ = object.bitmap_&8 != 0
	if present_ {
		if count > 0 {
			stream.WriteMore()
		}
		stream.WriteObjectField("key_ring_name")
		stream.WriteString(object.keyRingName)
		count++
	}
	present_ = object.bitmap_&16 != 0 && object.region != nil
	if present_ {
		if count > 0 {
			stream.WriteMore()
		}
		stream.WriteObjectField("region")
		writeCloudRegion(object.region, stream)
	}
	stream.WriteObjectEnd()
}

// UnmarshalCloudProviderData reads a value of the 'cloud_provider_data' type from the given
// source, which can be an slice of bytes, a string or a reader.
func UnmarshalCloudProviderData(source interface{}) (object *CloudProviderData, err error) {
	if source == http.NoBody {
		return
	}
	iterator, err := helpers.NewIterator(source)
	if err != nil {
		return
	}
	object = readCloudProviderData(iterator)
	err = iterator.Error
	return
}

// readCloudProviderData reads a value of the 'cloud_provider_data' type from the given iterator.
func readCloudProviderData(iterator *jsoniter.Iterator) *CloudProviderData {
	object := &CloudProviderData{}
	for {
		field := iterator.ReadObject()
		if field == "" {
			break
		}
		switch field {
		case "aws":
			value := readAWS(iterator)
			object.aws = value
			object.bitmap_ |= 1
		case "gcp":
			value := readGCP(iterator)
			object.gcp = value
			object.bitmap_ |= 2
		case "key_location":
			value := iterator.ReadString()
			object.keyLocation = value
			object.bitmap_ |= 4
		case "key_ring_name":
			value := iterator.ReadString()
			object.keyRingName = value
			object.bitmap_ |= 8
		case "region":
			value := readCloudRegion(iterator)
			object.region = value
			object.bitmap_ |= 16
		default:
			iterator.ReadAny()
		}
	}
	return object
}
