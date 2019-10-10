<?php

namespace Bank;

use Bank\Projection\Projection;

class EventStore
{
    /**
     * @var Projection[]
     */
    private $projections;

    public function __construct(Projection ...$projections)
    {
        $this->projections = $projections;
    }

    public function replay(string $filePath)
    {
        $jsonStream = \JsonMachine\JsonMachine::fromFile($filePath);

        foreach ($jsonStream as $data) {
            $this->project($data);
        }
    }

    private function project($event)
    {
        foreach ($this->projections as $projection) {
            $projection->accept($event);
        }
    }
}
