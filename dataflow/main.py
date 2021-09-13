#!/usr/bin/env python3

import json
import sys

import apache_beam as beam
from apache_beam import Map
from apache_beam.io import ReadFromPubSub
from apache_beam.io.gcp.datastore.v1new.datastoreio import WriteToDatastore
from apache_beam.io.gcp.datastore.v1new.types import Entity, Key
from apache_beam.options.pipeline_options import PipelineOptions

project_id = "pierre-test-321108"
collected_requests_subscription = "projects/{}/subscriptions/Result-sub".format(project_id)
result_table = "{}:result".format(project_id)


def create_entity(data):
    entity = Entity(key=Key(
        path_elements=['result', data['id']],
        project=project_id
    ))
    entity.set_properties({
        'WebhookID': data['webhook_id'],
        'Values': data['values'],
        'Result': data['result'],
        'Processed': True
    })
    return entity


def run():
    pipeline_options = PipelineOptions(streaming=True, save_main_session=True)

    with beam.Pipeline(options=pipeline_options) as p:
        (
                p

                | "Read from PubSub"
                >> ReadFromPubSub(subscription=collected_requests_subscription)
                .with_output_types(bytes)

                | "Data to utf-8"
                >> Map(lambda x: x.decode("utf-8"))

                | "Load json data"
                >> Map(json.loads)

                | "To Datastore Entity"
                >> Map(create_entity)

                | "Write to Datastore"
                >> WriteToDatastore(project_id)
        )


def main():
    run()
    return 0


if __name__ == "__main__":
    try:
        exit(main() or 0)
    except Exception:
        print("Unexpected error: ", sys.exc_info()[0])
        raise
