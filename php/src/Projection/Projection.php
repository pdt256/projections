<?php

namespace Bank\Projection;

interface Projection
{
    public function accept($event);
}
