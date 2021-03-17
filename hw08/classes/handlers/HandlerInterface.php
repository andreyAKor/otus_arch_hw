<?php

namespace classes\handlers;

/**
 * Interface HandlerInterface
 * @package classes\handlers
 */
interface HandlerInterface
{
    /**
     * @param HandlerInterface $h
     */
    public function setNext(HandlerInterface $h): void;

    /**
     * @param string $fileName
     */
    public function handle(string $fileName): void;
}
