<?php

namespace classes\writer;

/**
 * Class WriteFile
 * @package classes\writer
 */
final class WriteFile implements WriterInterface
{
    /**
     * @param string $fileName
     * @param array $strings
     * @throws \Exception
     */
    public function Write(string $fileName, array $strings): void
    {
        if (($handle = fopen($fileName, 'a')) == false) {
            throw new \Exception(sprintf('Open file is failing: %s', $fileName));
        }

        foreach ($strings as $row) {
            if (!fputs($handle, $row)) {
                throw new \Exception(sprintf('Write to file is failing: %s', $fileName));
            }
        }

        if (!fclose($handle)) {
            throw new \Exception(sprintf('Close file is failing: %s', $fileName));
        }
    }
}