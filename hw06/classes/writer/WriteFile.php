<?php

namespace classes\writer;

/**
 * Class WriteFile
 * @package classes\writer
 */
final class WriteFile implements WriterInterface
{
    /**
     * @var string
     */
    private $fileName = '';

    /**
     * ReadFile constructor.
     * @param string $fileName
     */
    public function __construct(string $fileName)
    {
        $this->fileName = $fileName;
    }

    /**
     * @param string $content
     */
    public function Write(array $matrix): void
    {
        if (($handle = fopen($this->fileName, 'w')) == false) {
            throw new \Exception(sprintf('Open csv-file is failing: %s', $this->fileName));
        }

        if (count($matrix) > 1) {
            foreach ($matrix as $row) {
                if (!fputcsv($handle, $row)) {
                    throw new \Exception(sprintf('Write to csv-file is failing: %s', $this->fileName));
                }
            }
        } else {
            if (!fputcsv($handle, $matrix)) {
                throw new \Exception(sprintf('Write to csv-file is failing: %s', $this->fileName));
            }
        }

        if (!fclose($handle)) {
            throw new \Exception(sprintf('Close csv-file is failing: %s', $this->fileName));
        }
    }
}