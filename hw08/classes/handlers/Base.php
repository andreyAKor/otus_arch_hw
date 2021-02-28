<?php

namespace classes\handlers;

use classes\reader\ReaderInterface;
use classes\writer\WriterInterface;

/**
 * Class Base
 * @package classes\handlers
 */
class Base implements HandlerInterface
{
    /**
     * @var HandlerInterface|null
     */
    private ?HandlerInterface $next = null;

    /**
     * @var ReaderInterface
     */
    protected ReaderInterface $reader;

    /**
     * @var WriterInterface
     */
    protected WriterInterface $writer;

    /**
     * @var string
     */
    protected string $outFileName;

    /**
     * Base constructor.
     * @param ReaderInterface $reader
     * @param WriterInterface $writer
     * @param string $outFileName
     */
    public function __construct(ReaderInterface $reader, WriterInterface $writer, string $outFileName)
    {
        $this->reader = $reader;
        $this->writer = $writer;
        $this->outFileName = $outFileName;
    }

    /**
     * @param HandlerInterface $h
     */
    public function setNext(HandlerInterface $h): void
    {
        $this->next = $h;
    }

    /**
     * @param string $fileName
     */
    public function handle(string $fileName): void
    {
        if ($this->next instanceof HandlerInterface) {
            $this->next->handle($fileName);
        }
    }

    /**
     * @param string $fileName
     * @param string $extension
     * @return bool
     */
    protected function canHandle(string $fileName, string $extension): bool
    {
        return strtoupper(pathinfo($fileName, PATHINFO_EXTENSION)) == strtoupper($extension);
    }
}
