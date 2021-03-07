<?php

namespace classes\writer;

/**
 * Class WriteFileMatrixes
 * @package classes\writer
 */
final class WriteFileMatrixes implements WriterInterface
{
    /**
     *
     */
    private const MATRIX_SERPARATOR = '---';

    /**
     * @var string
     */
    private string $fileName;

    /**
     * ReadFile constructor.
     * @param string $fileName
     */
    public function __construct(string $fileName)
    {
        $this->fileName = $fileName;
    }

    /**
     * @param array $matrixes
     * @throws \Exception
     */
    public function Write(array $matrixes): void
    {
        if (($handle = fopen($this->fileName, 'w')) == false) {
            throw new \Exception(sprintf('Open csv-file is failing: %s', $this->fileName));
        }

        foreach ($matrixes as $idx => $matrix) {
            if ($idx) {
                if (!fputs($handle, self::MATRIX_SERPARATOR . PHP_EOL)) {
                    throw new \Exception(sprintf('Write serparator to csv-file is failing: %s', $this->fileName));
                }
            }

            foreach ($matrix as $row) {
                if (!fputcsv($handle, $row)) {
                    throw new \Exception(sprintf('Write to csv-file is failing: %s', $this->fileName));
                }
            }
        }

        if (!fclose($handle)) {
            throw new \Exception(sprintf('Close csv-file is failing: %s', $this->fileName));
        }
    }
}