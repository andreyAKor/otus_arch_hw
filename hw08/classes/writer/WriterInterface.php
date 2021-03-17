<?php

namespace classes\writer;

/**
 * Interface WriterInterface
 * @package classes\writer
 */
interface WriterInterface
{
    /**
     * @param string $fileName
     * @param array $strings
     */
    public function Write(string $fileName, array $strings): void;
}
