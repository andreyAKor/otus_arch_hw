<?php

namespace classes\matrix;

use classes\writer\WriterInterface;
use Mockery;

/**
 * Class MatrixGenaratorTest
 * @package Classes\Matrix
 */
class MatrixGenaratorTest extends \Codeception\Test\Unit
{
    /**
     *
     */
    public function testRun()
    {
        $matrix = [
            [1, 2, 3],
            [4, 5, 6],
            [7, 8, 9]
        ];

        $writer = Mockery::mock(WriterInterface::class);
        $writer->allows()
            ->Write([
                $matrix,
                $matrix,
            ]);

        $matrixGenarator = Mockery::mock(MatrixGenarator::class)->makePartial();
        $matrixGenarator->__construct($writer);
        $matrixGenarator->shouldReceive('Generate')
            ->twice()
            ->andReturn($matrix);

        $matrixGenarator->Run();
    }
}