<?php

namespace classes\matrix;

use classes\reader\ReaderInterface;
use classes\writer\WriterInterface;

/**
 * Class AbstractTemplate
 * @package classes\matrix
 */
abstract class AbstractTemplate
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
     * AbstractTemplate constructor.
     * @param ReaderInterface $reader
     * @param WriterInterface $writer
     * @param string $operation
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
        $this->writer->Write($this->Do($this->reader->Read()));
    }

    /**
     * @param array $matrix
     * @return array
     */
    abstract public function Do(array $matrix): array;
}
