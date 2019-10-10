<?php

namespace Bank\Projection;

class Count implements Projection
{
    public $eventCount;

    public function accept($event)
    {
        $this->eventCount++;

        switch ($event['type']) {
            case 'MoneyWasDeposited':
            case 'MoneyWasWithdrawn':
        }
    }
}
