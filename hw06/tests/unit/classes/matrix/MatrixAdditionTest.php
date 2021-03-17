<?php

namespace classes\matrix;

use classes\reader\ReaderInterface;
use classes\writer\WriterInterface;
use Mockery;

/**
 * Class MatrixGenaratorTest
 * @package Classes\Matrix
 */
class MatrixAdditionTest extends \Codeception\Test\Unit
{
    /**
     *
     */
    public function testDo()
    {
        $reader = Mockery::mock(ReaderInterface::class);
        $writer = Mockery::mock(WriterInterface::class);

        $matrixAddition = new MatrixGenarator($reader, $writer);

        $matrixIn = [
            [1, 2, 3],
            [4, 5, 6],
            [7, 8, 9],
        ];
        $matrixActual = $matrixAddition->Do($matrixIn);

        $expectedMatrix = [
            [2, 4, 6],
            [8, 10, 12],
            [14, 16, 18],
        ];
        $this->assertEquals($expectedMatrix, $matrixActual);
    }
}