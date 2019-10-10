import sys
from os import getcwd

sys.path.insert(0, getcwd())

from eventstore.store import EventStore
from projection.count import Count

if __name__ == "__main__":
    count = Count()
    filepath = f'{getcwd()}/../../data/events.json'
    event_store = EventStore(count)
    event_store.replay(filepath)

    print(f'Total Events: {count.event_count}')
