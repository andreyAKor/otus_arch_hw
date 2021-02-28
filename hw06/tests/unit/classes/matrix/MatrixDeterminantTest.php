<?php

namespace classes\matrix;

use classes\reader\ReaderInterface;
use classes\writer\WriterInterface;
use Mockery;

/**
 * Class MatrixDeterminantTest
 * @package Classes\Matrix
 */
class MatrixDeterminantTest extends \Codeception\Test\Unit
{
    /**
     *
     */
    public function testDo()
    {
        $reader = Mockery::mock(ReaderInterface::class);
        $writer = Mockery::mock(WriterInterface::class);

        $matrixDeterminant = new MatrixDeterminant($reader, $writer);

        $pairs = [
            [
                'matrix' => [
                    [1, 2, 3],
                    [3, 2, 1],
                    [4, 5, 10],
                ],
                'determinat' => [-16],
            ],
            [
                'matrix' => [
                    [1, 2, 3],
                    [4, 5, 6],
                    [7, 8, 9],
                ],
                'determinat' => [0],
            ],
            [
                'matrix' => [
                    [1, 2],
                    [3, 4],
                ],
                'determinat' => [-2],
            ],
            [
                'matrix' => [
                    [1],
                ],
                'determinat' => [1],
            ],
        ];

        foreach ($pairs as $pair) {
            $matrixActual = $matrixDeterminant->Do($pair['matrix']);
            $this->assertEquals($pair['determinat'], $matrixActual);
        }
    }
}