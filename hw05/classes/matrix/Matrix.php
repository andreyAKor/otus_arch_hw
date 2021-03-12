<?php

namespace classes\matrix;

/**
 * Class Matrix
 * @package classes\matrix
 */
final class Matrix
{
    /**
     * @var MatrixOperationInterface
     */
    private MatrixOperationInterface $operation;

    /**
     * Matrix constructor.
     * @param MatrixOperationInterface $operation
     */
    public function __construct(MatrixOperationInterface $operation)
    {
        $this->operation = $operation;
    }

    /**
     *
     */
    public function Calc(): void
    {
        $this->operation->Run();
    }
}
