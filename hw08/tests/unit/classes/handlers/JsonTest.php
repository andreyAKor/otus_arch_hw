<?php namespace classes\handlers;

use \Codeception\Test\Unit;
use classes\reader\ReaderInterface;
use classes\writer\WriterInterface;
use Mockery;

/**
 * Class JsonTest
 * @package classes\handlers
 */
class JsonTest extends Unit
{
    private const READ_FILE = 'file.json';
    private const READ_CONTENT = 'some json content';
    private const OUT_FILE = 'out_file.txt';

    /**
     *
     */
    public function testDo()
    {
        $reader = Mockery::mock(ReaderInterface::class);
        $reader->allows()
            ->Read(self::READ_FILE)
            ->andReturns([self::READ_CONTENT]);

        $writer = Mockery::mock(WriterInterface::class);
        $writer->allows()
            ->Write(self::OUT_FILE, [self::READ_CONTENT]);

        $handler = new Json($reader, $writer, self::OUT_FILE);
        $handler->handle(self::READ_FILE);
    }
}