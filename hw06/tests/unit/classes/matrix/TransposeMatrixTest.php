<?php

namespace classes\matrix;

use classes\reader\ReaderInterface;
use classes\writer\WriterInterface;
use Mockery;

/**
 * Class TransposeMatrixTest
 * @package Classes\Matrix
 */
class TransposeMatrixTest extends \Codeception\Test\Unit
{
    /**
     *
     */
    public function testDo()
    {
        $reader = Mockery::mock(ReaderInterface::class);
        $writer = Mockery::mock(WriterInterface::class);

        $transposeMatrix = new TransposeMatrix($reader, $writer);

        $matrixIn = [
            [1, 2, 3],
            [4, 5, 6],
            [7, 8, 9],
        ];
        $matrixActual = $transposeMatrix->Do($matrixIn);

        $expectedMatrix = [
            [1, 4, 7],
            [2, 5, 8],
            [3, 6, 9],
        ];
        $this->assertEquals($expectedMatrix, $matrixActual);
    }
}