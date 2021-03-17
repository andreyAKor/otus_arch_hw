<?php

namespace classes\matrix;

use classes\writer\WriterInterface;

/**
 * Class MatrixGenarator
 * @package classes\matrix
 */
class MatrixGenarator
{
    /**
     * @var WriterInterface
     */
    private WriterInterface $writer;

    /**
     * MatrixGenarator constructor.
     * @param WriterInterface $writer
     */
    public function __construct(WriterInterface $writer)
    {
        $this->writer = $writer;
    }

    /**
     *
     */
    public function Run(): void
    {
        $this->writer->Write([
            $this->Generate(5, 0, 10),
            $this->Generate(5, 0, 10),
        ]);
    }

    /**
     * @param int $size
     * @param int $min
     * @param int $max
     * @return array
     * @throws \Exception
     */
    public function Generate(int $size, int $min, int $max): array
    {
        $matrix = [];

        for ($i = 0; $i < $size; $i++) {
            $matrix[$i] = [];

            for ($j = 0; $j < $size; $j++) {
                $matrix[$i][$j] = random_int($min, $max);
            }
        }

        return $matrix;
    }
}
