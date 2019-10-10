<?php

namespace Bank;

require __DIR__ . '/../vendor/autoload.php';

$count = new Projection\Count();
$eventStore = new EventStore(
    $count
);
$eventStore->replay(__DIR__ . '/../../data/events.json');

print('Total Events: ' . $count->eventCount . PHP_EOL);
