<?php

namespace classes\reader;

/**
 * Class ReadFile
 * @package classes\reader
 */
final class ReadFile implements ReaderInterface
{
    /**
     * @param string $fileName
     * @return array
     * @throws \Exception
     */
    public function Read(string $fileName): array
    {
        if (($handle = fopen($fileName, 'r')) == false) {
            throw new \Exception(sprintf('Open file is failing: %s', $fileName));
        }

        $strings = [];
        while (($data = fgets($handle)) !== false) {
            $strings[] = $data;
        }

        if (!fclose($handle)) {
            throw new \Exception(sprintf('Close file is failing: %s', $fileName));
        }

        return $strings;
    }
}