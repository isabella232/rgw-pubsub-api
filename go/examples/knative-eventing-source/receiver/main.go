/*
Copyright 2018 The Knative Authors

Licensed under the Apache License, Version 2.0 (the "License");
you may not use this file except in compliance with the License.
You may obtain a copy of the License at

    https://www.apache.org/licenses/LICENSE-2.0

Unless required by applicable law or agreed to in writing, software
distributed under the License is distributed on an "AS IS" BASIS,
WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
See the License for the specific language governing permissions and
limitations under the License.
*/

package main

import (
	"context"
	"log"
	"net/http"
	"time"

	"github.com/knative/pkg/cloudevents"

	"github.com/ceph/rgw-pubsub-api/go/pkg"
)

func handler(ctx context.Context, e *rgwpubsub.RGWEvent) {
	metadata := cloudevents.FromContext(ctx)
	log.Printf("[%s] %s %s. Object: %q  Bucket: %q", metadata.EventTime.Format(time.RFC3339), metadata.ContentType, metadata.Source, e.Info.Key.Name, e.Info.Bucket.Name)
}

func main() {
	log.Print("Ready and listening on port 8080")
	log.Fatal(http.ListenAndServe(":8080", cloudevents.Handler(handler)))
}
