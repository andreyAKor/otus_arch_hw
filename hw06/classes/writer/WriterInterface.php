<?php

namespace classes\writer;

/**
 * Interface WriterInterface
 * @package classes\writer
 */
interface WriterInterface
{
    /**
     * @param array $matrix
     */
    public function Write(array $matrix): void;
}
