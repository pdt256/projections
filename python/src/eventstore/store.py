import json


class EventStore:

    projections: list

    def __init__(self, projections):
        self.projections = [projections]

    def replay(self, filepath):
        with open(filepath) as event_file:
            json_data = json.load(event_file)
        for data in json_data:
            self.project(data)

    def project(self, event):
        for projection in self.projections:
            projection.accept(event)
