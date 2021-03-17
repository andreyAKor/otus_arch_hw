<?php

namespace classes\matrix;

use classes\reader\ReaderInterface;
use classes\writer\WriterInterface;

/**
 * Class MatrixAddition
 * @package classes\matrix
 */
final class MatrixAddition implements MatrixOperationInterface
{
    /**
     * @var ReaderInterface
     */
    private ReaderInterface $reader;

    /**
     * @var WriterInterface
     */
    private WriterInterface $writer;

    /**
     * MatrixAddition constructor.
     * @param ReaderInterface $reader
     * @param WriterInterface $writer
     */
    public function __construct(ReaderInterface $reader, WriterInterface $writer)
    {
        $this->reader = $reader;
        $this->writer = $writer;
    }

    /**
     *
     */
    public function Run(): void
    {
        $matrixes = $this->reader->Read();
        $this->writer->Write($this->Add($matrixes[0], $matrixes[1]));
    }

    /**
     * @param array $matrix1
     * @param array $matrix2
     * @return array
     */
    public function Add(array $matrix1, array $matrix2): array
    {
        foreach ($matrix1 as $i => $row) {
            foreach ($row as $j => $elem) {
                $matrix1[$i][$j] = $elem + $matrix2[$i][$j];
            }
        }

        return $matrix1;
    }
}
