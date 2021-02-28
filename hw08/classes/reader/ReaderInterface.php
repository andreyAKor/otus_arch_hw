<?php

namespace classes\reader;

/**
 * Interface ReaderInterface
 * @package classes\reader
 */
interface ReaderInterface
{
    /**
     * @param string $fileName
     * @return array
     */
    public function Read(string $fileName): array;
}
