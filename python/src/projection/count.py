from typing import List
from projection.projection import Projection


class Count(Projection):
    event_count: int
    event_types: List

    def __init__(self):
        self.event_count = 0
        self.event_types = {
            'MoneyWasDeposited': 'do_nothing',
            'MoneyWasWithdrawn': 'do_nothing'
        }

    def accept(self, event):

        self.event_count = self.event_count + 1

        type = event.get('type')

        if type in self.event_types:

            event_callable = self.event_types[type]
            getattr(self, event_callable)()

    def do_nothing(self):
        pass
