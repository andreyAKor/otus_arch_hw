<?php

namespace classes\reader;

/**
 * Class ReadFileMatrixes
 * @package classes\reader
 */
final class ReadFile implements ReaderInterface
{
    /**
     * @var string
     */
    private $fileName = '';

    /**
     * ReadFileMatrixes constructor.
     * @param string $fileName
     */
    public function __construct(string $fileName)
    {
        $this->fileName = $fileName;
    }

    /**
     * @return array
     */
    public function Read(): array
    {
        if (($handle = fopen($this->fileName, 'r')) == false) {
            throw new \Exception(sprintf('Open csv-file is failing: %s', $this->fileName));
        }

        $matrix = [];
        while (($data = fgetcsv($handle)) !== false) {
            $matrix[] = array_map(function (string $item) {
                return trim($item);
            }, $data);
        }

        if (!fclose($handle)) {
            throw new \Exception(sprintf('Close csv-file is failing: %s', $this->fileName));
        }

        return $matrix;
    }
}