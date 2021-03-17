<?php

namespace classes\reader;

/**
 * Class ReadFileMatrixes
 * @package classes\reader
 */
final class ReadFileMatrixes implements ReaderInterface
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

        $matrixes = [[]];
        $matrixIterator = 0;

        while (($data = fgetcsv($handle)) !== false) {
            if (isset($data[0]) && trim($data[0]) == self::MATRIX_SERPARATOR) {
                $matrixIterator++;
                $matrixes[$matrixIterator] = [];

                continue;
            }

            $matrixes[$matrixIterator][] = array_map(function (string $item) {
                return trim($item);
            }, $data);
        }

        if (!fclose($handle)) {
            throw new \Exception(sprintf('Close csv-file is failing: %s', $this->fileName));
        }

        return $matrixes;
    }
}